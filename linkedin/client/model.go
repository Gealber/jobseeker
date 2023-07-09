package client

import (
	"net/url"
	"strconv"
)

//	{
//	    "keyword": "Golang",
//	    "location": "United States",
//	    "position": 1,
//	    "pageNum": 0,
//	    "f_TPR": r604800,
//	    "f_WT": 2,
//	}
type SearchParam struct {
	// keyword of the search
	Keywords string
	// location of the jobs to appear
	Location string
	// time of job post, r604800
	FTPR string
	// type of job, onsite or remote
	// with values 1 and 2 respectively
	FWT      string
	Position int
	PageNum  int
}

func (s *SearchParam) Valid() bool {
	return s.FWT == "1" || s.FWT == "2"
}

func (s *SearchParam) Query() string {
	urlVal := url.Values{}
	urlVal.Add("keywords", s.Keywords)
	urlVal.Add("location", s.Location)
	urlVal.Add("f_TPR", s.FTPR)
	urlVal.Add("f_WT", s.FWT)
	urlVal.Add("pageNum", strconv.Itoa(s.PageNum))
	urlVal.Add("position", strconv.Itoa(s.Position))

	return urlVal.Encode()
}
