package categoryhandler

//ICategory to apply this interface across the category level
type ICategory interface {
	ApplyCategory(parServerity float32) bool
}

//CategoryStrategy to start implenet the stratgey across category level
type CategoryStrategy struct {
	Category ICategory
	Result   bool
}

//ApplyCategoryStrategy to be applied across category level
func (c *CategoryStrategy) ApplyCategoryStrategy(parServerity float32) bool {
	c.Result = c.Category.ApplyCategory(parServerity)
	return c.Result
}
