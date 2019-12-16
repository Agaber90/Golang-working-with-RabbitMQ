package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsFileExists(t *testing.T) {
	fileName := "../inputs/input_1.json"
	pvIsExist := IsFileExist(fileName)
	if pvIsExist {
		assert.True(t, pvIsExist, "File Exists")
	}
}

func Test_IsFileNotExists(t *testing.T) {
	fileName := "../inputs/input_1.json"
	pvIsExist := IsFileExist(fileName)
	if !pvIsExist {
		assert.False(t, pvIsExist, "File Is not Exists")
	}
}

func Test_CategoryValidation(t *testing.T) {
	parService := "A"
	var pvSeverity float32
	pvSeverity = 10
	pvIsValide := CategoryValidation(pvSeverity, parService)
	if pvIsValide {
		assert.True(t, pvIsValide, "Category is valid")
	}
}

func Test_CategoryIsNotValid(t *testing.T) {
	parService := "G"
	var pvSeverity float32
	pvSeverity = 25
	pvIsValide := CategoryValidation(pvSeverity, parService)
	if !pvIsValide {
		assert.False(t, pvIsValide, "Category is not valid")
	}
}
