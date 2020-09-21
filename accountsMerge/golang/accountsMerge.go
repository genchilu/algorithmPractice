package accountsMerge

import (
	"sort"
	//"fmt"
)
func accountsMerge(accounts [][]string) [][]string {
	if len(accounts) <= 0 {
		return accounts
	}

	for i, account := range accounts {
		sort.Strings(account[1:])
		deduplicate := []string{}
		for j:=0;j<len(account);{
			e := account[j]
			deduplicate = append(deduplicate, account[j])
			for j<len(account) && account[j] == e {
				j++
			}
		}
		accounts[i] = deduplicate
	}

	merged := [][]string{}
	for i, account1 := range accounts {
		if len(account1) == 0 {
			continue
		}

		for j:=i+1;j<len(accounts);j++ {
			account2 := accounts[j]
			if len(account2) == 0 {
				continue
			}
			shouldMerged := false
			for k:=1;k<len(account2);k++{
				email := account2[k]
				idx := sort.SearchStrings(account1[1:], email) +1
				if idx < len(account1) && account1[idx] == email {
					shouldMerged = true
					break
				}
			}
			if shouldMerged {
				for k:=1;k<len(account2);k++{
					email := account2[k]
					idx := sort.SearchStrings(account1[1:], email) + 1
					if idx == len(account1) {
						account1 = append(account1, email)
					} else if account1[idx] != email {
						tmp := append([]string{}, account1[idx:]...)
						account1 = append(account1[0:idx], email)
						account1 = append(account1, tmp...)
					}
				}

				accounts[j] = []string{}
				accounts[i] = account1
				j=i
			}
		}
		merged = append(merged, account1)
	}

	return merged
}