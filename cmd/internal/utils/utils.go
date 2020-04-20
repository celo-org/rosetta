package utils

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
)

func ExitOnError(err error) {
	if err != nil {
		log.Crit("Unknown error", "err", err)
	}
}

func PrettyPrint(value interface{}) {
	marshallRes, err := json.MarshalIndent(value, "", " ")
	ExitOnError(err)
	fmt.Println(string(marshallRes))
}
