package cmd

import (
	"fmt"
	"os"

	"github.com/winiceo/genv/internal/db"
	"github.com/spf13/cobra"
)

func newStatusCmd(s db.Store) *cobra.Command {
	statusDesc := "get current environment's status"

	statusLongDesc := `status - Get the current environment's status

Environments can be in different states:
- "ready": the environment is ready for use
- "error": the environment is in a bad state
- "off": the environment hasn't been created yet

To move from "off" to "ready" state, run "envctl create".

To fix "error" state, you can try recreating the environment with
"envctl destroy" followed by "envctl create".`

	statusReady := `The environment is ready!

Run "envctl login" to enter it.`

	statusError := `Something is wrong with the environment. :(

Try recreating it by running "envctl destroy", followed by "envctl create".`

	statusOff := `The environment is off.

Run "envctl create" to spin it up!`

	runStatus := func(cmd *cobra.Command, args []string) {
		env, err := s.Read()
		if err != nil {
			fmt.Printf("error reading data store: %v\n", err)
			os.Exit(1)
		}

		switch env.Status {
		case db.StatusReady:
			fmt.Println(statusReady)
		case db.StatusError:
			fmt.Println(statusError)
		case db.StatusOff:
			fmt.Println(statusOff)
		}
	}

	return &cobra.Command{
		Use:   "status",
		Short: statusDesc,
		Long:  statusLongDesc,
		Run:   runStatus,
	}
}
