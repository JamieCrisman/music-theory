// An interval is the difference between two pitches.
package note

// Ordered set of all the intervals, e.g. for stepping from the root of a chord outward to its other tones.
var IntervalOrder = []Interval{
	I2,
	I3,
	I4,
	I5,
	I6,
	I7,
	I8,
	I9,
	I10,
	I11,
	I12,
	I13,
	I14,
	I15,
	I16,
}

// Interval between notes.
type Interval int

const (
	I2  = Interval(2)
	I3  = Interval(3)
	I4  = Interval(4)
	I5  = Interval(5)
	I6  = Interval(6)
	I7  = Interval(7)
	I8  = Interval(8)
	I9  = Interval(9)
	I10 = Interval(10)
	I11 = Interval(11)
	I12 = Interval(12)
	I13 = Interval(13)
	I14 = Interval(14)
	I15 = Interval(15)
	I16 = Interval(15)
)
