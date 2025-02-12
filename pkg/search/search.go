package search

import (
	"context"
	"fmt"
	"log"

	"github.com/blugelabs/bluge"
)

// SearchDocuments performs a BM25-ranked search on the PDF index
func SearchDocuments(keyword string) {
	// Open Bluge index
	config := bluge.DefaultConfig("pdf_index")
	reader, err := bluge.OpenReader(config)
	if err != nil {
		log.Fatalf("Failed to open index: %v", err)
	}
	defer reader.Close()

	// Create a BM25 query
	q := bluge.NewMatchQuery(keyword).
		SetField("content").
		SetBoost(1.2) // BM25 boosting factor

	req := bluge.NewTopNSearch(10, q).
		WithStandardAggregations()

	// Execute search
	iter, err := reader.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	// Print results
	match, err := iter.Next()
	for err == nil && match != nil {
		fmt.Printf("Found match in document: %s\n", match)
		match, err = iter.Next()
	}
}
