// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: test.sql

package gen

import (
	"context"
)

const selectTest = `-- name: SelectTest :exec
SELECT 1
`

func (q *Queries) SelectTest(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, selectTest)
	return err
}