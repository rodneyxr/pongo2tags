package pongo2tags

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/flosch/pongo2"
)

var static_url = "/static/"

func init() {
	RegisterTag("static", staticTag)
}

func StaticURL(url string) {
	static_url = strings.TrimRight(url, "/ ")
}

type tagStaticNode struct {
	path string
}

func (node *tagStaticNode) Execute(ctx *ExecutionContext, writer TemplateWriter) *Error {
	writer.WriteString(static_url + node.path)
	return nil
}

func staticTag(doc *Parser, start *Token, arguments *Parser) (INodeTag, *Error) {
	static_node := &tagStaticNode{}

	if arguments.Count() != 1 {
		return nil, arguments.Error("static tag has "+strconv.Itoa(arguments.Count())+" arguments.", nil)
	}

	static_node.path = arguments.Get(0).Val
	fmt.Println(static_node.path)

	return static_node, nil
}
