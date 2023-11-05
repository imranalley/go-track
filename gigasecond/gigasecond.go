// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"fmt"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	dateInSeconds := t.Unix()

	dateInSeconds += 1000000000

	ds := fmt.Sprintf("%ds", dateInSeconds)
	d, _ := time.ParseDuration(ds)
	timein := t.Add(time.Second * time.Duration(d))

	// // days := d.Hours()

	// t = t.Add(time.Second * d)

	return timein
}
