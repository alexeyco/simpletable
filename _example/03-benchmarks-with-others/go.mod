module github.com/alexeyco/simpletable/_example/03-benchmarks-with-others

go 1.23

require (
	github.com/alexeyco/simpletable v0.0.0
	github.com/apcera/termtables v0.0.0-20170405184538-bcbc5dc54055
	github.com/gosuri/uitable v0.0.4
)

require (
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.31 // indirect
	golang.org/x/sys v0.29.0 // indirect
)

replace github.com/alexeyco/simpletable => ../..
