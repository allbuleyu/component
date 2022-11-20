package divide

import (
	"testing"
)

func TestDivide(t *testing.T) {
	opts := []Option{
		OptMulti(1),
		OptPage(1),
		OptPageSize(1),
		OptPageTotal(10),
	}
	d := NewDivide(opts...)

	//data := make([]int, n)
	//for i := 0; i < n; i++ {
	//	data[i] = i
	//}

	f := func(i int) error {
		t.Logf("%d", i)
		return nil
	}

	d.Run(f)
}
