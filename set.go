package main

type Set[T comparable] interface {
	Add(el T)
	Has(el T) bool
	Delete(el T)
	Size() int
	Differnce(set2 Set[T]) Set[T]
	Union(set2 Set[T]) Set[T]
	Intersection(set2 Set[T]) Set[T]
	IsSubSetOf(set2 Set[T]) bool
	IsSuperSetOf(set2 Set[T]) bool
	IsDisjointFrom(set2 Set[T]) bool
	SymmetricDiffernce(set2 Set[T]) Set[T]
	getElements() map[T]struct{}
}
type set[c comparable] struct {
	elements map[c]struct{}
	size     int
}

func NewSet[c comparable]() Set[c] {
	return &set[c]{}
}
func NewSetFromArr[c comparable](arr []c) Set[c] {
	set := NewSet[c]()
	for i := 0; i < len(arr); i++ {
		set.Add(arr[i])
	}
	return set
}
func (s *set[T]) Add(el T) {
	if s.elements == nil {
		s.elements = make(map[T]struct{})
	}
	if s.Has(el) {
		return
	}
	s.elements[el] = struct{}{}
	s.size++
}
func (s *set[T]) Has(el T) bool {
	_, ok := s.elements[el]
	return ok
}
func (s *set[T]) Delete(el T) {
	delete(s.elements, el)
	s.size--
}
func (s *set[T]) Size() int {
	return s.size
}
func (set1 *set[T]) Differnce(set2 Set[T]) Set[T] {
	diffSet := NewSet[T]()
	for val := range set1.getElements() {
		if !set2.Has(val) {
			diffSet.Add(val)
		}
	}
	return diffSet
}
func (set1 *set[T]) Union(set2 Set[T]) Set[T] {
	newSet := NewSet[T]()
	for val := range set1.getElements() {
		newSet.Add(val)
	}
	for val := range set2.getElements() {
		newSet.Add(val)
	}
	return newSet
}
func (set1 *set[T]) Intersection(set2 Set[T]) Set[T] {
	newSet := NewSet[T]()
	for val := range set1.getElements() {
		if set2.Has(val) {
			newSet.Add(val)
		}
	}
	for val := range set2.getElements() {
		if set1.Has(val) {
			newSet.Add(val)
		}
	}
	return newSet
}
func (set1 *set[T]) IsSubSetOf(set2 Set[T]) bool {
	for val := range set1.getElements() {
		if !set2.Has(val) {
			return false
		}
	}
	return true
}
func (set1 *set[T]) IsSuperSetOf(set2 Set[T]) bool {
	for val := range set2.getElements() {
		if !set1.Has(val) {
			return false
		}
	}
	return true
}
func (set1 *set[T]) IsDisjointFrom(set2 Set[T]) bool {
	interSet := set1.Intersection(set2)
	return interSet.Size() == 0
}
func (set1 *set[T]) SymmetricDiffernce(set2 Set[T]) Set[T] {
	return set1.Differnce(set2).Union(set2.Differnce(set1))

}
func (set1 *set[T]) getElements() map[T]struct{} {
	return set1.elements
}
