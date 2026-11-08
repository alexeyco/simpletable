## Purpose

This capability prepares cell text for display by measuring line widths, applying alignment, matching row heights, and accounting for cells that span columns.

## Requirements

### Requirement: Split and trim cell text lines

Cell content SHALL be split at newline characters, and each resulting line SHALL be trimmed with `strings.TrimSpace` before its width is measured.

#### Scenario: Prepare multiline cell content

**WHEN** cell text contains newline-separated lines with surrounding whitespace
**THEN** the content stores the lines separately and trims each line's surrounding whitespace

### Requirement: Measure display width

The display width of a cell line SHALL be measured with `runewidth.StringWidth` after ANSI CSI escape sequences are removed.

#### Scenario: Measure colored text

**WHEN** a cell line contains an ANSI CSI escape sequence
**THEN** that sequence does not contribute to the measured display width

### Requirement: Align cell lines

Cell lines SHALL be padded to the cell content width according to left, center, or right alignment.

#### Scenario: Align lines within a cell

**WHEN** a line is shorter than the cell content width
**THEN** left alignment pads on the right, right alignment pads on the left, and center alignment divides padding with the smaller half on the left

### Requirement: Match cell content heights

Cell content height SHALL equal its number of lines, and increasing its height SHALL append empty lines without shrinking existing content.

#### Scenario: Pad shorter content

**WHEN** the requested content height exceeds the number of existing lines
**THEN** empty lines are appended until the requested height is reached

### Requirement: Size columns and spanning cells

Column width SHALL be based on the widest non-spanned cell in that column, and a cell SHALL be treated as spanned when its `Span` value is greater than one.

#### Scenario: Resize a spanning cell

**WHEN** a cell has `Span` greater than one and its content exceeds the width of its covered columns and spacing
**THEN** the covered columns are widened to accommodate the cell content
