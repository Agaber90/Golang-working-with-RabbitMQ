package categoryhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsLowCategory(t *testing.T) {
	pvLowCatStrategy := CategoryStrategy{LowCategory{}, false}
	pvLowCatStrategy.ApplyCategoryStrategy(1)
	pvResult := pvLowCatStrategy.Result
	if pvResult {
		assert.True(t, pvResult, "Is Low")
	}
}

func Test_IsNotLowCategory(t *testing.T) {
	pvLowCatStrategy := CategoryStrategy{LowCategory{}, false}
	pvLowCatStrategy.ApplyCategoryStrategy(4)
	pvResult := pvLowCatStrategy.Result
	if !pvResult {
		assert.False(t, pvResult, "Is NOT Low")
	}
}

func Test_IsMedCategory(t *testing.T) {
	pvMedCatStrategy := CategoryStrategy{MediumCategory{}, false}
	pvMedCatStrategy.ApplyCategoryStrategy(4)
	pvResult := pvMedCatStrategy.Result
	if pvResult {
		assert.True(t, pvResult, "Is Medium")
	}
}

func Test_IsNotMedCategory(t *testing.T) {
	pvMedCatStrategy := CategoryStrategy{MediumCategory{}, false}
	pvMedCatStrategy.ApplyCategoryStrategy(8)
	pvResult := pvMedCatStrategy.Result
	if !pvResult {
		assert.False(t, pvResult, "Is NOT Medium")
	}
}

func Test_IsHighCategory(t *testing.T) {
	pvHighCatStrategy := CategoryStrategy{HighCategory{}, false}
	pvHighCatStrategy.ApplyCategoryStrategy(8)
	pvResult := pvHighCatStrategy.Result
	if pvResult {
		assert.True(t, pvResult, "Is High")
	}
}

func Test_IsNotHighCategory(t *testing.T) {
	pvHighCatStrategy := CategoryStrategy{HighCategory{}, false}
	pvHighCatStrategy.ApplyCategoryStrategy(11)
	pvResult := pvHighCatStrategy.Result
	if !pvResult {
		assert.False(t, pvResult, "Is NOT High")
	}
}
