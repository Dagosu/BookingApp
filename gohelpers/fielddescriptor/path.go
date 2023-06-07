package fielddescriptor

import (
	"regexp"
	"strconv"
	"strings"

	op "github.com/Dagosu/BookingAp/gohelpers/operations"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	replacer       = strings.NewReplacer("[", ".", "]", ".")
	fieldPathIndex = regexp.MustCompile(`\[.*?\]`)
	exportFields   = regexp.MustCompile(`\[[0-9]]*?\]`)
)

// NormalizePath will replace all parenthesis with dots
// eg: delay_codes[0]code => delay_codes.0.code
// eg: codeshares[0] => codeshares.0
func NormalizePath(fieldPath string) string {
	field := replacer.Replace(fieldPath)

	return strings.TrimSuffix(field, ".")
}

// MapPathToPolicyResource will match the fieldDescriptor path to the policy resource
// eg: delay_codes[0]code => delay_codes.code
// eg: codeshares[0] => codeshares
func MapPathToPolicyResource(fdPath string) string {
	paths := fieldPathIndex.FindAllString(fdPath, -1)

	if len(paths) >= 1 {
		for _, p := range paths {
			fdPath = strings.ReplaceAll(fdPath, p, ".")
		}
	}

	return strings.TrimSuffix(fdPath, ".")
}

// ExportedPath will return correctly mapped paths for exported files
// eg: vias[0]airport.iata => vias.airport.iata
// eg: calculated_fields[AEGS]datetime => calculated_fields.AEGS.datetime
func ExportedPath(fdPath string) string {
	return NormalizePath(exportFields.ReplaceAllString(fdPath, "."))
}

// UnsetValueByPath will unset a field (fieldDescriptor.path) from a proto
// currently works only for slices of objects
func UnsetValueByPath(pm protoreflect.Message, path string) {
	toRemove := path
	fieldPaths := strings.Split(toRemove, ".")
	fieldToRemove := fieldPaths[0]
	var subField, subFieldIndex string
	if len(fieldPaths) > 1 {
		subFieldIndex = fieldPaths[1]
	}
	if len(fieldPaths) > 2 {
		subField = fieldPaths[2]
	}

	pm.Range(func(field protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fieldName := string(field.Name())
		if fieldName == "id" {
			return true
		}
		if subFieldIndex != "" {
			toRemove = fieldToRemove + "." + subFieldIndex
		}
		if subField != "" {
			toRemove = subField
		}
		var intIndex int
		if i, err := strconv.Atoi(subFieldIndex); err == nil {
			intIndex = i
			if subField == "" {
				toRemove = fieldToRemove
			}
		}
		if fieldName == toRemove {
			pm.Clear(field)

			return true
		}
		if field.Kind() == protoreflect.MessageKind && fieldToRemove == fieldName && field.IsList() {
			UnsetValueByPath(v.List().Get(intIndex).Message(), subField)

			return true
		}

		return true
	})
}

// GetPaths will return all unique field descriptor paths (mongo compatible)
func GetPaths(fds map[int]*dt.FieldDescriptor) []string {
	fieldPaths := make([]string, 0, len(fds))

	for _, fd := range fds {
		field := MapPathToPolicyResource(fd.GetPath())

		if !op.Contains(field, fieldPaths) {
			fieldPaths = append(fieldPaths, field)
		}
	}

	return fieldPaths
}

// some of the field descriptors have the same path and the query will return error
// because there are duplicate paths in the $project section
// so we send the distinct paths to the GetFlights function
func GetExportFdsPaths(fds map[int]*dt.FieldDescriptor) []string {
	distinctNormalizedFds := []string{}

	for _, v := range fds {
		path := ExportedPath(v.GetPath())
		if !op.Contains(path, distinctNormalizedFds) {
			distinctNormalizedFds = append(distinctNormalizedFds, path)
		}
	}

	return distinctNormalizedFds
}
