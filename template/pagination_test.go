package template

import (
	"fmt"
	"testing"
)

const testPaginationResult1 = `<nav>
  <ul class="pagination">
    <li class="page-item disabled">
      <span class="page-link" aria-label="Previous">
        <span aria-hidden="true"><i class="fas fa-caret-left"></i></span>
      </span>
    </li>
    <li class="page-item active" aria-current="page">
      <span class="page-link" aria-label="1">
        1 <span class="sr-only">(current)</span>
      </span>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test1?page=2">
        2
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test1?page=3">
        3
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test1?page=4">
        4
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test1?page=5">
        5
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test1?page=2">
        <span aria-hidden="true"><i class="fas fa-caret-right"></i></span>
      </a>
    </li>
  </ul>
</nav>`

const testPaginationResult2 = `<nav>
  <ul class="pagination">
    <li class="page-item">
      <a class="page-link" href="/test2?page=16">
        <span aria-hidden="true"><i class="fas fa-caret-left"></i></span>
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test2?page=15">
        15
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test2?page=16">
        16
      </a>
    </li>
    <li class="page-item active" aria-current="page">
      <span class="page-link" aria-label="17">
        17 <span class="sr-only">(current)</span>
      </span>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test2?page=18">
        18
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test2?page=19">
        19
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test2?page=18">
        <span aria-hidden="true"><i class="fas fa-caret-right"></i></span>
      </a>
    </li>
  </ul>
</nav>`

const testPaginationResult3 = `<nav>
  <ul class="pagination">
    <li class="page-item">
      <a class="page-link" href="/test3?page=49">
        <span aria-hidden="true"><i class="fas fa-caret-left"></i></span>
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test3?page=46">
        46
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test3?page=47">
        47
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test3?page=48">
        48
      </a>
    </li>
    <li class="page-item">
      <a class="page-link" href="/test3?page=49">
        49
      </a>
    </li>
    <li class="page-item active" aria-current="page">
      <span class="page-link" aria-label="50">
        50 <span class="sr-only">(current)</span>
      </span>
    </li>
    <li class="page-item disabled">
      <span class="page-link" aria-label="Next">
        <span aria-hidden="true"><i class="fas fa-caret-right"></i></span>
      </span>
    </li>
  </ul>
</nav>`

func TestMakePagination(t *testing.T) {
	templates, err := New(nil)
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}
	err = addTestTemplates(templates)
	if err != nil {
		t.Errorf("adding tests: %s", err.Error())
		return
	}

	tables := []struct {
		config *PaginationConfig
		result string
	}{
		{
			&PaginationConfig{
				Count:         1000,
				DisplayCount:  20,
				HRef:          "/test1",
				MaxPagination: 5,
				Page:          1,
			},
			testPaginationResult1,
		},
		{
			&PaginationConfig{
				Count:         1000,
				DisplayCount:  20,
				HRef:          "/test2",
				MaxPagination: 5,
				Page:          17,
			},
			testPaginationResult2,
		},
		{
			&PaginationConfig{
				Count:         1000,
				DisplayCount:  20,
				HRef:          "/test3",
				MaxPagination: 5,
				Page:          50,
			},
			testPaginationResult3,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running making pagination", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			pag := MakePagination(table.config)

			result, err := testExecuteTemplate(templates, "test_pagination", pag)
			if err != nil {
				t.Errorf("unexpected error creating template: %s", err.Error())
				return
			}
			if result != table.result {
				t.Errorf("unexpected result\n\ngot:\n-------------\n%s\n\nwant:\n-------------\n%s\n", result, table.result)
				return
			}
		})
	}
}
