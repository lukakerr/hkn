package main

import (
	"fmt"
	"github.com/lukakerr/hkn"
)

var ids = []int{8869, 8908, 8881, 10403, 9125}

// Various examples of how to interact with the hkn module
// Uncomment each function call to see the result
func main() {
	client := hkn.NewClient()

	GetItem(client)
	// GetItems(client)
	// GetUser(client)
	// Login(client)
}

func GetItem(client *hkn.Client) {
	item, err := client.GetItem(ids[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", item)
}

func GetItems(client *hkn.Client) {
	items, err := client.GetItems(ids)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}

func GetUser(client *hkn.Client) {
	user, err := client.GetUser("jl")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", user)
}

func Login(client *hkn.Client) {
	// You'll need to use an actual username and password here
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cookie)
}
