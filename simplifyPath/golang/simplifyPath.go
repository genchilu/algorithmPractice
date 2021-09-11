package simplifyPath

import (
	"strings"
)

func simplifyPath(path string) string {

	//fmt.Printf("%v\n", strings.Split(path, "/"))
	result := []string{}
	for _, s := range strings.Split(path, "/") {
		if len(s) > 0 {
			if s == "." {
				continue
			} else if s == ".." {
				if len(result) > 0 {
					result = result[0 : len(result)-1]
				}
			} else {
				result = append(result, s)
			}

		}
	}
	return "/" + strings.Join(result[:], "/")
}
