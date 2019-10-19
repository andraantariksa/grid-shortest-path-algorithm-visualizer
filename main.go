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
		window_obj, err := builder.GetObject("main_window")
		errorCheck(err)

		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(window_obj)
		errorCheck(err)

		// Show the Window and all of its components.
		win.Show()
		application.AddWindow(win)

		/////////////////////////////////

		drawing_area_obj, err := builder.GetObject("graph_drawing_area")
		errorCheck(err)

		drawing_area, err := isDrawingArea(drawing_area_obj)
		errorCheck(err)

		drawing_area.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
			for y := 0.0; y < 15.0; y += 1.0 {
				for x := 0.0; x < 15.0; x += 1.0 {
					cr.SetSourceRGB(94, 184, 255)
					cr.Rectangle(25.0*x, 25.0*y, 25.0, 25.0)
					cr.Fill()
				}
			}
		})

		drawing_area.Widget.AddEvents(int(gdk.BUTTON_PRESS_MASK))
		drawing_area.Connect("button-press-event", func(da *gtk.DrawingArea, ev *gdk.Event) {
			log.Println("bruh")
		})

		// eventBoxObj, err := builder.GetObject("event_box")
		// eventBox := eventBoxObj.(*gtk.EventBox)

		// eventBox.Widget.AddEvents(int(gdk.BUTTON_PRESS_MASK))
		// eventBox.Connect("button-press-event", func(bo *gtk.EventBox, ev *gdk.Event) {
		// 	log.Println(bo.Show()
		// })
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
