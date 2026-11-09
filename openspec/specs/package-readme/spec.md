# Package Readme Specification

## Purpose

This capability defines how the repository README presents the package, its documented features, usage example, benchmarks, and comparison evidence. It helps readers assess the package quickly and gives maintainers a consistent way to keep that information current.

## Requirements

### Requirement: Order the README introduction

The README SHALL begin with the level-one heading `Simple tables in terminal with Go`, followed by the banner, badges, concise package description, feature comparison, and short Go example in that order.

#### Scenario: Read the README from the top

**WHEN** a reader opens the repository README
**THEN** the heading, banner, badges, package description, comparison, and short Go example appear in the required order, and the license section links to the license file

### Requirement: Reuse an abstract package banner

The reusable wide banner SHALL have a replaceable package-name wordmark, package-specific caption, and accent color. It SHALL use an asymmetrical broken-grid and row motif that suggests tabular structure without depicting a rendered table. The accent SHALL highlight selected strokes or cells. The simpletable banner SHALL use blue and display the caption `Simple tables in terminal with Go`.

#### Scenario: Reuse the banner for a Go package

**WHEN** a maintainer adapts the banner for another Go package
**THEN** the same wide abstract broken-grid and row motif displays the other package name and caption, accenting selected strokes or cells, while simpletable uses blue and the caption `Simple tables in terminal with Go`

### Requirement: Keep README examples and license accurate

The README SHALL use examples and feature counts that match the exported API and SHALL link to the license instead of reproducing its full text.

#### Scenario: Check API examples and style count

**WHEN** a reader follows an example or reads the built-in style count
**THEN** the example uses exported API fields and the count matches the seven built-in styles

#### Scenario: Match the example output

**WHEN** a reader runs the short Go example
**THEN** it produces the ASCII table shown immediately after the code block

### Requirement: Link comparison claims to public evidence

The README SHALL compare simpletable with termtables, uitable, tablewriter, and go-pretty using documented feature criteria. Each comparison SHALL link to the relevant public sources checked for that claim, use `✅` for `documented`, `⚠️` for `partial`, and `❌` for `not found in sources checked`, and show the date the comparison was checked. A legend SHALL explain these symbols. The `❌` status SHALL describe the sources reviewed and SHALL NOT assert that a library lacks the feature.

#### Scenario: Review comparison claims

**WHEN** a reader inspects a library and feature in the comparison
**THEN** the cell uses `✅`, `⚠️`, or `❌` according to the legend, links to the public sources checked, and the comparison shows its last-checked date

### Requirement: Extend the comparison table

Maintainers SHALL be able to extend the comparison by adding a library column or feature criterion without changing its format.

#### Scenario: Extend the comparison

**WHEN** a maintainer adds a library or a feature criterion
**THEN** the new column or row follows the existing comparison format and can include evidence links for its claims

### Requirement: Disclose current benchmark results in a collapsible section

The README SHALL keep benchmark results in a collapsible section and SHALL show results produced by the existing comparison harness on a MacBook. It SHALL identify the measured environment as `(MacBook)` and SHALL not present the historical Windows 10 / Intel Core i5-3470 measurements as MacBook results.

#### Scenario: Expand benchmark results

**WHEN** a reader expands the benchmark section
**THEN** the section shows the refreshed results for simpletable, termtables, and uitable, identifies the environment as `(MacBook)`, and links to the benchmark instructions

#### Scenario: Read the README without expanding benchmarks

**WHEN** a reader opens the README and leaves the benchmark section collapsed
**THEN** the benchmark results do not interrupt the primary description, comparison, and short example

### Requirement: Remind maintainers to refresh the comparison

The repository SHALL run a monthly scheduled GitHub Actions workflow that opens an issue to review and refresh the README comparison when no matching open reminder issue exists.

#### Scenario: Create the monthly reminder

**WHEN** the monthly schedule runs and no matching comparison-refresh issue is open
**THEN** the workflow opens an issue asking maintainers to check the comparison evidence and update the last-checked date

#### Scenario: Avoid duplicate open reminders

**WHEN** the monthly schedule runs while a matching comparison-refresh issue is already open
**THEN** the workflow does not open another reminder issue
