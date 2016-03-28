// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"testing"

	"gopkg.in/music-theory.v0/note"
	"github.com/stretchr/testify/assert"
)

func TestListAllForms(t *testing.T) {
	l := ListAllForms()
	assert.Equal(t, len(forms), len(l))
}

func TestFormMatchString(t *testing.T) {
	// TODO
}

func TestChordParseForms(t *testing.T) {
	c := Of("Cm nondominant -5 +6 +7 +9")
	assert.Equal(t, map[Interval]note.Class{
		I3: note.Ds,
		I6: note.A,
		I7: note.As,
		I9: note.D,
	}, c.Tones)
}

/*
 *
 private */

//func assertEquivalentForms(t *testing.T, expectForms []Form, actualForms []Form) {
//	for _, expectForm := range expectForms {
//		assert.Contains(t, actualForms, expectForm,
//			fmt.Sprintf("expect %v in actual %v", expectForms, actualForms))
//	}
//	for _, actualForm := range actualForms {
//		assert.Contains(t, expectForms, actualForm,
//			fmt.Sprintf("actual %v in expected %v", actualForms, expectForms))
//	}
//}
