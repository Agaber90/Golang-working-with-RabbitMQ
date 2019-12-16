package categoryhandler

//LowCategory for Low Category level
type LowCategory struct {
	LowValue bool
}

//ApplyCategory to implement the Low category
func (c LowCategory) ApplyCategory(value float32) bool {
	if value >= 0 && value <= 3.9 {
		c.LowValue = true
	}
	return c.LowValue
}
