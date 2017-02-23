package songo

import "strings"

type SongoQuery struct {
	query map[string][]string
	keys  []string
	size  int
}

func (s *SongoQuery) Get(key string) (v []string, ok bool) {
	if s.query == nil {
		s.query = make(map[string][]string)
	}
	v, ok = s.query[key]
	return
}

func (s *SongoQuery) Set(key, value string) {
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
	if s.Size() <= 0 {
		return []string{}
	}
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
