package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/profile"
	"github.com/yookoala/realpath"

	"github.com/VKCOM/php-parser/pkg/ast"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/errors"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"github.com/VKCOM/php-parser/pkg/visitor/dumper"
	"github.com/VKCOM/php-parser/pkg/visitor/nsresolver"
	"github.com/VKCOM/php-parser/pkg/visitor/printer"
	"github.com/VKCOM/php-parser/pkg/visitor/traverser"
)

var wg sync.WaitGroup
var phpVersion *version.Version
var profiler string
var dump *bool
var showResolvedNs *bool
var printBack *bool
var printPath *bool
var printErrors *bool
var printExecTime *bool

type file struct {
	path    string
	content []byte
}

type result struct {
	path     string
	rootNode ast.Vertex
	errors   []*errors.Error
}

func main() {
	start := time.Now()
	var phpVer string

	printExecTime = flag.Bool("time", false, "print execution time")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	printBack = flag.Bool("pb", false, "print AST back into the parsed file")
	printPath = flag.Bool("p", false, "print filepath")
	printErrors = flag.Bool("e", false, "print errors")
	dump = flag.Bool("d", false, "dump")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem, trace]")
	flag.StringVar(&phpVer, "phpver", "7.4", "php version")

	flag.Parse()

	var err error
	phpVersion, err = version.New(phpVer)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	switch profiler {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "trace":
		defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

	numCpu := runtime.GOMAXPROCS(0)

	fileCh := make(chan *file, numCpu)
	resultCh := make(chan result, numCpu)

	// run 4 concurrent parserWorkers
	for i := 0; i < numCpu; i++ {
		go parserWorker(fileCh, resultCh)
	}

	// run printer goroutine
	go printerWorker(resultCh)

	// process files
	processPath(flag.Args(), fileCh)

	// wait the all files done
	wg.Wait()
	close(fileCh)
	close(resultCh)

	elapsed := time.Since(start)
	if *printExecTime {
		log.Printf("took: %s", elapsed)
	}
}

func processPath(pathList []string, fileCh chan<- *file) {
	for _, path := range pathList {
		real, err := realpath.Realpath(path)
		checkErr(err)

		err = filepath.Walk(real, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(path) == ".php" {
				wg.Add(1)
				content, err := ioutil.ReadFile(path)
				checkErr(err)
				fileCh <- &file{path, content}
			}
			return nil
		})
		checkErr(err)
	}
}

func parserWorker(fileCh <-chan *file, r chan<- result) {
	for {
		f, ok := <-fileCh
		if !ok {
			return
		}

		var parserErrors []*errors.Error
		rootNode, err := parser.Parse(f.content, conf.Config{
			Version: phpVersion,
			ErrorHandlerFunc: func(e *errors.Error) {
				parserErrors = append(parserErrors, e)
			},
		})
		if err != nil {
			fmt.Println("Error:" + err.Error())
			os.Exit(1)
		}

		r <- result{path: f.path, rootNode: rootNode, errors: parserErrors}
	}
}

func printerWorker(r <-chan result) {
	var counter int

	for {
		res, ok := <-r
		if !ok {
			return
		}

		counter++

		if *printPath {
			_, _ = io.WriteString(os.Stderr, "==> ["+strconv.Itoa(counter)+"] "+res.path+"\n")
		}

		if *printErrors {
			for _, e := range res.errors {
				_, _ = io.WriteString(os.Stderr, "==> "+e.String()+"\n")
			}
		}

		if *printBack {
			o := bytes.NewBuffer([]byte{})
			p := printer.NewPrinter(o)
			res.rootNode.Accept(p)

			err := ioutil.WriteFile(res.path, o.Bytes(), 0644)
			checkErr(err)
		}

		if *showResolvedNs {
			v := nsresolver.NewNamespaceResolver()
			traverser.NewTraverser(v).Traverse(res.rootNode)
			for _, n := range v.ResolvedNames {
				_, _ = io.WriteString(os.Stderr, "===> "+n+"\n")
			}
		}

		if *dump == true {
			dumper.NewDumper(os.Stdout).WithPositions().WithTokens().Dump(res.rootNode)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
