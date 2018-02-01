package main

import (
	"github.com/caicloud/nirvana/cli"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/db"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cli.NewCommand(&cobra.Command{
		Use:   "schedule",
		Short: "schedule is a micro-service that handles scheduling CRUDs",
		Long:  "schedule is a micro-service that handles scheduling CRUDs",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("this if cool!")
		},
	})
	cmd.AddFlag(db.Flags()...)
	if err := cmd.Execute(); err != nil {
		log.Warningf("error execution command %v", err)
	}
}
