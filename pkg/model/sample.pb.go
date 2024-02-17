// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: protos/model/sample.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildDate        string             `protobuf:"bytes,2,opt,name=build_date,proto3" json:"build_date,omitempty" bson:"build_date,omitempty"`
	SimVersion       *string            `protobuf:"bytes,1,opt,name=sim_version,proto3,oneof" json:"sim_version,omitempty" bson:"sim_version,omitempty"`
	Modified         *bool              `protobuf:"varint,3,opt,name=modified,proto3,oneof" json:"modified,omitempty" bson:"modified,omitempty"`
	Config           string             `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	InitialCharacter string             `protobuf:"bytes,5,opt,name=initial_character,proto3" json:"initial_character,omitempty" bson:"initial_character,omitempty"`
	CharacterDetails []*Character       `protobuf:"bytes,6,rep,name=character_details,proto3" json:"character_details,omitempty" bson:"character_details,omitempty"`
	TargetDetails    []*Enemy           `protobuf:"bytes,7,rep,name=target_details,proto3" json:"target_details,omitempty" bson:"target_details,omitempty"`
	Seed             string             `protobuf:"bytes,8,opt,name=seed,proto3" json:"seed,omitempty" bson:"seed,omitempty"`
	Logs             []*structpb.Struct `protobuf:"bytes,9,rep,name=logs,proto3" json:"logs,omitempty" bson:"logs,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_protos_model_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

func (x *Sample) GetSimVersion() string {
	if x != nil && x.SimVersion != nil {
		return *x.SimVersion
	}
	return ""
}

func (x *Sample) GetModified() bool {
	if x != nil && x.Modified != nil {
		return *x.Modified
	}
	return false
}

func (x *Sample) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Sample) GetInitialCharacter() string {
	if x != nil {
		return x.InitialCharacter
	}
	return ""
}

func (x *Sample) GetCharacterDetails() []*Character {
	if x != nil {
		return x.CharacterDetails
	}
	return nil
}

func (x *Sample) GetTargetDetails() []*Enemy {
	if x != nil {
		return x.TargetDetails
	}
	return nil
}

func (x *Sample) GetSeed() string {
	if x != nil {
		return x.Seed
	}
	return ""
}

func (x *Sample) GetLogs() []*structpb.Struct {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_protos_model_sample_proto protoreflect.FileDescriptor

var file_protos_model_sample_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x06, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x69, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x11,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x34, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x69, 0x6d,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2f, 0x67,
	0x63, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_sample_proto_rawDescOnce sync.Once
	file_protos_model_sample_proto_rawDescData = file_protos_model_sample_proto_rawDesc
)

func file_protos_model_sample_proto_rawDescGZIP() []byte {
	file_protos_model_sample_proto_rawDescOnce.Do(func() {
		file_protos_model_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_sample_proto_rawDescData)
	})
	return file_protos_model_sample_proto_rawDescData
}

var file_protos_model_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_model_sample_proto_goTypes = []interface{}{
	(*Sample)(nil),          // 0: model.Sample
	(*Character)(nil),       // 1: model.Character
	(*Enemy)(nil),           // 2: model.Enemy
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
}
var file_protos_model_sample_proto_depIdxs = []int32{
	1, // 0: model.Sample.character_details:type_name -> model.Character
	2, // 1: model.Sample.target_details:type_name -> model.Enemy
	3, // 2: model.Sample.logs:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_model_sample_proto_init() }
func file_protos_model_sample_proto_init() {
	if File_protos_model_sample_proto != nil {
		return
	}
	file_protos_model_sim_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
	file_protos_model_sample_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_sample_proto_goTypes,
		DependencyIndexes: file_protos_model_sample_proto_depIdxs,
		MessageInfos:      file_protos_model_sample_proto_msgTypes,
	}.Build()
	File_protos_model_sample_proto = out.File
	file_protos_model_sample_proto_rawDesc = nil
	file_protos_model_sample_proto_goTypes = nil
	file_protos_model_sample_proto_depIdxs = nil
}
