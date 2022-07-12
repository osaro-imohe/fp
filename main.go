package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/gocarina/gocsv"
)

type Transaction struct {
	ID              string  `csv:"id"`
	Amount          float64 `csv:"amount"`
	BankCountryCode string  `csv:"bank_country_code"`
}

type Latencies map[string]int

type Result struct {
	ID         string
	Fraudulent bool
}

//go:embed latencies.json
var latenciescsv []byte

// naive approach
// sort by descending transaction amount volume
// greater come first in array
// return first transactions with a cummulative latence less than 1000ms

// better approach
// tag each transaction with a ratio of amount to transaction latency
// sort by transaction amount to latency ratio

func prioritizeTransactions(transactions []Transaction, totalTime int) []Transaction {
	var latencies Latencies
	err := json.Unmarshal(latenciescsv, &latencies)
	if err != nil {
		log.Println(err)
		return nil
	}
	sort.SliceStable(transactions, func(i, j int) bool {
		currRatio := transactions[i].Amount / float64(latencies[transactions[i].BankCountryCode])
		nextRatio := transactions[j].Amount / float64(latencies[transactions[j].BankCountryCode])
		return currRatio > nextRatio
	})
	var results []Transaction
	var index int
	var count float64
	var time int
	for time < totalTime {
		time += latencies[transactions[index].BankCountryCode]
		if time > totalTime {
			break
		}
		results = append(results, transactions[index])
		count += transactions[index].Amount
		index += 1
	}

	log.Printf("The max USD value that can be processed in %vms is $%v", totalTime, count)

	return results
}

func main() {
	results, err := os.Open("./transactions.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer results.Close()
	if len(os.Args) < 2 {
		log.Println("please input a total time in ms")
		return
	}
	var totalTime int
	totalTime, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println("please input a total time in ms")
		return
	}
	transactions := []Transaction{}
	if err := gocsv.UnmarshalFile(results, &transactions); err != nil {
		log.Println(err)
		return
	}

	res := prioritizeTransactions(transactions, totalTime)
	log.Println(res)
}
