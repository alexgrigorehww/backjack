package main

import (
	"log"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	mw := new(MyMainWindow)

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "BlackJack",
		MinSize:  Size{320, 240},
		Size:     Size{300, 300},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
}

func (mw *MyMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bounds := mw.paintWidget.ClientBounds()

	rectPen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(255, 0, 0))
	if err != nil {
		return err
	}
	defer rectPen.Dispose()

	if err := canvas.DrawRectangle(rectPen, bounds); err != nil {
		return err
	}

	ellipseBrush, err := walk.NewHatchBrush(walk.RGB(0, 255, 0), walk.HatchCross)
	if err != nil {
		return err
	}
	defer ellipseBrush.Dispose()

	if err := canvas.FillEllipse(ellipseBrush, bounds); err != nil {
		return err
	}

	linesBrush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 255))
	if err != nil {
		return err
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenDash, 8, linesBrush)
	if err != nil {
		return err
	}
	defer linesPen.Dispose()

	return nil
}
