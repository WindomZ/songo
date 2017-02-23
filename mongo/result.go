package mongo

import "github.com/WindomZ/songo"

type SongoResultMap map[string]interface{}

type SongoResult struct {
	songo    *Songo
	result   SongoResultMap
	lastMap  SongoResultMap
	lastName string
	last     interface{}
}

func (s *SongoResult) Reset() {
	s.lastMap = s.result
	s.lastName = ""
	s.last = s.result
}

func (s SongoResult) IsEmpty() bool {
	return s.result == nil || len(s.result) == 0
}

func (s *SongoResult) get(key, operator string) bool {
	//defer func() {
	//	println(fmt.Sprintf("get: %v - %v", key, operator))
	//	println(fmt.Sprintf("get: %#v", s.lastMap))
	//	println(fmt.Sprintf("get: %#v", s.lastName))
	//	println(fmt.Sprintf("get: %#v", s.last))
	//}()
	if songo.IsQueryOperatorGroup(operator) {
		// like: {"$or": [{"age": 23 }, {"name": "robin"}]}
		s.lastName = operator
		if v, ok := s.lastMap[operator]; ok {
			s.last = v // []SongoResultMap{}
		} else {
			s.last = []SongoResultMap{} // []SongoResultMap{}
			s.lastMap[operator] = s.last
		}
	} else if songo.IsQueryOperatorKV(operator) {
		// like: {"age": {"$gt" :20}}
		// like: {"age": {"$in": [20, 22, 25]}}
		if _, ok := s.last.([]SongoResultMap); ok {
			return true
		}
		if v, ok := s.lastMap[key]; ok {
			if v, ok := v.(SongoResultMap); ok {
				s.last = v // SongoResultMap{}
			} else {
				return false
			}
		} else {
			s.last = SongoResultMap{} // SongoResultMap{}
			s.lastMap[key] = s.last
		}
	} else if songo.IsQueryOperatorV(operator) {
		// like: {"age": 20}
		if _, ok := s.last.([]SongoResultMap); ok {
			return true
		}
		s.last = s.lastMap // SongoResultMap{}
	} else {
		return false
	}
	return true
}

func (s *SongoResult) set(key, operator string, value interface{}) bool {
	//defer func() {
	//	println(fmt.Sprintf("set: %v - %v - %#v", key, operator, value))
	//	println(fmt.Sprintf("set: %#v", s.lastMap))
	//	println(fmt.Sprintf("set: %#v", s.lastName))
	//	println(fmt.Sprintf("set: %#v", s.last))
	//}()
	if songo.IsQueryOperatorGroup(operator) {
		// like: {"$or": [{"age": 23 }, {"name": "robin"}]}
		return true
	} else if songo.IsQueryOperatorKV(operator) {
		// like: {"age": {"$gt" :20}}
		// like: {"age": {"$in": [20, 22, 25]}}
		if v, ok := s.last.([]SongoResultMap); ok {
			s.last = append(v, SongoResultMap{operator: value})
			s.lastMap[s.lastName] = s.last
		} else if v, ok := s.last.(SongoResultMap); ok {
			v[operator] = value
		}
	} else if songo.IsQueryOperatorV(operator) {
		// like: {"age": 20}
		if v, ok := s.last.([]SongoResultMap); ok {
			s.last = append(v, SongoResultMap{key: value})
			s.lastMap[s.lastName] = s.last
		} else if v, ok := s.last.(SongoResultMap); ok {
			v[key] = value
		}
	} else {
		return false
	}
	return true
}

func (s *SongoResult) Update(key string, operators []string, value interface{}) {
	s.Reset()
	for _, o := range operators {
		if !s.get(key, o) {
			break
		} else if !s.set(key, o, value) {
			break
		}
	}
}

func (s *SongoResult) Result() SongoResultMap {
	if s.songo == nil {
		return s.result
	}
	if s.songo.must != nil && len(s.songo.must) != 0 {
		for k, v := range s.songo.must {
			s.result[k] = v
		}
	}
	if s.songo.least != nil && s.IsEmpty() {
		for k, v := range s.songo.least {
			s.result[k] = v
		}
	}
	s.songo = nil
	return s.result
}
