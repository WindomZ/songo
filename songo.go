package songo

import (
	"net/url"
	"strconv"
	"strings"
)

type Songo struct {
	Limit int
	Page  int
	Sort  []string
	Query SongoQuery
}

func (s *Songo) ParseURL(u *url.URL) error {
	vs := u.Query()
	for k, v := range vs {
		if len(v) == 0 {
			continue
		}
		switch k {
		case "_limit":
			s.Limit, _ = strconv.Atoi(v[0])
		case "_page":
			s.Page, _ = strconv.Atoi(v[0])
		case "_sort":
			s.Sort = strings.Split(strings.Join(v, ","), ",")
		default:
			s.Query.set(k, v[0])
		}
	}
	return nil
}

func (s *Songo) ParseRawURL(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	return s.ParseURL(u)
}
