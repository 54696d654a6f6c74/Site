package services

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) []byte {
	p := parser.New()

	parsed := p.Parse(md)

	opts := html.RendererOptions{Flags: html.CommonFlags}
	r := html.NewRenderer(opts)

	return markdown.Render(parsed, r)
}
