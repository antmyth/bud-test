package issues

import (
	context "context"

	"github.com/antmyth/comix-lib/library"
	v "github.com/antmyth/comix-lib/viewmodel"
)

type Controller struct {
	// Dependencies...
	Comics *library.ComicsLib
	Img    *v.Image
}

// Index of issues
// GET /issues
func (c *Controller) Index(ctx context.Context) (issues []*v.Issue, err error) {
	return make([]*v.Issue, 0), nil
}

// Show issue
// GET /issues/:id
func (c *Controller) Show(ctx context.Context, id int) (issue *v.Issue, err error) {
	return c.Comics.GetIssueByID(id), nil
}
