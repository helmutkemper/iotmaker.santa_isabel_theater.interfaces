package iotmaker_santa_isabel_theater_interfaces

type ISpriteDraw interface {
	GetId() string
	Draw()
	GetCollisionBox(xEvent, yEvent float64) bool
}
