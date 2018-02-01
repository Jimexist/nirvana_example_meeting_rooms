package room

import "github.com/Masterminds/squirrel"

// Info models the room's basic information
type Info struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

const tableName = "rooms"

func FindByID(db squirrel.StatementBuilderType, ID int32) (*Info, error) {
	info := &Info{}
	err := db.Select(
		"id", "name", "location",
	).From(tableName,
	).Where(squirrel.Eq{"id": ID}).QueryRow().Scan(
		&info.ID,
		&info.Name,
		&info.Location,
	)
	if err != nil {
		return nil, err
	}
	return info, nil
}
