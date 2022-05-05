package template

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSidebar_ActivateFromPath(t *testing.T) {
	t.Parallel()

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
		sidebar := Sidebar{
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
				Children: Sidebar{
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
			t.Parallel()

			sidebar.ActivateFromPath(table.mastchStr)
			testSidebar(t, sidebar, table.results, table.mastchStr, i, 0, 0)
		})
	}
}

//revive:disable:argument-limit
func testSidebar(t *testing.T, sidebar Sidebar, expectations []map[bool]interface{}, matchStr string, tid, parent, depth int) {
	for i, s := range sidebar {
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
			testSidebar(t, s.Children, expectations[i][expected].([]map[bool]interface{}), matchStr, tid, i, depth+1)
		}
	}
} //revive:enable:argument-limit
