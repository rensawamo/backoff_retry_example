// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/some_service.proto

package proto

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

// リクエストメッセージ
type SomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SomeRequest) Reset() {
	*x = SomeRequest{}
	mi := &file_proto_some_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeRequest) ProtoMessage() {}

func (x *SomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_some_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeRequest.ProtoReflect.Descriptor instead.
func (*SomeRequest) Descriptor() ([]byte, []int) {
	return file_proto_some_service_proto_rawDescGZIP(), []int{0}
}

func (x *SomeRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// レスポンスメッセージ
type SomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *SomeResponse) Reset() {
	*x = SomeResponse{}
	mi := &file_proto_some_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeResponse) ProtoMessage() {}

func (x *SomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_some_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeResponse.ProtoReflect.Descriptor instead.
func (*SomeResponse) Descriptor() ([]byte, []int) {
	return file_proto_some_service_proto_rawDescGZIP(), []int{1}
}

func (x *SomeResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_proto_some_service_proto protoreflect.FileDescriptor

var file_proto_some_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6f, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x52, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x98, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x10, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0xca, 0x02, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02,
	0x17, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_some_service_proto_rawDescOnce sync.Once
	file_proto_some_service_proto_rawDescData = file_proto_some_service_proto_rawDesc
)

func file_proto_some_service_proto_rawDescGZIP() []byte {
	file_proto_some_service_proto_rawDescOnce.Do(func() {
		file_proto_some_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_some_service_proto_rawDescData)
	})
	return file_proto_some_service_proto_rawDescData
}

var file_proto_some_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_some_service_proto_goTypes = []any{
	(*SomeRequest)(nil),  // 0: some_service.SomeRequest
	(*SomeResponse)(nil), // 1: some_service.SomeResponse
}
var file_proto_some_service_proto_depIdxs = []int32{
	0, // 0: some_service.SomeService.SomeMethod:input_type -> some_service.SomeRequest
	1, // 1: some_service.SomeService.SomeMethod:output_type -> some_service.SomeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_some_service_proto_init() }
func file_proto_some_service_proto_init() {
	if File_proto_some_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_some_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_some_service_proto_goTypes,
		DependencyIndexes: file_proto_some_service_proto_depIdxs,
		MessageInfos:      file_proto_some_service_proto_msgTypes,
	}.Build()
	File_proto_some_service_proto = out.File
	file_proto_some_service_proto_rawDesc = nil
	file_proto_some_service_proto_goTypes = nil
	file_proto_some_service_proto_depIdxs = nil
}
