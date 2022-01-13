package commands

import (
    "utils"

    // "github.com/TwiN/go-color"

    "fmt"
    "strings"
)

type CommitCommand struct {
}

func (cmd CommitCommand) GetShortHelp() string {
    return ""
}

func (cmd CommitCommand) GetLongHelp() string {
    return ``
}

func (cmd CommitCommand) Execute(args ...string) {
    toDelete := -1
    for i := 0; i < len(args); i++ {
        if
    }
    // TODO: Remove item at index toDelete

    // TODO: Commit
    if !noPush {
        // TODO: Push
        fmt.Println("PUSH!")
    } else {
        fmt.Println("DON'T PUSH!")
    }
    // out, errout, err := utils.ExecuteInConsole("git clean -fdX " + strings.Join(args, " "))
    // out = strings.TrimSuffix(out, "\n")
    // errout = strings.TrimSuffix(errout, "\n")
    
    // fmt.Println(color.Ize(utils.InfoColor, out))
    // fmt.Println(color.Ize(utils.ErrorColor, errout))
}


func init() {
    cmd := CommitCommand{}
    
    utils.Commands = append(utils.Commands, cmd)
    
    // Add each alias
    utils.CommandAliases["commit"] = cmd
}
