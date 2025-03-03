package main

// HasValueInStringList checks if a given value exists in a list of strings.
// It takes a slice of strings (stringList) and a string value (value) as input parameters.
// It returns true if the value is found in the stringList, otherwise it returns false.
func HasValueInStringList(stringList []string, value string) bool {
	for _, string := range stringList {
		if string == value {
			return true
		}
	}
	return false
}
