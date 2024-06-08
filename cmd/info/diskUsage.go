/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

func getInGB(size uint64) uint64 {
	return size / 1024 / 1024 / 1024
}

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Prints disk usage of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")

		fmt.Printf("Total: %v GB\n", getInGB(usage.Size())) // total disk space
		fmt.Printf("Free: %v GB\n", getInGB(usage.Free()))  // disk space available
		fmt.Printf("Used: %v GB\n", getInGB(usage.Used()))  // disk space used
		fmt.Printf("Usage: %v\n", usage.Usage())            // percentage of disk space used

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	InfoCmd.AddCommand(diskUsageCmd)
}
