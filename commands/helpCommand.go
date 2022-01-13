package commands

import (
    "utils"

    "fmt"
)

type HelpCommand struct {
}

func (cmd HelpCommand) GetShortHelp() string {
    return ""
}

func (cmd HelpCommand) GetLongHelp() string {
    return ""
}

func (cmd HelpCommand) Execute(args ...string) {
    fmt.Println(utils.GetHelp(args...))
}


func init() {
    cmd := HelpCommand{}
    
    utils.Commands = append(utils.Commands, cmd)
    
    // Add each alias
    utils.CommandAliases["help"] = cmd
}
