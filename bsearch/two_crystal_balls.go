package search

import "math"

func search_crystal_balls(breaks []bool) int {
    jump := int(math.Floor(math.Sqrt(float64(len(breaks)))))
    i := 0
    for ; i < len(breaks); i = i + jump {
        if breaks[i] {
            break
        }
    }
    i -= jump

    for ; i < len(breaks); i = i + 1 {
        if breaks[i] {
            return i
        }
    }
    return -1
}
