package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	type Msg struct {
		message string
	}

	type URL struct {
		url   string
		depth int
	}
	type Quit struct{}

	ch := make(chan interface{}, 20)
	crawler := func(url string, depth int) {
		defer func() { ch <- Quit{} }()

		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			ch <- Msg{fmt.Sprintf("%s\n", err)}
			return
		}

		ch <- Msg{fmt.Sprintf("found: %s %q \n", url, body)}

		for _, u := range urls {
			ch <- URL{u, depth - 1}
		}
	}

	work := 0
	memo := make(map[string]bool)

	ch <- URL{url, depth}

	for {
		req := <-ch

		switch req := req.(type) {
		case Msg:
			fmt.Println(req.message)
		case URL:
			if req.depth > 0 && !memo[req.url] {
				memo[req.url] = true
				work++

				go crawler(req.url, req.depth)
			}
		case Quit:
			work--
		}

		if work <= 0 {
			break
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Second * 5)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programing Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/ps/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
