package main

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "io.gitlab.andraantariksa.grid-shortest-path-algorithm-visualizer"

type BoxState int32

const (
	BOX_STATE_EMPTY BoxState = 0
	BOX_STATE_WALL  BoxState = 1
)

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNewFromFile("gtk_ui.glade")
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		// Get the object with the id of "main_window".
		winObj, err := builder.GetObject("application_window_main")
		errorCheck(err)

		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(winObj)
		errorCheck(err)

		// Show the Window and all of its components.
		win.Show()
		application.AddWindow(win)

		aboutDialogObj, _ := builder.GetObject("about_dialog")
		aboutDialog := aboutDialogObj.(*gtk.AboutDialog)

		modelButtonAboutObj, _ := builder.GetObject("model_button_about")
		modelButtonAbout := modelButtonAboutObj.(*gtk.ModelButton)

		modelButtonAbout.Connect("clicked", func() {
			aboutDialog.Show()
		})

		windowPreferencesObj, _ := builder.GetObject("window_preferences")
		windowPreferences := windowPreferencesObj.(*gtk.Window)

		modelButtonPreferencesObj, _ := builder.GetObject("model_button_preferences")
		modelButtonPreferences := modelButtonPreferencesObj.(*gtk.ModelButton)

		modelButtonPreferences.Connect("clicked", func() {
			windowPreferences.Show()
		})

		drawingAreaObj, err := builder.GetObject("drawing_area_grid")
		errorCheck(err)

		drawingArea, err := isDrawingArea(drawingAreaObj)
		errorCheck(err)

		boxSize := 25.0

		boxState := [15][15]BoxState{}
		for y := 0; y < 15; y++ {
			for x := 0; x < 15; x++ {
				boxState[y][x] = BOX_STATE_EMPTY
			}
		}

		drawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
			for y := 0; y < 15; y++ {
				for x := 0; x < 15; x++ {
					switch boxState[y][x] {
					case BOX_STATE_EMPTY:
						cr.SetSourceRGB(255, 255, 255)
					case BOX_STATE_WALL:
						cr.SetSourceRGB(0, 0, 0)
					}
					cr.Rectangle(boxSize*float64(x), boxSize*float64(y), boxSize, boxSize)
					cr.Fill()
				}
			}
		})

		drawingArea.Widget.AddEvents(int(gdk.BUTTON_PRESS_MASK))
		drawingArea.Connect("button-press-event", func(da *gtk.DrawingArea, event *gdk.Event) {
			eventMotion := gdk.EventMotionNewFromEvent(event)
			posX, posY := eventMotion.MotionVal()
			boxPosX := int(posX / boxSize)
			boxPosY := int(posY / boxSize)
			boxState[boxPosY][boxPosX] = BOX_STATE_WALL

			eventButton := gdk.EventButtonNewFromEvent(event)
			buttonPressed := eventButton.ButtonVal()

			switch buttonPressed {
			case 1:
				boxState[boxPosY][boxPosX] = BOX_STATE_WALL
			case 3:
				boxState[boxPosY][boxPosX] = BOX_STATE_EMPTY
			}

			win.QueueDraw()
		})
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application closed")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func isWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func isDrawingArea(obj glib.IObject) (*gtk.DrawingArea, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.DrawingArea); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.DrawingArea")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
