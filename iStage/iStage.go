package iStage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

type IStage interface {
	AddEngine(engine engine.IEngine)
	GetEngine() engine.IEngine
	AddToIndexList(id string, index int, dimensions dimensions.Dimensions)
	RemoveFromIndexList(id string)
	AddToDraw(f func()) (string, int)
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
