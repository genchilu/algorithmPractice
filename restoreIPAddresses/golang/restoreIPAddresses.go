package restoreIPAddresses

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {

	found := make([][]bool, 4)
	cidrs := make([][][]string, 4)
	for i := 0; i < 4; i++ {
		cidrs[i] = make([][]string, len(s))
		found[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			found[i][j] = false
			cidrs[i][j] = []string{}
		}
	}

	return findAddress(s, 4, &found, &cidrs)
}

func findAddress(s string, idx int, found *[][]bool, cidrs *[][][]string) []string {

	if len(s) == 0 || idx*3 < len(s) {
		return []string{}
	}

	if (*found)[idx-1][len(s)-1] {
		return (*cidrs)[idx-1][len(s)-1]
	}

	if idx == 1 {
		if valid(s) {
			return []string{s}
		}

		return []string{}
	}

	//fmt.Printf("2222 %s, %d\n", s, idx)
	result := []string{}

	for i := 0; i < len(s); i++ {
		if valid(s[0 : i+1]) {
			cidrs := findAddress(s[i+1:], idx-1, found, cidrs)
			for _, cidr := range cidrs {
				result = append(result, s[0:i+1]+"."+cidr)
			}
		}
	}

	(*found)[idx-1][len(s)-1] = true
	(*cidrs)[idx-1][len(s)-1] = result

	return result
}

func valid(s string) bool {
	seg, _ := strconv.Atoi(s)
	if s[0] != '0' && seg <= 255 {
		return true
	}

	if s[0] == '0' && len(s) == 1 {
		return true
	}

	return false
}
