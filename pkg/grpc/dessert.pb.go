// versionの指定
// 指定がなければデフォルトで2になる

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: dessert.proto

// packageの宣言
// Protocol Buffersのパッケージ名。
//「パッケージ名.型名」という形で他のprotoファイル内の型を参照することができる

package grpc

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

// リクエストの定義
type DessertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// データ型 フィールド名 = フィールド番号
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DessertRequest) Reset() {
	*x = DessertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dessert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DessertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DessertRequest) ProtoMessage() {}

func (x *DessertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dessert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DessertRequest.ProtoReflect.Descriptor instead.
func (*DessertRequest) Descriptor() ([]byte, []int) {
	return file_dessert_proto_rawDescGZIP(), []int{0}
}

func (x *DessertRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DessertRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// レスポンスの定義
type DessertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DessertResponse) Reset() {
	*x = DessertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dessert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DessertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DessertResponse) ProtoMessage() {}

func (x *DessertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dessert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DessertResponse.ProtoReflect.Descriptor instead.
func (*DessertResponse) Descriptor() ([]byte, []int) {
	return file_dessert_proto_rawDescGZIP(), []int{1}
}

func (x *DessertResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DessertResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_dessert_proto protoreflect.FileDescriptor

var file_dessert_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x22, 0x34, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x73,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47,
	0x0a, 0x0f, 0x44, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x59, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dessert_proto_rawDescOnce sync.Once
	file_dessert_proto_rawDescData = file_dessert_proto_rawDesc
)

func file_dessert_proto_rawDescGZIP() []byte {
	file_dessert_proto_rawDescOnce.Do(func() {
		file_dessert_proto_rawDescData = protoimpl.X.CompressGZIP(file_dessert_proto_rawDescData)
	})
	return file_dessert_proto_rawDescData
}

var file_dessert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dessert_proto_goTypes = []interface{}{
	(*DessertRequest)(nil),  // 0: grpcApi.DessertRequest
	(*DessertResponse)(nil), // 1: grpcApi.DessertResponse
}
var file_dessert_proto_depIdxs = []int32{
	0, // 0: grpcApi.DessertService.GetDessertStream:input_type -> grpcApi.DessertRequest
	1, // 1: grpcApi.DessertService.GetDessertStream:output_type -> grpcApi.DessertResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dessert_proto_init() }
func file_dessert_proto_init() {
	if File_dessert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dessert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DessertRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dessert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DessertResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dessert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dessert_proto_goTypes,
		DependencyIndexes: file_dessert_proto_depIdxs,
		MessageInfos:      file_dessert_proto_msgTypes,
	}.Build()
	File_dessert_proto = out.File
	file_dessert_proto_rawDesc = nil
	file_dessert_proto_goTypes = nil
	file_dessert_proto_depIdxs = nil
}
