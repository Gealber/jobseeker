package client

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchParam struct {
	Keywords string
	Website  string
	// currently only last week available
	// with value 'qdr:w'
	Period string
}

func (s *SearchParam) Valid() bool {
	return s.Period == "qdr:w"
}

func (s *SearchParam) Query() string {
	urlVal := url.Values{}
	urlVal.Add("q", fmt.Sprintf("%s site:%s", strings.ToLower(s.Keywords), s.Website))
	urlVal.Add("tbs", "qdr:w")

	return urlVal.Encode()
}
