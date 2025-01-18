package commands

type Command interface {
	Execute(username string, args []string, isAdmin bool) (string, error)
	Name() string
}
