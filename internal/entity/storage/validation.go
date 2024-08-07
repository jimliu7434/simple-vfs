package storage

import (
	"regexp"
	"strings"
)

// IsValidUsername check if string only containes letters, numbers, and underscores
// and is between 3 and 20 characters long
func IsValidUsername(username string) bool {
	rgex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	if !rgex.MatchString(username) {
		return false
	}
	return true
}

// IsValidFoldername check if string not containes any special characters
// and is between 1 and 50 characters long
func IsValidFoldername(foldername string) bool {
	// TODO: allowd some symbols
	rgex := regexp.MustCompile(`^[a-zA-Z0-9_]{1,50}$`)
	if !rgex.MatchString(foldername) {
		return false
	}
	return true
}

// IsValidFilename check if string not containes any special characters
// and is between 1 and 50 characters long
func IsValidFilename(filename string) bool {
	// TODO: allowd some symbols
	rgex := regexp.MustCompile(`^[a-zA-Z0-9_]{1,50}$`)
	if !rgex.MatchString(filename) {
		return false
	}
	return true
}

// IsValidSortType check if string is one of the allowed values
// "asc" or "desc" (ASC or DESC are also allowed)
func IsValidSortType(sortType string) bool {
	sortType = strings.ToLower(sortType)
	if sortType != "asc" && sortType != "desc" {
		return false
	}
	return true
}
