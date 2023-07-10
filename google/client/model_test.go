package client

import (
	"fmt"
	"testing"
)

func Test_Query(t *testing.T) {
	t.Run("creating query", func(t *testing.T) {
		s := SearchParam{
			Keywords: "go backend",
			Website:  "app.otta.com",
		}
		if s.Query() != "q=go+backend+site%3Aapp.otta.com" {
			t.Fatal(fmt.Sprintf("unexpected query: '%s'\n", s.Query()))
		}
	})
}
