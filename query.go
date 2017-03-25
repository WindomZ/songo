package songo

import (
	"errors"
	"strings"
)

type SongoQuery struct {
	includes map[string]bool
	query    map[string][]string
	keys     []string
	size     int
}

func (s *SongoQuery) Include(key string) {
	if s.includes == nil {
		s.includes = make(map[string]bool)
	}
	s.includes[key] = true
}

func (s *SongoQuery) Exclude(key string) {
	if s.includes == nil {
		s.includes = make(map[string]bool)
	}
	s.includes[key] = false
}

func (s *SongoQuery) Analyze() error {
	if s.includes != nil && len(s.includes) != 0 {
		for k, v := range s.includes {
			if !v {
			} else if _, ok := s.query[k]; !ok {
				return errors.New("songo: missing key: " + k)
			}
		}
	}
	return nil
}

func (s *SongoQuery) Get(key string) (v []string, ok bool) {
	if s.query == nil {
		s.query = make(map[string][]string)
		return nil, false
	}
	v, ok = s.query[key]
	return
}

func (s *SongoQuery) Set(key, value string) {
	if s.includes != nil {
		if v, ok := s.includes[key]; ok && !v {
			return
		}
	}
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)
	if len(key) == 0 || len(value) == 0 {
		return
	}
	if strings.Contains(value, ",$") ||
		strings.Contains(value, ", $") {
		vs := strings.Split(value, ",")
		for _, v := range vs {
			s.Set(key, v)
		}
		return
	}
	if i := strings.LastIndex(key, "$"); i >= 0 {
		value = key[:i] + value
		key = key[i+1:]
	}
	if v, ok := s.Get(key); ok {
		s.query[key] = append(v, value)
	} else {
		s.query[key] = []string{value}
		s.keys = append(s.keys, key)
	}
	s.size++
}

func (s SongoQuery) Size() int {
	return s.size
}

func (s SongoQuery) GetKeys() []string {
	return s.keys
}

func (s SongoQuery) GetValues(key string) ([]SongoQueryValue, bool) {
	if strs, ok := s.Get(key); ok {
		r := make([]SongoQueryValue, 0, len(strs))
		for _, str := range strs {
			if qv, ok := SplitQueryValue(str); ok {
				r = append(r, qv)
			}
		}
		return r, len(r) != 0
	}
	return nil, false
}

func (s SongoQuery) GetQuery(key string, indexes ...int) (operators []string, value interface{}, ok bool) {
	if qvs, _ok := s.GetValues(key); _ok {
		index := 0
		if len(indexes) != 0 {
			index = indexes[0]
		}
		if index < 0 || index >= len(qvs) {
			return
		}
		operators, value = qvs[index].GetQuery()
		ok = true
	}
	return
}
