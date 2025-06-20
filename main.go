package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const creditScoreMin = 500
const creditScoreMax = 900

type creditRating struct {
	CreditRating int `json:"credit_rating"`
}

func getCreditScore(w http.ResponseWriter, r *http.Request) {
	var creditRating = creditRating{
		CreditRating: rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(creditRating)
	if err != nil {
		return
	}
}

func handleRequests() {
	http.Handle("/creditscore", http.HandlerFunc(getCreditScore))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
