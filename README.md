# hkn

[![GoDoc](https://godoc.org/github.com/lukakerr/hkn?status.svg)](https://godoc.org/github.com/lukakerr/hkn)

A go module for interacting with Hacker News.

### Features

A ticked checkbox indicates the feature currently exists in `master`.

An item refers to either a story, comment, ask, job, poll or poll part

- [x] Get a single item
- [x] Get multiple items
- [x] Get largest item id
- [x] Get top 500 new, top and best stories (or a number >= 0, <= 500)
- [x] Get top 200 ask, show and job stories (or a number >= 0, <= 200)
- [x] Get changed items and profiles
- [x] Get a user
- [ ] Get a user's submissions
- [ ] Get a user's comments
- [ ] Get a user's hidden items
- [ ] Get a user's upvoted items
- [ ] Get a user's favorited items
- [x] Login a user
- [x] Upvote an item
- [x] Unvote a comment
- [ ] Downvote a comment
- [x] Create a story
- [ ] Create a poll
- [x] Create a comment
- [ ] Flag an item
- [ ] Hide an item
- [ ] Favorite an item
- [ ] Edit an item
- [ ] Delete an item
- [ ] Search
  - [ ] Full text
  - [ ] By tag
  - [ ] By created at date
  - [ ] By points
  - [ ] By number of comments
  - [ ] By page number
  - [ ] Sorted by relevance, then points, then number of comments
  - [ ] Sorted by most recent

### Usage

First get `hkn`:

```bash
$ go get github.com/lukakerr/hkn
```

Import into your project:

```go
import "github.com/lukakerr/hkn"

// or

import (
        "github.com/lukakerr/hkn"
)
```

#### Methods

Examples of all methods on the client can be found in [example/main.go](./example/main.go).

First create a client:

```go
client := hkn.NewClient()
```

Various methods can then be then called on the client:

**Get a single item by id**

```go
// Returns (Item, error)
item, err := client.GetItem(8869)
```

**Get multiple items by ids**

```go
// Returns ([]Item, error)
items, err := client.GetItems([]int{8869, 8908, 8881, 10403, 9125})
```

**Get max item id**

```go
// Returns (int, error)
id, err := client.GetMaxItemID()
```

**Get the latest item and profile updates**

```go
// Returns (Updates, error)
updates, err := client.GetUpdates()
```

**Get top stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetTopStories(20)
```

**Get new stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetNewStories(20)
```

**Get best stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetBestStories(20)
```

**Get latest ask stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetLatestAskStories(20)
```

**Get latest show stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetLatestShowStories(20)
```

**Get latest job stories given a number**

```go
// Returns ([]int, error)
stories, err := client.GetLatestJobStories(20)
```

**Get a user by id**

```go
// Returns (User, error)
user, err := client.GetUser("jl")
```

**Login a user with a username and password**

```go
// The cookie returned is used for actions that require a user to be logged in
// Returns (*http.Cookie, error)
cookie, err := client.Login("username", "password")
```

**Upvote an item**

> A cookie is required to upvote, get this from logging in

```go
// Returns (bool, error)
upvoted, err := client.Upvote(8869, cookie)
```

**Unvote a comment**

> A cookie is required to unvote, get this from logging in

```go
// Returns (bool, error)
unvoted, err := client.Unvote(8869, cookie)
```

**Create a comment**

> A cookie is required to create a comment, get this from logging in

```go
// Returns (bool, error)
content := "Really cool."
commented, err := client.Comment(8869, content, cookie)
```

**Create a story with a title and URL**

> A cookie is required to create a story, get this from logging in

```go
// Returns (bool, error)
title := "A title."
URL := "https://a.url.com"
created, err := client.CreateStoryWithURL(title, URL, cookie)
```

**Create a story with a title and text**

> A cookie is required to create a story, get this from logging in

```go
// Returns (bool, error)
title := "A title."
text := "Some text."
created, err := client.CreateStoryWithText(title, text, cookie)
```

### Running

To run the example locally:

```bash
$ go run example/main.go
```

### Testing

```bash
$ go test
```
