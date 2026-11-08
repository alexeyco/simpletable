## Purpose

This capability exposes built-in and caller-supplied style values that describe border, section-divider, and cell-separator components.

## Requirements

### Requirement: Provide built-in table styles

The package SHALL provide seven built-in styles: Default, Compact, Compact Lite, Compact Classic, Markdown, Rounded, and Unicode.

#### Scenario: Access a built-in style

**WHEN** a caller accesses one of the package's built-in style values
**THEN** the style value is available as a `*Style`

### Requirement: Define style border and separator components

A table style SHALL expose border components, divider components, and the string used to separate cells.

#### Scenario: Configure style components

**WHEN** a caller accesses a `Style` value
**THEN** it exposes `Border`, `Divider`, and `Cell` fields, with border and divider components defined by their respective structs

### Requirement: Apply a caller-supplied style

The table's `SetStyle` method SHALL set the active style to the supplied style pointer.

#### Scenario: Set a custom style

**WHEN** a caller invokes `SetStyle` with a style pointer
**THEN** the table's active style pointer is the supplied pointer
