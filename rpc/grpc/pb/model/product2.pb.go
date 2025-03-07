// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/product2.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items       []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Order2) Reset() {
	*x = Order2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order2) ProtoMessage() {}

func (x *Order2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order2.ProtoReflect.Descriptor instead.
func (*Order2) Descriptor() ([]byte, []int) {
	return file_proto_product2_proto_rawDescGZIP(), []int{0}
}

func (x *Order2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order2) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order2) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order2) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order2) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

var File_proto_product2_proto protoreflect.FileDescriptor

var file_proto_product2_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01,
	0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x51, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x12, 0x3d, 0x0a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product2_proto_rawDescOnce sync.Once
	file_proto_product2_proto_rawDescData = file_proto_product2_proto_rawDesc
)

func file_proto_product2_proto_rawDescGZIP() []byte {
	file_proto_product2_proto_rawDescOnce.Do(func() {
		file_proto_product2_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product2_proto_rawDescData)
	})
	return file_proto_product2_proto_rawDescData
}

var file_proto_product2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_product2_proto_goTypes = []interface{}{
	(*Order2)(nil),                 // 0: proto.Order2
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_proto_product2_proto_depIdxs = []int32{
	1, // 0: proto.OrderManagement2.searchOrders:input_type -> google.protobuf.StringValue
	0, // 1: proto.OrderManagement2.searchOrders:output_type -> proto.Order2
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_product2_proto_init() }
func file_proto_product2_proto_init() {
	if File_proto_product2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order2); i {
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
			RawDescriptor: file_proto_product2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product2_proto_goTypes,
		DependencyIndexes: file_proto_product2_proto_depIdxs,
		MessageInfos:      file_proto_product2_proto_msgTypes,
	}.Build()
	File_proto_product2_proto = out.File
	file_proto_product2_proto_rawDesc = nil
	file_proto_product2_proto_goTypes = nil
	file_proto_product2_proto_depIdxs = nil
}
