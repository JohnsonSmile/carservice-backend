package main

import (
	"carservice/client"
	"log"
)

func main() {

	err := client.NewChainClient()
	if err != nil {
		panic(err)
	}
	log.Printf("chain client initialize successfully")
	user, err := client.GetUser("cmtestuser1")
	if err != nil {
		log.Fatal(err)
	}
	client1AddrInt, client1EthAddr, client1AddrSki, err := client.MakeAddrAndSkiFromCrtFilePath(user.SignCrtPath)
	client.Client.TokenContractBalanceOf(client1AddrInt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ethaddr => %s\n", client1EthAddr)
	log.Printf("skiaddr => %s\n", client1AddrSki)
}
