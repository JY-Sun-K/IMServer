// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.0
// source: letter.proto

package pb

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

type MsgPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressIp string `protobuf:"bytes,1,opt,name=AddressIp,proto3" json:"AddressIp,omitempty"`
	From      int64  `protobuf:"varint,2,opt,name=From,proto3" json:"From,omitempty"`
	To        int64  `protobuf:"varint,3,opt,name=To,proto3" json:"To,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=Msg,proto3" json:"Msg,omitempty"`
	SendTime  string `protobuf:"bytes,5,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
}

func (x *MsgPoint) Reset() {
	*x = MsgPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_letter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPoint) ProtoMessage() {}

func (x *MsgPoint) ProtoReflect() protoreflect.Message {
	mi := &file_letter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPoint.ProtoReflect.Descriptor instead.
func (*MsgPoint) Descriptor() ([]byte, []int) {
	return file_letter_proto_rawDescGZIP(), []int{0}
}

func (x *MsgPoint) GetAddressIp() string {
	if x != nil {
		return x.AddressIp
	}
	return ""
}

func (x *MsgPoint) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MsgPoint) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *MsgPoint) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MsgPoint) GetSendTime() string {
	if x != nil {
		return x.SendTime
	}
	return ""
}

type WriteStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MP *MsgPoint `protobuf:"bytes,1,opt,name=MP,proto3" json:"MP,omitempty"`
}

func (x *WriteStreamRequest) Reset() {
	*x = WriteStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_letter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteStreamRequest) ProtoMessage() {}

func (x *WriteStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_letter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteStreamRequest.ProtoReflect.Descriptor instead.
func (*WriteStreamRequest) Descriptor() ([]byte, []int) {
	return file_letter_proto_rawDescGZIP(), []int{1}
}

func (x *WriteStreamRequest) GetMP() *MsgPoint {
	if x != nil {
		return x.MP
	}
	return nil
}

type WriteStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err  string `protobuf:"bytes,1,opt,name=Err,proto3" json:"Err,omitempty"`
	Code int64  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *WriteStreamResponse) Reset() {
	*x = WriteStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_letter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteStreamResponse) ProtoMessage() {}

func (x *WriteStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_letter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteStreamResponse.ProtoReflect.Descriptor instead.
func (*WriteStreamResponse) Descriptor() ([]byte, []int) {
	return file_letter_proto_rawDescGZIP(), []int{2}
}

func (x *WriteStreamResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *WriteStreamResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_letter_proto protoreflect.FileDescriptor

var file_letter_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x7a, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32,
	0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x02, 0x4d, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02,
	0x4d, 0x50, 0x22, 0x3b, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x55, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_letter_proto_rawDescOnce sync.Once
	file_letter_proto_rawDescData = file_letter_proto_rawDesc
)

func file_letter_proto_rawDescGZIP() []byte {
	file_letter_proto_rawDescOnce.Do(func() {
		file_letter_proto_rawDescData = protoimpl.X.CompressGZIP(file_letter_proto_rawDescData)
	})
	return file_letter_proto_rawDescData
}

var file_letter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_letter_proto_goTypes = []interface{}{
	(*MsgPoint)(nil),            // 0: pb.MsgPoint
	(*WriteStreamRequest)(nil),  // 1: pb.WriteStreamRequest
	(*WriteStreamResponse)(nil), // 2: pb.WriteStreamResponse
}
var file_letter_proto_depIdxs = []int32{
	0, // 0: pb.WriteStreamRequest.MP:type_name -> pb.MsgPoint
	1, // 1: pb.StreamService.WriteStream:input_type -> pb.WriteStreamRequest
	2, // 2: pb.StreamService.WriteStream:output_type -> pb.WriteStreamResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_letter_proto_init() }
func file_letter_proto_init() {
	if File_letter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_letter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPoint); i {
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
		file_letter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteStreamRequest); i {
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
		file_letter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteStreamResponse); i {
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
			RawDescriptor: file_letter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_letter_proto_goTypes,
		DependencyIndexes: file_letter_proto_depIdxs,
		MessageInfos:      file_letter_proto_msgTypes,
	}.Build()
	File_letter_proto = out.File
	file_letter_proto_rawDesc = nil
	file_letter_proto_goTypes = nil
	file_letter_proto_depIdxs = nil
}