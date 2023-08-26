package pdfgenerator

import (
	response "github.com/NicholasLiem/ModulAjar_Backend/http"
	"github.com/go-pdf/fpdf"
	"net/http"
	"os"
)

func GenPDFHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	outputFolder := os.Getenv("OUTPUT_PATH")
	outputPath := outputFolder + "/hello.pdf"

	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.Mkdir(outputFolder, os.ModePerm)
		if err != nil {
			response.ErrorResponse(rw, http.StatusInternalServerError, "Could not create output folder: "+err.Error())
			return
		}
	}

	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		response.ErrorResponse(rw, http.StatusInternalServerError, "Could not generate pdf: "+err.Error())
		return
	}

	response.SuccessResponse(rw, http.StatusCreated, "A new pdf is generated", outputPath)
	return
}
