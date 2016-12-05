func find(l, r, ans, beg int, s string) (int, int) {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l -= 1
        r += 1
    }
    if r-l-1 > ans {
        ans = r-l-1
        beg = l+1
    }
    return ans, beg
}

func longestPalindrome(s string) string {
    sz, ans, beg := len(s), 0, 0
    for i := 0; i < sz; i += 1 {
        ans, beg = find(i, i+1, ans, beg, s);
        ans, beg = find(i-1, i+1, ans, beg, s);
    }
    return s[beg:beg+ans]
}
