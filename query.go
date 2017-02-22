package songo

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
	if s.query != nil {
		return len(s.query)
	}
	return 0
}

func (s SongoQuery) GetKeys() []string {
	return s.keys
}

func (s SongoQuery) GetValue(key string) (*SongoQueryValue, bool) {
	if str, ok := s.get(key); ok {
		if qv, ok := SplitQueryValue(str); ok {
			return &qv, true
		}
	}
	return nil, false
}

func (s SongoQuery) GetQuery(key string) (operators []string, value interface{}, ok bool) {
	var qv *SongoQueryValue
	if qv, ok = s.GetValue(key); ok {
		operators, value = qv.GetQuery()
	}
	return
}
