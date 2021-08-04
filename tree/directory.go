package tree

import (
	"encoding/json"
	"strings"
)

type treeDbStore struct{
	id int
	tree *treeNode
	nameMap map[int]*treeNode
}

type treeNode struct {
	Id int                 `json:"id"`
	Name       string      `json:"name"`
	ParentId int   `json:"parentId"`
	SubTrees   []*treeNode `json:"subTrees"`
}

//假设是线程安全的
var cache=map[string]*treeDbStore{}

func getDBStore(owner string) *treeDbStore {
	if cache["test"]==nil{
		//todo init from db
		rootNode := new(treeNode)
		rootNode.Id=-1
		rootNode.Name="root"
		rootNode.SubTrees =make([]*treeNode,0)

		tepDB := new(treeDbStore)
		tepDB.nameMap=make(map[int]*treeNode,1)
		tepDB.nameMap[-1]=rootNode
		tepDB.tree=rootNode
		cache["test"]=tepDB
	}
	return cache["test"]
}
func getDirectory(owner string,queryName string) (*treeNode,error) {
	dbStore := getDBStore(owner)
	if len(queryName)>0 {
		var copyOfTree = new(treeNode)
		b, err := json.Marshal(dbStore.tree)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, copyOfTree)
		fuzzyQuery(copyOfTree,queryName)
		return copyOfTree,nil
	}
	return dbStore.tree,nil
}
func fuzzyQuery(root *treeNode,queryName string) bool{
	if root == nil{
		return false
	}
	hasQueryName := false
	subHasQueryName :=false
	//todo 替换成kmp？
	if strings.Contains(root.Name,queryName) {
		hasQueryName=true
	}
	for index := range root.SubTrees {
		tmpHasQueryName:=fuzzyQuery(root.SubTrees[index],queryName)
		subHasQueryName= subHasQueryName || tmpHasQueryName
		if !tmpHasQueryName{
			//cut subtree
			root.SubTrees[index]=nil
		}
	}
	return hasQueryName||subHasQueryName
}
func renameNode(owner string,nodeId int,name string)  {
	dbStore := getDBStore(owner)
	dbStore.nameMap[nodeId].Name = name
	//todo store in db
}
func moveNode(owner string,nodeId int,toParentId int){
	dbStore := getDBStore(owner)

	node := dbStore.nameMap[nodeId]
	//断掉原有指针
	originParentNode := dbStore.nameMap[node.ParentId]
	for index, tree := range originParentNode.SubTrees {
		if tree == node{
			originParentNode.SubTrees = append(originParentNode.SubTrees[:index], originParentNode.SubTrees[index+1:]...)
			break
		}
	}
	// 连接上
	node.ParentId= toParentId
	dbStore.nameMap[toParentId].SubTrees = append(dbStore.nameMap[toParentId].SubTrees, node)
}
func deleteNode(owner string,nodeId int){
	//todo recursive delete
	dbStore := getDBStore(owner)
	node := dbStore.nameMap[nodeId]
	originParentNode := dbStore.nameMap[node.ParentId]
	for index, tree := range originParentNode.SubTrees {
		if tree == node{
			originParentNode.SubTrees = append(originParentNode.SubTrees[:index], originParentNode.SubTrees[index+1:]...)
			break
		}
	}
	dbStore.nameMap[nodeId]=nil
}
func addNode(owner string,nodeId int,parentId int,name string)*treeNode {
	dbStore := getDBStore(owner)
	parentNode := dbStore.nameMap[parentId]

	tempNode := new(treeNode)
	tempNode.Id=nodeId
	tempNode.Name=name
	tempNode.ParentId=parentId
	tempNode.SubTrees =make([]*treeNode,0)
	parentNode.SubTrees =append(parentNode.SubTrees, tempNode)

	dbStore.nameMap[nodeId]=tempNode
	return tempNode
}
func addData(){

}
func deleteData(){

}
func moveData(){

}