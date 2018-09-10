package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"log"
	"os"

	"github.com/balacode/one-file-pdf"

	"github.com/wcharczuk/go-chart"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	fmt.Println(`Generating a "Hello World" PDF...`)

	// create a new PDF using 'A4' page size
	var pdf = pdf.NewPDF("A4")

	// set the measurement units to centimeters
	pdf.SetUnits("cm")

	// draw a grid to help us align stuff (just a guide, not necessary)
	pdf.DrawUnitGrid()

	// draw the word 'HELLO' in orange, using 100pt bold Helvetica font
	// - text is placed on top of, not below the Y-coordinate
	// - you can use method chaining
	pdf.SetFont("Helvetica-Bold", 100).
		SetXY(5, 5).
		SetColor("Orange").
		DrawText("HELLO")

	// draw the word 'WORLD' in blue-violet, using 100pt Helvetica font
	// note that here we use the colo(u)r hex code instead
	// of its name, using the CSS/HTML format: #RRGGBB
	pdf.SetXY(5, 9).
		SetColor("#8A2BE2").
		SetFont("Helvetica", 100).
		DrawText("WORLD!")

	// draw a flower icon using 300pt Zapf-Dingbats font
	pdf.SetX(7).SetY(17).
		SetColorRGB(255, 0, 0).
		SetFont("ZapfDingbats", 300).
		DrawText("a")

	pdf.AddPage()

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	//buffer := bytes.NewBuffer([]byte{})

	f, err := os.Create("test_reporter.png")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	//defer f.Close()

	//graph.Render(chart.PNG, buffer)
	f_err := graph.Render(chart.PNG, w)

	check(f_err)
	w.Flush()

	/*
	   p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	       {0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
	   })
	   p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	       {0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
	   })

	   pts := XYs{{0, 0}, {0, 1}, {0.5, 1}, {0.5, 0.6}, {0, 0.6}}
	   line, err := NewLine(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   scatter, err := NewScatter(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   p.Add(line, scatter)

	   pts = XYs{{1, 0}, {0.75, 0}, {0.75, 0.75}}
	   line, err = NewLine(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   scatter, err = NewScatter(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   p.Add(line, scatter)

	   pts = XYs{{0.5, 0.5}, {1, 0.5}}
	   line, err = NewLine(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   scatter, err = NewScatter(pts)
	   if err != nil {
	       log.Panic(err)
	   }
	   p.Add(line, scatter)

	   err = p.Save(100, 100, "plotLogo.png")
	   if err != nil {
	       log.Panic(err)
	   }
	*/

	pdf.DrawImage(3, 18, 6, "test_reporter.png")

	// save the file:
	// if the file exists, it will be overwritten
	// if the file is in use, prints an error message
	pdf.SaveFile("hello.pdf")
} // main
