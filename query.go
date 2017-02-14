package songo

import (
	"strconv"
	"strings"
)

type SongoQuery struct {
	query map[string]string
	keys  []string
}

func (s *SongoQuery) get(key string) (v string, ok bool) {
	if s.query == nil {
		s.query = make(map[string]string)
	}
	v, ok = s.query[key]
	return
}

func (s *SongoQuery) set(key, value string) string {
	v, ok := s.get(key)
	if !ok {
		s.keys = append(s.keys, key)
	}
	s.query[key] = value
	return v
}

func (s SongoQuery) Size() int {
	if s.query == nil {
		return 0
	}
	return len(s.query)
}

func (s SongoQuery) GetQueryKeys() []string {
	return s.keys
}

func (s SongoQuery) GetQuery(key string) (tag string, value interface{}) {
	if v, ok := s.get(key); ok {
		vs := strings.Split(v, "$")
		if len(vs) <= 2 {
			return
		}
		tag = "$" + vs[1]
		vs[2] = strings.TrimSpace(vs[2])
		if v, err := strconv.ParseBool(vs[2]); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseInt(vs[2], 10, 64); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseFloat(vs[2], 64); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseUint(vs[2], 10, 64); err == nil {
			value = v
			return
		}
		if strings.Contains(vs[2], ",") {
			value = strings.Split(vs[2], ",")
		} else {
			value = vs[2]
		}
	}
	return
}
