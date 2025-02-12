package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/faruqii/FastPDF/pkg/search"
	"rsc.io/pdf"
)

type PDF struct {
	Path    string // PDF file path
	Content string // Extracted text content
}

// Read extracts text content from the PDF file
func (p *PDF) Read() error {
	file, err := os.Open(p.Path)
	if err != nil {
		return fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	reader, err := pdf.NewReader(file, stat.Size())
	if err != nil {
		return fmt.Errorf("failed to create PDF reader: %w", err)
	}

	var textBuffer bytes.Buffer
	numPages := reader.NumPage()

	for i := 1; i <= numPages; i++ {
		page := reader.Page(i)
		if page.V.IsNull() {
			continue
		}

		content := page.Content()

		for _, text := range content.Text {
			textBuffer.WriteString(text.S + " ")
		}
		textBuffer.WriteString("\n")
	}

	p.Content = textBuffer.String()

	err = search.IndexDocument(p.Path, p.Content)
	if err != nil {
		return fmt.Errorf("failed to index document: %w", err)
	}

	return nil
}
