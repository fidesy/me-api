package main

import (
	"context"
	"log"

	"github.com/fidesy/me-api/pkg/api"
)

func main() {
	c := api.New(&api.Config{})

	stats, err := c.GetActivities(context.TODO(), "primates", 0, 5)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range stats {
		log.Println(s)
	}
}
