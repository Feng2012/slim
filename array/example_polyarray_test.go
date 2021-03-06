package array_test

import (
	"fmt"

	"github.com/openacid/slim/array"
)

func ExampleNewDense() {

	// Only 10 bits/elt to store a serias of incremental int64
	nums := []int32{
		10000387, 10000393, 10000399, 10000404, 10000407, 10000410, 10000415,
		10000418, 10000420, 10000422, 10000426, 10000430, 10000434, 10000439,
		10000444, 10000446, 10000448, 10000451, 10000456, 10000459, 10000462,
		10000465, 10000470, 10000473, 10000479, 10000482, 10000488, 10000490,
		10000494, 10000500, 10000506, 10000509, 10000513, 10000519, 10000521,
		10000528, 10000530, 10000534, 10000537, 10000540, 10000544, 10000546,
		10000551, 10000556, 10000560, 10000566, 10000568, 10000572, 10000574,
		10000576, 10000580, 10000585, 10000588, 10000592, 10000594, 10000600,
		10000603, 10000606, 10000608, 10000610,
	}

	a := array.NewDense(nums)

	fmt.Println("last elt is:", a.Get(int32(a.Len()-1)))

	st := a.Stat()
	for _, k := range []string{
		"elt_width",
		"mem_elts",
		"bits/elt"} {
		fmt.Printf("%10s : %d\n", k, st[k])
	}

	// Unordered output:
	// last elt is: 10000610
	//  elt_width : 2
	//   mem_elts : 32
	//   bits/elt : 11
}
