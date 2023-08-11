package kstats

// Mode returns the mode [most frequent value(s)] of a slice of numbers.
func Mode[N Number](numbers ...N) []N {
	switch len(numbers) {
	case 0:
		return nil
	case 1:
		return numbers
	default:
		return mode(numbers)
	}
}

func mode[N Number](numbers []N) []N {
	var (
		size   = 5
		clone  = SortedCloneDif(numbers)
		mode   = make([]N, size)
		cnt    = 1
		maxCnt = 1
		length = len(numbers)
	)

	for inx := 1; inx < length; inx++ {
		switch {
		case clone[inx] == clone[inx-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, clone[inx-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], clone[inx-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}

	mode, maxCnt = modeAndMaxCnt(mode, clone, cnt, maxCnt)
	if maxCnt == 1 || len(mode)*maxCnt == length && maxCnt != length {
		return nil
	}

	return mode
}

func modeAndMaxCnt[N Number](mode, clone []N, cnt, maxCnt int) ([]N, int) {
	length := len(clone)

	switch {
	case cnt == maxCnt:
		return append(mode, clone[length-1]), cnt
	case cnt > maxCnt:
		return append(mode[:0], clone[length-1]), cnt
	}

	return mode, maxCnt
}
