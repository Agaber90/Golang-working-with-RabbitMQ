package categoryhandler

//MediumCategory for medium Category level
type MediumCategory struct {
	MediumValue bool
}

//ApplyCategory to implement the medium category
func (m MediumCategory) ApplyCategory(value float32) bool {
	m.MediumValue = false
	if value >= 4 && value <= 7.9 {
		m.MediumValue = true

	}
	return m.MediumValue
}
