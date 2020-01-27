package iotmaker_santa_isabel_theater_interfaces

type ISpriteDraw interface {
	// en: Get a unique element id
	//
	// pt_br: Retorna o id único do elemento
	GetId() string

	// en: Draw function to re draw element at every frame
	//
	// pt_br: Função de desenho para redesenhar o elemento a cada frame
	Draw()

	// en: Get an information about (x, y) is in element boxe
	//
	// pt_br: Retorna a informação se (x, y) está dentro da caxa do elemento
	GetCollisionBox(xEvent, yEvent float64) bool
}
