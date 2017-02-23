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
	values := u.Query()
	for k, vs := range values {
		if len(vs) == 0 {
			continue
		}
		switch k {
		case "_limit":
			s.Limit, _ = strconv.Atoi(vs[0])
		case "_page":
			s.Page, _ = strconv.Atoi(vs[0])
		case "_sort":
			s.Sort = strings.Split(strings.Join(vs, ","), ",")
		default:
			for _, v := range vs {
				s.Query.Set(k, v)
			}
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
