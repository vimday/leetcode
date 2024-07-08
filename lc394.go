package xl_leetcode

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
			ptr++
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	res := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; (*ptr)++ {
		res += string(s[*ptr])
	}
	return res
}

func getString(v []string) string {
	res := ""
	for _, s := range v {
		res += s
	}
	return res
}
