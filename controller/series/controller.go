package series

import (
	context "context"
	"log"

	data "github.com/antmyth/buddy/internal/data"
)

type Controller struct {
	// Dependencies...
	Comics *data.ComicsLib
}

// Index of series
// GET /series/
func (c *Controller) Index(ctx context.Context, reload *bool) (series []*data.Series, err error) {
	if reload != nil && *reload {
		log.Printf("Should reload data?%v\n", true)
		c.Comics = c.Comics.Reload()
	}
	return c.Comics.SeriesPointers(), nil
}

// Show series
// GET /series/:id
func (c *Controller) Show(ctx context.Context, id int) (series *data.Series, err error) {
	res := c.Comics.SeriesList[id]
	return &res, nil
}
