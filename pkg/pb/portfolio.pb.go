// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/portfolio.proto

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

type SkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SkillRequest) Reset() {
	*x = SkillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_portfolio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillRequest) ProtoMessage() {}

func (x *SkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_portfolio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillRequest.ProtoReflect.Descriptor instead.
func (*SkillRequest) Descriptor() ([]byte, []int) {
	return file_pb_portfolio_proto_rawDescGZIP(), []int{0}
}

// The response message containing the skills
type SkillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Skill       []*Skill `protobuf:"bytes,2,rep,name=skill,proto3" json:"skill,omitempty"`
}

func (x *SkillResponse) Reset() {
	*x = SkillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_portfolio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillResponse) ProtoMessage() {}

func (x *SkillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_portfolio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillResponse.ProtoReflect.Descriptor instead.
func (*SkillResponse) Descriptor() ([]byte, []int) {
	return file_pb_portfolio_proto_rawDescGZIP(), []int{1}
}

func (x *SkillResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SkillResponse) GetSkill() []*Skill {
	if x != nil {
		return x.Skill
	}
	return nil
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Proficient int32  `protobuf:"varint,2,opt,name=proficient,proto3" json:"proficient,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_portfolio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_pb_portfolio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_pb_portfolio_proto_rawDescGZIP(), []int{2}
}

func (x *Skill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Skill) GetProficient() int32 {
	if x != nil {
		return x.Proficient
	}
	return 0
}

var File_pb_portfolio_proto protoreflect.FileDescriptor

var file_pb_portfolio_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x59, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x3b, 0x0a, 0x05, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x4e, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x6e, 0x69, 0x31, 0x39, 0x39, 0x38, 0x2f,
	0x73, 0x68, 0x61, 0x6e, 0x69, 0x31, 0x39, 0x39, 0x38, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_portfolio_proto_rawDescOnce sync.Once
	file_pb_portfolio_proto_rawDescData = file_pb_portfolio_proto_rawDesc
)

func file_pb_portfolio_proto_rawDescGZIP() []byte {
	file_pb_portfolio_proto_rawDescOnce.Do(func() {
		file_pb_portfolio_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_portfolio_proto_rawDescData)
	})
	return file_pb_portfolio_proto_rawDescData
}

var file_pb_portfolio_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_portfolio_proto_goTypes = []interface{}{
	(*SkillRequest)(nil),  // 0: portfolio.SkillRequest
	(*SkillResponse)(nil), // 1: portfolio.SkillResponse
	(*Skill)(nil),         // 2: portfolio.Skill
}
var file_pb_portfolio_proto_depIdxs = []int32{
	2, // 0: portfolio.SkillResponse.skill:type_name -> portfolio.Skill
	0, // 1: portfolio.portfolio.ListSkills:input_type -> portfolio.SkillRequest
	1, // 2: portfolio.portfolio.ListSkills:output_type -> portfolio.SkillResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_portfolio_proto_init() }
func file_pb_portfolio_proto_init() {
	if File_pb_portfolio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_portfolio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillRequest); i {
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
		file_pb_portfolio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillResponse); i {
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
		file_pb_portfolio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
			RawDescriptor: file_pb_portfolio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_portfolio_proto_goTypes,
		DependencyIndexes: file_pb_portfolio_proto_depIdxs,
		MessageInfos:      file_pb_portfolio_proto_msgTypes,
	}.Build()
	File_pb_portfolio_proto = out.File
	file_pb_portfolio_proto_rawDesc = nil
	file_pb_portfolio_proto_goTypes = nil
	file_pb_portfolio_proto_depIdxs = nil
}