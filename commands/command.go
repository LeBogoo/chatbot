package commands

type Command interface {
	Execute(args []string, isAdmin bool) (string, error)
	Name() string
}
