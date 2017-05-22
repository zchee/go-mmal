// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-clang/v3.9/clang"
)

const (
	prefix = "MMAL"
)

var builtinTypes = map[string]string{
	"int8_t":       "int8",
	"int16_t":      "int16",
	"int32_t":      "int32",
	"int64_t":      "int64",
	"uint8_t":      "uint8",
	"uint16_t":     "uint16",
	"uint32_t":     "uint32",
	"uint64_t":     "uint64",
	"unsigned int": "uint",
}

var header = os.Args[1]
var templ = `package ` + strings.ToLower(prefix) + `

/*
#include <` + trimIncludeDir(header) + `>
*/
import "C"`

func trimIncludeDir(d string) string {
	var dd []string
	for _, sep := range []string{"userland/", "include/"} {
		if strings.Contains(d, sep) {
			dd = strings.Split(d, sep)
		}
	}
	return dd[1]
}

func main() {
	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(header, nil, nil, clang.DefaultEditingTranslationUnitOptions()|clang.TranslationUnit_DetailedPreprocessingRecord|clang.TranslationUnit_KeepGoing)
	defer tu.Dispose()

	var buf bytes.Buffer
	buf.WriteString(templ + "\n\n")

	seen := make(map[string]bool)
	firstEnum := false
	visitor := func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		spelling := cursor.Spelling()
		if cursor.IsNull() {
			return clang.ChildVisit_Continue
		}
		if file, _, _, _ := cursor.Location().FileLocation(); file.Name() != header {
			return clang.ChildVisit_Continue
		}

		switch kind := cursor.Kind(); kind {
		case clang.Cursor_FunctionDecl:
			buf.WriteString(fmt.Sprintf("func %s(", UpperCase(spelling)))
			numArgs := cursor.NumArguments()
			for i := uint32(0); i < uint32(numArgs); i++ {
				buf.WriteString(fmt.Sprintf("%s %s", cursor.Argument(i).Spelling(), formatType(cursor.Argument(i).Type().Spelling())))
				if i+1 < uint32(numArgs) {
					buf.WriteString(", ")
				}
			}
			resType := UpperCase(cursor.ResultType().Spelling())
			buf.WriteString(fmt.Sprintf(") %s {\n", resType))
			buf.WriteString(fmt.Sprintf("\treturn "+resType+"(C."+cursor.Spelling()) + "(")
			for i := uint32(0); i < uint32(numArgs); i++ {
				arg := cursor.Argument(i).Spelling()
				switch arg {
				case "port":
					arg = "port.c"
				}
				buf.WriteString(arg)
				if i+1 < uint32(numArgs) {
					buf.WriteString(", ")
				}
			}
			buf.WriteString("))\n")
			buf.WriteString("}\n\n")
		case clang.Cursor_StructDecl:
			if seen[cursor.USR()] {
				return clang.ChildVisit_Continue
			}
			buf.WriteString("type " + UpperCase(spelling) + " struct {\n\tc C." + spelling + "\n}\n\n")
			seen[cursor.USR()] = true
		case clang.Cursor_FieldDecl:
			if ptype := parent.Spelling(); ptype != "" {
				var builtin bool
				var retv string
				if t, ok := builtinTypes[cursor.Type().Spelling()]; ok {
					retv = t
					builtin = true
				} else {
					retv = UpperCase(cursor.Type().Spelling())
				}
				buf.WriteString("func (" + ReceiverName(ptype) + " " + UpperCase(ptype) + ") " + UpperCase(spelling) + "() " + retv + " {\n")
				buf.WriteString("\treturn " + retv)
				if builtin {
					buf.WriteString("(" + ReceiverName(ptype) + ".c." + spelling + ")\n}\n\n")
				} else {
					buf.WriteString("{" + ReceiverName(ptype) + ".c." + spelling + "}\n}\n\n")
				}
			}
		case clang.Cursor_TypeRef:
			// noting to do
		case clang.Cursor_EnumDecl:
			if seen[cursor.USR()] {
				buf.WriteString(")\n\n")
				return clang.ChildVisit_Continue
			}
			if spelling != "" {
				buf.WriteString("type " + UpperCase(spelling) + " C." + spelling + "\n\n")
				firstEnum = true
			}
			seen[cursor.USR()] = true
			buf.WriteString("const (\n")
		case clang.Cursor_EnumConstantDecl:
			if firstEnum {
				buf.WriteString("\t" + UpperCase(spelling) + " " + UpperCase(parent.Spelling()) + " = iota\n")
				firstEnum = false
			} else {
				buf.WriteString("\t" + UpperCase(spelling) + "\n")
			}
		case clang.Cursor_IntegerLiteral:
			if spelling != "0" {
				fmt.Println(spelling)
			}
		case clang.Cursor_MacroDefinition:
			// fmt.Printf("kind: %s, spelling: %s\n", cursor.Kind(), spelling)
			// WriteConst(&buf, cursor)
		default:
			// fmt.Printf("kind: %s, spelling: %s\n", cursor.Kind(), spelling)
		}

		return clang.ChildVisit_Recurse
	}

	tu.TranslationUnitCursor().Visit(visitor)
	fmt.Println(buf.String())
}

func WriteConst(buf io.Writer, cursor clang.Cursor) {
	buf.Write([]byte("const " + cursor.Spelling() + " = " + cursor.ResultType().Spelling() + "\n"))
}

func ReceiverName(s string) string {
	return string(bytes.ToLower([]byte{TrimLibPrefix(s)[0]}))
}

func TrimLibPrefix(s string) string {
	return strings.TrimPrefix(s, prefix+"_")
}

func formatType(s string) string {
	s = strings.TrimPrefix(s, prefix+"_")
	if v, ok := builtinTypes[s]; ok {
		return v
	}
	if strings.HasSuffix(s, " *") {
		s = strings.Replace(s, " *", "", 1)
	}
	return UpperCase(s)
}

// LowerCase returns s with the first character lower-cased. LowerCase
// assumes s is an ASCII-represented string.
func LowerCase(s string) string {
	if len(s) == 0 {
		return s
	}
	var buf bytes.Buffer
	ss := strings.Split(s, "_")
	for _, n := range ss {
		if n == prefix {
			continue
		}
		if len(n) == 1 {
			buf.WriteString(string(n[0] | ' '))
			continue
		}
		buf.WriteString(string(n[0]|' ') + strings.ToLower(n[1:]))
	}
	return buf.String()
	if len(s) == 0 {
		return s
	}
	return buf.String()
}

// UpperCase returns s with the first character upper-cased. UpperCase
// assumes s is an ASCII-represented string.
func UpperCase(s string) string {
	if len(s) == 0 {
		return s
	}
	if v, ok := builtinTypes[s]; ok {
		return v
	}
	var buf bytes.Buffer
	ss := strings.Split(s, "_")
	for _, n := range ss {
		if n == prefix || n == strings.ToLower(prefix) {
			continue
		}
		if len(n) == 1 {
			if n == "T" {
				buf.WriteString("Type")
				continue
			}
			buf.WriteString(string(n[0] &^ ' '))
			continue
		}
		buf.WriteString(string(n[0]&^' ') + strings.ToLower(n[1:]))
	}
	return buf.String()
}
