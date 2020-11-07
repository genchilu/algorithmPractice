package searchSuggestionsSystem

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	m := make(map[string][]string)

	for _, product := range products {
		for i := 0; i < len(product); i++ {
			if _, ok := m[product[0:i+1]]; !ok {
				m[product[0:i+1]] = []string{}
			}
			m[product[0:i+1]] = append(m[product[0:i+1]], product)
		}
	}

	results := [][]string{}
	for i := 0; i < len(searchWord); i++ {
		if result, ok := m[searchWord[0:i+1]]; ok {
			sort.Strings(result)
			if len(result) > 3 {
				result = result[0:3]
			}
			results = append(results, result)
		} else {
			results = append(results, []string{})
		}
	}
	return results
}
