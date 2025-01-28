package cmd

import (
	"encoding/csv"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all your todos",
	Long:  `use this command to list your todos in the current todo list`,
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func printData(Data [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	for _, row := range Data {
		for i, col := range row {
			if i > 0 {
				w.Write([]byte("\t "))
			}
			w.Write([]byte(col))
		}
		w.Write([]byte("\n"))
	}
	w.Flush()
}

func list(cmd *cobra.Command, args []string) {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	printData(records)
}
