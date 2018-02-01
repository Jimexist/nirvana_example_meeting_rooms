package meeting

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/caicloud/nirvana_example_meeting_rooms/pkg/room"
)

// Info models the meeting's basic information
type Info struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	Room      *room.Info `json:"room"`
}

const tableName = "meetings"

func findByIDSql(db squirrel.StatementBuilderType, ID int32) squirrel.SelectBuilder {
	return db.Select("id", "name", "start_time", "end_time").From(tableName).Where(squirrel.Eq{"id": ID})
}

func FindByID(db squirrel.StatementBuilderType, ID int32) (*Info, error) {
	info := &Info{}
	err := findByIDSql(db, ID).QueryRow().Scan(
		&info.ID,
		&info.Name,
		&info.StartTime,
		&info.EndTime,
	)
	if err != nil {
		return nil, err
	}
	return info, nil
}
