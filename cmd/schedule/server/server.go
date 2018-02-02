package server

import (
	"github.com/caicloud/kubernetes-admin/api"
	"github.com/caicloud/nirvana"
)

// Run creates a new server and run blockingly
func Run() error {
	serverConfig := nirvana.NewDefaultConfig("0.0.0.0", 8000)
	server := nirvana.NewServer(serverConfig)
	serverConfig.Configure(nirvana.Descriptor(api.GetVersionDescriptor("rooms")))
	return server.Serve()
}
