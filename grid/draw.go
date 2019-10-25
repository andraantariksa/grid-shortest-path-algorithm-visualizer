package grid

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (grd *Grid) Draw(da *gtk.DrawingArea, cr *cairo.Context) {
	for y := uint(0); y < grd.Size; y++ {
		for x := uint(0); x < grd.Size; x++ {
			switch grd.BoxesState[y][x] {
			case BOX_STATE_EMPTY:
				cr.SetSourceRGB(1.0, 1.0, 1.0)
			case BOX_STATE_START:
				cr.SetSourceRGB(0, 0.84, 1.0)
			case BOX_STATE_END:
				cr.SetSourceRGB(0.21, 0.28, 0.7)
			case BOX_STATE_WALL:
				cr.SetSourceRGB(0.0, 0.0, 0.0)
			}
			cr.Rectangle(grd.BoxSize*float64(x), grd.BoxSize*float64(y), grd.BoxSize, grd.BoxSize)
			cr.Fill()
		}
	}
}

func (grd *Grid) BoxPressed(da *gtk.DrawingArea, event *gdk.Event) {
	eventMotion := gdk.EventMotionNewFromEvent(event)
	posX, posY := eventMotion.MotionVal()
	boxPosX := int(posX / grd.BoxSize)
	boxPosY := int(posY / grd.BoxSize)

	eventButton := gdk.EventButtonNewFromEvent(event)
	buttonPressed := eventButton.ButtonVal()

	switch buttonPressed {
	// Left click
	case 1:
		grd.BoxesState[boxPosY][boxPosX] = BOX_STATE_WALL
		// Right click
	case 3:
		grd.BoxesState[boxPosY][boxPosX] = BOX_STATE_EMPTY
	}

	grd.Window.QueueDraw()
}
