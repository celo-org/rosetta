package examples

import (
	"context"
	"log"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/rosetta/kliento/client"
	"github.com/celo-org/rosetta/kliento/contracts"
	"github.com/celo-org/rosetta/kliento/contracts/helpers"
	"github.com/celo-org/rosetta/kliento/registry"
)

// Watch stable token transfers to specific address
func WatchStableTokenTransfers(client *client.CeloClient, recipient string) {
	recipientAddress := common.HexToAddress(recipient)

	registry, err := registry.New(client)
	if err != nil {
		log.Fatalf("registry instantiation failed %e", err)
	}

	ctx := context.Background()
	stableToken, err := registry.GetStableTokenContract(ctx, nil)
	if err != nil {
		log.Fatalf("stableToken instantiation failed %e", err)
	}

	transferChannel := make(chan *contracts.StableTokenTransfer, 100)
	subscription, err := stableToken.WatchTransfer(&bind.WatchOpts{
		Context: ctx,
		Start:   nil,
	}, transferChannel, nil, []common.Address{recipientAddress})
	if err != nil {
		log.Fatalf("watchTransfer subscription failed %e", err)
	}
	defer subscription.Unsubscribe()
	defer close(transferChannel)

	nextTransfer := func() (*contracts.StableTokenTransfer, error) {
		select {
		case err := <-subscription.Err():
			return nil, err
		case transfer := <-transferChannel:
			return transfer, nil
		}
	}

	log.Printf("Watching transfers...")
	for i := 0; ; i++ {
		transfer, err := nextTransfer()
		if err != nil {
			log.Fatalf("subscription err %e", err)
		}
		transferSlice, err := helpers.EventToSlice(transfer)
		if err != nil {
			log.Fatalf("decoding err %e", err)
		}

		log.Printf("[%d] transfer %v", i, transferSlice)
	}
}
