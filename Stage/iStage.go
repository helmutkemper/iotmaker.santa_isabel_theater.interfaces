package Stage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

type IStage interface {
	AddEngine(engine engine.IEngine)
	GetEngine() engine.IEngine
	AddToIndexList(id string, index int, dimensions genericTypes.Dimensions)
	RemoveFromIndexList(id string) error
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
