package lib

import (
	"bytes"
	"github.com/yuin/goldmark"
	"log/slog"
)

func renderMD(md []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(md, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func RenderMD(md string) string {
	m, err := renderMD([]byte(md))
	if err != nil {
		slog.Error("Fatal error parsing markdown", "err", err)
		return md
	}
	return `<div class="markdown-content">` + string(m) + `</div>`
}
