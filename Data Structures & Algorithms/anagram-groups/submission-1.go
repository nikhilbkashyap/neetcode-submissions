
func groupAnagrams(strs []string) [][]string {

	res := make(map[string][]string)

	for _, s := range strs {
		chars := strings.Split(s, "")
		sort.Strings(chars)
		sortedS := strings.Join(chars, "")
		res[sortedS] = append(res[sortedS], s)
	}
	result := make([][]string, 0, len(res))
	for _, group := range res {
		result = append(result, group)
	}

	return result
}