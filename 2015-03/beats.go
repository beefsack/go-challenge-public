package drum

import "bytes"

const (
	beatSep  byte = '|'
	beatPlay byte = 'x'
	beatRest byte = '-'
)

// Beats is a definition of a 16 beat drum loop, usually for a track.
type Beats [16]bool

func (b Beats) String() string {
	buf := bytes.NewBuffer([]byte{beatSep})
	for i, beat := range b {
		c := beatRest
		if beat {
			c = beatPlay
		}
		buf.WriteByte(c)
		if (i+1)%4 == 0 {
			buf.WriteByte(beatSep)
		}
	}
	return buf.String()
}
