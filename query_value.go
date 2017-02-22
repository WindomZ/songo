package songo

import "strings"

type SongoQueryValue struct {
	Operator string
	Value    interface{}
	Done     bool
}

func (s *SongoQueryValue) Split(value string) bool {
	if strings.HasPrefix(value, "$") {
		value = value[1:]
	}
	vs := strings.Split(value, "$")
	if len(vs) <= 1 {
		s.Done = false
		return false
	}
	s.Operator = "$" + strings.TrimSpace(vs[0])
	if !VerifyQueryOperator(s.Operator) {
		s.Done = false
		return false
	}
	if len(vs) == 2 {
		s.Value = strings.TrimSpace(vs[1])
		s.Done = true
	} else {
		var ss SongoQueryValue
		s.Value = &ss
		s.Done = ss.Split(strings.Join(vs[1:], "$"))
	}
	return s.Done
}

func (s *SongoQueryValue) ValueString() string {
	if s.Done && s.Value != nil {
		if qv, ok := s.Value.(SongoQueryValue); ok {
			return qv.String()
		}
		return s.Value.(string)
	}
	return ""
}

func (s *SongoQueryValue) String() string {
	if s.Done && s.Value != nil {
		if qv, ok := s.Value.(SongoQueryValue); ok {
			return s.Operator + qv.String()
		}
		return s.Operator + s.Value.(string)
	}
	return ""
}

func SplitQueryValue(value string) (v SongoQueryValue, ok bool) {
	ok = v.Split(value)
	return
}

func VerifyQueryValue(value string) bool {
	_, ok := SplitQueryValue(value)
	return ok
}
