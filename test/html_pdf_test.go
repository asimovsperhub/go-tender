package test

import (
	pdf "github.com/adrg/go-wkhtmltopdf"
	"log"
	"os"
	"testing"
)

func Test_HtmlToPdf(t *testing.T) {
	// Initialize library.
	if err := pdf.Init(); err != nil {
		log.Fatal(err)
	}
	defer pdf.Destroy()

	// Create object from file.
	//object, err := pdf.NewObject("sample1.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//object.Header.ContentCenter = "[title]"
	//object.Header.DisplaySeparator = true

	// Create object from URL.
	//object2, err := pdf.NewObject("https://google.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//object2.Footer.ContentLeft = "[date]"
	//object2.Footer.ContentCenter = "Sample footer information"
	//object2.Footer.ContentRight = "[page]"
	//object2.Footer.DisplaySeparator = true

	// Create object from reader.
	inFile, err := os.Open("sample2.html")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	object3, err := pdf.NewObjectFromReader(inFile)
	if err != nil {
		log.Fatal(err)
	}
	object3.Zoom = 1.5
	object3.TOC.Title = "Table of Contents"

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()
	converter.Add(object3)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create("out.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Run converter.
	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
}

func Test_PdfToWord(t *testing.T) {
}
