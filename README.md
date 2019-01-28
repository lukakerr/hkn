# hkn

A go module for interacting with Hacker News.

### Features

A ticked checkbox indicates the feature currently exists in `master`.

An item refers to either a story, comment, ask, job, poll or poll part

- [x] Get a single item
- [x] Get multiple items
- [ ] Get largest item id
- [ ] Get top 500 new, top and best stories
- [ ] Get top 200 ask, show and job stories
- [ ] Get changed items and profiles
- [x] Get a user
- [x] Login a user
- [ ] Logout a user
- [ ] Upvote an item
- [ ] Unvote an item (only comments)
- [ ] Create an item
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

```go
// Get a single item by id
// Returns (Item, error)
item, err := client.GetItem(8869)

// Get multiple items by ids
// Returns ([]Item, error)
items, err := client.GetItems([]int{8869, 8908, 8881, 10403, 9125})

// Get a user by id
// Returns (User, error)
user, err := client.GetUser("jl")

// Login a user with a username and password
// The cookie returned is used for actions that require a user to be logged in
// Returns (*http.Cookie, error)
cookie, err := client.Login("username", "password")
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
