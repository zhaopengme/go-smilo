// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

const errorFile = "../../../gonum/blas/gonum/errors.go"

func main() {
	path, err := filepath.Abs(errorFile)
	if err != nil {
		log.Fatalf("no absolute path for %q: %v", errorFile, err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("failed to parse %q: %v", path, err)
	}

	dst := filepath.Base(errorFile)
	o, err := os.Create(dst)
	if err != nil {
		log.Fatalf("failed to create %q: %v", dst, err)
	}
	defer o.Close()

	fmt.Fprintln(o, header)
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			fmt.Fprintln(o, c.Text)
		}
		break
	}
	fmt.Fprintln(o, pkg)
	p := printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 8,
	}
	// Remove comment associated with the const block.
	for _, d := range f.Decls {
		if d, ok := d.(*ast.GenDecl); ok {
			d.Doc = nil
		}
	}
	p.Fprint(o, fset, f.Decls)
	fmt.Fprintln(o)
}

const (
	header = `// Code generated by "go generate gonum.org/v1/netlib/blas/netlib”; DO NOT EDIT.
`
	pkg = `
package netlib

// Copied from gonum/blas/gonum. Keep in sync.`
)
