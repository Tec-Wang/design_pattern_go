package prototype

import "testing"

func TestPrototype(t *testing.T) {
	// 文件结构
	// outerFoler-insideFoler, file1
	// 			- file 2
	//          - file 3
	file1 := &File{name: "File 1"}
	file2 := &File{name: "File 2"}
	file3 := &File{name: "File 3"}
	insideFolder := &Folder{name: "insideFolder", childern: []Inode{file2, file3}}
	outerFolder := &Folder{name: "outerFolder", childern: []Inode{file1, insideFolder}}

	// 打印外层文件夹
	outerFolder.Print("外层文件夹打印")
	// 克隆一个外层文件夹，并且打印
	outerFolder.Clone().Print("克隆结果")
}
