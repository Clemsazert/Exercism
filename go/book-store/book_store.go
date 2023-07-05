package bookstore

const price = 8

var reductions = map[int]int{
	1: 0,
	2: 5,
	3: 10,
	4: 20,
	5: 25,
}

func GenerateBooksMap(books []int) map[int]int {
	booksMap := map[int]int{}
	for _, book := range books {
		booksMap[book] += 1
	}
	return booksMap
}

func ComputeCost(booksGroups [][]int) int {
	result := 0
	for _, group := range booksGroups {
		result += price * (100 - reductions[len(group)]) * len(group)
	}
	return result
}

func GenerateGroups(booksMap map[int]int) [][]int {
	var groups [][]int
	for book, amount := range booksMap {
		for i := 0; i < amount; i += 1 {
			if i < len(groups) {
				groups[i] = append(groups[i], book)
			} else {
				groups = append(groups, []int{book})
			}
		}
	}
	return groups
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func GenerateGroupsRec(groups [][]int, newBook int) [][]int {
	// Try to put newBook in a bag
	for i := 0; i < len(groups); i++ {
		if len(groups[i]) < 5 && !contains(groups[i], newBook) {
			groups[i] = append(groups[i], newBook)
			return groups
		}
	}
	groups = append(groups, []int{newBook})
	// TODO: Try to rearrange 2 groups
	return groups
}

func Cost(books []int) int {
	// booksMap := GenerateBooksMap(books)
	// groups := GenerateGroups(booksMap)
	var groups [][]int
	for _, book := range books {
		groups = GenerateGroupsRec(groups, book)
	}
	return ComputeCost(groups)
}
