package operations

// UnsetElement unsets a certain value from a slice
func UnsetElement(arr []string, remove string) []string {
	for i, elem := range arr {
		if elem == remove {
			arr = removeIndex(arr, i)
		}
	}

	return arr
}

// AppendElement returns the slice if val is already inside it, else it will be appended
func AppendElement(arr []string, val string) []string {
	if Contains(val, arr) {
		return arr
	}

	arr = append(arr, val)

	return arr
}

// Contains checks for existence of an element inside a slice
func Contains(search string, list []string) bool {
	for _, v := range list {
		if v == search {
			return true
		}
	}

	return false
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func Intersection(a []string, b []string) (str []string) {
	if len(a) > len(b) {
		a, b = b, a
	}
	mapA, mapB := make(map[string]bool), make(map[string]bool)
	for _, e := range a {
		mapA[e] = true
	}
	for _, e := range b {
		mapB[e] = true
	}
	for _, e := range a {
		_, okA := mapA[e]
		_, okB := mapB[e]
		if okA && okB {
			str = append(str, e)
		}
	}

	return str
}

// Diff returns the elements in a that are not found in b
func Diff(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}
