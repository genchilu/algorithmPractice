package longestHappyString

func longestDiverseString(a int, b int, c int) string {
	m := make(map[byte]int)
	m[byte('a')] = a
	m[byte('b')] = b
	m[byte('c')] = c

	bs := []byte{}

	next := find(m, byte(0))
	for m[next] > 0 {
		if m[next] == 1 {
			bs = append(bs, next)
			m[next] = m[next] - 1
		} else {
			bs = append(bs, next)
			bs = append(bs, next)
			m[next] = m[next] - 2
		}

		next = find(m, next)
		if m[next] == 0 {
			break
		}
		bs = append(bs, next)
		m[next] = m[next] - 1

		next = find(m, next)
	}

	next = find(m, byte(0))
	l := len(bs) - 1
	for l > 0 && m[next] > 0 {
		if bs[l] == next {
			if bs[l-1] == next {
				l--
			} else {
				bs = append(bs, byte(0))
				copy(bs[l+1:], bs[l:len(bs)-1])
				bs[l] = next
				m[next] = m[next] - 1

			}
		}
		l--
	}
	return string(bs)
}

func find(m map[byte]int, pre byte) (next byte) {
	max := 0

	for k, v := range m {
		if k != pre {
			if v >= max {
				max = v
				next = k
			}
		}
	}

	return next
}
