package commands

import (
    "utils"

    "github.com/TwiN/go-color"

    "fmt"
    "strings"
)

type CleanIgnoredCommand struct {
}

func (cmd CleanIgnoredCommand) GetShortHelp() string {
    return "Deletes all ignored files."
}

func (cmd CleanIgnoredCommand) GetLongHelp() string {
    return `Clean Ignored:
Deletes all files specified in the .gitignore file.
Usage: simplegit cleanignored
Aliases: clearignored`
}

func (cmd CleanIgnoredCommand) Execute(args ...string) {
    if utils.Exists(".gitignore") {
        out, errout, _ := utils.ExecuteInConsole("git clean -fdX " + strings.Join(args, " "))
        out = strings.TrimSuffix(out, "\n")
        errout = strings.TrimSuffix(errout, "\n")
        
        fmt.Println(color.Ize(utils.InfoColor, out))
        fmt.Println(color.Ize(utils.ErrorColor, errout))
    } else {
        fmt.Println(color.Ize(utils.ErrorColor, "SimpleGit: .gitignore file not found!"))
    }
}


func init() {
    cmd := CleanIgnoredCommand{}
    
    utils.Commands = append(utils.Commands, cmd)
    
    // Add each alias
    utils.CommandAliases["cleanignored"] = cmd
    utils.CommandAliases["clearignored"] = cmd
}
