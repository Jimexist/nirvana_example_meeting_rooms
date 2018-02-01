package api

import (
	"context"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/version"
)

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

var DefineVersion = definition.Descriptor{
	Path:     "/version",
	Produces: []string{definition.MIMEJSON},
	Definitions: []definition.Definition{
		{
			Method: definition.Get,
			Results: []definition.Result{
				{
					Description: "version info",
					Destination: definition.Data,
				},
			},
			Function: func(ctx context.Context) Version {
				return Version{
					Name:    "rooms",
					Version: version.Version,
					Commit:  version.Commit,
				}
			},
		},
	},
}
