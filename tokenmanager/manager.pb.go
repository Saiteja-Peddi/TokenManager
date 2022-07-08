// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: manager.proto

package tokenmanagerpb

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

type NormalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *NormalRequest) Reset() {
	*x = NormalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalRequest) ProtoMessage() {}

func (x *NormalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalRequest.ProtoReflect.Descriptor instead.
func (*NormalRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *NormalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Low  uint64 `protobuf:"varint,3,opt,name=Low,proto3" json:"Low,omitempty"`
	Mid  uint64 `protobuf:"varint,4,opt,name=Mid,proto3" json:"Mid,omitempty"`
	High uint64 `protobuf:"varint,5,opt,name=High,proto3" json:"High,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WriteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteRequest) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WriteRequest) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WriteRequest) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message uint64 `protobuf:"varint,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ServerResponse) GetMessage() uint64 {
	if x != nil {
		return x.Message
	}
	return 0
}

func (x *ServerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Low   uint64 `protobuf:"varint,3,opt,name=Low,proto3" json:"Low,omitempty"`
	Mid   uint64 `protobuf:"varint,4,opt,name=Mid,proto3" json:"Mid,omitempty"`
	High  uint64 `protobuf:"varint,5,opt,name=High,proto3" json:"High,omitempty"`
	Final uint64 `protobuf:"varint,6,opt,name=Final,proto3" json:"Final,omitempty"`
	Ts    uint64 `protobuf:"varint,7,opt,name=Ts,proto3" json:"Ts,omitempty"`
}

func (x *CopyRequest) Reset() {
	*x = CopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyRequest) ProtoMessage() {}

func (x *CopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyRequest.ProtoReflect.Descriptor instead.
func (*CopyRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *CopyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CopyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CopyRequest) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *CopyRequest) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *CopyRequest) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *CopyRequest) GetFinal() uint64 {
	if x != nil {
		return x.Final
	}
	return 0
}

func (x *CopyRequest) GetTs() uint64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

var File_manager_proto protoreflect.FileDescriptor

var file_manager_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x1f, 0x0a,
	0x0d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x6a,
	0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x4c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x4d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x22, 0x3a, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f,
	0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x4c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x4d, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x48, 0x69,
	0x67, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x54, 0x73, 0x32, 0xec, 0x02, 0x0a, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x44, 0x72,
	0x6f, 0x70, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x70, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x4f, 0x53, 0x5f, 0x50, 0x52, 0x4a, 0x32, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_manager_proto_rawDescOnce sync.Once
	file_manager_proto_rawDescData = file_manager_proto_rawDesc
)

func file_manager_proto_rawDescGZIP() []byte {
	file_manager_proto_rawDescOnce.Do(func() {
		file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_proto_rawDescData)
	})
	return file_manager_proto_rawDescData
}

var file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_manager_proto_goTypes = []interface{}{
	(*NormalRequest)(nil),  // 0: tokenmanager.NormalRequest
	(*WriteRequest)(nil),   // 1: tokenmanager.WriteRequest
	(*ServerResponse)(nil), // 2: tokenmanager.ServerResponse
	(*CopyRequest)(nil),    // 3: tokenmanager.CopyRequest
}
var file_manager_proto_depIdxs = []int32{
	0, // 0: tokenmanager.TokenManager.Create:input_type -> tokenmanager.NormalRequest
	1, // 1: tokenmanager.TokenManager.Write:input_type -> tokenmanager.WriteRequest
	0, // 2: tokenmanager.TokenManager.Read:input_type -> tokenmanager.NormalRequest
	0, // 3: tokenmanager.TokenManager.Drop:input_type -> tokenmanager.NormalRequest
	3, // 4: tokenmanager.TokenManager.CopyToken:input_type -> tokenmanager.CopyRequest
	2, // 5: tokenmanager.TokenManager.Create:output_type -> tokenmanager.ServerResponse
	2, // 6: tokenmanager.TokenManager.Write:output_type -> tokenmanager.ServerResponse
	2, // 7: tokenmanager.TokenManager.Read:output_type -> tokenmanager.ServerResponse
	2, // 8: tokenmanager.TokenManager.Drop:output_type -> tokenmanager.ServerResponse
	2, // 9: tokenmanager.TokenManager.CopyToken:output_type -> tokenmanager.ServerResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_manager_proto_init() }
func file_manager_proto_init() {
	if File_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalRequest); i {
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
		file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerResponse); i {
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
		file_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyRequest); i {
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
			RawDescriptor: file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_proto_goTypes,
		DependencyIndexes: file_manager_proto_depIdxs,
		MessageInfos:      file_manager_proto_msgTypes,
	}.Build()
	File_manager_proto = out.File
	file_manager_proto_rawDesc = nil
	file_manager_proto_goTypes = nil
	file_manager_proto_depIdxs = nil
}
