## MODIFIED Requirements

### Requirement: Size columns and spanning cells

Column width SHALL be based on the widest non-spanned cell in that column, and a cell SHALL be treated as spanned when its `Span` value is greater than one. The layout SHALL resize spanning cells in two passes: first widening the covered columns until every spanning cell's content fits its total covered width, then setting each spanning cell's width to the final total width of its covered columns and spacing. The rendered table SHALL not depend on the order of spanning cells.

#### Scenario: Resize a spanning cell

**WHEN** a cell has `Span` greater than one and its content exceeds the width of its covered columns and spacing
**THEN** the covered columns are widened to accommodate the cell content

#### Scenario: Later spanning cell widens shared columns

**WHEN** a spanning cell's columns are widened while processing a different spanning cell that appears later in the table
**THEN** the first spanning cell's width equals the final total width of its covered columns and spacing, and every rendered row line has the same length
