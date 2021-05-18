package database

import "sort"

func findNewId(repo TodoRepository) int {
	sort.Sort(repo.data)
	for i := range repo.data {
		if i != repo.data[i].Id {
			return i
		}
	}
	return repo.data.Len()
}
