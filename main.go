package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
)

func main() {
	//define the data

	//read the template
	var templ *template.Template
	var err error

	templ, err = template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	// apply the parsed HTML template data and keep the result in a Buffer
	var doc bytes.Buffer
	err = templ.Execute(&doc, ReturnData())
	if err != nil {
		panic(err)
	}

	//convert html to pdf
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(&doc))
	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	//save the pdf
	err = pdfg.WriteFile("output.pdf")
	if err != nil {
		panic(err)
	}

	//convert pdf to base64
	base := base64.StdEncoding.EncodeToString(pdfg.Bytes())
	fmt.Println(base)

	//output the pdf
	err = pdfg.WriteFile("output.pdf")
	if err != nil {
		panic(err)
	}
}
