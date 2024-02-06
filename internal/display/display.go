package display

import "fmt"

func ListPizzas(m map[string]int) string {
	s := ""
	for k, v := range m {
		s += fmt.Sprintf("  %s : %d\n", k, v)
	}
	return s
}

