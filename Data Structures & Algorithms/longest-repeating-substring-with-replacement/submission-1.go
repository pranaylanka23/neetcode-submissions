func characterReplacement(s string, k int) int {
    count := make(map[byte]int)
    l, maxf := 0, 0

    for r := 0; r < len(s); r++ {
        count[s[r]]++
        if count[s[r]] > maxf {
            maxf = count[s[r]]
        }

        if (r - l + 1) - maxf > k {
            count[s[l]]--
            l++
        }
    }
	
    return len(s) - l
}