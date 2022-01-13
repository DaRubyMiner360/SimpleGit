module main

require command v0.0.0

replace command v0.0.0 => ./command

require utils v0.0.0

replace utils v0.0.0 => ./utils

require (
	commands v0.0.0
	github.com/TwiN/go-color v1.0.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
)

replace commands v0.0.0 => ./commands

go 1.16
