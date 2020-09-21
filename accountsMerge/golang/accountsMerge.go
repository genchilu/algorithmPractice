package accountsMerge

import (
	"sort"
	//"fmt"
)
func accountsMerge(accounts [][]string) [][]string {
	if len(accounts) <= 0 {
		return accounts
	}

	mailMailMap := make(map[string]map[string]bool)

	for _, acc:=range accounts {
		for i:=1;i<len(acc);i++{
			if _,ok:=mailMailMap[acc[1]];!ok {
				mailMailMap[acc[1]]= make(map[string]bool)
			}
			mailMailMap[acc[1]][acc[i]] = true

			if _,ok:=mailMailMap[acc[i]];!ok {
				mailMailMap[acc[i]]= make(map[string]bool)
			}
			mailMailMap[acc[i]][acc[1]] = true
		}

	}

	result := [][]string{}
	used := make(map[string]bool)
	for _, acc := range accounts {
		if len(acc) == 1 {
			result = append(result, []string{acc[0]})
		} else {
			if _,ok:=used[acc[1]];!ok {
				mergeAcc := []string{acc[0]}
				_dfs(mailMailMap, used, acc[1], &mergeAcc)
				sort.Strings(mergeAcc[1:])
				result = append(result, mergeAcc)
			}
		}
	}

	return result
}

func _dfs(mailMailMap map[string]map[string]bool, used map[string]bool, mail string, mails *[]string) {
	if _,ok:=used[mail];!ok {
		used[mail] = true
		*mails = append(*mails, mail)
		for e,_:= range mailMailMap[mail] {
			_dfs(mailMailMap, used, e, mails)
		}
		//fmt.Printf("%v\n", mails)
	}
} 