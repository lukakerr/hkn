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

	getItem(client)
	// getItems(client)
	// getMaxItemID(client)
	// getUpdates(client)
	// getUser(client)
	// login(client)
	// getTopStories(client)
	// getNewStories(client)
	// getBestStories(client)
	// getLatestAskStories(client)
	// getLatestShowStories(client)
	// getLatestJobStories(client)
	// upvote(client)
	// unvote(client)
	// comment(client)
	// createStoryWithURL(client)
	// createStoryWithText(client)
}

func getItem(client *hkn.Client) {
	item, err := client.GetItem(ids[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", item)
}

func getItems(client *hkn.Client) {
	items, err := client.GetItems(ids)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}

func getMaxItemID(client *hkn.Client) {
	id, err := client.GetMaxItemID()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
}

func getUpdates(client *hkn.Client) {
	updates, err := client.GetUpdates()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", updates)
}

func getUser(client *hkn.Client) {
	user, err := client.GetUser("jl")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", user)
}

func login(client *hkn.Client) {
	// You'll need to use an actual username and password here
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cookie)
}

func getTopStories(client *hkn.Client) {
	stories, err := client.GetTopStories(50)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func getNewStories(client *hkn.Client) {
	stories, err := client.GetNewStories(50)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func getBestStories(client *hkn.Client) {
	stories, err := client.GetBestStories(50)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func getLatestAskStories(client *hkn.Client) {
	stories, err := client.GetLatestAskStories(50)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func getLatestShowStories(client *hkn.Client) {
	stories, err := client.GetLatestShowStories(50)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func getLatestJobStories(client *hkn.Client) {
	stories, err := client.GetLatestJobStories(30)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stories)
}

func upvote(client *hkn.Client) {
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	upvoted, err := client.Upvote(ids[0], cookie)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(upvoted)
}

func unvote(client *hkn.Client) {
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	unvoted, err := client.Unvote(ids[0], cookie)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(unvoted)
}

func comment(client *hkn.Client) {
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	content := "Really cool."
	commented, err := client.Comment(ids[0], content, cookie)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commented)
}

func createStoryWithURL(client *hkn.Client) {
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	title := "test"
	URL := "https://github.com"
	created, err := client.CreateStoryWithURL(title, URL, cookie)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(created)
}

func createStoryWithText(client *hkn.Client) {
	cookie, err := client.Login("username", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	title := "A title"
	text := "Some text"
	created, err := client.CreateStoryWithText(title, text, cookie)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(created)
}
