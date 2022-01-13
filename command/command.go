package command

type Command interface {
    GetShortHelp() string
    GetLongHelp() string
    Execute(args ...string)
}
