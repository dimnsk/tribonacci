package main

import (
	"math/big"
	//"fmt"
)

// Calculate tribonacci value for n-number
func trib(n *big.Int) *big.Int {
	t0 := big.NewInt(0)
	t1 := big.NewInt(0)
	t2 := big.NewInt(1)

	if n.Cmp(big.NewInt(3)) == -1 {
		return t0;
	}

	for i := big.NewInt(3); n.Cmp(i) > 0; i.Add(i, big.NewInt(1)) {
		t2_old := new(big.Int)
		t2_old.Set(t2)
		t2.Add(t2.Add(t2, t0), t1)
		t1, t0 = t2_old, t1
		//fmt.Printf("i=%d : t2 = %d  t1 = %d  t0 = %d t2_old = %d\n", i, t2, t1, t0, t2_old)
	}
	return t2
}
