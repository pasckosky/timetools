package main

import (
	"image/color"
	"log"
	"time"
	tm "times/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Applicazione struct {
	gui  fyne.App
	win  fyne.Window
	cont *fyne.Container
}

func (a *Applicazione) logLifecycle() {
	a.gui.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.gui.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.gui.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.gui.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func (a *Applicazione) Add(obj fyne.CanvasObject) {
	a.cont.Add(obj)
}

func (a *Applicazione) Setup() {

	clock := canvas.NewText("---", color.RGBA{R: 10, G: 20, B: 255, A: 255})
	clock.TextSize = 64
	clock.Alignment = fyne.TextAlignCenter
	a.cont.Add(clock)

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		var now time.Time

		now = time.Now()
		for {
			time_str := func() string {
				var f string
				if now.Second()%2 == 0 {
					f = "15:04"
				} else {
					f = "15 04"
				}
				return now.Format(f)
			}

			clock.Text = time_str()

			fyne.Do(clock.Refresh)

			now = <-ticker.C
		}
	}()
}

func Init(id string, title string, width float32, height float32) Applicazione {
	var this Applicazione

	this.gui = app.NewWithID("org.pasckosky.times")
	this.logLifecycle()

	this.win = this.gui.NewWindow("Container")
	this.win.SetMaster()
	this.cont = container.NewVBox()
	this.win.SetContent(this.cont)
	this.win.Resize(fyne.NewSize(width, height))

	fyne.CurrentApp().Settings().SetTheme(tm.PanelTheme())

	return this
}

var MainApp Applicazione

func main() {
	//fmt.Printf("Startup at %s\n", time.Now().Local().Format(time.RFC3339))

	MainApp = Init("org.pasckosky.times", "Container", 600, 1024)
	//MainApp.setClock()
	MainApp.Setup()
	//astrotime.CalcSunrise(time.Now(), 0, 0)

	MainApp.win.ShowAndRun()
}
