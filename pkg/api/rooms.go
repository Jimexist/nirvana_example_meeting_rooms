package api

import (
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/room"
)

func GetRoomListDescriptor() definition.Descriptor {
	return
	definition.ListDefinitionFor(room.Info{}, "list rooms",
	).Produce(definition.MIMEJSON,
	).Result(definition.DataResultFor("").Operator(func() {

	}),

	)
}
