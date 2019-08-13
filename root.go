package main

import (
	"io"

	"github.com/spf13/cobra"
)

type GlobalOptions struct {
	Verbose 	bool
	Insecure	bool
	userAgent	string
	Input		io.Reader
	Output		io.Writer
}

type RootCommand struct {
	cmd  *cobra.Command
	opts *GlobalOptions
}

func NewRootCommand(r io.Reader, w io.Writer) *RootCommand {
	c := &RootCommand{
		cmd: &cobra.Command{
			Use:   "grpcurl",
			Short: "A handy and universal gRPC command line client",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
		},
		opts: &GlobalOptions{
			Input:  r,
			Output: w,
		},
	}
	c.cmd.PersistentFlags().BoolVarP(&c.opts.Verbose, "verbose", "v", false, "verbose output")
	c.cmd.PersistentFlags().StringVarP(&c.opts.userAgent, "user-agent", "u", "", "with user-agent")
	c.cmd.PersistentFlags().BoolVarP(&c.opts.Insecure, "insecure", "k", false, "with insecure")
	c.cmd.AddCommand(NewListServicesCommand(c.opts).Command())
	c.cmd.AddCommand(NewCallCommand(c.opts).Command())
	return c
}

func (c *RootCommand) Command() *cobra.Command {
	return c.cmd
}
