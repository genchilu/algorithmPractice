package numberOfProvinces

func findCircleNum(isConnected [][]int) int {

	isVistited := make([]bool, len(isConnected))
	for i, _ := range isVistited {
		isVistited[i] = false
	}

	c := 0
	for i, _ := range isConnected {
		component := []int{}
		dfs(i, isConnected, &isVistited, &component)
		if len(component) > 0 {
			//fmt.Printf("%v\n", component)
			c++
		}
	}
	return c
}

func dfs(cur int, isConnected [][]int, isVistited *[]bool, component *[]int) {
	if !(*isVistited)[cur] {
		(*isVistited)[cur] = true
		*component = append(*component, cur)
		for i, v := range isConnected[cur] {
			if v == 1 {
				dfs(i, isConnected, isVistited, component)
			}
		}
	}
}
