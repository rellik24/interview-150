package interview150

import (
	"fmt"
	"sort"
	"strings"
)

// Group Anagrams
// https://leetcode.com/problems/group-anagrams/submissions/1496007466/?envType=study-plan-v2&envId=top-interview-150
func GroupAnagrams() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	m := make(map[string][]string)
	for _, v := range strs {
		tmp := strings.Split(v, "")
		sort.Strings(tmp)
		val := strings.Join(tmp, "")

		m[val] = append(m[val], v)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
