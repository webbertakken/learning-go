package main

import (
	"fmt"
	arg "github.com/alexflint/go-arg"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"
)

func discoverDirs(wl *worklist.WorkList, path string) {
	dirEnts, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}

	for _, dirEnt := range dirEnts {
		path := filepath.Join(path, dirEnt.Name())
		if dirEnt.IsDir() {
			discoverDirs(wl, path)
		} else {
			wl.Add(worklist.NewJob(path))
		}
	}
}

var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional"`
}

// Run with `go run ./mgrep func .`
func main() {
	arg.MustParse(&args)

	var workersGroup sync.WaitGroup
	workList := worklist.New(100)

	results := make(chan worker.Result, 100)
	numWorkers := 10

	workersGroup.Add(1)
	go func() {
		defer workersGroup.Done()
		discoverDirs(&workList, args.SearchDir)
		workList.Finalise(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersGroup.Add(1)

		go func() {
			defer workersGroup.Done()

			for {
				workEntry := workList.Next()
				if workEntry.Path == "" {
					return
				}

				workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
				if workerResult != nil {
					for _, result := range workerResult.Inner {
						results <- result
					}
				}
			}
		}()
	}

	blockWorkersGroup := make(chan struct{})
	go func() {
		workersGroup.Wait()
		close(blockWorkersGroup)
	}()

	var displayGroup sync.WaitGroup

	displayGroup.Add(1)
	go func() {
		for {
			select {
			case result := <-results:
				fmt.Printf("%v[%v]:%v\n", result.Path, result.LineNum, result.Line)
			case <-blockWorkersGroup:
				if len(results) == 0 {
					displayGroup.Done()
					return
				}
			}
		}
	}()

	displayGroup.Wait()
}
