package helper

import (
	"fmt"
	"os"
	"regexp"
)

//IsFileExist to validate file is Exists
func IsFileExist(parFileName string) bool {
	pvFileInfo, pvError := os.Stat(parFileName)
	if os.IsNotExist(pvError) {
		return false
	}
	return !pvFileInfo.IsDir()
}

//CategoryValidation to validate the following
//1-validate everity between 0.0 and 10.0
//2-validate Service from A to F
func CategoryValidation(parSeverity float32, parService string) bool {
	pvMatched, pvError := regexp.MatchString(`[A-F]`, parService)
	if pvError != nil {
		fmt.Println(pvError)
	}
	if (parSeverity >= 0.0 && parSeverity <= 10.0) && pvMatched {
		return true
	}

	return false
}
