package server

import (
	"github.com/caicloud/nirvana"
)

func Run() error {
	serverConfig := nirvana.NewDefaultConfig("0.0.0.0", 8000)
	server := nirvana.NewServer(serverConfig)
	return server.Serve()
}
