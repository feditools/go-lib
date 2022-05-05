package lib

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetPaginationFromURL(t *testing.T) {
	t.Parallel()

	url1, _ := url.Parse("http://localhost/page")
	url2, _ := url.Parse("http://localhost/page?page=6")
	url3, _ := url.Parse("http://localhost/page?count=30")
	url4, _ := url.Parse("http://localhost/page?count=25&page=16")
	url5, _ := url.Parse("http://localhost/page?page=66&count=15")

	//revive:disable:add-constant
	tables := []struct {
		input            *url.URL
		defaultCount     int
		outputPage       int
		outputCount      int
		outputCountFound bool
	}{
		{url1, 10, 1, 10, false},
		{url2, 10, 6, 10, false},
		{url3, 10, 1, 30, true},
		{url4, 20, 16, 25, true},
		{url5, 100, 66, 15, true},
	}
	//revive:enable:add-constant

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running GetPaginationFromURL on %s", i, table.input.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			page, count, countFound := GetPaginationFromURL(table.input, table.defaultCount)
			if page != table.outputPage {
				t.Errorf("[%d] invalid page, got: '%d', want: '%d'", i, page, table.outputPage)
			}
			if count != table.outputCount {
				t.Errorf("[%d] invalid count, got: '%d', want: '%d'", i, count, table.outputCount)
			}
			if countFound != table.outputCountFound {
				t.Errorf("[%d] invalid countFound, got: '%v', want: '%v'", i, countFound, table.outputCountFound)
			}
		})
	}
}
