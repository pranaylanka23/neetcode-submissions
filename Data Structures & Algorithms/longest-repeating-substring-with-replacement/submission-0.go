func characterReplacement(s string, k int) int {
    count := make(map[byte]int)
    l, maxf := 0, 0

    for r := 0; r < len(s); r++ {
        count[s[r]]++
        if count[s[r]] > maxf {
            maxf = count[s[r]]
        }

        // Changed from 'for' to 'if'
        // If the window becomes invalid, we just shift the whole window right by 1
        if (r - l + 1) - maxf > k {
            count[s[l]]--
            l++
        }
    }

    // The window size at the end represents the maximum valid window found
    return len(s) - l
}