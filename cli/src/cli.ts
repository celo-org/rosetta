import {Command, flags} from '@oclif/command'
import fetch from 'node-fetch'
import Web3 from 'web3'

class Query extends Command {
  static flags = {
    web3: flags.string({ description: 'web3 provider' }),
    rosetta: flags.string({ description: 'rosetta url' }),
    transaction: flags.string({ description: 'transaction to query' })
  }

  async run() {
    const res = this.parse(Query)
    if (!res.flags.web3) {
      throw new Error('No web3')
    }
    const web3 = new Web3(res.flags.web3)

    if (!res.flags.transaction) {
      throw new Error('No transaction provided')
    }
    const receipt = await web3.eth.getTransactionReceipt(res.flags.transaction)
    console.log('Receipt:')
    console.log(receipt)
    const transaction = await web3.eth.getTransaction(res.flags.transaction)
    console.log('Transaction:')
    console.log(transaction)

    if (!res.flags.rosetta) {
      throw new Error('No rosetta')
    }
    const query = {
      network_identifier: {
        blockchain: "celo",
        network: "200312"
      },
      block_identifier: {
        index: transaction.blockNumber,
        hash: transaction.blockHash,
      },
      transaction_identifier: {
        hash: res.flags.transaction
      }
    }
    const resp = await fetch(`${res.flags.rosetta}/block/transaction`,
      { method: 'POST', body: JSON.stringify(query) }
    )
    const jsonResp = await resp.json()
    console.log('Rosetta response:')
    console.log(JSON.stringify(jsonResp, null, '  '))
  }
}

Query.run()
