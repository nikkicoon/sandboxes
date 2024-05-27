package main

import (
	"fmt"
	"context"
	"github.com/google/go-github/github"
)

func main () {
	client := github.NewClient(nil)
	orgs, _, _ := client.Organizations.List(context.Background(), "nikicoon", nil)
	fmt.Printf("orgs: %s\n", orgs)
}
