# Simple scripts for investigating reconciliation failures.
# Usage help: python3 scripts/reconciliation_debugging.py --help

import requests
import argparse


def wrap_response(response):
    if response.status_code != 200:
        raise ValueError("Request failed: ", response.text)
    return response.json()


def network_req_from_chain_id(chain_id):
    return {"blockchain": "celo", "network": f"{chain_id}"}


def request_account_balance(node_url, network_id, address, block_number):
    return wrap_response(
        requests.post(
            f"{node_url}/account/balance",
            json={
                "network_identifier": network_id,
                "account_identifier": {
                    "address": address,
                },
                "block_identifier": {"index": block_number},
            },
        )
    )


def request_block(node_url, network_id, block_number):
    """
    Query blockscout API for block_tx
    """
    req_json = {
        "network_identifier": network_id,
        "block_identifier": {"index": block_number},
    }
    return wrap_response(requests.post(f"{node_url}/block", json=req_json))


def request_block_tx(node_url, network_id, address, block_number, block_hash, tx_hash):
    return wrap_response(
        requests.post(
            f"{node_url}/block/transaction",
            json={
                "network_identifier": network_id,
                "account_identifier": {
                    "address": address,
                },
                "block_identifier": {"index": block_number, "hash": block_hash},
                "transaction_identifier": {"hash": tx_hash},
            },
        )
    )


def account_balance(
    block_number,
    node_url,
    reconcile_address,
    network_id,
):
    # get account/balance balance
    resp = request_account_balance(
        node_url, network_id, reconcile_address, block_number
    )

    balances = resp["balances"]

    # if not, let's investigate this
    assert len(balances) == 1
    return int(balances[0]["value"])


def block_tx_balance_change(
    block_number,
    block_hash,
    tx_hash,
    node_url,
    reconcile_address,
    network_id,
):
    resp = request_block_tx(
        node_url,
        network_id,
        reconcile_address,
        block_number,
        block_hash,
        tx_hash,
    )
    try:
        ops = resp["transaction"].get("operations", [])  # or []
        return sum(
            [
                int(op["amount"]["value"])
                for op in ops
                if op["account"]["address"] == reconcile_address
                and op["status"] == "success"
            ]
        )

    except Exception as e:
        print(f"Exception for response: {resp}")
        raise e


# Useful when trying to shrink the range to search over and reconcile,
# i.e. when it's not yet known when exactly the reconciliation error occurs.
def scan_account_balance_changes(
    start_block,
    end_block,
    step,
    node_url,
    reconcile_address,
    network_id,
):
    balance = None
    balance_changes = []
    for block_number in range(start_block, end_block, step):
        curr_balance = account_balance(
            block_number, node_url, reconcile_address, network_id
        )
        if balance != curr_balance:
            print(f"balance change at: {block_number}, curr_balance: {curr_balance}")
            balance_changes.append(balance)
        balance = curr_balance
    return balance_changes


# Performs targetted balance reconciliation (same idea as Rosetta CLI checks)
# over a restricted range and for a particular address.
def reconcile_range(
    start_block,
    end_block,
    node_url,
    reconcile_address,
    network_id,
):
    assert start_block > 0
    curr_balance = account_balance(
        start_block - 1, node_url, reconcile_address, network_id
    )
    print(f"starting balance: {curr_balance}")

    for i in range(start_block, end_block):
        print(f"beginning block {i}")
        # fetch account_balance changes
        new_balance = account_balance(i, node_url, reconcile_address, network_id)
        block = request_block(node_url, network_id, i)
        block_hash = block["block"]["block_identifier"]["hash"]

        sum_ops_change = sum(
            [
                block_tx_balance_change(
                    i,
                    block_hash,
                    tx["hash"],
                    node_url,
                    reconcile_address,
                    network_id,
                )
                for tx in block.get("other_transactions", [])
            ]
        )
        assert (
            curr_balance + sum_ops_change == new_balance
        ), f"failure for block {i}, curr_balance: {curr_balance}, new_balance: {new_balance}, sum_ops_change: {sum_ops_change}"
        curr_balance = new_balance


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("start", type=int, help="starting block in range")
    parser.add_argument("end", type=int, help="ending block in range")
    parser.add_argument("chain_id", type=int, help="chain ID in genesis.json")
    parser.add_argument(
        "address", type=str, help="address whose balance to reconcile over range"
    )
    parser.add_argument(
        "-n", "--node_url", default="http://localhost:8080", help="rosetta RPC URL"
    )
    args = parser.parse_args()
    network_req = network_req_from_chain_id(args.chain_id)

    reconcile_range(args.start, args.end, args.node_url, args.address, network_req)
