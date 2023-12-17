package sphinx_client

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

func (d *sphinx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.SelectContext(ctx, dest, query, args...)
}
