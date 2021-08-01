> This is a fork of the [z7zmey](https://github.com/z7zmey) [parser](https://github.com/VKCOM/php-parser) that adds PHP 8 support.

PHP Parser written in Go
========================

<img src="./parser.jpg" alt="PHP Parser written in Go" width="980"/>

[![GoDoc](https://godoc.org/github.com/VKCOM/php-parser?status.svg)](https://godoc.org/github.com/VKCOM/php-parser)
[![Build Status](https://github.com/VKCOM/php-parser/workflows/Go/badge.svg)](https://github.com/VKCOM/php-parser/workflows/Go)
[![Go Report Card](https://goreportcard.com/badge/github.com/VKCOM/php-parser)](https://goreportcard.com/report/github.com/VKCOM/php-parser)

This project uses [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) and [ragel](https://www.colm.net/open-source/ragel/) tools to create PHP parser. It parses source code into [AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree). It can be used to write static analysis, refactoring, metrics, code style formatting tools.

Features
---------

- Fully support PHP 5, PHP 7 and PHP 8.0 syntax
- Abstract syntax tree (AST) representation
- Traversing AST
- Resolving namespace names
- Parsing syntax-invalid PHP files
- Saving and printing free-floating comments and whitespaces

Who Uses
--------

- [VKCOM/noverify](https://github.com/VKCOM/noverify) — pretty fast linter for PHP
- [VKCOM/nocolor](https://github.com/VKCOM/nocolor) — architecture validation tool for PHP based on the [*concept of colored functions*](https://github.com/VKCOM/nocolor/blob/master/docs/introducing_colors.md)
- [quasilyte/phpgrep](https://github.com/quasilyte/phpgrep) — tool for syntax-aware PHP code search

Usage example
-------

```Golang
package main

import (
	"log"
	"os"

	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/errors"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"github.com/VKCOM/php-parser/pkg/visitor/dumper"
)

func main() {
	src := []byte(`<?php echo "Hello world";`)

	// Error handler

	var parserErrors []*errors.Error
	errorHandler := func(e *errors.Error) {
		parserErrors = append(parserErrors, e)
	}

	// Parse

	rootNode, err := parser.Parse(src, conf.Config{
		Version:          &version.Version{Major: 8, Minor: 0},
		ErrorHandlerFunc: errorHandler,
	})

	if err != nil {
		log.Fatal("Error:" + err.Error())
	}
	
	if len(parserErrors) > 0 {
		for _, e := range parserErrors {
			log.Println(e.String())
		}
		os.Exit(1)
	}

	// Dump

	goDumper := dumper.NewDumper(os.Stdout).
		WithTokens().
		WithPositions()

	rootNode.Accept(goDumper)
}
```

Install
-------

```
go get github.com/VKCOM/php-parser/cmd/php-parser
```

CLI
---

```
php-parser [flags] <path> ...
```

| flag       | type     | description                         |
| ---------- | -------- | ----------------------------------- |
| `--p`      | `bool`   | Print file paths                    |
| `--e`      | `bool`   | Print errors                        |
| `--d`      | `bool`   | Dump AST in Golang format           |
| `--r`      | `bool`   | Resolve names                       |
| `--prof`   | `string` | Start profiler: `[cpu, mem, trace]` |
| `--phpver` | `string` | PHP version (default: 8.0)          |

Namespace resolver
------------------

Namespace resolver is a visitor that resolves nodes fully qualified name and saves into `map[node.Node]string` structure

- For `Class`, `Interface`, `Trait`, `Function`, `Constant` nodes it saves name with current namespace.
- For `Name`, `Relative`, `FullyQualified` nodes it resolves `use` aliases and saves a fully qualified name.
