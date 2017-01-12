package wildcard

import (
	"testing"
)

func TestExactMatch(t *testing.T) {
	testExactMatch("3wo6", "3wo6", true, t)
	testExactMatch("*", "3wo6", true, t)
	testExactMatch("*6", "3wo6", true, t)
	testExactMatch("3*", "3wo6", true, t)
	testExactMatch("3*6", "3wo6", true, t)
	testExactMatch("?wo?", "3wo6", true, t)
	testExactMatch("3?6", "3wo6", false, t)
	testExactMatch("????", "3wo6", true, t)
	testExactMatch("[36]wo[36]", "3wo6", true, t)
	testExactMatch("?[wo]?", "3wo6", false, t)
	testExactMatch("?[wo][ok]?", "3wo6", true, t)
}

func testExactMatch(p, s string, b bool, t *testing.T) {
	if ExactMatch(p, s) != b {
		t.Fatalf("%v -- %v : %v", p, s, !b)
	}
}
