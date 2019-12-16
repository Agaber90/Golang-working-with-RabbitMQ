package model

//CategoriesBody to hole the Categories date
type CategoriesBody struct {
	Categories Categories `json:"categories"`
}

//Categories to hold the category in case High Medium or Low
type Categories struct {
	High   High   `json:"high"`
	Medium Medium `json:"medium"`
	Low    Low    `json:"low"`
}

//High to holf the High Category date
type High struct {
	Amount      int      `json:"amount"`
	MaxSeverity float32  `json:"max_severity"`
	Services    []string `json:"services"`
}

//Medium to hold the Medium Category data
type Medium struct {
	Amount      int      `json:"amount"`
	MaxSeverity float32  `json:"max_severity"`
	Services    []string `json:"services"`
}

//Low to hold the Low Category data
type Low struct {
	Amount      int      `json:"amount"`
	MaxSeverity float32  `json:"max_severity"`
	Services    []string `json:"services"`
}
