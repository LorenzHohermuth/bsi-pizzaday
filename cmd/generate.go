/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/lorenzhohermuth/bsi-pizzaday/internal/display"
	"github.com/lorenzhohermuth/bsi-pizzaday/internal/dotfile"
	"github.com/lorenzhohermuth/bsi-pizzaday/internal/spread"
	"github.com/lorenzhohermuth/bsi-pizzaday/pkg/csv"
	"github.com/spf13/cobra"
)

type PizzaSlot struct {
	Slot string
	Type string
}

type SlotSpreadMapping struct {
	Slot PizzaSlot
	PizzaMap map[string]int
}

var availableSlots []string
var availablePizzaTypes []string
var opt dotfile.Options

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
		opt = dotfile.ParseFile(os.Getenv("DOT_DIR") + "\\.pizzaconfig")
		availableSlots = opt.Slots
		availablePizzaTypes = opt.PizzaTypes

		if fileFlag != "" {

			spm := []SlotSpreadMapping{}
			ps := generatePizzaSlots()
			mat := csv.Decode(fileFlag, opt.Location)			

			for _, v := range ps {
				amount, m := count(mat, v)	
				spm = append(spm, 
					SlotSpreadMapping{
						Slot: v,
						PizzaMap: spread.SpreadPizzas(m, amount),
					})
			}
			printOrder(spm)
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

func count(mat [][]string, pizzaSlot PizzaSlot) (int, map[string]int) {
	m := map[string]int{}
	counter := 0
	for _,v := range mat {
		if v[1] == pizzaSlot.Slot && v[2] == pizzaSlot.Type {
			counter++
		}
	}
	if pizzaSlot.Type == availablePizzaTypes[0] {
		m = spread.CreateMap(opt.MeatPizzas)
	}else{
		m = spread.CreateMap(opt.VegiPizzas)
	}
	return counter, m
}

func printOrder(spm []SlotSpreadMapping) {
	s1 := opt.PickUp1 + "\n"
	s1 += display.ListPizzas(spm[0].PizzaMap)
	s1 += display.ListPizzas(spm[1].PizzaMap)
	s2 := "\n" + opt.PickUp2 + "\n"
	s2 += display.ListPizzas(spm[2].PizzaMap)
	s2 += display.ListPizzas(spm[3].PizzaMap)
	fmt.Println(s1 + s2)
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
