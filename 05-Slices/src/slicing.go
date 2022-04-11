package src

import s "github.com/inancgumus/prettyslice"

func Slicing() {
	msg := []byte{'a', 'b', 'c', 'd', 'e'}
	msg1 := msg[2:5]
	s.Show("alphabets", msg1)
}
