package array

func IntersectSet[T comparable](firstArray []T, secondArray []T) []T {
	var result []T
	if len(firstArray) != 0 && len(secondArray) != 0 {
		table := map[T]bool{}
		for i := range firstArray {
			table[firstArray[i]] = true
		}

		for i := range secondArray {
			if val, ok := table[secondArray[i]]; ok && val {
				result = append(result, secondArray[i])
				table[secondArray[i]] = false
			}
		}
	}

	return result
}

func DifferenceSet[T comparable](firstArray []T, secondArray []T) []T {
	var result []T
	if len(firstArray) == 0 && len(secondArray) == 0 {
		return result
	} else if len(firstArray) == 0 {
		return CopyArray(secondArray)
	} else if len(secondArray) == 0 {
		return CopyArray(firstArray)
	}
	table := map[T]bool{}
	for i := range secondArray {
		table[secondArray[i]] = true
	}

	for i := range firstArray {
		if _, ok := table[firstArray[i]]; !ok {
			result = append(result, firstArray[i])
			table[firstArray[i]] = true
		}
	}

	return result
}

func DifferenceUnionSet[T comparable](firstArray []T, secondArray []T) []T {
	firstDiff := DifferenceSet(firstArray, secondArray)
	secondDiff := DifferenceSet(secondArray, firstArray)
	return Union(firstDiff, secondDiff)
}
