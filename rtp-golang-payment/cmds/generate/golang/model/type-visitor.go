package model

import "strings"

type TypeVisitor struct {
	NumberOfLeaves int
	Types          map[string]TypeVisitedItem
}

type TypeVisitedItem struct {
	Name     string
	IsStruct bool
	IsPtr    bool
	IsArray  bool
}

func (sv *TypeVisitor) Visit(name, aType string, isStruct, isPrt, isArray bool) (string, error) {
	item := TypeVisitedItem{Name: aType, IsStruct: isStruct, IsPtr: isPrt, IsArray: isArray}
	if !isStruct {
		sv.NumberOfLeaves++
	}

	n := aType
	if isStruct {
		// In here in the map you could get different keys to the same type if it is used as a pointer or as an array or as a struct,
		// This is wrong for simple types but I keep here for complex types for future use...
		var sb strings.Builder
		if isPrt {
			sb.WriteString("*")
		}

		if isArray {
			sb.WriteString("[]")
		}

		sb.WriteString(aType)
		if sv.Types == nil {
			sv.Types = make(map[string]TypeVisitedItem)
		}

		n = sb.String()
	}
	sv.Types[n] = item
	return "", nil
}

func (sv *TypeVisitor) Pop() (string, error) {
	return "", nil
}

// TypeWithPackageTrimmed trim the type if the package matches the parameter. Used to avoid self references to package.
func (si TypeVisitedItem) TypeWithPackageTrimmed(currentPackage string) string {

	n := si.Name
	if currentPackage != "" && strings.HasPrefix(si.Name, currentPackage) {
		n = si.Name[len(currentPackage)+1:]
	}

	return n
}

// NameOfTypeFunction returns the package if it is specified and doesn't match the one passed as param.
// optionally you can set the flag to add a trailing dot in the string
func (si TypeVisitedItem) NameOfTypeFunction(functNamePrefix string, currentPackage string) string {

	var sb strings.Builder

	p := ""
	n := si.Name
	ndx := strings.Index(n, ".")
	if ndx > 0 {
		p = si.Name[0:ndx]
		n = si.Name[ndx+1:]
	}

	if p != currentPackage {
		sb.WriteString(p)
		sb.WriteString(".")
	}

	sb.WriteString(functNamePrefix)
	sb.WriteString(n)
	return sb.String()
}
