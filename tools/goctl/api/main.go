package main

import (
	"fmt"
	"os"

	"github.com/wjames2000/go-zero/core/logx"
	"github.com/wjames2000/go-zero/tools/goctl/api/parser"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	p, err := parser.NewParser(os.Args[1])
	logx.Must(err)
	api, err := p.Parse()
	logx.Must(err)
	fmt.Println(api)
}
