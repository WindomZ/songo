package songo

import (
	"net/url"
	"strconv"
	"strings"
)

type SongoParser struct {
	Limit int
	Page  int
	Sort  []string
	Query SongoQuery
}

func (s *SongoParser) ParseURL(u *url.URL) error {
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
	return s.Query.Analyze()
}

func (s *SongoParser) ParseRawURL(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	return s.ParseURL(u)
}

func (s *SongoParser) Least(key string, value interface{}) Songo {
	if value != nil {
		s.Query.Set(key, value.(string))
	}
	return s
}

func (s *SongoParser) Must(key string, value interface{}) Songo {
	if value != nil {
		s.Query.Set(key, value.(string))
	}
	return s
}

func (s *SongoParser) Include(key string) Songo {
	s.Query.Include(key)
	return s
}

func (s *SongoParser) Exclude(key string) Songo {
	s.Query.Exclude(key)
	return s
}
