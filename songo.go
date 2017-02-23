package songo

import (
	"net/url"
	"strconv"
	"strings"
)

type ISongo interface {
	ParseURL(u *url.URL) error
	ParseRawURL(rawurl string) error

	Least(key string, value interface{})
	Must(key string, value interface{})
}

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
			s.Query.Set(k, v[0])
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

func (s *Songo) Least(key string, value interface{}) {
	if value != nil {
		s.Query.Set(key, value.(string))
	}
}

func (s *Songo) Must(key string, value interface{}) {
	if value != nil {
		s.Query.Set(key, value.(string))
	}
}
