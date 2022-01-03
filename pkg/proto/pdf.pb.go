// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pdf/proto/pdf.proto

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

type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie string `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	Data  string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Time  string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Row   int32  `protobuf:"varint,5,opt,name=row,proto3" json:"row,omitempty"`
	Place int32  `protobuf:"varint,6,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdf_proto_pdf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pdf_proto_pdf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_pdf_proto_pdf_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRequest) GetMovie() string {
	if x != nil {
		return x.Movie
	}
	return ""
}

func (x *GenerateRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GenerateRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *GenerateRequest) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *GenerateRequest) GetPlace() int32 {
	if x != nil {
		return x.Place
	}
	return 0
}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdf_proto_pdf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pdf_proto_pdf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_pdf_proto_pdf_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponse) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

var File_pdf_proto_pdf_proto protoreflect.FileDescriptor

var file_pdf_proto_pdf_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x64, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x64, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x64, 0x66, 0x22, 0x77, 0x0a, 0x0f, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x32, 0x49, 0x0a, 0x0c, 0x50,
	0x44, 0x46, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x64, 0x66, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x64, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pdf_proto_pdf_proto_rawDescOnce sync.Once
	file_pdf_proto_pdf_proto_rawDescData = file_pdf_proto_pdf_proto_rawDesc
)

func file_pdf_proto_pdf_proto_rawDescGZIP() []byte {
	file_pdf_proto_pdf_proto_rawDescOnce.Do(func() {
		file_pdf_proto_pdf_proto_rawDescData = protoimpl.X.CompressGZIP(file_pdf_proto_pdf_proto_rawDescData)
	})
	return file_pdf_proto_pdf_proto_rawDescData
}

var file_pdf_proto_pdf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pdf_proto_pdf_proto_goTypes = []interface{}{
	(*GenerateRequest)(nil),  // 0: pdf.GenerateRequest
	(*GenerateResponse)(nil), // 1: pdf.GenerateResponse
}
var file_pdf_proto_pdf_proto_depIdxs = []int32{
	0, // 0: pdf.PDFGenerator.Generate:input_type -> pdf.GenerateRequest
	1, // 1: pdf.PDFGenerator.Generate:output_type -> pdf.GenerateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pdf_proto_pdf_proto_init() }
func file_pdf_proto_pdf_proto_init() {
	if File_pdf_proto_pdf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pdf_proto_pdf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_pdf_proto_pdf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
			RawDescriptor: file_pdf_proto_pdf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pdf_proto_pdf_proto_goTypes,
		DependencyIndexes: file_pdf_proto_pdf_proto_depIdxs,
		MessageInfos:      file_pdf_proto_pdf_proto_msgTypes,
	}.Build()
	File_pdf_proto_pdf_proto = out.File
	file_pdf_proto_pdf_proto_rawDesc = nil
	file_pdf_proto_pdf_proto_goTypes = nil
	file_pdf_proto_pdf_proto_depIdxs = nil
}
