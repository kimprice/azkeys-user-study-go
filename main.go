package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
)

var (
	ctx = context.Background()
)

func getClient() *azkeys.Client {
	client, err := azkeys.NewClient("vaultURL", nil, nil)
	if err != nil {
		panic(err)
	}

	return client
}

func task1(keyName string) {
	// write code that creates a new ECHSM key
	// use getClient() to obtain the client

}

func task2(keyName string) {
	// write code that creates a new RSAHSM key
	// use getClient() to obtain the client

}

func task3(keyName string) {
	// write code that retrieves a key and prints a message if the key's type is not EC
	// use getClient() to obtain the client

}

func task4(keyName string) {
	// write code that retrieves a key and prints a message if the key's deletion level is recoverable
	// use getClient() to obtain the client

}

func task5() {
	// write code that detects if the value "unwrap" is a predefined value for JSONWebKeyOperation

}

func main() {
	task1("key1")
	task2("key1")
	task3("key1")
	task4("key1")
	task5()
}
