// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: shooters/nexor/v1/nexor.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductCreated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SupplierId    string                 `protobuf:"bytes,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductCreated) Reset() {
	*x = ProductCreated{}
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreated) ProtoMessage() {}

func (x *ProductCreated) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreated.ProtoReflect.Descriptor instead.
func (*ProductCreated) Descriptor() ([]byte, []int) {
	return file_shooters_nexor_v1_nexor_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCreated) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *ProductCreated) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type SayHelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_shooters_nexor_v1_nexor_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_nexor_v1_nexor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_shooters_nexor_v1_nexor_proto_rawDescGZIP(), []int{2}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_shooters_nexor_v1_nexor_proto protoreflect.FileDescriptor

const file_shooters_nexor_v1_nexor_proto_rawDesc = "" +
	"\n" +
	"\x1dshooters/nexor/v1/nexor.proto\x12\x11shooters.nexor.v1\"t\n" +
	"\x0eProductCreated\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1f\n" +
	"\vsupplier_id\x18\x03 \x01(\tR\n" +
	"supplierId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\x03R\tcreatedAt\"%\n" +
	"\x0fSayHelloRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\",\n" +
	"\x10SayHelloResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessageB\xba\x01\n" +
	"\x15com.shooters.nexor.v1B\n" +
	"NexorProtoP\x01Z/github.com/shoot3rs/nexor/gen/shooters/nexor/v1\xa2\x02\x03SNX\xaa\x02\x11Shooters.Nexor.V1\xca\x02\x11Shooters\\Nexor\\V1\xe2\x02\x1dShooters\\Nexor\\V1\\GPBMetadata\xea\x02\x13Shooters::Nexor::V1b\x06proto3"

var (
	file_shooters_nexor_v1_nexor_proto_rawDescOnce sync.Once
	file_shooters_nexor_v1_nexor_proto_rawDescData []byte
)

func file_shooters_nexor_v1_nexor_proto_rawDescGZIP() []byte {
	file_shooters_nexor_v1_nexor_proto_rawDescOnce.Do(func() {
		file_shooters_nexor_v1_nexor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shooters_nexor_v1_nexor_proto_rawDesc), len(file_shooters_nexor_v1_nexor_proto_rawDesc)))
	})
	return file_shooters_nexor_v1_nexor_proto_rawDescData
}

var file_shooters_nexor_v1_nexor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shooters_nexor_v1_nexor_proto_goTypes = []any{
	(*ProductCreated)(nil),   // 0: shooters.nexor.v1.ProductCreated
	(*SayHelloRequest)(nil),  // 1: shooters.nexor.v1.SayHelloRequest
	(*SayHelloResponse)(nil), // 2: shooters.nexor.v1.SayHelloResponse
}
var file_shooters_nexor_v1_nexor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shooters_nexor_v1_nexor_proto_init() }
func file_shooters_nexor_v1_nexor_proto_init() {
	if File_shooters_nexor_v1_nexor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shooters_nexor_v1_nexor_proto_rawDesc), len(file_shooters_nexor_v1_nexor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shooters_nexor_v1_nexor_proto_goTypes,
		DependencyIndexes: file_shooters_nexor_v1_nexor_proto_depIdxs,
		MessageInfos:      file_shooters_nexor_v1_nexor_proto_msgTypes,
	}.Build()
	File_shooters_nexor_v1_nexor_proto = out.File
	file_shooters_nexor_v1_nexor_proto_goTypes = nil
	file_shooters_nexor_v1_nexor_proto_depIdxs = nil
}
