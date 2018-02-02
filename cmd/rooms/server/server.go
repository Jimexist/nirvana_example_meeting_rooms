package server

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/api"
)

// Run creates a new server and run blockingly
func Run() error {
	serverConfig := nirvana.NewDefaultConfig("0.0.0.0", 8000)
	serverConfig.Configure(nirvana.Descriptor(api.GetVersionDescriptor("rooms")))
	server := nirvana.NewServer(serverConfig)
	return server.Serve()
}
