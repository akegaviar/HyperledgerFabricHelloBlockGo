package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions to put Hello Block on ledger
type SmartContract struct {
	contractapi.Contract
}

// Hello allows for strings to put Hello Block on ledger
type Hello struct {
	Block string `json:"block"`
}

// runHello puts Hello Block on ledger
func (s *SmartContract) runHello(ctx contractapi.TransactionContextInterface, helloRecord string, block string) error {
	hello := Hello{
		Block: block,
	}

	helloAsBytes, _ := json.Marshal(hello)

	return ctx.GetStub().PutState(helloRecord, helloAsBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error creating hellBlock chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting hellBlock chaincode: %s", err.Error())
	}
}
