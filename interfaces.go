package sphinx_client

import (
	"context"
)

type Sphinx interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
