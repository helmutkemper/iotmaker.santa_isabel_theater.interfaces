package iStage

import (
	iotmaker_santa_isabel_theater_interfaces "github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iEngine"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
)

type IStage interface {
	AddEngine(engine iEngine.IEngine)
	GetEngine() iEngine.IEngine
	GetStage() IStage
	GetCanvas() iotmaker_platform_IDraw.IDraw
	GetScratchPad() iotmaker_platform_IDraw.IDraw
	GetCache() iotmaker_platform_IDraw.IDraw
	AddToIndexList(id string, index int, collision func(x, y int) bool)
	RemoveFromIndexList(id string)
	IsDraggable(x, y int) string
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
	SetWidth(width int)
	SetHeight(height int)
	Clear()
}
