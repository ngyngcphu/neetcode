func groupAnagrams(strs []string) [][]string {
    sA := make(map[[26]int][]string)
    for _, s := range strs {
        var count [26]int
        for _, c := range []rune(s) {
            count[c - 'a']++
        }
        sA[count] = append(sA[count], s)
    }

    var res [][]string
    for _, v := range sA {
        res = append(res, v)
    }
    return res
}
