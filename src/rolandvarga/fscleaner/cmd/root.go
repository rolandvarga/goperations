// Copyright © 2018 Roland Varga <roland.varga.can@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
    "fmt"
    "os"

    // homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    // log "github.com/sirupsen/logrus"
)

var srcDir string
var keepCnt int
var sftDel, hrdDel, rvsList bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "fscleaner",
    Short: "fscleaner is a CLI utility designed for housekeeping.",
    Long: `fscleaner is a CLI utility designed for housekeeping. It allows
for quick deletion or cleanup in your folders. Gives you the option of
moving your items in your '.Trash' & '/tmp' folders, or to get rid of 
your files for good. Please see the available commands for further 
information.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    rootCmd.PersistentFlags().StringVarP(&srcDir, "dir", "d", "", "Directory to cleanse or delete")
    rootCmd.PersistentFlags().IntVarP(&keepCnt, "keep", "", 0, "Number of child objects to keep in directory")
    rootCmd.PersistentFlags().BoolVarP(&sftDel, "soft", "", true, "Enables soft deletion")
    rootCmd.PersistentFlags().BoolVarP(&hrdDel, "hard", "", false, "Enables hard deletion")
    rootCmd.PersistentFlags().BoolVarP(&rvsList, "reverse", "", false, "Reverse list of child objects (delete new first)")
}

