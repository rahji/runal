package main

import (
	"context"

	"github.com/emprcl/runal"
)

func main() {
	runal.Run(context.Background(), setup, draw, nil, nil)
}

func setup(c *runal.Canvas) {
	c.NoLoop()
}

func draw(c *runal.Canvas) {
}

// func onKey(c *runal.Canvas, e runal.KeyEvent) {
// }
// func onMouse(c *runal.Canvas, e runal.MouseEvent) {
// }
