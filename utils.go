package l00p8

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultP   = 1
	defaultIpp = 10
)

type ListParams struct {
	Page         int
	ItemsPerPage int
	OrderField   string
	OrderAsc     bool
}

func (params ListParams) SQLOrderAndPaging() string {
	q := ""
	if params.OrderField != "" {
		q += "ORDER BY " + params.OrderField
		if params.OrderAsc {
			q += " ASC"
		} else {
			q += " DESC"
		}
	}
	if params.ItemsPerPage > 0 {
		q += " LIMIT " + strconv.Itoa(params.ItemsPerPage)
		q += " OFFSET " + strconv.Itoa(params.ItemsPerPage*(params.Page-1))
	}
	return q
}

// ParseFromRequest parses query object from http.Request's URL
// and also extracts paging parameters from URL's query parameters (ipp - items per page, p - page)
// Example: /tasks?q={status:"todo"}&ipp=10&p=2
// Tries to unmarshal q parameter to query object. q parameter must be in bson/json format.
// This function receives http.Request object in order to be able to work with http request body in some cases.
func ParseFromRequest(r *http.Request, query interface{}) *ListParams {
	return parseFromQuery(r.URL.Query(), query)
}

func parseFromQuery(values url.Values, query interface{}) *ListParams {
	var ipp, p int
	var err error
	lp := &ListParams{}
	q := values.Get("q")
	pStr := values.Get("p")
	ippStr := values.Get("ipp")
	_ = bson.UnmarshalJSON([]byte(q), query)
	if pStr == "" {
		p = defaultP
	} else {
		p, err = strconv.Atoi(pStr)
		if err != nil {
			p = defaultP
		}
	}

	if ippStr == "" {
		ipp = defaultIpp
	} else {
		ipp, err = strconv.Atoi(ippStr)
		if err != nil {
			ipp = defaultIpp
		}
	}
	lp.Page = p
	lp.ItemsPerPage = ipp
	lp.OrderAsc = true

	orderField := values.Get("order")
	if orderField != "" {
		if orderField[0] == '-' {
			lp.OrderAsc = false
			lp.OrderField = orderField[1:]
		} else {
			lp.OrderField = orderField
		}
	}

	return lp
}
