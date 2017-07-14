package example

import (
	"github.com/wcharczuk/go-chart"
	"log"
	"net/http"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}
	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
