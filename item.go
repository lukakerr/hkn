package hkn

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Item struct {
	Id      int    `json:"id"`
	Deleted bool   `json:"deleted"`
	Type    string `json:"type"`
	By      string `json:"by"`
	Time    int32  `json:"time"`
	Text    string `json:"text"`
	Dead    bool   `json:"dead"`
	Parent  int    `json:"parent"`
	Kids    []int  `json:"kids"`
	URL     string `json:"url"`
	Score   int    `json:"score"`
	Title   string `json:"title"`
	Parts   []int  `json:"parts"`
}

type Items []Item

// Get an item given an id
func GetItem(id int, url string) (Item, error) {
	reqUrl := fmt.Sprintf("%s/%s/%d", url, "item", id)

	resp, err := GetBody(reqUrl)

	var item Item

	if err != nil {
		return item, err
	}

	err = json.Unmarshal(resp, &item)
	return item, err
}

// Get items given a slice of ids
// This function is parallelised and thus does not guarantee order
func GetItems(ids []int, url string) (Items, error) {
	var (
		items Items
		wg    sync.WaitGroup
	)

	// Add length of ids to the WaitGroup
	wg.Add(len(ids))

	for _, id := range ids {
		// Spawn a thread for each item
		go func(id int, url string) {
			defer wg.Done()

			item, err := GetItem(id, url)

			if err != nil {
				return
			}

			items = append(items, item)
		}(id, url)
	}

	// Wait until all threads are done
	wg.Wait()

	return items, nil
}