package drum

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()
	return Decode(f)
}

// Decode decodes a Pattern from a reader.
func Decode(r io.Reader) (p *Pattern, err error) {
	p = NewPattern()
	if p.Version, p.Tempo, err = decodeHeader(r); err != nil {
		err = fmt.Errorf("unable to decode header, %v", err)
		return
	}
	for {
		var t *Track
		if t, err = decodeTrack(r); err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return
		}
		p.Tracks = append(p.Tracks, t)
	}
	return
}

func decodeHeader(r io.Reader) (version string, tempo float32, err error) {
	// Check that the header is correct.
	p := make([]byte, 14)
	if _, err = r.Read(p); err != nil {
		err = fmt.Errorf("unable to read header bytes, %v", err)
		return
	}
	if string(p[:6]) != "SPLICE" {
		err = errors.New("expected to start with 'SPLICE'")
		return
	}

	// Read out the version.
	p = make([]byte, 32)
	if _, err = r.Read(p); err != nil {
		err = fmt.Errorf("unable to read version bytes, %v", err)
		return
	}
	version = string(bytes.Replace(p, []byte{0}, []byte{}, -1))

	// Read out the tempo.
	if err = binary.Read(r, binary.LittleEndian, &tempo); err != nil {
		err = fmt.Errorf("unable to read tempo, %v", err)
		return
	}
	return
}

func decodeTrack(r io.Reader) (t *Track, err error) {
	t = NewTrack()

	if binary.Read(r, binary.LittleEndian, &t.ID); err != nil {
		if err != io.EOF {
			err = fmt.Errorf("unable to read ID, %v", err)
		}
		return
	}

	var nLen byte
	if binary.Read(r, binary.LittleEndian, &nLen); err != nil {
		if err != io.EOF {
			err = fmt.Errorf("unable to read name length, %v", err)
		}
		return
	}

	p := make([]byte, nLen)
	if _, err = r.Read(p); err != nil {
		if err != io.EOF {
			err = fmt.Errorf("unable to read name, %v", err)
		}
		return
	}
	t.Name = string(p)

	t.Beats, err = decodeBeats(r)
	return
}

func decodeBeats(r io.Reader) (b Beats, err error) {
	b = Beats{}
	p := make([]byte, len(b))
	if _, err = r.Read(p); err != nil {
		if err != io.EOF {
			err = fmt.Errorf("unable to read beats, %v", err)
		}
	}

	for i, v := range p {
		b[i] = v > 0
	}

	return
}
