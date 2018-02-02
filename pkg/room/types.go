package room

import "github.com/Masterminds/squirrel"

// Info models the room's basic information
type Info struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

const tableName = "rooms"

func findByIDSql(db squirrel.StatementBuilderType, ID int32) squirrel.SelectBuilder {
	return db.Select(
		"id",
		"name",
		"location",
		"created_at",
		"updated_at",
	).From(tableName).Where(squirrel.Eq{"id": ID})
}

func FindByID(db squirrel.StatementBuilderType, ID int32) (*Info, error) {
	info := &Info{}
	err := findByIDSql(db, ID).QueryRow().Scan(
		&info.ID,
		&info.Name,
		&info.Location,
		&info.CreatedAt,
		&info.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return info, nil
}
