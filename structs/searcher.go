package structs

import (
	"reflect"
	"strconv"
	"strings"
)

// Searcher is a interface for search fields in struct.
type Searcher interface {
	SearchField(string) []reflect.Value
	GetField(string) []reflect.Value
}

type searcher struct {
	Origin interface{}
	Values []reflect.Value

	delitimter string
}

// NewSearcher return a new searcher with struct s.
func NewSearcher(s interface{}, opts ...Option) Searcher {
	sc := &searcher{Origin: s, delitimter: ".", Values: []reflect.Value{reflect.ValueOf(s)}}

	for _, opt := range opts {
		opt(sc)
	}
	return sc
}

func (s *searcher) SearchField(path string) []reflect.Value {
	var (
		vs     []reflect.Value
		fields = strings.Split(path, s.delitimter)
	)
	for _, field := range fields {
		vs = s.GetField(field)
		s.Values = []reflect.Value{}
		for _, v := range vs {
			switch v.Kind() {
			case reflect.Ptr:
				s.Values = append(s.Values, reflect.Indirect(v))
			default:
				s.Values = append(s.Values, v)
			}
		}
	}
	return s.Values
}

func (s *searcher) GetField(field string) (result []reflect.Value) {
	for _, v := range s.Values {
		switch v.Kind() {
		case reflect.Slice:
			idx, err := strconv.Atoi(field)
			if err == nil {
				if idx >= v.Len() {
					continue
				}
				result = append(result, v.Index(idx))
				continue
			}
			if field == "*" {
				for idx := 0; idx < v.Len(); idx++ {
					result = append(result, v.Index(idx))
				}
			}

		case reflect.Ptr:
			v = reflect.Indirect(v)
			fallthrough
		default:
			result = append(result, v.FieldByName(field))
		}
	}

	return result
}
