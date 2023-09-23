package builder

type Builder interface {
	reset()
	setWindowType()
	setDoorType()
	setNumFloor()
	// getResult 最终的产品类型不一定相关，所以可以将获取方法放到具体的builder中
}
