package storage

import "sort"

func sortAscByName[T ISortable](a []T, reverse bool) {
	sortFunc := func(i, j int) bool {
		return a[i].GetName() < a[j].GetName()
	}
	if reverse {
		sortFunc = func(i, j int) bool {
			return a[i].GetName() > a[j].GetName()
		}
	}
	sort.SliceStable(a, sortFunc)
	return
}

func sortAscByTime[T ISortable](a []T, reverse bool) {
	sortFunc := func(i, j int) bool {
		return a[i].GetCreatedAt().Before(a[j].GetCreatedAt())
	}
	if reverse {
		sortFunc = func(i, j int) bool {
			return a[i].GetCreatedAt().After(a[j].GetCreatedAt())
		}
	}
	sort.SliceStable(a, sortFunc)
	return
}
