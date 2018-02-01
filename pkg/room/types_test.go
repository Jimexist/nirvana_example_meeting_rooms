package room

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func Test_findByIDSql(t *testing.T) {
	type args struct {
		db squirrel.StatementBuilderType
		ID int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "id=42",
			args: args{
				ID: 42,
				db: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
			},
			want: `SELECT id, name, location, created_at, updated_at FROM rooms WHERE id = $1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, args, err := findByIDSql(tt.args.db, tt.args.ID).ToSql()
			assert.Nil(t, err)
			assert.Equal(t, []interface{}{
				interface{}(int32(42)),
			}, args)
			assert.Equal(t, got, tt.want)
		})
	}
}
