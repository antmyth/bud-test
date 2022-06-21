package series

import (
	context "context"

	"github.com/antmyth/buddy/internal/data"
	"github.com/antmyth/comix-lib/library"
	v "github.com/antmyth/comix-lib/viewmodel"
)

type Controller struct {
	// Dependencies...
	Comics *library.ComicsLib
	Img    *v.Image
}

// Index of series
// GET /series/
func (c *Controller) Index(ctx context.Context, page *int, size *int) (series []*v.Series, err error) {
	pg := 0
	if page != nil {
		pg = *page
	}
	sz := 100
	if size != nil {
		sz = *size
	}
	ss := c.Comics.GetAllSeriesPaginated(pg, sz)
	out := data.AsSeriesPointers(ss)
	return out, nil
}

// Show series
// GET /series/:id
func (c *Controller) Show(ctx context.Context, id int) (series *v.Series, err error) {
	res := c.Comics.GetSeriesByIDWithIssues(id)
	return res, nil
}
