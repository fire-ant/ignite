package cmdutil

import "github.com/spf13/pflag"

// This file contains a collection of common flag adders
func AddNameFlag(fs *pflag.FlagSet, name *string) {
	fs.StringVarP(name, "name", "n", *name, "Specify the name")
}

func AddInteractiveFlag(fs *pflag.FlagSet, interactive *bool) {
	fs.BoolVarP(interactive, "interactive", "i", *interactive, "Attach to the VM after starting")
}

func AddImportKernelFlags(fs *pflag.FlagSet, kernelName *string) {
	fs.StringVarP(kernelName, "import-kernel", "k", *kernelName, "Import a new kernel from the image with the specified name")
}