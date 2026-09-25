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

// brief error implementation
type myerror struct {
	s string
}

func (e *myerror) Error() string {
	return e.s
}

func myNew(text string) error {
	return &myerror{text}
}

func Atoi(s string) (int, error) {

	if len(s) == 0 {
		return 0, myNew("empty string")
	}

	sign := 1
	start := 0

	//check for the leading sign
	// if there is a sign, the iteration starts from the next position
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start == len(s) {
		return 0, myNew("only sign entred")
	}

	//convert every char to corresponding integer
	n := 0
	for i := start; i < len(s); i++ {
		ch := s[i]

		//check is character is not a valid int
		if ch < '0' || ch > '9' {
			return 0, myNew("One of the characters does not become a valid int")
		}

		//obtain the integer value
		integer_value := int(ch - '0')

		//shift the previous digits to the left by multipying by 10
		n *= 10

		//add the next integer value
		n += integer_value

	}
	return n * sign, nil
}
