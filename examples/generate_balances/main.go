package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	bootstrapHeader = "AccountIdentifier_address,Amount_value,Currency_symbol,Currency_decimals"
)

type genesis struct {
	Alloc map[string]genesisAllocation `json:"alloc"`
}
type genesisAllocation struct {
	Balance string `json:"balance"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_balances <genesisURL>")
		os.Exit(1)
	}

	genesisURL := os.Args[1]
	// Get Genesis File
	resp, err := http.Get(genesisURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	genesisAllocations := &genesis{}
	if err := json.Unmarshal(body, genesisAllocations); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", genesisAllocations)
	// Prepare data
	data := [][]string{strings.Split(bootstrapHeader, ",")}
	for k, v := range genesisAllocations.Alloc {
		if v.Balance == "0" {
			continue
		}
		data = append(data, []string{k, v.Balance, "cGLD", "18"})
	}
	// Write to CSV
	file, err := os.Create("bootstrap_balances.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
