package main

import (
	"flag"
	"fmt"
	"github.com/jclohmann/mapgen"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
)

func main() {
	source := flag.String("source", "", "Sourcefile to parse.")
	flag.Parse()

	if *source == "" {
		fmt.Println("You have to specify a source-file.")
		return
	}

	fset := token.NewFileSet()
	sourceFile, err := parser.ParseFile(fset, *source, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	commentRegexp := regexp.MustCompile("\\+mapgen[ ]*(type=([a-zA-Z0-9]+))")
	for _, comment := range sourceFile.Comments {
		text := comment.Text()
		match := commentRegexp.FindStringSubmatch(text)
		if len(match) > 0 {
			typeName := match[2]
			if typeName == "" {
				fmt.Println("Missing type")
				return
			}
			mg := mapgen.NewMapGen(filepath.Dir(*source), sourceFile.Name.String(), typeName)
			err := mg.Generate()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println("done.")
}
