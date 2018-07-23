package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/dotnews/conduit/pipeline"
	"github.com/dotnews/conduit/queue"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("Running conduit")
	cRoot := os.Getenv("CRAWLER_ROOT")

	for _, absPath := range findPipelines(cRoot) {
		glog.Infof("Loading pipeline: %s", absPath)

		pipeline.New(
			filepath.Dir(absPath),
			absPath,
			queue.New(1*time.Second),
		).Run()
	}

	select {}
}

func findPipelines(cRoot string) []string {
	pipelines := []string{}

	err := filepath.Walk(cRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "pipeline.yml" {
			pipelines = append(pipelines, path)
		}

		return err
	})

	if err != nil {
		glog.Fatalf("Failed walking directory: %s", cRoot)
	}

	return pipelines
}
