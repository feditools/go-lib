package template

import (
	"fmt"
	"regexp"
	"testing"
)

type testActivableSlices []testActivableSlice
type testActivableSlice struct {
	Matcher *regexp.Regexp

	Active bool

	Children testActivableSlices
}

// GetChildren returns the children of the node or nil if no children.
func (s *testActivableSlices) GetChildren(i int) ActivableSlice {
	if len(*s) == 0 {
		return nil
	}
	return &(*s)[i].Children
}

// GetMatcher returns the matcher of the node or nil if no matcher.
func (s *testActivableSlices) GetMatcher(i int) *regexp.Regexp {
	return (*s)[i].Matcher
}

// SetActive sets the active bool based on the match regex.
func (s *testActivableSlices) SetActive(i int, a bool) {
	(*s)[i].Active = a
}

// Len returns the matcher of the node or nil if no matcher.
func (s *testActivableSlices) Len() int {
	return len(*s)
}

// tests

func TestSetActive(t *testing.T) {
	tables := []struct {
		mastchStr string
		results   []map[bool]interface{}
	}{
		{
			"/test1",
			[]map[bool]interface{}{
				{
					true: nil,
				},
				{
					false: nil,
				},
				{
					false: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test1/sub1",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					false: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test2",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					true: nil,
				},
				{
					false: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test2/sub1",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					true: nil,
				},
				{
					false: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test3",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					true: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test3/sub1",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					true: []map[bool]interface{}{
						{
							true: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test3/sub1/foo",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					true: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							false: nil,
						},
					},
				},
			},
		},
		{
			"/test3/sub2",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					true: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							true: nil,
						},
					},
				},
			},
		},
		{
			"/test3/sub2/foo",
			[]map[bool]interface{}{
				{
					false: nil,
				},
				{
					false: nil,
				},
				{
					true: []map[bool]interface{}{
						{
							false: nil,
						},
						{
							true: nil,
						},
					},
				},
			},
		},
	}

	for i, table := range tables {
		i := i
		table := table
		slices := testActivableSlices{
			{
				Matcher: regexp.MustCompile(`^/test1$`),
				Active:  false,
			},
			{
				Matcher: regexp.MustCompile(`^/test2`),
				Active:  false,
			},
			{
				Matcher: regexp.MustCompile(`^/test3`),
				Active:  false,
				Children: testActivableSlices{
					{
						Matcher: regexp.MustCompile(`^/test3/sub1$`),
						Active:  false,
					},
					{
						Matcher: regexp.MustCompile(`^/test3/sub2`),
						Active:  false,
					},
				},
			},
		}

		name := fmt.Sprintf("[%d] Running activation test on %s", i, table.mastchStr)
		t.Run(name, func(t *testing.T) {
			//t.Parallel()

			SetActive(&slices, table.mastchStr)
			testSlices(t, slices, table.results, table.mastchStr, i, 0, 0)
		})
	}
}

//revive:disable:argument-limit
func testSlices(t *testing.T, slices testActivableSlices, expectations []map[bool]interface{}, matchStr string, tid, parent, depth int) {
	for i, s := range slices {
		if parent == 0 {
			t.Logf("[%d][%d][%d] checking activation", tid, depth, i)
		} else {
			t.Logf("[%d][%d][%d->%d] checkign activation", tid, depth, parent, i)
		}

		// get keys
		j := 0
		keys := make([]bool, len(expectations[i]))
		for k := range expectations[i] {
			keys[j] = k
			j++
		}

		// test active
		expected := keys[0]
		if s.Active != expected {
			if parent == 0 {
				t.Errorf("[%d][%d][%d] got wrong activation state got: %v, want: %v (str: %s, re: %s)", tid, depth, i, s.Active, expected, matchStr, s.Matcher.String())
			} else {
				t.Errorf("[%d][%d][%d->%d] got wrong activation state got: %v, want: %v (str: %s, re: %s)", tid, depth, parent, i, s.Active, expected, matchStr, s.Matcher.String())
			}
		}

		// run on children
		if len(s.Children) > 0 {
			testSlices(t, s.Children, expectations[i][expected].([]map[bool]interface{}), matchStr, tid, i, depth+1)
		}
	}
} //revive:enable:argument-limit
