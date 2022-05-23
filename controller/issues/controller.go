package issues

import (
	context "context"

	"github.com/antmyth/buddy/internal/data"
)

type Controller struct {
	// Dependencies...
	Comics *data.ComicsLib
}

// Index of issues
// GET /issues
func (c *Controller) Index(ctx context.Context) (issues []*data.Issue, err error) {
	return c.Comics.IssuesPointers(), nil
}

// Show issue
// GET /issues/:id
func (c *Controller) Show(ctx context.Context, id int) (issue *data.Issue, err error) {
	return c.Comics.IssuesPointers()[id], nil
}
