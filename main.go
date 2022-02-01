package main

import (
	"context"
	"fmt"

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

	client := getClient()

	_, err := client.CreateKey(ctx, keyName, azkeys.KeyTypes.ECHSM(), nil)
	if err != nil {
		panic(err)
	}
}

func task2(keyName string) {
	// write code that creates a new RSAHSM key
	// use getClient() to obtain the client

	client := getClient()

	_, err := client.CreateKey(ctx, keyName, azkeys.KeyTypes.RSAHSM(), nil)
	if err != nil {
		panic(err)
	}
}

func task3(keyName string) {
	// write code that retrieves a key and prints a message if the key's type is not EC
	// use getClient() to obtain the client

	client := getClient()

	key, err := client.GetKey(ctx, keyName, nil)
	if err != nil {
		panic(err)
	}

	if *key.Key.KeyType != azkeys.KeyTypes.EC() {
		fmt.Println("key is not an EC")
	}
}

func task4(keyName string) {
	// write code that retrieves a key and prints a message if the key's deletion level is recoverable
	// use getClient() to obtain the client

	client := getClient()

	key, err := client.GetKey(ctx, keyName, nil)
	if err != nil {
		panic(err)
	}

	if *key.Attributes.RecoveryLevel != azkeys.DeletionRecoveryLevels.Recoverable() {
		fmt.Println("key is not recoverable")
	}
}

func task5() {
	// write code that detects if the value "unwrap" is a predefined value for JSONWebKeyOperation

	for _, val := range azkeys.JSONWebKeyOperations.Values() {
		if val == "unwrap" {
			fmt.Println("unwrap is a predefined value")
			break
		}
	}
}

func main() {
	task1("key1")
	task2("key1")
	task3("key1")
	task4("key1")
	task5()
}
