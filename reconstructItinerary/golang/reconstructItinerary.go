package reconstructItinerary

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	m := make(map[string]map[string]int)
	cities := []string{}
	total := 0

	for _, ticket := range tickets {
		if _, ok := m[ticket[0]]; !ok {
			m[ticket[0]] = make(map[string]int)
		}
		if _, ok := m[ticket[1]]; !ok {
			m[ticket[1]] = make(map[string]int)
		}
		m[ticket[0]][ticket[1]]++
		total++
	}

	for city, _ := range m {
		cities = append(cities, city)
	}

	sort.Strings(cities)

	itinerary := []string{"JFK"}

	if dfs("JFK", m, total, &itinerary) {
		return itinerary
	}

	return []string{}
}

func dfs(c string, m map[string]map[string]int, total int, itinerary *[]string) bool {
	tocs := []string{}
	for toc, _ := range m[c] {
		tocs = append(tocs, toc)
	}

	sort.Strings(tocs)

	for _, toc := range tocs {
		m[c][toc]--
		total--
		if m[c][toc] == 0 {
			delete(m[c], toc)
		}

		*itinerary = append(*itinerary, toc)
		if dfs(toc, m, total, itinerary) {
			return true
		}

		m[c][toc]++
		total++
		*itinerary = (*itinerary)[0 : len(*itinerary)-1]
	}

	return total == 0
}
