// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tags.sql

package gen

import (
	"context"

	"github.com/lib/pq"
)

const selectTags = `-- name: SelectTags :many
SELECT id, project_id, value
FROM key_tags
WHERE project_id = $1
  AND LOWER(value) = ANY ($2::varchar[])
`

type SelectTagsParams struct {
	ProjectID int64    `db:"project_id"`
	Column2   []string `db:"column_2"`
}

func (q *Queries) SelectTags(ctx context.Context, arg SelectTagsParams) ([]KeyTag, error) {
	rows, err := q.db.QueryContext(ctx, selectTags, arg.ProjectID, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeyTag
	for rows.Next() {
		var i KeyTag
		if err := rows.Scan(&i.ID, &i.ProjectID, &i.Value); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
