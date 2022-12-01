package utils

type Void interface{}

var void Void

type IntSet map[int]Void
type StringSet map[string]Void

func (s IntSet) Add(value int) (added bool) {
	if !s.Has(value) {
		s[value] = void
		added = true
	}
	return
}

func (s IntSet) Remove(value int) (removed bool) {
	removed = s.Has(value)
	delete(s, value)
	return
}

func (s IntSet) Has(value int) (has bool) {
	_, has = s[value]
	return
}

func (s StringSet) Add(value string) (added bool) {
	if !s.Has(value) {
		s[value] = void
		added = true
	}
	return
}

func (s StringSet) Remove(value string) (removed bool) {
	removed = s.Has(value)
	delete(s, value)
	return
}

func (s StringSet) Has(value string) (has bool) {
	_, has = s[value]
	return
}

func (s StringSet) Copy() StringSet {
	c := make(StringSet, len(s))
	for k := range s {
		c[k] = void
	}
	return c
}
