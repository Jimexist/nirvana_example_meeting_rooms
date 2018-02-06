package api

import (
	"context"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/version"
)

// Version is the info returned during version query.
type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

// GetVersionDescriptor will return a version description customized with app name.
func GetVersionDescriptor(name string) definition.Descriptor {
	return definition.Descriptor{
		Path:     "/version",
		Produces: []string{definition.MIMEJSON},
		Definitions: []definition.Definition{
			{
				Method:      definition.Get,
				Description: "returns version information about this service",
				Results: []definition.Result{
					{
						Description: "version info",
						Destination: definition.Data,
					},
				},
				Function: func(_ context.Context) Version {
					return Version{
						Name:    name,
						Version: version.Version,
						Commit:  version.Commit,
					}
				},
			},
		},
	}
}
