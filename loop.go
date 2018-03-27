package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot square root negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	if x == 0 {
		return 0, nil
	}

	z := 1.0
	do := true

	for i := 1; do; i++ {
		next := z - (z*z-x)/(2*z)
		if dif := next - z; dif < 0 {
			dif = -dif
			if dif < 0.000001 {
				do = false
			}
		}
		z = next
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(0))
}
