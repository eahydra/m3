// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package index

import (
	"time"

	m3ninxindex "github.com/m3db/m3/src/m3ninx/index"
	"github.com/m3db/m3/src/m3ninx/index/segment"
)

type activeSegmentState byte

const (
	mutableActiveSegmentState activeSegmentState = iota
	fstActiveSegmentState
)

func (a activeSegmentState) String() string {
	switch a {
	case mutableActiveSegmentState:
		return "mutableActiveSegment"
	case fstActiveSegmentState:
		return "fstActiveSegment"
	}
	return "unknownActiveSegment"
}

// activeSegment starts out backed by a mutable segment, which is rotated
// to a FST segment based on size constraints.
type activeSegment struct {
	creationTime   time.Time
	state          activeSegmentState
	mutableSegment segment.MutableSegment
	fstSegment     segment.Segment
}

func (a *activeSegment) Reader() (m3ninxindex.Reader, error) {
	if a.state == fstActiveSegmentState {
		return a.fstSegment.Reader()
	}
	return a.mutableSegment.Reader()
}

func (a *activeSegment) Size() int64 {
	if a.state == fstActiveSegmentState {
		return a.fstSegment.Size()
	}
	return a.mutableSegment.Size()
}

func (a *activeSegment) Close() error {
	if a.state == mutableActiveSegmentState {
		return a.mutableSegment.Close()
	}

	// TODO(prateek): handle rotatingActiveSegmentState
	return a.fstSegment.Close()
}
