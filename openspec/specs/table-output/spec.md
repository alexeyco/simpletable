## Purpose

This capability converts a table's header, body, and footer into terminal text and provides methods for writing that text to standard output.

## Requirements

### Requirement: Initialize table sections

The package SHALL provide a constructor that creates a table with empty header, body, and footer sections.

#### Scenario: Create a new table

**WHEN** a caller creates a table with `New`
**THEN** the table has empty header, body, and footer cell slices

### Requirement: Render table sections in order

The table renderer SHALL place a non-empty header before body rows and a non-empty footer after body rows.

#### Scenario: Render header, body, and footer

**WHEN** a table has a non-empty header, body rows, and a non-empty footer
**THEN** its rendered output places the header first, body rows next, and footer last

### Requirement: Render section dividers and table borders

The table renderer SHALL include a section divider after a non-empty header or before a non-empty footer only when the active style's divider center glyph is non-empty. It SHALL include top and bottom borders when their rendered border lines are non-empty.

#### Scenario: Render configured dividers and borders

**WHEN** a table has a non-empty header or footer and the active style has a non-empty divider center glyph
**THEN** the output includes the divider after the header or before the footer; top and bottom borders are included when their rendered lines are non-empty

### Requirement: Return newline-separated table text

The table's `String` method SHALL return its rendered lines joined with newline characters.

#### Scenario: Convert rendered lines to a string

**WHEN** a caller invokes `String` on a table
**THEN** the returned table lines are separated by newline characters

### Requirement: Write rendered table text to standard output

The table's `Print` method SHALL write the rendered string without adding a newline, and `Println` SHALL write it with a trailing newline.

#### Scenario: Print with and without a trailing newline

**WHEN** a caller invokes `Print` or `Println`
**THEN** `Print` writes the rendered string as-is and `Println` adds a trailing newline

### Requirement: Bound per-render allocations

The table renderer SHALL allocate materially less on repeated `String()`
calls while producing byte-identical output and an unchanged exported API,
including `Print` and `Println`. The package SHALL NOT add any guarantee for
concurrent `String()` calls on one shared table.

#### Scenario: Output stays byte identical

**WHEN** the golden-output test renders representative tables covering all 7
built-in styles plus multiline, ANSI-colored, CJK-wide, and spanned-cell
inputs
**THEN** each rendered string matches its committed `testdata/*.golden`
fixture byte for byte

#### Scenario: Allocation gate met per benchmark

**WHEN** the seven root `BenchmarkStyle*` benchmarks run via `make bench`
after the change
**THEN** each benchmark reports at least 30 percent lower `allocs/op` and at
least 15 percent lower `B/op` than its recorded baseline (allocs/op / B/op:
Default 671/28785, Compact 635/27841, CompactLite 635/27841, CompactClassic
587/25559, Markdown 663/28289, Rounded 671/28785, Unicode 691/32429),
evaluated per benchmark rather than as an aggregate

#### Scenario: Recycled cell content renders identically

**WHEN** the same table is rendered twice in a row
**THEN** both renders produce identical output matching the golden fixture

#### Scenario: Distinct tables render concurrently

**WHEN** different tables are rendered from separate goroutines under the
race detector
**THEN** each table renders correctly, no data race is reported, and no
same-table concurrency guarantee is added or removed
