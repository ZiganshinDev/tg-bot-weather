// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: services/proto/cityservice.proto

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

type CityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityName string `protobuf:"bytes,1,opt,name=cityName,proto3" json:"cityName,omitempty"`
}

func (x *CityRequest) Reset() {
	*x = CityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_cityservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRequest) ProtoMessage() {}

func (x *CityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_cityservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityRequest.ProtoReflect.Descriptor instead.
func (*CityRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_cityservice_proto_rawDescGZIP(), []int{0}
}

func (x *CityRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

type CoordinatesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *CoordinatesReply) Reset() {
	*x = CoordinatesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_cityservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinatesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinatesReply) ProtoMessage() {}

func (x *CoordinatesReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_cityservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinatesReply.ProtoReflect.Descriptor instead.
func (*CoordinatesReply) Descriptor() ([]byte, []int) {
	return file_services_proto_cityservice_proto_rawDescGZIP(), []int{1}
}

func (x *CoordinatesReply) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CoordinatesReply) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_services_proto_cityservice_proto protoreflect.FileDescriptor

var file_services_proto_cityservice_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x29, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x32, 0x5c, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x69, 0x67, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x44, 0x65,
	0x76, 0x2f, 0x74, 0x67, 0x2d, 0x62, 0x6f, 0x74, 0x2d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_cityservice_proto_rawDescOnce sync.Once
	file_services_proto_cityservice_proto_rawDescData = file_services_proto_cityservice_proto_rawDesc
)

func file_services_proto_cityservice_proto_rawDescGZIP() []byte {
	file_services_proto_cityservice_proto_rawDescOnce.Do(func() {
		file_services_proto_cityservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_cityservice_proto_rawDescData)
	})
	return file_services_proto_cityservice_proto_rawDescData
}

var file_services_proto_cityservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_proto_cityservice_proto_goTypes = []interface{}{
	(*CityRequest)(nil),      // 0: cityservice.CityRequest
	(*CoordinatesReply)(nil), // 1: cityservice.CoordinatesReply
}
var file_services_proto_cityservice_proto_depIdxs = []int32{
	0, // 0: cityservice.CityService.GetCityCoordinates:input_type -> cityservice.CityRequest
	1, // 1: cityservice.CityService.GetCityCoordinates:output_type -> cityservice.CoordinatesReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_cityservice_proto_init() }
func file_services_proto_cityservice_proto_init() {
	if File_services_proto_cityservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_cityservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityRequest); i {
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
		file_services_proto_cityservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinatesReply); i {
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
			RawDescriptor: file_services_proto_cityservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_cityservice_proto_goTypes,
		DependencyIndexes: file_services_proto_cityservice_proto_depIdxs,
		MessageInfos:      file_services_proto_cityservice_proto_msgTypes,
	}.Build()
	File_services_proto_cityservice_proto = out.File
	file_services_proto_cityservice_proto_rawDesc = nil
	file_services_proto_cityservice_proto_goTypes = nil
	file_services_proto_cityservice_proto_depIdxs = nil
}
