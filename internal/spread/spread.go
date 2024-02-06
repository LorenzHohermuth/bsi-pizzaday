package spread


func SpreadPizzas(m map[string]int, amount int) map[string]int {
	return m
}

func createMap(arr []string) map[string]int {
	m := make(map[string]int)
	for _, v := range arr {
		m[v] = 0
	}
	return m
}
