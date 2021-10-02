package stevepro

import "testing"

func TestAdd(t *testing.T) {
	const expect = 5
	actual := Add(2, 3)
	if actual != expect {
		t.Errorf("sum expect: %d actual: %d", expect, actual)
	}
}
