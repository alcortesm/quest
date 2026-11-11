package cli

import (
	"fmt"
	"io"
	"runtime/debug"
)

// Version subcommand: it prints the cli version.
func Version(out io.Writer) error {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		_, err := fmt.Fprintln(out, "unknown: no build info")
		return err
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

	if revision != "" {
		if modified == "true" {
			_, err := fmt.Fprintf(out, "%s-dirty\n", revision)
			return err
		}
		_, err := fmt.Fprintln(out, revision)
		return err
	}

	if info.Main.Version != "" {
		_, err := fmt.Fprintln(out, info.Main.Version)
		return err
	}

	_, err := fmt.Fprintln(out, "unknown: build info has no version or revision data")
	return err
}
