package data

import (
	"log"

	lib "github.com/antmyth/comix-lib/library"
	v "github.com/antmyth/comix-lib/viewmodel"
)

const fileLocation = "/Users/antonionascimento/dev/sandboxes/go-sandbox/comix-utils/output-lib/lib.json"

func New() *lib.ComicsLib {
	res, err := lib.New()
	if err != nil {
		log.Panic(err)
	}
	return res
}

func AsSeriesPointers(series []v.Series) []*v.Series {
	res := make([]*v.Series, len(series))
	for i := range series {
		res[i] = &series[i]
	}
	return res
}

func AsIssuesPointers(issues []v.Issue) []*v.Issue {
	res := make([]*v.Issue, len(issues))
	for i := range issues {
		res[i] = &issues[i]
	}
	return res
}

func AsPublishersPointers(publishers []v.Publisher) []*v.Publisher {
	res := make([]*v.Publisher, len(publishers))
	for i := range publishers {
		res[i] = &publishers[i]
	}
	return res
}
