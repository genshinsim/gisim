package main

// SOURCE: https://github.com/arkavo-com/pb-go-tag-bson
// MIT LICENSED

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	var directory string
	var verbose bool
	flag.StringVar(&directory, "dir", ".", "directory to search for pb.go files")
	flag.BoolVar(&verbose, "verbose", false, "verbosity")
	flag.Parse()
	filepathErr := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ignore all files that have not been generated by pb
		if !strings.HasSuffix(path, ".pb.go") {
			return nil
		}
		originalSize := info.Size()
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("could not parse go file: %v", err)
		}
		var v visitor
		// debug
		// spew.Dump(node)
		// modify tags
		ast.Walk(v, node)
		// output file
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("could not open output file: %v", err)
		}
		defer f.Close()
		err = format.Node(f, fset, node)
		if err != nil {
			return fmt.Errorf("could not write output file: %v", err)
		}
		if verbose {
			newInfo, err := os.Stat(path)
			if err != nil {
				fmt.Println(err)
			}
			if originalSize != newInfo.Size() {
				fmt.Print("modified:")
			}
			fmt.Println(path)
		}
		return nil
	})
	if filepathErr != nil {
		fmt.Printf("could not correctly handle directory: %v", filepathErr)
	}
}

// visitor to handle every ast node
type visitor struct{}

// Visit find fields starting with XXX_ and add bson tag
// If field has json tag, then replicate it as bson tag
func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	_, ok := n.(*ast.Field)
	if ok {
		ast.Inspect(n, func(subn ast.Node) bool {
			switch subd := subn.(type) {
			case *ast.Field:
				if f, ok := n.(*ast.Field); ok {
					if len(f.Names) == 0 {
						return false
					}
					// skip non public fields
					r := []rune(f.Names[0].Name)
					if len(r) == 0 {
						return false
					}
					ru, _ := utf8.DecodeRuneInString(f.Names[0].Name)
					return unicode.IsUpper(ru)
				}
			case *ast.BasicLit:
				existingTagValue := strings.Trim(subd.Value, "`")
				bsonTags := []string{"bson:\"-\""}
				// if no tag original, add ignore flag
				if existingTagValue == "" {
					subd.Value = fmt.Sprintf("`%s`", strings.Join(bsonTags, " "))
					return true
				}
				// from: https://github.com/golang/protobuf/issues/1388
				// the json tags should be ignored
				// instead, need to extract it out of the proto fields, specifically json=xxx
				existingTags := strings.Split(existingTagValue, " ")
				name := "-"
				for _, existingTag := range existingTags {
					existingTagParts := strings.Split(existingTag, ":")
					if existingTagParts[0] == "protobuf" {
						protoTagsValue := strings.Trim(existingTagParts[1], `"`)
						protoTags := strings.Split(protoTagsValue, ",")
						for _, existingProtoTag := range protoTags {
							if strings.HasPrefix(existingProtoTag, "name=") && name == "-" {
								name = strings.TrimPrefix(existingProtoTag, "name=")
							}
							if strings.HasPrefix(existingProtoTag, "json=") {
								name = strings.TrimPrefix(existingProtoTag, "json=")
							}
						}
					}
				}
				bsonTags = []string{fmt.Sprintf(`bson:"%v,omitempty"`, name)}
				mergedTags := existingTags
				mergedTags = append(mergedTags, bsonTags...)
				subd.Value = fmt.Sprintf("`%s`", strings.Join(mergedTags, " "))
			}
			return true
		})
	}
	return v
}
