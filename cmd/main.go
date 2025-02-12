package main

import (
	"fmt"

	"github.com/faruqii/FastPDF/pkg/parser"
	"github.com/faruqii/FastPDF/pkg/search"
)

func main() {
	pdf := parser.PDF{Path: "example.pdf"}
	err := pdf.Read()
	if err != nil {
		fmt.Println("failed to read document")
		return
	}

	search.SearchDocuments("Golang")
}
