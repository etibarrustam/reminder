package client

import (
	"flag"
	"fmt"
	"net/mail"
	"os"
	"time"
)

type BackendHTTPClient interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
	Edit(id, title, message string, duration time.Duration) ([]byte, error)
	Fetch(ids []string) ([]byte, error)
	Delete(ids []string) error
	Healthy(host string) bool
}

func NewSwitch(uri string) Switch  {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client: httpClient,
		backendAPIURL: uri,
	}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit": s.edit,
		"fetch": s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}

type Switch struct {
	client BackendHTTPClient
	backendAPIURL string
	commands map[string]func() func(string) error
}

func (s Switch) Switch() error  {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]

	if !ok {
		return fmt.Errorf("invalid command: '%s'\n", cmdName)
	}

	return cmd()(cmdName)
}

func (s Switch) Help()  {
	var help string

	for name := range s.commands {
		help += name + "\t --help\n"
	}

	fmt.Printf("Usage of: %s:\n <command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}
		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *m, *d)

		if err != nil {
			return wrapError("could not create reminder", err)
		}

		fmt.Printf("Reminder created successfully: \n%s", string(res))

		fmt.Println("create reminder")
		return nil
	}
}
func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(editCmd)

		if err := s.checkArgs(2); err != nil {
			return err
		}
		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *m, *d)

		if err != nil {
			return wrapError("could not edit reminder", err)
		}

		fmt.Printf("Reminder rfirf successfully: \n%s", string(res))

		fmt.Println("create reminder")
		return nil
	}
}
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fmt.Println("fetch reminder")
		return nil
	}
}
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		fmt.Println("delete reminder")
		return nil
	}
}
func (s Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("health reminder")
		return nil
	}
}

func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)

	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&t, "message", "", "Reminder message")
	f.StringVar(&t, "m", "", "Reminder message")
	f.StringVar(&t, "duration", "", "Reminder time")
	f.StringVar(&t, "d", "", "Reminder time")

	return &t, &m, &d
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])

	if err != nil {
		return wrapError("could not parse '" + cmd.Name() +"'", err)
	}

	return nil
}

func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args) - 2 < minArgs {
		fmt.Printf(
			"incorrect use of %s\n%s %s --help\n",
			os.Args[1], os.Args[0], os.Args[1],
		)
	}

	return fmt.Errorf(
		"%s excepts as least %d arg(s), %d provided",
		os.Args[1], minArgs, len(os.Args) - 2,
	)
}