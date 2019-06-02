package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "voy",
    Short: "A CLI for the Voyager Index project.",
    Long: `Search for cities and ranking from the comfort of your shell.
Easily send POST requests to http://voyager-index.herokuapp.com/.`,
    Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
