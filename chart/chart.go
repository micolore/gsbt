package chart

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func CreateChat() {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	f, _ := os.Create("/Users/kubrick/go/src/moppo.com/moppo-webhook-server/bar.html")
	fmt.Println("create bat.html ok!")
	bar.Render(f)
}

func CreateBar(title, desc string, category []string, data []opts.BarData,url string) {
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    title,
		Subtitle: desc,
	}))

	bar.SetXAxis(category).
		AddSeries("Category", data)
	f, _ := os.Create(url)
	fmt.Println("create bat.html ok!")
	bar.Render(f)
}
