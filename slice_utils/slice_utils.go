package slice_utils

func Index[T comparable](s []T, wanted T) int {
	// return the index of the first apperance of wanted in s
	// if wanted does not exist in s, return -1

	for index, elem := range s {
		if elem == wanted {
			return index
		}
	}

	return -1
}
