package main

import (
	"context"
	"fmt"
	"image/color"
	"log"
	"sync"
	"time"
	tm "times/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type timeEntry struct {
	widget.Entry
}

func (n *timeEntry) Keyboard() mobile.KeyboardType {
	return mobile.SingleLineKeyboard
}

func newTimeEntry() *timeEntry {
	e := &timeEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`2\d{3}-[012]?\d-([012]\d|3[01]) [012]?\d:[012345]?\d`, "Must contain a date time")
	return e
}

type Applicazione struct {
	gui       fyne.App
	win       fyne.Window
	cont      *fyne.Container
	ctx       context.Context
	ctxCancel context.CancelFunc
	wg        *sync.WaitGroup

	converting bool

	entries map[string]*timeEntry
	tz      map[string]*time.Location
	dist    map[string]*widget.Label
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

func (a *Applicazione) toTime(tz, dateString string) time.Time {
	layout := "2006-01-02 15:04"

	location := a.tz[tz]
	parsedTime, _ := time.ParseInLocation(layout, dateString, location)

	return parsedTime
}

func (a *Applicazione) setTimes(from string, value string) {
	if a.converting {
		return
	}

	a.converting = true
	defer func() { a.converting = false }()

	labelText := func(name string, t time.Time, vin bool) string {
		_, zs := t.Zone()

		m := "+"
		if zs < 0 {
			m = "-"
			zs = -zs
		}
		oh := zs / 3600
		om := (zs / 60) % 60

		mark := ""
		if vin {
			mark = " *"
		}

		return fmt.Sprintf("%s: %s%02d:%02d%s", name, m, oh, om, mark)
	}

	c := a.entries[from]

	if c.Validate() != nil {
		return
	}

	fmt.Printf("Valid %s: %s\n", from, value)

	t0 := a.toTime(from, value)

	for k, v := range a.entries {
		if k == from {
			a.dist[k].SetText(labelText(k, t0, true))
			fyne.Do(a.dist[k].Refresh)
			continue
		}

		vv := t0.In(a.tz[k])
		tv := vv.Format("2006-01-02 15:04")
		v.SetText(tv)
		a.dist[k].SetText(labelText(k, vv, false))

		fyne.Do(func() {
			v.Refresh()
			a.dist[k].Refresh()
		})
	}

}

func (a *Applicazione) Setup() {

	clock := canvas.NewText("---", color.RGBA{R: 10, G: 20, B: 255, A: 255})
	clock.TextSize = 64
	clock.Alignment = fyne.TextAlignCenter
	a.cont.Add(clock)

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()

		showTime := func(t time.Time) {
			var f string
			if t.Second()%2 == 0 {
				f = "15:04"
			} else {
				f = "15 04"
			}
			clock.Text = t.Format(f)
			fyne.Do(clock.Refresh)
		}

		showTime(time.Now())
		for {
			select {
			case <-a.ctx.Done():
				return
			case now := <-time.After(500 * time.Millisecond):
				showTime(now)
			}
		}
	}()

	newEntry := func(placeholder string) (*timeEntry, *widget.Label) {
		l := widget.NewLabel(placeholder)
		a.cont.Add(l)
		c := newTimeEntry()
		c.SetPlaceHolder("Insert time YYYY-MM-DD hh:mm")
		// add it now
		a.cont.Add(c)
		c.OnChanged = func(s string) {
			a.setTimes(placeholder, s)
		}
		return c, l
	}

	//var err error

	a.tz = make(map[string]*time.Location)
	a.tz["UTC"], _ = time.LoadLocation("UTC")
	a.tz["Italia"], _ = time.LoadLocation("Europe/Rome")
	a.tz["India"], _ = time.LoadLocation("Asia/Kolkata")
	a.tz["USA"], _ = time.LoadLocation("America/New_York")
	a.tz["Australia"], _ = time.LoadLocation("Australia/Adelaide")

	a.entries = make(map[string]*timeEntry)
	a.dist = make(map[string]*widget.Label)
	a.entries["UTC"], a.dist["UTC"] = newEntry("UTC")
	a.entries["Italia"], a.dist["Italia"] = newEntry("Italia")
	a.entries["India"], a.dist["India"] = newEntry("India")
	a.entries["USA"], a.dist["USA"] = newEntry("USA")
	a.entries["Australia"], a.dist["Australia"] = newEntry("Australia")
}

func Init(id string, title string, width float32, height float32) *Applicazione {
	var a Applicazione

	a.ctx, a.ctxCancel = context.WithCancel(context.Background())
	a.wg = &sync.WaitGroup{}
	a.gui = app.NewWithID(id)
	a.logLifecycle()

	a.win = a.gui.NewWindow(title)
	a.win.SetMaster()
	a.cont = container.NewVBox()
	a.win.SetContent(a.cont)
	a.win.Resize(fyne.NewSize(width, height))

	fyne.CurrentApp().Settings().SetTheme(tm.PanelTheme())

	return &a
}

func main() {
	MainApp := Init("org.pasckosky.times", "Time Toolbox", 600, 800)
	MainApp.Setup()

	MainApp.win.ShowAndRun()
	MainApp.ctxCancel()

	// wait for end
	MainApp.wg.Wait()
}
