package slices

import "sort"

func SortStringsInPlace(s []string) []string {
	r := s
	sort.Strings(r)
	return r
}

func SortStringsPure(s []string) []string {
	r := s
	sort.Strings(r)
	return r
}

type User struct {
	FirstName string
	LastName  string
}

func SortUsersPure(s []User) []User {
	r := s
	sort.Slice(r, func(i, j int) bool {
		if r[i].FirstName == r[j].FirstName {
			return r[i].LastName < r[j].LastName
		}

		return r[i].FirstName < r[j].FirstName
	})
	return r
}

func SortUsersPureDesc(s []User) []User {
	r := s
	sort.Slice(r, func(i, j int) bool {
		return r[i].FirstName > r[j].FirstName
	})
	return r
}
