package list

func StringPos(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func ContainString(slice []string, element string) bool {
	return !(StringPos(slice, element) == -1)
}
