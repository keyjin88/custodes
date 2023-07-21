package app

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"time"
)

type Menu interface {
	Show()
}

type MainMenu struct {
}

func NewMainMenu() *MainMenu {
	return &MainMenu{}
}

func (m MainMenu) Show() {
	err := ui.Init()
	if err != nil {
		fmt.Printf("Ошибка при инициализации termui: %v\n", err)
		return
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "Text Box"
	p.Text = "КОМПЛЕКС \"СТРАЖ\""
	p.SetRect(0, 0, 50, 5)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	barchartData := []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.SetRect(50, 5, 75, 15)
	bc.Labels = []string{"C0", "C1", "C2", "C3", "C4", "C5"}
	bc.BarColors[0] = ui.ColorGreen
	bc.NumStyles[0] = ui.NewStyle(ui.ColorBlack)

	draw := func(count int) {
		bc.Data = barchartData[count/2%10:]
		ui.Render(p, bc)
	}

	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			//updateParagraph(tickerCount)
			draw(tickerCount)
			tickerCount++
		}
	}
}
