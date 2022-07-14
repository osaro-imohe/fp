package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

type Transaction struct {
	ID              string  `csv:"id"`
	Amount          float64 `csv:"amount"`
	BankCountryCode string  `csv:"bank_country_code"`
}

type Latencies map[string]int

//go:embed latencies.json
var latenciescsv []byte

// naive approach
// sort by descending transaction amount volume
// greater come first in array
// return first transactions with a cummulative latence less than 1000ms

// better approach
// tag each transaction with a ratio of amount to transaction latency
// sort by transaction amount to latency ratio

// best approach
// dynamic programming 2d array
// store the solutions of solved subproblems in 2d array
// compare (maxValueIn2dArray + curr) against maxValueIn2dArray
// max value is the last element in the matrix e.g values[len(transactions)][totalTime]

func prioritizeTransactions(transactions []Transaction, totalTime int) []Transaction {
	var latencies Latencies
	err := json.Unmarshal(latenciescsv, &latencies)
	if err != nil {
		log.Println(err)
		return nil
	}

	values := make([][]float64, len(transactions)+1)
	for i := range values {
		values[i] = make([]float64, totalTime+1)
	}

	keep := make([][]int, len(transactions)+1)
	for i := range keep {
		keep[i] = make([]int, totalTime+1)
	}

	var results []Transaction

	for i := int64(0); i < int64(totalTime)+1; i++ {
		values[0][i] = 0
		keep[0][i] = 0
	}

	for i := 0; i < len(transactions)+1; i++ {
		values[i][0] = 0
		keep[i][0] = 0
	}

	for i := 1; i <= len(transactions); i++ {
		for j := int(1); j <= int(totalTime); j++ {
			maxValWithoutCurr := values[i-1][j]
			maxValWithCurr := float64(0)

			weightOfCurr := latencies[transactions[i-1].BankCountryCode]

			if j >= weightOfCurr {
				maxValWithCurr = transactions[i-1].Amount
				remainingCapacity := j - weightOfCurr
				maxValWithCurr += values[i-1][remainingCapacity]
			}

			if maxValWithCurr > maxValWithoutCurr {
				values[i][j] = maxValWithCurr
				keep[i][j] = 1
			} else {
				values[i][j] = maxValWithoutCurr
				keep[i][j] = 0
			}
		}
	}

	log.Printf("The max USD value that can be processed in %vms is $%v", totalTime, values[len(transactions)][totalTime])

	n := len(transactions)
	c := totalTime

	for n > 0 {
		if keep[n][c] == 1 {
			results = append(results, transactions[n-1])
			c -= latencies[transactions[n-1].BankCountryCode]
		}
		n--
	}
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
