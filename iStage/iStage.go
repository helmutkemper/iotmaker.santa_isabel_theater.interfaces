package iStage

import (
	iotmaker_santa_isabel_theater_interfaces "github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

type IStage interface {
	AddEngine(engine engine.IEngine)
	GetEngine() engine.IEngine
	AddToIndexList(id string, index int, collision func(x, y float64) bool)
	RemoveFromIndexList(id string)
	IsDraggable(x, y float64) string
	AddToDraw(sprite iotmaker_santa_isabel_theater_interfaces.ISpriteDraw) (string, int)
	RemoveFromDraw(id string)
	AddToCalc(f func()) (string, int)
	RemoveFromCalc(id string)
	AddToHighLatency(f func()) (string, int)
	RemoveFromHighLatency(id string)
	AddToSystem(f func()) (string, int)
	RemoveFromSystem(id string)
	CursorHide()
	CursorShow()
	SetCursorDrawFunc(function func())
	SetCursorStageId(id string)
	SetWidth(width float64)
	SetHeight(height float64)
	Clear()
}
