package drum

import "bytes"

const (
	stepSep  byte = '|'
	stepPlay byte = 'x'
	stepRest byte = '-'
)

// Steps is a definition of a 16 step drum loop, usually for a track.
type Steps [16]bool

func (b Steps) String() string {
	buf := bytes.NewBuffer([]byte{stepSep})
	for i, step := range b {
		c := stepRest
		if step {
			c = stepPlay
		}
		buf.WriteByte(c)
		if (i+1)%4 == 0 {
			buf.WriteByte(stepSep)
		}
	}
	return buf.String()
}
