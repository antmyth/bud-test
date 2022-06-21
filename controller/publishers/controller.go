package publishers

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

// Index of publishers
// GET /publishers
func (c *Controller) Index(ctx context.Context) (publishers []*v.Publisher, err error) {
	ss := c.Comics.GetAllPublishers()
	out := data.AsPublishersPointers(ss)
	return out, nil
}

// Show publisher
// GET /publishers/:id
func (c *Controller) Show(ctx context.Context, id int) (publisher *v.Publisher, err error) {
	out := c.Comics.GetPublisherByID(id)
	out.Series = c.Comics.GetSeriesByPublisher(*out)
	return out, nil
}
