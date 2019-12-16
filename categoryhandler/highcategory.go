package categoryhandler

//HighCategory to handle value if matched
type HighCategory struct {
	HighValue bool
}

//ApplyCategory to implement the High category
func (c HighCategory) ApplyCategory(value float32) bool {
	if value >= 8 && value <= 10 {
		c.HighValue = true
	}
	return c.HighValue
}
