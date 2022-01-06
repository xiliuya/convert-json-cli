package model

import (
	"fmt"
	"os"
	"strings"
)

type GoWriter struct {
	outputPath string
	nested     bool
}

func (gw *GoWriter) SetOutputPath(outputPath string) {
	gw.outputPath = outputPath
}

func (gw *GoWriter) SetNested(nested bool) {
	gw.nested = nested
}

func (gw *GoWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(gw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("package model\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(gw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (gw *GoWriter) GetStruct(abstractStruct Struct) string {
	var nested string = ""
	if abstractStruct.Name == "AutoGenerated" && gw.nested {
		nested = "[]"
	}

	return "type " + abstractStruct.Name + " " + nested + "struct {\n" + gw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (gw *GoWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += gw.GetField(i, field)
	}

	return result
}

func (gw *GoWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "[]" + gw.GetTypeName(field.TypeName[:len(field.TypeName)-2])
	} else {
		typeName = gw.GetTypeName(field.TypeName)
	}

	return "\t" + strings.Title(field.Index) + "\t" + typeName + "\t" + fmt.Sprintf("`json:\"%s\"`", field.Index) + "\n"
}

func (gw *GoWriter) GetTypeName(typeName string) string {
	return typeName
}
