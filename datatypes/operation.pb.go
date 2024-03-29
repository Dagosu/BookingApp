// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: operation.proto

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

type OperationType int32

const (
	OperationType_UNKNOWN_UNSPECIFIED OperationType = 0
	OperationType_INSERT              OperationType = 1
	OperationType_REPLACE             OperationType = 2
	OperationType_UPDATE              OperationType = 3
	OperationType_DELETE              OperationType = 4
	// initial data send completed, but server remains connected
	// so it can send further updates
	OperationType_READY OperationType = 5
	// on client display the `error` field
	OperationType_ERROR OperationType = 6
	// used to notify client about a long operation status
	OperationType_PROGRESS OperationType = 7
	// operation was finished, server should disconnect after sending this
	OperationType_FINISHED OperationType = 8
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "UNKNOWN_UNSPECIFIED",
		1: "INSERT",
		2: "REPLACE",
		3: "UPDATE",
		4: "DELETE",
		5: "READY",
		6: "ERROR",
		7: "PROGRESS",
		8: "FINISHED",
	}
	OperationType_value = map[string]int32{
		"UNKNOWN_UNSPECIFIED": 0,
		"INSERT":              1,
		"REPLACE":             2,
		"UPDATE":              3,
		"DELETE":              4,
		"READY":               5,
		"ERROR":               6,
		"PROGRESS":            7,
		"FINISHED":            8,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_operation_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_operation_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{0}
}

var File_operation_proto protoreflect.FileDescriptor

var file_operation_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x8b, 0x01, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52,
	0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x08, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x67, 0x6f, 0x73, 0x75, 0x2f,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operation_proto_rawDescOnce sync.Once
	file_operation_proto_rawDescData = file_operation_proto_rawDesc
)

func file_operation_proto_rawDescGZIP() []byte {
	file_operation_proto_rawDescOnce.Do(func() {
		file_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_operation_proto_rawDescData)
	})
	return file_operation_proto_rawDescData
}

var file_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_operation_proto_goTypes = []interface{}{
	(OperationType)(0), // 0: operation.OperationType
}
var file_operation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_operation_proto_init() }
func file_operation_proto_init() {
	if File_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_operation_proto_goTypes,
		DependencyIndexes: file_operation_proto_depIdxs,
		EnumInfos:         file_operation_proto_enumTypes,
	}.Build()
	File_operation_proto = out.File
	file_operation_proto_rawDesc = nil
	file_operation_proto_goTypes = nil
	file_operation_proto_depIdxs = nil
}
