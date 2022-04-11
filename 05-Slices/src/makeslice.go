package src

import s "github.com/inancgumus/prettyslice"

func MakeSlice() {
	msg := make([]byte, 0, 5)
	msg = append(msg, 'a')
	s.Show("msg", msg)
}
