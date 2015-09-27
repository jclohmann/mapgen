package main

import (
	"flag"
	"fmt"
	"github.com/jclohmann/mapgen"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"
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

	commentRegexp := regexp.MustCompile("\\+mapgen[ ]*type=([a-zA-Z0-9]+)([ ]*(targets=(.+))?)")
	for _, comment := range sourceFile.Comments {
		text := comment.Text()
		match := commentRegexp.FindStringSubmatch(text)

		if len(match) > 0 {
			typeName := match[1]
			if typeName == "" {
				fmt.Println("Missing type")
				return
			}
			var targets []string
			if len(match) > 3 {
				targets = strings.Split(match[4], ",")
			}
			mg := mapgen.NewMapGen(filepath.Dir(*source), sourceFile.Name.String(), typeName, targets)
			err := mg.Generate()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println("done.")
}
