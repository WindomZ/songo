package mongo

import "github.com/WindomZ/songo"

type Songo struct {
	songo.Songo
	least SongoResultMap
	must  SongoResultMap
}

func (s *Songo) init() *Songo {
	if s.least == nil {
		s.least = make(SongoResultMap)
	}
	if s.must == nil {
		s.must = make(SongoResultMap)
	}
	return s
}

func (s *Songo) Least(key string, value interface{}) {
	if len(key) != 0 && value != nil {
		s.init().least[key] = value
	}
}

func (s *Songo) Must(key string, value interface{}) {
	if len(key) != 0 && value != nil {
		s.init().must[key] = value
	}
}

func (s *Songo) songoResult() *SongoResult {
	return &SongoResult{
		songo:  s,
		result: make(SongoResultMap, s.Query.Size()+len(s.init().must)),
	}
}

func (s *Songo) Result() SongoResultMap {
	r := s.songoResult() // get new SongoResult
	for _, k := range s.Query.GetKeys() {
		if vs, ok := s.Query.GetValues(k); ok {
			for _, v := range vs {
				os, v := v.GetQuery()
				r.Update(k, os, v)
			}
		}
	}
	return r.Result()
}
