package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

const eps = 1e-9

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    z := 1.0
    p := z

    for {
        z = z - (z * z - x) / (2 * z)
        if math.Abs(z - p) < eps {
            break
        }
        p = z
    }

    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
