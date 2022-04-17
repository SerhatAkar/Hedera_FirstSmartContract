package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	accountID, err := hedera.AccountIDFromString(os.Getenv("MY_ACCOUNT_ID"))
	if err != nil {
		panic(err.Error())
	}

	privateKey, err := hedera.PrivateKeyFromString(os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err.Error())
	}

	var testnetClient *hedera.Client = hedera.ClientForTestnet()
	testnetClient.SetOperator(accountID, privateKey)

	type contractStruct struct {
		Object    string `json:"object"`
		OpCodes   string `json:"opcodes"`
		SourceMap string `json:"sourceMap"`
	}
	type bc struct {
		Bytecode contractStruct `json:"bytecode"`
	}
	type doc struct {
		Data bc `json:"data"`
	}

	var d doc = doc{}

	helloHedera, err := ioutil.ReadFile("./bin/LookupContract.json")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("JSON READING PASSED")
	}
	if err := json.Unmarshal([]byte(helloHedera), &d); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("UNMARSHAL FUNCTION PASSED")
	}

	bytecode := []byte(d.Data.Bytecode.Object)

	fileTx, err := hedera.NewFileCreateTransaction().
		SetContents([]byte(bytecode)).Execute(testnetClient)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("NEWFILECREATETRANSACTION PASSED")
	}

	fileReceipt, err := fileTx.GetReceipt(testnetClient)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("GET RECEIPT PASSED")
	}

	bytecodefileID := *fileReceipt.FileID

	fmt.Printf("Contract bytecode file: %v\n", bytecodefileID)

}
