package drum

import (
	"bytes"
	"fmt"
)

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	Version string
	Tempo   float32
	Tracks  []*Track
}

// NewPattern creates a new Pattern instance.
func NewPattern() *Pattern {
	return &Pattern{Tracks: []*Track{}}
}

func (p Pattern) String() string {
	buf := bytes.NewBufferString(fmt.Sprintf(
		"Saved with HW Version: %s\nTempo: %v\n",
		p.Version,
		p.Tempo,
	))

	// Output each track.
	for _, t := range p.Tracks {
		buf.WriteString(fmt.Sprintf("%s\t%s\n", t, t.Beats.String()))
	}

	return buf.String()
}
