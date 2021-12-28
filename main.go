package main

import (
	"fmt"
	"os"
	"time"

	"github.com/motemen/go-pocket/api"
)

func main() {
	consumerKey := os.Getenv("POCKET_API_CONSUMER_KEY")
	accessKey := os.Getenv("POCKET_API_ACCESS_KEY")

	client := api.NewClient(consumerKey, accessKey)

	options := &api.RetrieveOption{
		Sort:        api.SortOldest,
		ContentType: api.ContentTypeArticle,
		Since:       int(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Unix()),
		DetailType:  api.DetailTypeComplete,
	}
	res, err := client.Retrieve(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	itemsByTag := map[string][]api.Item{}
	itemsByTag["untagged"] = []api.Item{}

	for _, item := range res.List {
		if len(item.Tags) == 0 {
			itemsByTag["untagged"] = append(itemsByTag["untagged"], item)
			continue
		}

		for tag, _ := range item.Tags {
			if _, ok := itemsByTag[tag]; ok {
				itemsByTag[tag] = append(itemsByTag[tag], item)
			} else {
				itemsByTag[tag] = []api.Item{item}
			}
			// 記事が重複してほしくないので一回でbreakする
			break
		}
	}

	for tag, items := range itemsByTag {
		fmt.Printf("### %s\n", tag)

		for _, item := range items {
			fmt.Printf("- [%s](%s)\n", item.ResolvedTitle, item.ResolvedURL)
		}
	}

}
