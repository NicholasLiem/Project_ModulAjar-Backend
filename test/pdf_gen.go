package main

import (
	"github.com/go-pdf/fpdf"
)

var gofpdfDir string

func main() {
	pdf := fpdf.New(fpdf.OrientationPortrait, "mm", "A4", "")
	pdf.SetCompression(true)
	//
	//template := pdf.CreateTemplate(func(tpl *fpdf.Tpl) {
	//	tpl.SetMargins(10, 30, 10)
	//	tpl.Image("avatar.png", 10, 10, 30, 0, false, "", 0, "")
	//	tpl.SetFont("Arial", "B", 16)
	//	tpl.Text(40, 20, "Template says hello")
	//	tpl.SetDrawColor(0, 100, 200)
	//	tpl.SetLineWidth(2.5)
	//	tpl.Line(92, 12, 105, 22)
	//})
	//
	//template2 := pdf.CreateTemplate(func(tpl *fpdf.Tpl) {
	//	tpl.UseTemplate(template)
	//	subtemplate := tpl.CreateTemplate(func(tpl2 *fpdf.Tpl) {
	//		tpl2.Image("avatar.png", 6, 86, 30, 0, false, "", 0, "")
	//		tpl2.SetFont("Arial", "B", 16)
	//		tpl2.Text(40, 100, "Subtemplate says hello")
	//		tpl2.SetDrawColor(0, 200, 100)
	//		tpl2.SetLineWidth(2.5)
	//		tpl2.Line(102, 92, 112, 102)
	//	})
	//	tpl.UseTemplate(subtemplate)
	//})
	//
	//_, tplSize := template.Size()
	//
	//b, _ := template2.Serialize()
	//template3, _ := fpdf.DeserializeTemplate(b)
	//
	//pdf.AddPage()
	//pdf.UseTemplate(template3)
	//pdf.UseTemplateScaled(template3, fpdf.PointType{X: 0, Y: 30}, tplSize)

	const (
		fontSize = 12
		halfX    = 105
	)

	pdf.AddPage()
	pdf.SetFont("Arial", "", fontSize)
	_, lineHt := pdf.GetFontSize()

	pdf.Write(lineHt, "Hello World!")
	pdf.SetX(halfX)
	pdf.Write(lineHt, "This is standard text.\n")
	pdf.Ln(lineHt * 2)

	pdf.SubWrite(10, "H", 33, 0, 0, "")
	pdf.Write(10, "ello World!")
	pdf.SetX(halfX)
	pdf.Write(10, "This is text with a capital first letter.\n")
	pdf.Ln(lineHt * 2)

	pdf.SubWrite(lineHt, "Y", 6, 0, 0, "")
	pdf.Write(lineHt, "ou can also begin the sentence with a small letter. And word wrap also works if the line is too long, like this one is.")
	pdf.SetX(halfX)
	pdf.Write(lineHt, "This is text with a small first letter.\n")
	pdf.Ln(lineHt * 2)

	pdf.Write(lineHt, "The world has a lot of km")
	pdf.SubWrite(lineHt, "2", 6, 4, 0, "")
	pdf.SetX(halfX)
	pdf.Write(lineHt, "This is text with a superscripted letter.\n")
	pdf.Ln(lineHt * 2)

	pdf.Write(lineHt, "The world has a lot of H")
	pdf.SubWrite(lineHt, "2", 6, -3, 0, "")
	pdf.Write(lineHt, "O")
	pdf.SetX(halfX)
	pdf.Write(lineHt, "This is text with a subscripted letter.\n")

	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		return
	}

}
