package cli

type CliError string

const (
	NoArgsError CliError = "no arguments passed"
)

func (e CliError) Error() string {
	return string(e)
}

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run() error {
	return NoArgsError
}
