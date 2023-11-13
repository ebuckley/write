package lib

import (
	"github.com/go-pdf/fpdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"io"
)

type pdfRenderer struct {
	pdf *fpdf.Fpdf
}

func (p *pdfRenderer) Render(w io.Writer, source []byte, n ast.Node) error {

	err := ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		var err error
		if entering {
			switch m := n.(type) {
			case *ast.Heading:
				p.pdf.SetFont("Arial", "", 24)
				p.pdf.MultiCell(0, 5, string(m.Text(source)), "", "", false)
				p.pdf.MultiCell(0, 8, "", "", "", false)
			case *ast.Paragraph:
				p.pdf.SetFont("Arial", "", 14)
				p.pdf.MultiCell(0, 8, string(m.Text(source)), "", "", false)
			case *ast.CodeBlock:
				p.pdf.SetFont("Courier", "", 12)
				p.pdf.MultiCell(0, 5, string(m.Text(source)), "", "", false)
			default:
				// I'm not sure how to handle this... yet
			}
			err = p.pdf.Error()
		}
		return s, err
	})
	return err
}

func (p *pdfRenderer) AddOptions(option ...renderer.Option) {
	//TODO implement me
	panic("implement me")
}

func Pdf(path string, content string) error {

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	rdr := goldmark.New(goldmark.WithRenderer(&pdfRenderer{pdf}))

	var wr io.Writer
	err := rdr.Convert([]byte(content), wr)
	if err != nil {
		return err
	}
	err = pdf.OutputFileAndClose(path)
	return err
}
