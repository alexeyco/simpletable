# Reposition the simpletable README

## Why now

The current README is long and does not lead with a clear package summary, a decision-focused comparison, or a short example. Some of its examples and feature counts also disagree with the exported API and the existing style specification. The benchmark figures are from an old Windows machine and need a fresh, clearly labeled run before they are presented as current.

## What changes

Create a wide, abstract banner that can be reused across Go package repositories by replacing the package name, package-specific caption, and accent color; use blue for simpletable and the caption `Simple tables in terminal with Go`. Restore the `# Simple tables in terminal with Go` heading above the banner. Follow it with badges, a concise description, an extensible feature comparison using check, warning, and cross symbols with an evidence legend, and a short Go example. Keep benchmark results in a collapsible section and refresh them with the existing benchmark harness on a MacBook. Replace the full license text with a link, and correct examples and feature counts to match the package API.

Add a monthly GitHub Actions workflow that opens an issue reminding maintainers to check and refresh the comparison. Keep the displayed Quick start output consistent with the code that produces it, and keep the reminder logic in `scripts/reminder.js`, called by `.github/workflows/reminder.yml`.

## Impact on capabilities

Add the `package-readme` capability for the README presentation, evidence-based feature comparison, benchmark disclosure, and scheduled comparison reminder. Table rendering, cell layout, and built-in styles do not change.

ASSUMPTION: The monthly reminder skips opening a new issue while a matching comparison-refresh issue remains open (untested — a duplicate issue is created when a matching open issue already exists).
