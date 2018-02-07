package api

import (
	"context"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/room"
)

// GetRoomDescriptor returns routes regarding rooms.
func GetRoomDescriptor() definition.Descriptor {
	idParam := definition.Parameter{
		Name:        "id",
		Source:      definition.Path,
		Description: "the id of the room",
	}
	return definition.Descriptor{
		Path:        "/rooms",
		Description: "rooms CRUD",
		Produces:    []string{definition.MIMEJSON},
		Children: []definition.Descriptor{
			{
				Path:        "/{id}",
				Description: "room",
				Definitions: []definition.Definition{
					{
						Method:      definition.Get,
						Description: "get by room",
						Parameters:  []definition.Parameter{idParam},
						Results: []definition.Result{
							{
								Description: "the room info",
								Destination: definition.Data,
								Operators:   []definition.Operator{},
							},
						},
						Function: func(ctx context.Context, ID int32) *room.Info {
							log.Infof("getting room by ID %v", ID)
							return nil
						},
					},
					{
						Method:      definition.Patch,
						Description: "update room info",
						Consumes:    []string{definition.MIMEJSON},
						Parameters: []definition.Parameter{idParam,
							{
								Name:        "name",
								Source:      definition.Body,
								Description: "the name to update",
							},
						},
						Results: []definition.Result{
							{
								Description: "the updated room info",
								Destination: definition.Data,
								Operators:   []definition.Operator{},
							},
						},
						Function: func(ctx context.Context, ID int32, name string) *room.Info {
							log.Infof("updating room %v with name %v", ID, name)
							return nil
						},
					},
				},
			},
		},
		Definitions: []definition.Definition{
			{
				Method:      definition.List,
				Description: "list all rooms",
				Parameters: []definition.Parameter{
					{
						Name:        "name",
						Source:      definition.Query,
						Description: "the name of the rooms to search for",
					},
				},
				Results: []definition.Result{
					{
						Description: "the rooms info",
						Destination: definition.Data,
						Operators:   []definition.Operator{},
					},
				},
				Function: func(ctx context.Context, name string) []room.Info {
					log.Infof("listing rooms with name %v", name)
					return []room.Info{}
				},
			},
		},
	}
}
