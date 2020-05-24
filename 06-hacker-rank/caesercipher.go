package lib

func CaeserCipher(s string, k int) string {
	kr := rune(k)
	o := make([]rune, 0)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			o = append(o, (((r-'a')+kr)%26)+'a')
		} else if r >= 'A' && r <= 'Z' {
			o = append(o, (((r-'A')+kr)%26)+'A')
		} else {
			o = append(o, r)
		}
	}
	return string(o)
}
