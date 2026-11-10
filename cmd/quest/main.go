package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	vFlag := flag.Bool("v", false, "print version")
	versionFlag := flag.Bool("version", false, "print version")
	flag.Parse()

	if *vFlag || *versionFlag {
		fmt.Println(getVersion())
		os.Exit(0)
	}
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	var revision, modified string
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			revision = setting.Value
			if len(revision) > 7 {
				revision = revision[:7] // short hash
			}
		case "vcs.modified":
			modified = setting.Value
		}
	}

	if revision == "" {
		return info.Main.Version
	}

	if modified == "true" {
		return fmt.Sprintf("%s-dirty", revision)
	}

	return revision
}
