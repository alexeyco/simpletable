package grid

// Collection of cells.
type Collection []*Cell

// Append cell to collection.
func (c *Collection) Append(cell *Cell) {
	if !c.Exists(cell) {
		*c = append(*c, cell)
	}
}

// Exists returns true if cell with the same row and column already exists in collection.
func (c *Collection) Exists(cell *Cell) bool {
	for _, collectionCell := range *c {
		if collectionCell.Row == cell.Row && collectionCell.Col == cell.Col {
			return true
		}
	}

	return false
}
