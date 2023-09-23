package prototype

// 文件和文件夹都属于Inode节点，同时文件夹可以包含文件夹
// 也就是Inode中有Inode，这就是组合-----向下-》向上收集
type Inode interface {
	cloneable
	Print(indentation string)
}
