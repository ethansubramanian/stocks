package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
	gocb "gopkg.in/couchbase/gocb.v1"
)

type Stock struct {
	stockPrice float32 `json:"Stock Price"`
}

func getPrice() (price float32) {
	q, _ := quote.Get("AAPL")
	curPrice := float32(q.RegularMarketPrice)
	return curPrice
}

var bucket *gocb.Bucket
var bucketName string

func main() {
	//Connecting to the database
	cluster, _ := gocb.Connect("couchbase://127.0.0.1") //Connects to the database (My server)
	//To resolve any authentication errors while logging into the database (My database)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "AAPL",
		Password: "!23456L8",
	})
	bucketName = "AAPLstock"
	bucket, _ := cluster.OpenBucket(bucketName, "")
	//-----------------------------------------------
	//Gets stock price
	thePrice := getPrice()
	theStock := Stock{stockPrice: thePrice}
	//jsonFile ,_ := json.Marshal(theStock)
	var priceParams []interface{}
	priceParams = append(priceParams, theStock.stockPrice)
	fmt.Println(thePrice)  //Simply prints the price of stock
	_, err := bucket.Upsert("Price", &priceParams, 0)
	if err != nil {
		fmt.Println("BUCKET ERROR: ", err)
	}

}
