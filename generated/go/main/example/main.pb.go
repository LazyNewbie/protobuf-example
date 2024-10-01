// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto_files/main.proto

package example

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         int32  `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Time         int32  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Country      string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	IdLanguage   int32  `protobuf:"varint,4,opt,name=idLanguage,proto3" json:"idLanguage,omitempty"`
	SiteHostName string `protobuf:"bytes,5,opt,name=siteHostName,proto3" json:"siteHostName,omitempty"`
	IdAdvertiser int32  `protobuf:"varint,6,opt,name=idAdvertiser,proto3" json:"idAdvertiser,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_files_main_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Event) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Event) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Event) GetIdLanguage() int32 {
	if x != nil {
		return x.IdLanguage
	}
	return 0
}

func (x *Event) GetSiteHostName() string {
	if x != nil {
		return x.SiteHostName
	}
	return ""
}

func (x *Event) GetIdAdvertiser() int32 {
	if x != nil {
		return x.IdAdvertiser
	}
	return 0
}

var File_proto_files_main_proto protoreflect.FileDescriptor

var file_proto_files_main_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xb1,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x69, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69,
	0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x64, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x64, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_main_proto_rawDescOnce sync.Once
	file_proto_files_main_proto_rawDescData = file_proto_files_main_proto_rawDesc
)

func file_proto_files_main_proto_rawDescGZIP() []byte {
	file_proto_files_main_proto_rawDescOnce.Do(func() {
		file_proto_files_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_main_proto_rawDescData)
	})
	return file_proto_files_main_proto_rawDescData
}

var file_proto_files_main_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_files_main_proto_goTypes = []any{
	(*Event)(nil), // 0: main.Event
}
var file_proto_files_main_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_files_main_proto_init() }
func file_proto_files_main_proto_init() {
	if File_proto_files_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_files_main_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_proto_files_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_main_proto_goTypes,
		DependencyIndexes: file_proto_files_main_proto_depIdxs,
		MessageInfos:      file_proto_files_main_proto_msgTypes,
	}.Build()
	File_proto_files_main_proto = out.File
	file_proto_files_main_proto_rawDesc = nil
	file_proto_files_main_proto_goTypes = nil
	file_proto_files_main_proto_depIdxs = nil
}
