package arrayList

import "log"

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	if len(s) == 0 {
		return ""
	}
	str := []byte(s)
	i := len(str) - 1
	j := i
	for i >= 0 {
		if str[j] == '#' && str[i] == '#' {
			i --
			continue
		} else if str[i] != '#' && str[j] != '#' && str[i] != str[j] {
			j --
			continue
		} else if str[i] != '#' && str[j] == '#' {
			str[i] = ' '
		}
		i --
		j --
	}
	ret := make([]byte, 0)
	for _, v := range str {
		if (v == '#') || (v == ' ') {
			continue
		}
		ret = append(ret, v)
	}
	log.Printf("debug: %v", string(ret))
	return string(ret)
}
