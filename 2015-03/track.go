package drum

import "fmt"

// Track is a track with a 16 beat loop.
type Track struct {
	ID    int32
	Name  string
	Beats Steps
}

// NewTrack creates a new Track instance.
func NewTrack() *Track {
	return &Track{
		Beats: Steps{},
	}
}

func (t Track) String() string {
	return fmt.Sprintf("(%d) %s", t.ID, t.Name)
}
