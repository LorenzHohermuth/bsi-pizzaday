package spread

import "fmt"

func SpreadPizzas(m map[string]int, amount int) map[string]int {
	var avr int = amount / len(m)
	delta := amount - (avr * len(m))

	i := 0
	for k := range m {
		m[k] = avr
		if i < delta {
			m[k]++
		}
		i++
	}
	fmt.Println(m)
	return m
}

func CreateMap(arr []string) map[string]int {
	m := make(map[string]int)
	for _, v := range arr {
		m[v] = 0
	}
	return m
}

func Count(mat [][]string, slot string, pizzatype string) int {
	counter := 0
	for _, v := range mat {
		if slot == v[1] && pizzatype == v[2] {
			counter++
		}
	}
	return counter
}
