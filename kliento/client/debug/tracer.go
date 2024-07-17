// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package debug

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/common/hexutil"
	"github.com/celo-org/celo-blockchain/eth/tracers"
)

var transferTracerTimeout = `50s`

var TransferTracer = `
// Geth Tracer that outputs cGLD transfers.
//
// Additional details (e.g. transaction hash & gas used) can be obtained from 
// the block at the corresponding transaction index.

{
  callStack: [ { transfers: [] } ],
  statusRevert: 'revert',
  statusSuccess: 'success',

  topCall() {
    return this.callStack[this.callStack.length - 1];
  },

  assertStackHeightEquals(expected, shouldThrow, info) {
    const msg = "Unexpected stack height. Expected: " + expected + " Actual: " + this.callStack.length + " Additional info: " + info
    if (this.callStack.length != expected) {
      if (shouldThrow) { 
        throw new Error(msg);
      }
      console.warn(msg);
    }
  },

  pushTransfers(targetTransfers, sourceTransfers, transferStatus) {
    for (var index in sourceTransfers) {
      const transfer = sourceTransfers[index];
      // Successful transfers become reverted if any ancestor call reverts.
      if (transfer.status != this.statusRevert) {
        transfer.status = transferStatus;
      }
      targetTransfers.push(transfer);
    }
  },

  // fault() is invoked when the actual execution of an opcode fails.
  fault(log, db) {
    this.assertStackHeightEquals(log.getDepth(), true, "");
    this.topCall().reverted = true;
  },

  // step() is invoked for every opcode that the VM executes.
  step(log, db) {
    const depth = log.getDepth()

    if (this.callStack.length > depth) {
      const finishedCall = this.callStack.pop();
      const ret = log.stack.peek(0);

      // Execution of opcode failed and returned 0
      if (ret.equals(0)) {
        finishedCall.reverted = true;
      }

      // Find to address for nested contract create with value.
      if ((finishedCall.op == 'CREATE' || finishedCall.op == 'CREATE2') &&
          finishedCall.transfers.length > 0 && !finishedCall.transfers[0].to) {
        const createTransfer = finishedCall.transfers[0];
        createTransfer.to = toHex(toAddress(ret.toString(16)));
        createTransfer.status = finishedCall.reverted ? this.statusRevert : this.statusSuccess;
      }

      // Propagate transfers made during the successful call.
      this.pushTransfers(this.topCall().transfers, finishedCall.transfers,
                         finishedCall.reverted ? this.statusRevert : this.statusSuccess);
    }

    this.assertStackHeightEquals(depth, true, "");

    // Capture any errors immediately.
    const error = log.getError();
    if (error !== undefined) {
      this.fault(log, db);
    } else {
      const op = log.op.toString();
      switch (op) {
        case 'REVERT':
          this.fault(log, db);
          break;

        case 'CREATE':
        case 'CREATE2':
          this.callStack.push({ op, transfers: [] })
          this.handleCreate(log, op);
          break;

        case 'SELFDESTRUCT':
          this.handleDestruct(log, db);
          break;

        case 'CALL':
        case 'CALLCODE':
        case 'STATICCALL':
        case 'DELEGATECALL':
          this.callStack.push({ transfers: [] })
          if (op != 'STATICCALL') {
            this.handleCall(log, op);
          }
          break;
      }
    }
  },
  
  handleCreate(log, op) {
    valueBigInt = bigInt(log.stack.peek(0));
    if (valueBigInt.gt(0)) {
      this.topCall().transfers.push({
        type: 'nested cGLD create contract transfer',
        from: toHex(log.contract.getAddress()),
        value: '0x' + valueBigInt.toString(16),
      });
    }
  },

  handleDestruct(log, db) {
    const contractAddress = log.contract.getAddress();
    const valueBigInt = db.getBalance(contractAddress)
    if (valueBigInt.gt(0)) {
      this.topCall().transfers.push({
        type: 'cGLD destroy contract transfer',
        from: toHex(contractAddress),
        to: toHex(toAddress(log.stack.peek(0).toString(16))),
        value: '0x' + valueBigInt.toString(16),
      });
    }
  },

  handleCall(log, op) {
    const to = toAddress(log.stack.peek(1).toString(16));
    if (!isPrecompiled(to)) {
      if (op != 'DELEGATECALL') {
        valueBigInt = bigInt(log.stack.peek(2));
        if (valueBigInt.gt(0)) {
          this.topCall().transfers.push({
            type: 'cGLD nested transfer',
            from: toHex(log.contract.getAddress()),
            to: toHex(to),
            value: '0x' + valueBigInt.toString(16),
          });
        }
      }
    } else if (toHex(to) == '0x00000000000000000000000000000000000000fd') {
      // This is the transfer precompile "address", inspect its arguments.
      const stackOffset = 1;
      const inputOffset = log.stack.peek(2 + stackOffset).valueOf();
      const inputLength = log.stack.peek(3 + stackOffset).valueOf();
      const inputEnd = inputOffset + inputLength;
      const input = toHex(log.memory.slice(inputOffset, inputEnd));
      const valueBigInt = bigInt(input.slice(2+64*2, 2+64*3), 16);

      this.topCall().transfers.push({
        type: 'cGLD transfer precompile',
        from: '0x'+input.slice(2+24, 2+64),
        to: '0x'+input.slice(2+64+24, 2+64*2),
        value: '0x'+valueBigInt.toString(16),
      });
    }
  },

  // result() is invoked when all the opcodes have been iterated over and returns
  // the final result of the tracing.
  //
  // Note that no errors should be thrown from this function because they will
  // override any other errors thrown from this script and timeout errors
  // managed by geth.
  result(ctx, db) {
    const rootCall = this.topCall();
    const transfers = []
    this.pushTransfers(transfers, rootCall.transfers,
                       rootCall.reverted ? this.statusRevert : this.statusSuccess);

    const create = ctx.type == 'CREATE' || ctx.type == 'CREATE2';
    if (ctx.type == 'CALL' || create) {
      valueBigInt = bigInt(ctx.value.toString());
      if (valueBigInt.gt(0)) {
        transfers.unshift({
          type: create ? 'cGLD create contract transfer' : 'cGLD transfer',
          from: toHex(ctx.from),
          to: toHex(ctx.to),
          value: '0x' + valueBigInt.toString(16),
          status: rootCall.reverted ? this.statusRevert : this.statusSuccess,
        });
      }
    }

    // Return in same format as callTracer: -calls, +transfers.
    return {
      type:      ctx.type,
      from:      toHex(ctx.from),
      to:        toHex(ctx.to),
      value:     '0x' + ctx.value.toString(16),
      gas:       '0x' + bigInt(ctx.gas).toString(16),
      gasUsed:   '0x' + bigInt(ctx.gasUsed).toString(16),
      block:     ctx.block,
      time:      ctx.time,
      transfers: transfers,
    };
  },
}`

type TransferTracerResponse struct {
	Transfers []Transfer `json:"transfers"`
}

type TransferStatus string

const (
	TransferStatusSuccess TransferStatus = "success"
	TransferStatusRevert  TransferStatus = "revert"
)

func (ts TransferStatus) String() string { return string(ts) }

type Transfer struct {
	From   common.Address `json:"from"`
	To     common.Address `json:"to"`
	Value  *big.Int       `json:"value"`
	Status TransferStatus `json:"status"`
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transfer) UnmarshalJSON(input []byte) error {
	type Transfer struct {
		From   common.Address `json:"from"`
		To     common.Address `json:"to"`
		Value  *hexutil.Big   `json:"value"`
		Status TransferStatus `json:"status"`
	}
	var dec Transfer
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	t.From = dec.From
	t.To = dec.To
	t.Value = (*big.Int)(dec.Value)
	t.Status = dec.Status
	if dec.Value == nil {
		return errors.New("missing required field 'value' for Transfer")
	}
	return nil
}

// TransactionTransfers is deprecated, use DebugClient.TraceTransaction with TransferTracer and TransferTracerResponse instead.
// e.g.
//  var res TransferTracerResponse
//  var timeout = "50s"
//  conf := &eth.TraceConfig{Tracer: &TransferTracer, Timeout: &timeout}
//  err := dc.TraceTransaction(ctx, &res, txhash, conf)
func (dc *DebugClient) TransactionTransfers(ctx context.Context, txhash common.Hash) ([]Transfer, error) {
	tracerConfig := &tracers.TraceConfig{Timeout: &transferTracerTimeout, Tracer: &TransferTracer}
	var response TransferTracerResponse

	err := dc.TraceTransaction(ctx, &response, txhash, tracerConfig)
	if err != nil {
		return nil, err
	}
	return response.Transfers, nil
}
