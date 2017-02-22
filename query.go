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

func (s SongoQuery) GetQuery(key string) (operator string, value interface{}) {
	if str, ok := s.get(key); ok {
		qv, ok := SplitQueryValue(str)
		if !ok {
			return
		}
		operator = qv.Operator
		str = qv.ValueString()
		//println(key, operator, str, ok)
		if v, err := strconv.ParseBool(str); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseInt(str, 10, 64); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseFloat(str, 64); err == nil {
			value = v
			return
		} else if v, err := strconv.ParseUint(str, 10, 64); err == nil {
			value = v
			return
		}
		if strings.Contains(str, ",") {
			value = strings.Split(str, ",")
		} else {
			value = str
		}
	}
	return
}
