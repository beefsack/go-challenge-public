package drum

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"testing"
)

func TestPattern_Encode(t *testing.T) {
	for _, f := range []string{
		"pattern_1.splice",
		"pattern_2.splice",
		"pattern_3.splice",
		"pattern_4.splice",
		"pattern_5.splice",
	} {
		raw, err := ioutil.ReadFile(path.Join("fixtures", f))
		if err != nil {
			t.Fatalf("unable to read %s, %v", f, err)
		}

		decoded, err := Decode(bytes.NewBuffer(raw))
		if err != nil {
			t.Fatalf("unable to decode %s, %v", f, err)
		}

		encoded := bytes.NewBuffer([]byte{})
		if err := decoded.Encode(encoded); err != nil {
			t.Fatalf("unable to encode %s, %v", f, err)
		}

		rawEncoded := encoded.Bytes()
		if !bytes.HasPrefix(raw, rawEncoded) {
			t.Errorf(
				"encoded did not match raw for %s.\nExpected:\n\n%s\n\nActual:\n\n%s",
				f,
				hex.Dump(raw),
				hex.Dump(rawEncoded),
			)
		}
	}
}

func ExampleAddCowbells() {
	raw, err := ioutil.ReadFile(path.Join("fixtures", "pattern_2.splice"))
	if err != nil {
		log.Fatalf("unable to read file, %v", err)
	}

	decoded, err := Decode(bytes.NewBuffer(raw))
	if err != nil {
		log.Fatalf("unable to decode, %v", err)
	}

	if l := len(decoded.Tracks); l != 4 {
		log.Fatalf("expected there to be 4 tracks, got %d", l)
	}

	// Add cowbells
	for i := 0; i < 16; i += 4 {
		decoded.Tracks[3].Steps[i] = true
	}
	decoded.Tracks[3].Steps[6] = true
	decoded.Tracks[3].Steps[14] = true

	pipe := bytes.NewBuffer([]byte{})
	if err := decoded.Encode(pipe); err != nil {
		log.Fatalf("unable to encode, %v", err)
	}
	redecoded, err := Decode(pipe)
	if err != nil {
		log.Fatalf("unable to decode for the second time, %v", err)
	}

	fmt.Print(redecoded)
	// Output:
	// Saved with HW Version: 0.808-alpha
	// Tempo: 98.4
	// (0) kick	|x---|----|x---|----|
	// (1) snare	|----|x---|----|x---|
	// (3) hh-open	|--x-|--x-|x-x-|--x-|
	// (5) cowbell	|x---|x-x-|x---|x-x-|
}
