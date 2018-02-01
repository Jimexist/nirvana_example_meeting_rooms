package server

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana_example_meeting_rooms/cmd/rooms/api"
)

func Run() error {
	serverConfig := nirvana.NewDefaultConfig("0.0.0.0", 8000)
	serverConfig.Configure(func(c *nirvana.Config) error {
		c.Descriptors = append(c.Descriptors, api.DefineVersion)
		return nil
	})
	server := nirvana.NewServer(serverConfig)
	return server.Serve()
}
