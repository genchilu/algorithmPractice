package groupShiftedStrings

//import "fmt"

func groupStrings(strings []string) [][]string {
	normailStrMap := make(map[string][]string)
	for _, s := range strings {
		normailStr := normailizeString(s)
		if _, ok := normailStrMap[normailStr];!ok {
			normailStrMap[normailStr] = []string{}
		}
		normailStrMap[normailStr] = append(normailStrMap[normailStr], s)
	}

	result := make([][]string, len(normailStrMap))

	count := 0
	for _, v:=range normailStrMap {
		result[count] = v
		count++
	}
    return result
}

func normailizeString(s string) string {
	c := make([]byte, len(s))
	for i:=0;i<len(s);i++{
		
		c[i] = s[i] - s[0] + 97
		if (s[i] < s[0]) {
			c[i] = 26 - (s[0] - s[i]) + 97
		}
	}
	return string(c)
}