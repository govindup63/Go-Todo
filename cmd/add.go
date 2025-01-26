package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: add,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("priority", "p", false, "Set the TO-DO to the top of TO-DO list")
}

func add(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	text := ""
	if len(args) > 0 {
		text = args[0]
	}
	for text == "" {
		color.Red("No input Detected Plz Enter your TO-DO")
		input, _ := reader.ReadString('\n')
		if len(input) > 0 && input != "\n" {
			text = input
		}
	}

	isPriority, _ := cmd.Flags().GetBool("priority")

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader2 := csv.NewReader(file)

	records, err := reader2.ReadAll()
	if err != nil {
		panic(err)
	}
	Id := NewId()

	data := [][]string{{strconv.Itoa(Id), strings.TrimSpace(text)}}
	var TodoList [][]string

	if isPriority {
		TodoList = append(data, records...)
	} else {
		TodoList = append(records, data...)
	}

	file, err = os.Create("data.csv")
	if err != nil {
		fmt.Println("Error reopening file for writing:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.WriteAll(TodoList)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
	writer.Flush()
	color.Green("data added✅")
}
