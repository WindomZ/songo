package songo

import (
	"strconv"
	"strings"
)

type SongoQueryValue struct {
	Operator string
	Value    interface{}
}

func (s *SongoQueryValue) Split(value string) bool {
	if strings.HasPrefix(value, "$") {
		value = value[1:]
	}
	vs := strings.Split(value, "$")
	if len(vs) <= 1 {
		return false
	}
	s.Operator = "$" + strings.TrimSpace(vs[0])
	if !VerifyQueryOperator(s.Operator) {
		return false
	}
	if len(vs) == 2 {
		s.Value = strings.TrimSpace(vs[1])
		return true
	}
	var ss SongoQueryValue
	s.Value = &ss
	return ss.Split(strings.Join(vs[1:], "$"))
}

func (s *SongoQueryValue) Next() (qv *SongoQueryValue, ok bool) {
	if s.Value != nil {
		qv, ok = s.Value.(*SongoQueryValue)
	}
	if !ok {
		qv = s
	}
	return
}

func (s *SongoQueryValue) HasNext() (ok bool) {
	_, ok = s.Next()
	return
}

func (s SongoQueryValue) GetValue() interface{} {
	if str, ok := s.Value.(string); ok {
		return StringToValue(str)
	}
	return s.Value
}

func (s SongoQueryValue) GetQuery() (operators []string, value interface{}) {
	qv := &s
	for ok := true; ok; {
		operators = append(operators, s.Operator)
		qv, ok = qv.Next()
	}
	value = qv.GetValue()
	return
}

func (s *SongoQueryValue) ValueString() string {
	if s.Value != nil {
		if qv, ok := s.Next(); ok {
			return qv.String()
		}
		return s.Value.(string)
	}
	return ""
}

func (s *SongoQueryValue) ValueStrings() (values []string) {
	values = append(values, s.Operator)
	if qv, ok := s.Next(); ok {
		values = append(values, qv.ValueStrings()...)
	} else {
		values = append(values, s.ValueString())
	}
	return
}

func (s *SongoQueryValue) String() string {
	if s.Value != nil {
		if qv, ok := s.Next(); ok {
			return s.Operator + qv.String()
		}
		return s.Operator + s.Value.(string)
	}
	return ""
}

func StringToValue(str string) interface{} {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return v
	} else if v, err := strconv.ParseFloat(str, 64); err == nil {
		return v
	} else if v, err := strconv.ParseUint(str, 10, 64); err == nil {
		return v
	} else if v, err := strconv.ParseBool(str); err == nil {
		return v
	} else if strings.Contains(str, ",") {
		return strings.Split(str, ",")
	}
	return str
}

func SplitQueryValue(value string) (v SongoQueryValue, ok bool) {
	ok = v.Split(value)
	return
}

func VerifyQueryValue(value string) bool {
	_, ok := SplitQueryValue(value)
	return ok
}
