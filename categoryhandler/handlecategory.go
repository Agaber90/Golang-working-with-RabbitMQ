package categoryhandler

import (
	"cs-backend-golang-finding/helper"
	"cs-backend-golang-finding/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//HandleCategory to start apply the change to the category struct and move the result to the output folder
func HandleCategory(parFinding *model.FindingBody, parCategory *model.CategoriesBody, parFolder string) {
	pvLowCatStrategy := CategoryStrategy{LowCategory{}, false}
	pvMediumCatStrategy := CategoryStrategy{MediumCategory{}, false}
	pvHighCatStrategy := CategoryStrategy{HighCategory{}, false}
	
	if len(parFinding.Findings) > 0 {
		for _, pvFinding := range parFinding.Findings {
			if helper.CategoryValidation(pvFinding.Severity, pvFinding.Service) {
				pvLowCatStrategy.ApplyCategoryStrategy(pvFinding.Severity)
				pvLowResult := pvLowCatStrategy.Result
				pvMediumCatStrategy.ApplyCategoryStrategy(pvFinding.Severity)
				pvMediumResult := pvMediumCatStrategy.Result
				pvHighCatStrategy.ApplyCategoryStrategy(pvFinding.Severity)
				pvHighResult := pvHighCatStrategy.Result

				if pvLowResult {
					parCategory.Categories.Low.Services = append(parCategory.Categories.Low.Services, pvFinding.Service)
					parCategory.Categories.Low.Amount = len(parCategory.Categories.Low.Services)
					if parCategory.Categories.Low.MaxSeverity > pvFinding.Severity {
						parCategory.Categories.Low.MaxSeverity = pvFinding.Severity
					} else {
						parCategory.Categories.Low.MaxSeverity = pvFinding.Severity
					}
				} else if pvMediumResult {
					parCategory.Categories.Medium.Services = append(parCategory.Categories.Medium.Services, pvFinding.Service)
					parCategory.Categories.Medium.Amount = len(parCategory.Categories.Medium.Services)
					if parCategory.Categories.Medium.MaxSeverity > pvFinding.Severity {
						parCategory.Categories.Medium.MaxSeverity = pvFinding.Severity
					} else {
						parCategory.Categories.Medium.MaxSeverity = pvFinding.Severity
					}
				} else if pvHighResult {
					parCategory.Categories.High.Services = append(parCategory.Categories.High.Services, pvFinding.Service)
					parCategory.Categories.High.Amount = len(parCategory.Categories.High.Services)
					if parCategory.Categories.High.MaxSeverity > pvFinding.Severity {
						parCategory.Categories.High.MaxSeverity = pvFinding.Severity
					} else {
						parCategory.Categories.High.MaxSeverity = pvFinding.Severity
					}
				}
			}
		}
		pvResultFile := "result_" + fmt.Sprint(parFinding.ID) + ".json"
		if helper.IsFileExist(pvResultFile) {
			os.Remove(pvResultFile)
		}
		pvCategoryResult, pvError := json.Marshal(parCategory)
		if pvError != nil {
			log.Fatal(pvError)
			panic(pvError)
		}

		pvError = ioutil.WriteFile(parFolder+pvResultFile, pvCategoryResult, 0644)
	}
}
