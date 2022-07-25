package model

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type SimpleVisitedItem struct {
	Name     string
	Type     string
	IsStruct bool
	IsPtr    bool
	IsArray  bool
}

// TypeWithPackageTrimmed trim the type if the package matches the parameter. Used to avoid self references to package.
func (si SimpleVisitedItem) TypeWithPackageTrimmed(currentPackage string) string {

	n := si.Type
	if currentPackage != "" && strings.HasPrefix(si.Type, currentPackage) {
		n = si.Type[len(currentPackage)+1:]
	}

	return n
}

// NameOfTypeFunction returns the package if it is specified and doesn't match the one passed as param.
// optionally you can set the flag to add a trailing dot in the string
func (si SimpleVisitedItem) NameOfTypeFunction(functNamePrefix string, currentPackage string) string {

	var sb strings.Builder

	p := ""
	n := si.Type
	ndx := strings.Index(n, ".")
	if ndx > 0 {
		p = si.Type[0:ndx]
		n = si.Type[ndx+1:]
	}

	if p != currentPackage {
		sb.WriteString(p)
		sb.WriteString(".")
	}

	sb.WriteString(functNamePrefix)
	sb.WriteString(n)
	return sb.String()
}

func (si SimpleVisitedItem) String() string {
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

type SimpleVisitorPath []SimpleVisitedItem

func (p SimpleVisitorPath) Identifier() string {
	v := strings.TrimPrefix(p.Value(), "/Doc/")
	v = strings.Replace(v, ".", "_", -1)
	v = strings.Replace(v, "/", "_", -1)
	v = strings.Replace(v, "*", "", -1)
	v = strings.Replace(v, "[]", "", -1)
	return v
}

func (p SimpleVisitorPath) ItemPathReference(itemNdx int) string {
	var sb strings.Builder

	if itemNdx == 0 || itemNdx >= len(p) {
		return fmt.Sprintf("out of index %d (len = %d)", itemNdx, len(p))
	}
	// Take the intermediate nodes skipping the first.
	for i := 1; i <= (itemNdx - 1); i++ {
		el := p[i]
		if i > 1 {
			sb.WriteString(".")
		}
		sb.WriteString(el.Name)
		if el.IsArray {
			sb.WriteString("[0]")
		}
	}

	if itemNdx > 1 {
		sb.WriteString(".")
	}
	sb.WriteString(p[itemNdx].Name)
	return sb.String()
}

func (p SimpleVisitorPath) Value() string {
	var sb strings.Builder
	for ndx, item := range p {

		switch ndx {
		case 0:
			// first element is _self of type document
			// sb.WriteString("/Doc")
		case 1:
			/*
				if item.IsPtr {
					sb.WriteString("*")
				}
				if item.IsArray {
					sb.WriteString("[]")
				}
			*/
			sb.WriteString(item.Name)
			if item.IsArray {
				sb.WriteString("[]")
			}
		default:
			sb.WriteString(".")
			sb.WriteString(item.Name)
			if item.IsArray {
				sb.WriteString("[]")
			}
		}

	}

	return sb.String()
}

func (p SimpleVisitorPath) LastItem() SimpleVisitedItem {
	return p[len(p)-1]
}

func (p SimpleVisitorPath) LastItemIndex() int {
	return len(p) - 1
}

type SimpleVisitor struct {
	NumberOfLeaves int
	currentPath    SimpleVisitorPath
	Paths          []SimpleVisitorPath
}

func (sv SimpleVisitor) currentPathToString() string {
	return sv.pathToString(sv.currentPath)
}

func (sv SimpleVisitor) pathToString(p []SimpleVisitedItem) string {

	var sb strings.Builder
	for _, item := range p {
		sb.WriteString("/")
		sb.WriteString(item.String())
	}

	return sb.String()
}

func (sv *SimpleVisitor) Visit(name, aType string, isStruct, isPrt, isArray bool) (string, error) {
	item := SimpleVisitedItem{Name: name, Type: aType, IsStruct: isStruct, IsPtr: isPrt, IsArray: isArray}
	if !isStruct {
		sv.NumberOfLeaves++
	}
	sv.currentPath = append(sv.currentPath, item)

	var arr SimpleVisitorPath
	arr = append(arr, sv.currentPath...)
	sv.Paths = append(sv.Paths, arr)
	return sv.currentPathToString(), nil
}

func (sv *SimpleVisitor) Pop() (string, error) {

	if len(sv.currentPath) > 0 {
		sv.currentPath = sv.currentPath[0 : len(sv.currentPath)-1]
	} else {
		return "", fmt.Errorf("cannot pop from path: %s", sv.currentPathToString())
	}

	return sv.currentPathToString(), nil
}

func (sv *SimpleVisitor) ShowInfo(currentPackage string) error {
	for i, p := range sv.Paths {
		log.Trace().Int("path-num", i).Msg(sv.pathToString(p))
	}

	return nil
}
