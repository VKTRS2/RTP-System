package model

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type TreeVisitorNode struct {
	Name     string
	Type     string
	IsStruct bool
	IsPtr    bool
	IsArray  bool
	Children []*TreeVisitorNode
}

func (si TreeVisitorNode) String() string {
	modifier := ""
	if si.IsPtr {
		modifier = "*"
	}

	if si.IsArray {
		modifier = "[]"
	}

	path := fmt.Sprintf("%s %s%s", si.Name, modifier, si.Type)
	return path
}

type TreeVisitor struct {
	Root *TreeVisitorNode
	Path []*TreeVisitorNode
}

func NewTreeVisitor(name, aType string) GoModelVisitor {

	rootNode := &TreeVisitorNode{
		Name: name,
		Type: aType,
	}

	v := &TreeVisitor{
		Root: rootNode,
	}

	v.Path = append(v.Path, rootNode)
	return v

}

func (tv TreeVisitor) currentPathToString() string {

	var sb strings.Builder
	for _, item := range tv.Path {
		sb.WriteString("/")
		sb.WriteString(item.String())
	}

	return sb.String()
}

func (tv *TreeVisitor) Visit(name, aType string, isStruct, isPrt, isArray bool) (string, error) {
	item := &TreeVisitorNode{Name: name, Type: aType, IsStruct: isStruct, IsPtr: isPrt, IsArray: isArray}

	if tv.Root == nil {
		tv.Root = item
		tv.Path = append(tv.Path, item)
	} else {
		currentPathIndex := len(tv.Path) - 1
		tv.Path[currentPathIndex].Children = append(tv.Path[currentPathIndex].Children, item)
		tv.Path = append(tv.Path, item)
	}

	return tv.currentPathToString(), nil
}

func (tv *TreeVisitor) Pop() (string, error) {

	if len(tv.Path) > 0 {
		tv.Path = tv.Path[0 : len(tv.Path)-1]
	} else {
		return "", fmt.Errorf("cannot pop from path: %s", tv.currentPathToString())
	}

	return tv.currentPathToString(), nil
}

func (tv *TreeVisitor) ShowInfo(currentPackage string) error {

	if tv.Root == nil {
		log.Info().Msg("the tree is empty")
		return nil
	}

	return ShowNodeInfo(currentPackage, tv.Root, true)
}

func ShowNodeInfo(currentPackage string, n *TreeVisitorNode, isRoot bool) error {

	var sb strings.Builder
	sb.WriteString(n.Name)

	if isRoot {
		sb.WriteString(":= ")
	} else {
		sb.WriteString(": ")
	}

	if n.IsPtr {
		sb.WriteString("&")
	}

	if n.IsArray {
		sb.WriteString("[]")
	}

	if currentPackage != "" && strings.HasPrefix(n.Type, currentPackage) {
		sb.WriteString(n.Type[len(currentPackage)+1:])
	} else {
		sb.WriteString(n.Type)
	}

	if n.IsStruct || n.IsArray {
		sb.WriteString(" {")
	} else {
		sb.WriteString("Example(),")
	}

	fmt.Println(sb.String())
	if n.IsArray && n.IsStruct {
		fmt.Println("{")
	}
	for _, child := range n.Children {
		ShowNodeInfo(currentPackage, child, false)
	}
	if n.IsArray && n.IsStruct {
		fmt.Println("},")
	}

	if n.IsStruct || n.IsArray {
		sb.Reset()
		sb.WriteString(" }")
		if !isRoot {
			sb.WriteString(" ,")
		}
		fmt.Println(sb.String())
	}

	return nil
}
