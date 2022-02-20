package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return nil
	}
	return restoreIpAddressesGreedy(s)
}

func restoreIpAddressesGreedy(s string) []string {
	length := len(s)
	sb := []byte(s)
	results := make([]string, 0)
	path := make([]string, 0)
	var restoreIpAddressesCore func(int, int)
	restoreIpAddressesCore = func(begin, end int) {
		if begin > end && len(path) == 4 {
			tmp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			results = append(results, tmp)
			return
		}
		for i := begin; i <= end; i ++ {
			if isInvalidIpAddr(sb, begin, i) {
				path = append(path, string(sb[begin:i + 1]))
				restoreIpAddressesCore(i + 1, end)
				// greedy
				path = path[:len(path) - 1]
			} else {
				break
			}
		}
	}
	restoreIpAddressesCore(0, length - 1)
	return results
}

func isInvalidIpAddr(in []byte, begin, end int) bool {
	if begin > end {
		return false
	}
	if in[begin] == '0' && end != begin {
		return false
	}
	num := 0
	for i := begin; i <= end; i ++ {
		if in[i] < '0' || in[i] > '9' {
			return false
		}
		num = num * 10 + int(in[i] - '0')
		if num > 255 {
			return false
		}
	}
	return true
}


func TestRestoreIpAddresses(t *testing.T) {
	in := "25525511135"
	t.Log(restoreIpAddresses(in))
}

func TestIsInvalidIpAddr(t *testing.T) {
	in := "128"
	assert.Equal(t, true, isInvalidIpAddr([]byte(in), 0, len(in) - 1))

	in = "255"
	assert.Equal(t, true, isInvalidIpAddr([]byte(in), 0, len(in) - 1))

	in = "256"
	assert.Equal(t, false, isInvalidIpAddr([]byte(in), 0, len(in) - 1))

	in = "1"
	assert.Equal(t, true, isInvalidIpAddr([]byte(in), 0, len(in) - 1))
}
