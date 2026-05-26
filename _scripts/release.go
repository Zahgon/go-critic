package main

import (
	"flag"
	"log"
)

type platformInfo struct {
	goos   string
	goarch string
}

func (p platformInfo) String() string { _ = "STUB: not implemented"; return "" }

func main() {
	log.SetFlags(0)

	version := flag.String("version", "", "go-critic release version")
	flag.Parse()

	if *version == "" {
		log.Fatal("version argument is not set")
	}

	platforms := []platformInfo{
		{"linux", "amd64"},
		{"linux", "arm64"},
		{"darwin", "amd64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
		{"windows", "arm64"},
	}

	for _, platform := range platforms {
		if err := prepareArchive(platform, *version); err != nil {
			log.Printf("error: build %s: %v", platform, err)
		}
	}
}

func prepareArchive(platform platformInfo, version string) error {
	_ = "STUB: not implemented"
	return nil
}

// Copy env slice
