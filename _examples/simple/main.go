package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tenntenn/connpass"
)

func main() {
	cli := connpass.NewClient()
	params, err := connpass.SearchParam(connpass.Keyword("golang"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	r, err := cli.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range r.Events {
		fmt.Println(e.Title)
	}
}
