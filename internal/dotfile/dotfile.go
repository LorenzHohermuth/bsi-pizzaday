package dotfile

import (
	"os"
	"strings"
)

var opt Options

type Options struct {
	VegiPizzas	[]string
	MeatPizzas	[]string
	PizzaTypes	[]string
	Slots				[]string
	Location		string
	PickUp1			string
	PickUp2			string
}

func ParseFile(dotfilePath string) Options {
	dat, err := os.ReadFile(dotfilePath)
	check(err)
	lines := strings.Split(string(dat), "\n")
	for _,v := range lines {
		setValue(v)
	}
	return opt
}

func setValue(line string) {
	if strings.TrimSpace(line) == "" {
		return
	} 
	lineSplit := strings.Split(line, "=")
	key := strings.TrimSpace(lineSplit[0])
	value := strings.TrimSpace(lineSplit[1]) 
	readValue(key, value)
}

func readValue(key string, value string){
	switch key {
	case "pizza.vegi":
		value = strings.Trim(value, "[]")
		arr := strings.Split(value, ",")
		for i := range arr {
			arr[i] = strings.TrimSpace(arr[i])	
		}
		opt.VegiPizzas = arr

	case "pizza.meat":
		value = strings.Trim(value, "[]")
		arr := strings.Split(value, ",")
		for i := range arr {
			arr[i] = strings.TrimSpace(arr[i])	
		}
		opt.MeatPizzas = arr

	case "available.slots":
		value = strings.Trim(value, "[]")
		arr := strings.Split(value, ",")
		for i := range arr {
			arr[i] = strings.TrimSpace(arr[i])	
		}
		opt.Slots = arr

	case "available.pizzaTypes":
		value = strings.Trim(value, "[]")
		arr := strings.Split(value, ",")
		for i := range arr {
			arr[i] = strings.TrimSpace(arr[i])	
		}
		opt.PizzaTypes = arr

	case "bsi.location":
		opt.Location = value
	
	case "slot1.timeToPickUp":
		opt.PickUp1 = value

	case "slot2.timeToPickUp":
		opt.PickUp2 = value
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
