package series

import (
	context "context"

	data "github.com/antmyth/buddy/internal/data"
)

type Controller struct {
	// Dependencies...
	Comics *data.ComicsLib
}

// Index of series
// GET /series/
func (c *Controller) Index(ctx context.Context) (series []*data.Series, err error) {
	return c.Comics.SeriesPointers(), nil
}

// Show series
// GET /series/:id
func (c *Controller) Show(ctx context.Context, id int) (series *data.Series, err error) {
	res := c.Comics.SeriesList[id]
	return &res, nil
}

// func LoadSeriesFromLibrary() []*Series {
// 	strLib := ReadFile("/Users/antonionascimento/dev/sandboxes/go-sandbox/comix-utils/output-lib/lib.json")
// 	json.Unmarshal([]byte(strLib), &lib)
// 	res := make([]*Series, len(lib.SeriesList))
// 	for i := range lib.SeriesList {
// 		res[i] = &lib.SeriesList[i]
// 	}
// 	return res
// }

// // -----
// type ComicsLib struct {
// 	SeriesCount int      `json:"seriesCount"`
// 	SeriesList  []Series `json:"seriesList"`
// }

// type Series struct {
// 	Series    string  `json:"series"`
// 	Volume    string  `json:"volume,omitempty"`
// 	Publisher string  `json:"publisher,omitempty"`
// 	Count     int     `json:"count"`
// 	Issues    []Issue `json:"issues"`
// }

// type Issue struct {
// 	Title     string `json:"title"`
// 	Series    string `json:"series"`
// 	Number    string `json:"number"`
// 	Volume    string `json:"volume,omitempty"`
// 	Publisher string `json:"publisher,omitempty"`
// }

// func (issue Issue) ToString() string {
// 	return fmt.Sprintf("%s|%s", issue.Series, issue.Volume)
// }

// func AsSeriesMap(issues []Issue) map[string][]Issue {
// 	res := make(map[string][]Issue, 0)
// 	for _, v := range issues {
// 		ss, e := res[v.ToString()]
// 		if e {
// 			ss = append(ss, v)
// 			res[v.ToString()] = ss
// 		} else {
// 			res[v.ToString()] = []Issue{v}
// 		}
// 	}
// 	return res
// }

// //ReadFile reads the file into a string
// func ReadFile(filename string) string {
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(-1)
// 	}
// 	return string(bs)
// }
