package lib

import (
	"github.com/go-pdf/fpdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"io"
)

type rFunc func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error)
type pdfRenderer struct {
	pdf         *fpdf.Fpdf
	renderFuncs map[ast.NodeKind]rFunc
}

func (p *pdfRenderer) Render(w io.Writer, source []byte, n ast.Node) error {

	err := ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		if cb, ok := p.renderFuncs[n.Kind()]; ok {
			return cb(source, n, entering)
		}
		return s, nil
	})
	return err
}

func (p *pdfRenderer) AddOptions(option ...renderer.Option) {
	//TODO implement me
	panic("implement me")
}

const p1 = 12
const baseHeight = 6
const fontSize = 14

func Pdf(path string, content string) error {

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	lineIndentLevel := 0

	baseMargin := pdf.GetX()

	renderFuncs := map[ast.NodeKind]rFunc{
		ast.KindHeading: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				pdf.SetFont("Arial", "", 24)
				pdf.MultiCell(0, baseHeight, string(node.Text(source)), "", "", false)
			} else {
				pdf.MultiCell(0, p1, "", "", "", false)
				pdf.SetFont("Arial", "", fontSize)
			}
			return ast.WalkSkipChildren, nil
		},
		ast.KindLink: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				// TODO make link clickable
				pdf.SetTextColor(0, 0, 255)
				pdf.SetFont("Arial", "U", fontSize)
				href := node.(*ast.Link).Destination
				pdf.WriteLinkString(baseHeight, string(node.Text(source)), string(href))
			} else {
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("Arial", "", fontSize)
			}
			return ast.WalkSkipChildren, nil
		},
		ast.KindText: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				pdf.Write(baseHeight, string(node.Text(source)))
			}
			return ast.WalkSkipChildren, nil
		},
		ast.KindParagraph: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if !entering {
				pdf.MultiCell(0, p1, "", "", "", false)
			}
			return ast.WalkContinue, nil
		},
		ast.KindCodeBlock: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				pdf.SetFont("Courier", "", fontSize)
				pdf.MultiCell(0, p1, "", "", "", false)
				pdf.MultiCell(0, baseHeight, string(node.Text(source)), "", "", false)
				pdf.MultiCell(0, p1, "", "", "", false)
			}
			return ast.WalkContinue, nil
		},
		ast.KindList: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				// if it's previous sibling is a paragraph then delete that paragraph margin!!
				if node.PreviousSibling() != nil && node.PreviousSibling().Kind() == ast.KindParagraph {
					pdf.SetY(pdf.GetY() - 1.5*baseHeight)
				}
				lineIndentLevel += 1
			} else {
				pdf.SetLeftMargin(baseMargin)
				lineIndentLevel -= 1
			}
			return ast.WalkContinue, nil
		},
		ast.KindListItem: func(source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
			if entering {
				pdf.Ln(baseHeight)
				pdf.SetLeftMargin(baseMargin + float64(lineIndentLevel)*2)
				pdf.Write(baseHeight, "- ")
			}
			return ast.WalkContinue, nil
		},
	}
	rdr := goldmark.New(goldmark.WithRenderer(&pdfRenderer{pdf, renderFuncs}))

	var wr io.Writer
	err := rdr.Convert([]byte(content), wr)
	if err != nil {
		return err
	}
	err = pdf.OutputFileAndClose(path)
	return err
}
