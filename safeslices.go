package safeslices

import "golang.org/x/exp/constraints"

func MaxSafe[T constraints.Ordered](s []T) (T, bool) {
    var zero T
    if len(s) == 0 {
        return zero, false
    }

    max := s[0]
    for _, v := range s[1:] {
        if v > max {
            max = v
        }
    }
    return max, true
}