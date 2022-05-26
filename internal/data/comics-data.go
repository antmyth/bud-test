package data

import (
	"encoding/json"
	"log"
	"sort"

	commonfunctions "github.com/antmyth/buddy/internal/common-functions"
)

const fileLocation = "/Users/antonionascimento/dev/sandboxes/go-sandbox/comix-utils/output-lib/lib.json"

func New() *ComicsLib {
	res := ComicsLib{}
	return res.Reload()
}

func (lib ComicsLib) Reload() *ComicsLib {
	log.Printf("Reloading ComicsDB from file?%s\n", fileLocation)
	LoadLibraryData(&lib)
	series := lib.SeriesList
	sort.SliceStable(series, func(i, j int) bool {
		if series[i].Series < series[j].Series {
			return true
		} else if series[i].Series > series[j].Series {
			return false
		} else {
			return series[i].Volume < series[j].Volume
		}
	})
	return &lib
}

type ComicsLib struct {
	SeriesCount int      `json:"seriesCount"`
	SeriesList  []Series `json:"seriesList"`
}

type Series struct {
	Series      string  `json:"series"`
	Volume      string  `json:"volume,omitempty"`
	Publisher   string  `json:"publisher,omitempty"`
	Count       int     `json:"count"`
	TotalCount  int     `json:"totalcount"`
	ID          string  `json:"id,omitempty"`
	Web         string  `json:"web,omitempty"`
	Images      Image   `json:"image,omitempty"`
	Issues      []Issue `json:"issues"`
	Location    string  `json:"location,omitempty"`
	Description string  `json:"description,omitempty"`
}

type Issue struct {
	ID             int    `json:"id,omitempty"`
	Title          string `json:"title"`
	Series         string `json:"series"`
	Number         string `json:"number"`
	Volume         string `json:"volume,omitempty"`
	Publisher      string `json:"publisher,omitempty"`
	Web            string `json:"web,omitempty"`
	VolumeAPI      string `json:"volume_api,omitempty"`
	Images         Image  `json:"image,omitempty"`
	SeriesLocation string `json:"seriesLocation,omitempty"`
	Location       string `json:"location,omitempty"`
}

type Image struct {
	SmallUrl    string `json:"small_url,omitempty"`
	ThumbUrl    string `json:"thumb_url,omitempty"`
	TinyUrl     string `json:"tiny_url,omitempty"`
	OriginalUrl string `json:"original_url,omitempty"`
}

func toPointers(series []Series) []*Series {
	res := make([]*Series, len(series))
	for i := range series {
		res[i] = &series[i]
	}
	return res
}

func toPointers2(issues []Issue) []*Issue {
	res := make([]*Issue, len(issues))
	for i := range issues {
		res[i] = &issues[i]
	}
	return res
}

func LoadLibraryData(lib *ComicsLib) {
	strLib := commonfunctions.ReadFile(fileLocation)
	json.Unmarshal([]byte(strLib), &lib)
}

func (lib ComicsLib) IssuesPointers() []*Issue {
	il := make([]Issue, 0)
	for _, s := range lib.SeriesList {
		for _, i := range s.Issues {
			il = append(il, i)
		}
	}
	return toPointers2(il)
}

func (lib ComicsLib) SeriesPointers() []*Series {
	return toPointers(lib.SeriesList)
}
