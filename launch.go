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

	var d contract

	helloHedera, err := ioutil.ReadFile("bin/LookupContract.json")
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
		SetContents(bytecode).Execute(testnetClient)
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
	fmt.Printf("Sol contract address is : %v\n", bytecodefileID.ToSolidityAddress())

	contractTx, err := hedera.NewContractCreateTransaction().
		SetBytecodeFileID(bytecodefileID).SetGas(100000).
		SetConstructorParameters(hedera.NewContractFunctionParameters().AddString("Yo").AddUint32(334343)).
		Execute(testnetClient)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("NEW CONTRACT CREATE TRANSACTION PASSED")
	}

	txRecord, err := contractTx.GetRecord(testnetClient)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("GET RECORD  PASSED")
	}

	newContractID := *txRecord.Receipt.ContractID
	nodeId, _ := hedera.AccountIDFromString("0.0.3")

	contractQuery, err := hedera.NewContractCallQuery().
		SetContractID(newContractID).
		SetGas(100000).
		SetFunction("getMobileNumber", hedera.NewContractFunctionParameters().
			AddString("Yo")).
		SetNodeAccountIDs([]hedera.AccountID{nodeId}).
		Execute(testnetClient)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("GET CONTRACT GET CALL PASSED")
	}

	getMessage := contractQuery.GetUint32(0)
	fmt.Println(getMessage)

	hedera.NewContractExecuteTransaction().
		SetGas(1000000).SetFunction("setMobileNumber", hedera.NewContractFunctionParameters().AddString("Test").AddUint32(334343)).SetNodeAccountIDs([]hedera.AccountID{nodeId}).Execute(testnetClient)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("SET NUMBER CALL PASSED")
	}

	getNumber, err := hedera.NewContractCallQuery().
		SetContractID(newContractID).
		SetGas(1000000).
		SetFunction("getMobileNumber", hedera.NewContractFunctionParameters().
			AddString("Yo")).
		SetNodeAccountIDs([]hedera.AccountID{nodeId}).
		Execute(testnetClient)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(getNumber.GetUint32(0))
	}
}
