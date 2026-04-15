func groupAnagrams(strs []string) [][]string {
    hm := make(map[string][]string)

    for _, s := range strs {
        characters := []rune(s)
        sort.Slice(characters, func(i, j int) bool {
            return characters[i] < characters[j]
        })
        hm[string(characters)] = append(hm[string(characters)], s)
    }

    var res [][]string
    for _, v := range hm {
        res = append(res, v)
    }
    return res
}
