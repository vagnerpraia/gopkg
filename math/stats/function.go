package gpstats

import (
	"sort"

	gpparity "github.com/vagnerpraia/gopkg/math/parity"
)

func IntsMean(ints []int) float64 {

	sum := 0

	for _, add := range ints {
		sum += add
	}

	return float64(sum) / float64(len(ints))
}

func Float64sMean(floats []float64) float64 {

	sum := 0.0

	for _, add := range floats {
		sum += add
	}

	return sum / float64(len(floats))
}

func IntsMedian(ints []int) int {

	length := len(ints)

	if length == 0 {
		return 0.0
	} else if length == 1 {
		return ints[0]
	}

	dataCopy := make([]int, len(ints))
	copy(dataCopy, ints)

	sort.Ints(dataCopy)

	if length > 0 {
		if gpparity.IntIsOdd(length) {
			return (dataCopy[length/2-1] + dataCopy[length/2]) / 2
		} else {
			return dataCopy[length/2]
		}
	}

	return 0
}

func Float64sMedian(floats []float64) float64 {

	length := len(floats)

	if length == 0 {
		return 0.0
	} else if length == 1 {
		return floats[0]
	}

	dataCopy := make([]float64, len(floats))
	copy(dataCopy, floats)

	sort.Float64s(dataCopy)

	if length > 0 {
		if gpparity.IntIsOdd(length) {
			return (dataCopy[length/2-1] + dataCopy[length/2]) / 2
		} else {
			return dataCopy[length/2]
		}
	}

	return 0.0
}

func IntsMode(ints []int) int {

	var mode int

	countMap := make(map[int]int)
	for _, value := range ints {
		countMap[value] += 1
	}

	max := 0

	for _, key := range ints {
		freq := countMap[key]

		if freq > max {
			mode = key
			max = freq
		}
	}

	return mode
}

func Float64sMode(floats []float64) float64 {

	var mode float64

	countMap := make(map[float64]int)
	for _, value := range floats {
		countMap[value] += 1
	}

	max := 0

	for _, key := range floats {
		freq := countMap[key]

		if freq > max {
			mode = key
			max = freq
		}
	}

	return mode
}

func IntsModes(ints []int) []int {

	l := len(ints)
	if l <= 1 {
		return ints
	}

	c := make([]int, len(ints))
	copy(c, ints)

	sort.Ints(c)

	mode := make([]int, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case c[i] == c[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, c[i-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], c[i-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, c[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], c[l-1])
		maxCnt = cnt
	}

	if maxCnt == 1 || len(mode)*maxCnt == l && maxCnt != l {
		return []int{}
	}

	return mode
}

func Float64sModes(floats []float64) []float64 {

	l := len(floats)
	if l <= 1 {
		return floats
	}

	c := make([]float64, len(floats))
	copy(c, floats)

	sort.Float64s(c)

	mode := make([]float64, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case c[i] == c[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, c[i-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], c[i-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, c[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], c[l-1])
		maxCnt = cnt
	}

	if maxCnt == 1 || len(mode)*maxCnt == l && maxCnt != l {
		return []float64{}
	}

	return mode
}
