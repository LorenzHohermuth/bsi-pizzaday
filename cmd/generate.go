/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/lorenzhohermuth/bsi-pizzaday/internal/dotfile"
	"github.com/lorenzhohermuth/bsi-pizzaday/internal/spread"
	"github.com/lorenzhohermuth/bsi-pizzaday/pkg/csv"
	"github.com/spf13/cobra"
)

type PizzaSlot struct {
	Slot string
	Type string
	Pizzas map[string]int
}

var availableSlots []string
var availablePizzaTypes []string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, _ := cmd.Flags().GetString("file")
		opt := dotfile.ParseFile(os.Getenv("DOT_DIR") + "\\.pizzaconfig")
		availableSlots = opt.Slots
		availablePizzaTypes = opt.PizzaTypes
		if fileFlag != "" {
			ps := generatePizzaSlots()
			mat := csv.Decode(fileFlag, opt.Location)			
			//vegiMap := spread.CreateMap(opt.VegiPizzas)
			meatMap := spread.CreateMap(opt.MeatPizzas)
			spread.SpreadPizzas(meatMap, 703)
			fmt.Println(mat)
		}
	},
}

func generatePizzaSlots() []PizzaSlot {
	arr := []PizzaSlot{}
	for _, slots := range availableSlots {
		for _, pizzaType := range availablePizzaTypes {
			ps := PizzaSlot {
				Slot: slots,
				Type: pizzaType,
			}
			arr = append(arr, ps)
		}
	}
	return arr
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("file", "", "CSV file to generate Pizzalist")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
