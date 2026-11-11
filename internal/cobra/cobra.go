package cobra

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/alcortesm/quest/internal/cli"
)

// Do is the main entry point of the quest cli.
func Do(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	if args == nil {
		args = []string{} // nil makes cobra fallback to os.Args
	}

	root := newRootCmd()
	root.SetArgs(args[1:])
	root.SetIn(stdin)
	root.SetOut(stdout)
	root.SetErr(stderr)

	return root.Execute()
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "quest",
		Short:         "Manage quest directories for you.",
		SilenceUsage:  true, // don't dump usage on every runtime error
		SilenceErrors: true, // main prints the error, not cobra
	}

	root.AddCommand(newVersionCmd())

	return root
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cli.Version(cmd.OutOrStdout())
		},
	}

	return cmd
}
