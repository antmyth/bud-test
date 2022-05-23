package controller

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Story struct
type Story struct {
	// Fields...
	Head  string
	Tails []string
}

// Index of stories
// GET /
func (c *Controller) Index(ctx context.Context) (stories []*Story, err error) {
	var sss []*Story
	s := Story{
		Head:  "head",
		Tails: []string{"t", "tt"},
	}
	sss = append(sss, &s)
	return sss, nil
}

// Show story
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (story *Story, err error) {

	return &Story{}, nil
}
