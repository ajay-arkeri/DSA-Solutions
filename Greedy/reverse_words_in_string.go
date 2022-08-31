package greedy

//GFG ---> reverse words in a string ---> i.love.u --> u.love.i

func ReverseWords(s string) string {
	var temp, ans []rune

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			for i2, j := 0, len(temp)-1; i2 < j; i2, j = i2+1, j-1 {
				temp[i2], temp[j] = temp[j], temp[i2]
			}
			ans = append(ans, temp...)
			ans = append(ans, '.')
			temp = []rune{}
		} else {
			temp = append(temp, rune(s[i]))
		}
	}
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	ans = append(ans, temp...)

	return string(ans)
}
