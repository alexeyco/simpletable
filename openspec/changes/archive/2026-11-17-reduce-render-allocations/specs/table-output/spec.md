## ADDED Requirements

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
