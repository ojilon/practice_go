package myfmt

func Uint64Len(n uint64) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		count++
		n /= 10
	}

	return count
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var value uint64 //enure only positive value -> eliminate the negative sign
	var buflen int
	var isnegative bool

	if n < 0 {
		isnegative = true
		value = ^uint64(n) + 1 //safe casting to handle math.MinInt
		buflen = Uint64Len(value)
		buflen += 1 //extra byte for the negative
	} else {
		value = uint64(n)
		buflen = Uint64Len(value)
	}

	//temporary buffer
	buf := make([]byte, buflen)

	//fill the buffer backwards
	// map each int digit to corresponding ascii character value
	for value > 0 {
		buflen--

		buf[buflen] = byte('0' + (value % 10))
		value /= 10
	}

	if isnegative {
		buflen--
		buf[buflen] = '-'
	}

	return string(buf)
}

func Atoi(s string) int {
	return 0
}
