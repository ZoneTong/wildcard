//github.com/ZoneTong/wildcard
package wildcard

import (
	"strings"
)

// wildcard like in Makefile
// simple support rules: * ? [...]
// not support: '\'
func ExactMatch(pattern, str string) bool {
	lp := len(pattern)
	ls := len(str)
	if lp == 0 && ls == 0 {
		return true
	} else if lp == 0 || ls == 0 {
		return false
	}

	p := pattern[0]
	if p == '*' {
		for i := 0; i <= ls; i++ {
			if ExactMatch(pattern[1:], str[i:]) {
				return true
			}
		}
	} else if p == '?' {
		return ExactMatch(pattern[1:], str[1:])
	} else if p == '[' { //rule: [...]
		i := strings.Index(pattern, "]")
		if i >= 0 && strings.Contains(pattern[1:i], string(str[0])) {
			return ExactMatch(pattern[i+1:], str[1:])
		}

	} else if p == str[0] {
		return ExactMatch(pattern[1:], str[1:])
	}
	return false
}
