package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTreeLocal(t *testing.T){
	addNode("test",1,-1,"女装")
	addNode("test",2,1,"裙装")
	addNode("test",3,1,"衬衫")
	addNode("test",4,2,"连衣裙")
	addNode("test",5,2,"T百褶裙")
	addNode("test",6,3,"T恤")
	addNode("test",7,3,"短袖")
	addNode("test",8,7,"夏威夷短袖")

	//获取完整菜单树
	node, _ := getDirectory("test","")
	bytes, _ := json.Marshal(node)
	fmt.Printf("%s\n", bytes)
	//模糊查询 裙
	node, _ = getDirectory("test","裙")
	bytes, _ = json.Marshal(node)
	fmt.Printf("%s\n", bytes)
	//模糊查询 T
	node, _ = getDirectory("test","T")
	bytes, _ = json.Marshal(node)
	fmt.Printf("%s\n", bytes)
	//短袖移动到裙装下面
	moveNode("test",7,2)
	node, _ = getDirectory("test","")
	bytes, _ = json.Marshal(node)
	fmt.Printf("%s\n", bytes)
	//删除短袖
	deleteNode("test",7)
	node, _ = getDirectory("test","")
	bytes, _ = json.Marshal(node)
	fmt.Printf("%s\n", bytes)

}
