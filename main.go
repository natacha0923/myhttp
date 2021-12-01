package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

func app() error {
	limit := flag.Int("parallel", 10, "limit the number of parallel requests")
	flag.Parse()

	args := flag.Args()
	httpClient := NewHTTPRequester()

	results := Run(args, *limit, httpClient)
	for res := range results {
		if res.Err != nil {
			fmt.Println("ERROR", res.Err)
			continue
		}

		fmt.Println(res.Resp.Url, res.Resp.Hash)
	}

	return nil
}

func Run(args []string, limit int, requester Requester) chan Result {
	if limit < 1 {
		panic("invalid limit")
	}

	results := make(chan Result, limit)
	semaphore := make(chan struct{}, limit)

	client := Client{Requester: requester}

	go func() {
		for _, arg := range args {
			semaphore <- struct{}{}
			go func(url string) {
				defer func() {
					<-semaphore
				}()
				results <- NewResult(client.GetRequestWithMD5(url))
			}(arg)
		}

		for i := 0; i < limit; i++ {
			semaphore <- struct{}{}
		}
		close(results)
	}()

	return results
}

type Result struct {
	Resp Response
	Err  error
}

func NewResult(resp Response, err error) Result {
	return Result{Resp: resp, Err: err}
}
