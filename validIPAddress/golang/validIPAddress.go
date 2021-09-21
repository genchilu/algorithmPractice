package validIPAddress

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	groups := strings.Split(IP, ".")
	if len(groups) == 4 {
		if isIPV4(groups) {
			return "IPv4"
		}
		return "Neither"
	}

	groups = strings.Split(IP, ":")
	if len(groups) == 8 {
		if isIPV6(groups) {
			return "IPv6"
		}
		return "Neither"
	}

	return "Neither"
}

func isIPV4(groups []string) bool {

	for _, group := range groups {
		if len(group) == 0 {
			return false
		}

		if len(group) > 1 && group[0] == '0' {
			return false
		}

		v, err := strconv.Atoi(group)
		if err != nil {
			return false
		}

		if v > 255 {
			return false
		}
	}
	return true
}

func isIPV6(groups []string) bool {
	for _, group := range groups {
		if len(group) == 0 {
			return false
		}

		if len(group) > 4 {
			return false
		}

		for i := 0; i < len(group); i++ {
			if group[i] >= 48 && group[i] <= 57 {
				continue
			}

			if group[i] >= 65 && group[i] <= 70 {
				continue
			}

			if group[i] >= 97 && group[i] <= 102 {
				continue
			}

			return false
		}
	}
	return true
}
