package sendkeys

import "testing"

func Test_rng(t *testing.T) {
	for n := 0; n != 100; n++ {
		zero := rng(55555)
		one := rng(55555)
		t.Logf("Random0: %d Random1: %d", zero, one)
		if zero == one {
			t.Fatal("rng hit a duplicate")
		}
		zero = 0
		one = 0
	}

}
