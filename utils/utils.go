package utils

import (
    "command"

    "github.com/TwiN/go-color"
    
    "bytes"
	"errors"
	"strings"
    "fmt"
	// "flag"
	"os"
    "os/exec"
)

var Commands = []command.Command{}
var CommandAliases = make(map[string]command.Command)

var InfoColor = "\033[94m"
// var InfoEmphasisColor = "\033[34m"
var InfoEmphasisColor = InfoColor
// var WarningColor = color.Yellow
var ErrorColor = "\033[91m" + color.Bold
// var ErrorEmphasisColor = "\033[31m" + color.Bold
var ErrorEmphasisColor = ErrorColor

func Exists(name string) bool {
    _, err := os.Stat(name)
    if err == nil {
        return true
    }
    if errors.Is(err, os.ErrNotExist) {
        return false
    }
    return false
}

func Contains(s []command.Command, e command.Command) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func Prepend(x []string, y string) []string {
    x = append(x, "")
    copy(x[1:], x)
    x[0] = y
    return x
}

func ExecuteInConsole(command string) (string, string, error) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return stdout.String(), stderr.String(), err
}

func GetHelp(args ...string) string {
	help := ""
	if len(args) > 0 && args[0] != "" {
        value, found := CommandAliases[args[0]]
        found = found && value.GetLongHelp() != ""
        if found {
            help = color.Ize(InfoColor, value.GetLongHelp())
        } else {
            out, eo, e := ExecuteInConsole(strings.Join(Prepend(args, "git")[:], " "))
            out = strings.TrimSuffix(out, "\n")
            eo = strings.TrimSuffix(eo, "\n")

            if e != nil {
                if strings.Contains(eo, "is not a git command.") {
                    help = color.Ize(ErrorColor, "SimpleGit: '" + color.Reset + ErrorEmphasisColor + args[0] + color.Reset + ErrorColor + "' is not a Git or SimpleGit command. See 'simplegit --help'.")
                } else {
                    help = color.Ize(ErrorColor, eo)
                }
            } else {
                help = out
            }
        }
	} else {
        help = color.Ize(InfoColor, "See " + InfoEmphasisColor + "simplegit help <command>" + color.Reset + InfoColor + " to read about a specific subcommand.\n\n")
        usedCmds := []command.Command{}
        for key, value := range CommandAliases {
            if !Contains(usedCmds, value) && value.GetShortHelp() != "" {
                help += InfoEmphasisColor + key + color.Reset + InfoColor + strings.Repeat(" ", 24 - len(key)) + value.GetShortHelp() + "\n"
                usedCmds = append(usedCmds, value)
            }
        }
        help += color.Ize(InfoColor, `

These are common Git commands used in various situations:

start a working area (see also: ` + InfoEmphasisColor + `simplegit help tutorial` + color.Reset + InfoColor + `)
   ` + InfoEmphasisColor + `clone` + color.Reset + InfoColor + `             Clone a repository into a new directory
   ` + InfoEmphasisColor + `init` + color.Reset + InfoColor + `              Create an empty Git repository or reinitialize an existing one

work on the current change (see also: ` + InfoEmphasisColor + `simplegit help everyday` + color.Reset + InfoColor + `)
   ` + InfoEmphasisColor + `add` + color.Reset + InfoColor + `               Add file contents to the index
   ` + InfoEmphasisColor + `mv` + color.Reset + InfoColor + `                Move or rename a file, a directory, or a symlink
   ` + InfoEmphasisColor + `restore` + color.Reset + InfoColor + `           Restore working tree files
   ` + InfoEmphasisColor + `rm` + color.Reset + InfoColor + `                Remove files from the working tree and from the index
   ` + InfoEmphasisColor + `sparse-checkout` + color.Reset + InfoColor + `   Initialize and modify the sparse-checkout

examine the history and state (see also: ` + InfoEmphasisColor + `simplegit help revisions` + color.Reset + InfoColor + `)
   ` + InfoEmphasisColor + `bisect` + color.Reset + InfoColor + `            Use binary search to find the commit that introduced a bug
   ` + InfoEmphasisColor + `diff` + color.Reset + InfoColor + `              Show changes between commits, commit and working tree, etc
   ` + InfoEmphasisColor + `grep` + color.Reset + InfoColor + `              Print lines matching a pattern
   ` + InfoEmphasisColor + `log` + color.Reset + InfoColor + `               Show commit logs
   ` + InfoEmphasisColor + `show` + color.Reset + InfoColor + `              Show various types of objects
   ` + InfoEmphasisColor + `status` + color.Reset + InfoColor + `            Show the working tree status

grow, mark and tweak your common history
   ` + InfoEmphasisColor + `branch` + color.Reset + InfoColor + `            List, create, or delete branches
   ` + InfoEmphasisColor + `commit` + color.Reset + InfoColor + `            Record changes to the repository
   ` + InfoEmphasisColor + `merge` + color.Reset + InfoColor + `             Join two or more development histories together
   ` + InfoEmphasisColor + `rebase` + color.Reset + InfoColor + `            Reapply commits on top of another base tip
   ` + InfoEmphasisColor + `reset` + color.Reset + InfoColor + `             Reset current HEAD to the specified state
   ` + InfoEmphasisColor + `switch` + color.Reset + InfoColor + `            Switch branches
   ` + InfoEmphasisColor + `tag` + color.Reset + InfoColor + `               Create, list, delete or verify a tag object signed with GPG

collaborate (see also: ` + InfoEmphasisColor + `simplegit help workflows` + color.Reset + InfoColor + `)
   ` + InfoEmphasisColor + `fetch` + color.Reset + InfoColor + `             Download objects and refs from another repository
   ` + InfoEmphasisColor + `pull` + color.Reset + InfoColor + `              Fetch from and integrate with another repository or a local branch
   ` + InfoEmphasisColor + `push` + color.Reset + InfoColor + `              Update remote refs along with associated objects

` + InfoEmphasisColor + `simplegit help -a` + color.Reset + InfoColor + ` and ` + InfoEmphasisColor + `simplegit help -g` + color.Reset + InfoColor + ` list available subcommands and some
concept guides. See ` + InfoEmphasisColor + `simplegit help <command>` + color.Reset + InfoColor + ` or ` + InfoEmphasisColor + `simplegit help <concept>` + color.Reset + InfoColor + `
to read about a specific subcommand or concept.
See ` + InfoEmphasisColor + `simplegit help git` + color.Reset + InfoColor + ` for an overview of the system.`)
	}
	return help
}

func ExecuteCommand(args []string) (string, string, error) {
	var out string
    var errout string
    var err error
	value, found := CommandAliases[args[0]]
	if found {
        cmdArgs := args[1:]
		value.Execute(cmdArgs...)
	} else {
		out, eo, e := ExecuteInConsole(strings.Join(Prepend(args, "git")[:], " "))
        out = strings.TrimSuffix(out, "\n")
        eo = strings.TrimSuffix(eo, "\n")

		if e != nil {
            if strings.Contains(eo, "is not a git command.") {
                fmt.Println("SimpleGit: '" + args[0] + "' is not a Git or SimpleGit command. See 'simplegit --help'.")
            } else {
                fmt.Println(eo)
            }
		} else {
            fmt.Println(out)
		}
	}
	return out, errout, err
}
