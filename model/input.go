package model

//FindingBody struct to hold the whole findings data
type FindingBody struct {
	ID       int       `json:"id"`
	Findings []Finding `json:"findings"`
}

//Finding struct to hold finding data
type Finding struct {
	Name        string       `json:"name"`
	Severity    float32      `json:"severity"`
	Service     string       `json:"service"`
	SubFindings []SubFinding `json:"sub_findings"`
}

//SubFinding to hold the sub finding
type SubFinding struct {
	Name     string  `json:"name"`
	Severity float32 `json:"severity"`
}
