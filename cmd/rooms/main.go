package main

import (
	"github.com/caicloud/nirvana/cli"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana_example_meeting_rooms/cmd/rooms/server"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/db"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cli.NewCommand(&cobra.Command{
		Use:   "rooms",
		Short: "rooms is a micro-service that handles room CRUDs",
		Long:  "rooms is a micro-service that handles room CRUDs",
		Run: func(cmd *cobra.Command, args []string) {
			if err := server.Run(); err != nil {
				log.Error(err)
			} else {
				log.Info("server exited without an error")
			}
		},
	})
	if err := cmd.AddFlag(db.Flags()...); err != nil {
		log.Errorf("failed to add flags: %v", err)
	} else if err := cmd.Execute(); err != nil {
		log.Warningf("error execution command %v", err)
	}
}
