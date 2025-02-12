package search

import (
	"log"

	"github.com/blugelabs/bluge"
)

func IndexDocument(path, content string) error {
	cfg := bluge.DefaultConfig("pdf_index")
	writer, err := bluge.OpenWriter(cfg)
	if err != nil {
		return err
	}
	defer writer.Close()

	// create a doc
	doc := bluge.NewDocument(path).
		AddField(bluge.NewTextField("content", content).SearchTermPositions())

	// insert doc to idx
	err = writer.Update(doc.ID(), doc)
	if err != nil {
		log.Printf("Failed to index document %v", err)
		return err
	}

	return nil
}
