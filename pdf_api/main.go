package main

import (
	"bytes"
	"github.com/go-chi/chi"
	"github.com/signintech/gopdf"
	"io"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/generate-pdf", generatePDF)

	_ = http.ListenAndServe(":8185", r)
}

func generatePDF(w http.ResponseWriter, r *http.Request) {
	// Create a new PDF document

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// Set font and size
	fontPath := "./fonts/wts43.ttf" // Replace with the path to a TrueType font file
	if err := pdf.AddTTFFont("customFont", fontPath); err != nil {
		http.Error(w, "Error adding font", http.StatusInternalServerError)
		return
	}
	_ = pdf.SetFont("customFont", "", 14)

	// Add content to the PDF
	_ = pdf.Cell(nil, "Hello, World!")

	// Create a buffer to store the PDF content
	var pdfBuffer bytes.Buffer

	// Implement io.WriterTo interface to write the PDF to the buffer
	if _, err := pdf.WriteTo(&pdfBuffer); err != nil {
		http.Error(w, "Error generating PDF", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Disposition", "attachment; filename=sample1.pdf")
	w.Header().Set("Content-Type", "application/pdf")

	// Write the PDF buffer to the response
	_, err := io.Copy(w, &pdfBuffer)
	if err != nil {
		http.Error(w, "Error writing PDF to response", http.StatusInternalServerError)
		return
	}
}
