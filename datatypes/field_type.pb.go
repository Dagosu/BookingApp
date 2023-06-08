// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: field_type.proto

package datatypes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FieldType details on how to render a certain field.
// The types should be limited to primitives. Expansion on custom types should
// be indicated by field type = CUSTOM + field `type_hint`
type FieldType int32

const (
	// FIELD_TYPE_UNSPECIFIED is the default value, it should not be used
	FieldType_FIELD_TYPE_UNSPECIFIED FieldType = 0
	// FIELD_TYPE_STRING is used for text
	FieldType_FIELD_TYPE_STRING FieldType = 1
	// FIELD_TYPE_INTEGER is used for numbers without decimals
	FieldType_FIELD_TYPE_INTEGER FieldType = 2
	// FIELD_TYPE_DATETIME is used for absolute timestamps
	FieldType_FIELD_TYPE_DATETIME FieldType = 3
	// FIELD_TYPE_BOOLEAN is used for boolean values
	FieldType_FIELD_TYPE_BOOLEAN FieldType = 4
	// TDS: Move this to somewhere else
	// https://airportlabs.atlassian.net/browse/AODBSC-266
	FieldType_FIELD_TYPE_MDM_RESOURCE FieldType = 5
	// FIELD_TYPE_CUSTOM is used for custom types
	FieldType_FIELD_TYPE_CUSTOM FieldType = 6
	// FIELD_TYPE_DURATION is used for relative time durations
	FieldType_FIELD_TYPE_DURATION FieldType = 7
	// TDS: Move this to somewhere else
	// https://airportlabs.atlassian.net/browse/AODBSC-266
	FieldType_FIELD_TYPE_CALCULATED_FIELD FieldType = 8
	// FIELD_TYPE_ENUM is used for enums
	FieldType_FIELD_TYPE_ENUM FieldType = 9
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0: "FIELD_TYPE_UNSPECIFIED",
		1: "FIELD_TYPE_STRING",
		2: "FIELD_TYPE_INTEGER",
		3: "FIELD_TYPE_DATETIME",
		4: "FIELD_TYPE_BOOLEAN",
		5: "FIELD_TYPE_MDM_RESOURCE",
		6: "FIELD_TYPE_CUSTOM",
		7: "FIELD_TYPE_DURATION",
		8: "FIELD_TYPE_CALCULATED_FIELD",
		9: "FIELD_TYPE_ENUM",
	}
	FieldType_value = map[string]int32{
		"FIELD_TYPE_UNSPECIFIED":      0,
		"FIELD_TYPE_STRING":           1,
		"FIELD_TYPE_INTEGER":          2,
		"FIELD_TYPE_DATETIME":         3,
		"FIELD_TYPE_BOOLEAN":          4,
		"FIELD_TYPE_MDM_RESOURCE":     5,
		"FIELD_TYPE_CUSTOM":           6,
		"FIELD_TYPE_DURATION":         7,
		"FIELD_TYPE_CALCULATED_FIELD": 8,
		"FIELD_TYPE_ENUM":             9,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_field_type_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_field_type_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_field_type_proto_rawDescGZIP(), []int{0}
}

var File_field_type_proto protoreflect.FileDescriptor

var file_field_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x8a,
	0x02, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x44, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x43, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x10, 0x09, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x67, 0x6f, 0x73, 0x75,
	0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_field_type_proto_rawDescOnce sync.Once
	file_field_type_proto_rawDescData = file_field_type_proto_rawDesc
)

func file_field_type_proto_rawDescGZIP() []byte {
	file_field_type_proto_rawDescOnce.Do(func() {
		file_field_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_type_proto_rawDescData)
	})
	return file_field_type_proto_rawDescData
}

var file_field_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_type_proto_goTypes = []interface{}{
	(FieldType)(0), // 0: field_type.FieldType
}
var file_field_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_field_type_proto_init() }
func file_field_type_proto_init() {
	if File_field_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_proto_goTypes,
		DependencyIndexes: file_field_type_proto_depIdxs,
		EnumInfos:         file_field_type_proto_enumTypes,
	}.Build()
	File_field_type_proto = out.File
	file_field_type_proto_rawDesc = nil
	file_field_type_proto_goTypes = nil
	file_field_type_proto_depIdxs = nil
}