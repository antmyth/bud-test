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
func (c *Controller) Index(ctx context.Context, reload *bool) (series []*v.Series, err error) {
	ss := c.Comics.GetAllSeries()
	out := data.AsSeriesPointers(ss)
	return out, nil
}

// Show series
// GET /series/:id
func (c *Controller) Show(ctx context.Context, id int) (series *v.Series, err error) {
	res := c.Comics.GetSeriesByIDWithIssues(id)
	return res, nil
}
