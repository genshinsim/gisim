// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/model/db.proto

package model

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

type DBEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//basic info
	Key         string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CreateDate  uint64            `protobuf:"varint,2,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	RunDate     uint64            `protobuf:"varint,3,opt,name=run_date,json=runDate,proto3" json:"run_date,omitempty"`
	SimDuration *DescriptiveStats `protobuf:"bytes,4,opt,name=sim_duration,json=simDuration,proto3" json:"sim_duration,omitempty"`
	Config      string            `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	Hash        string            `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	//indexing data
	TotalDamage      *DescriptiveStats `protobuf:"bytes,7,opt,name=total_damage,json=totalDamage,proto3" json:"total_damage,omitempty"`
	CharNames        []string          `protobuf:"bytes,8,rep,name=char_names,json=charNames,proto3" json:"char_names,omitempty"`
	TargetCount      int32             `protobuf:"varint,9,opt,name=target_count,json=targetCount,proto3" json:"target_count,omitempty"`
	MeanDpsPerTarget float64           `protobuf:"fixed64,10,opt,name=mean_dps_per_target,json=meanDpsPerTarget,proto3" json:"mean_dps_per_target,omitempty"`
	//detailed results
	Team        []*Character                 `protobuf:"bytes,11,rep,name=team,proto3" json:"team,omitempty"`
	DpsByTarget map[string]*DescriptiveStats `protobuf:"bytes,12,rep,name=dps_by_target,json=dpsByTarget,proto3" json:"dps_by_target,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DBEntry) Reset() {
	*x = DBEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBEntry) ProtoMessage() {}

func (x *DBEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBEntry.ProtoReflect.Descriptor instead.
func (*DBEntry) Descriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{0}
}

func (x *DBEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DBEntry) GetCreateDate() uint64 {
	if x != nil {
		return x.CreateDate
	}
	return 0
}

func (x *DBEntry) GetRunDate() uint64 {
	if x != nil {
		return x.RunDate
	}
	return 0
}

func (x *DBEntry) GetSimDuration() *DescriptiveStats {
	if x != nil {
		return x.SimDuration
	}
	return nil
}

func (x *DBEntry) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *DBEntry) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *DBEntry) GetTotalDamage() *DescriptiveStats {
	if x != nil {
		return x.TotalDamage
	}
	return nil
}

func (x *DBEntry) GetCharNames() []string {
	if x != nil {
		return x.CharNames
	}
	return nil
}

func (x *DBEntry) GetTargetCount() int32 {
	if x != nil {
		return x.TargetCount
	}
	return 0
}

func (x *DBEntry) GetMeanDpsPerTarget() float64 {
	if x != nil {
		return x.MeanDpsPerTarget
	}
	return 0
}

func (x *DBEntry) GetTeam() []*Character {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *DBEntry) GetDpsByTarget() map[string]*DescriptiveStats {
	if x != nil {
		return x.DpsByTarget
	}
	return nil
}

var File_protos_model_db_proto protoreflect.FileDescriptor

var file_protos_model_db_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb0, 0x04, 0x0a, 0x07, 0x44, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73,
	0x69, 0x6d, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x70, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x6d, 0x65, 0x61, 0x6e, 0x44, 0x70, 0x73, 0x50, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x24, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x43, 0x0a, 0x0d, 0x64, 0x70, 0x73, 0x5f, 0x62, 0x79,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44, 0x70,
	0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x64, 0x70, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x57, 0x0a, 0x10, 0x44,
	0x70, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_db_proto_rawDescOnce sync.Once
	file_protos_model_db_proto_rawDescData = file_protos_model_db_proto_rawDesc
)

func file_protos_model_db_proto_rawDescGZIP() []byte {
	file_protos_model_db_proto_rawDescOnce.Do(func() {
		file_protos_model_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_db_proto_rawDescData)
	})
	return file_protos_model_db_proto_rawDescData
}

var file_protos_model_db_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_model_db_proto_goTypes = []interface{}{
	(*DBEntry)(nil),          // 0: model.DBEntry
	nil,                      // 1: model.DBEntry.DpsByTargetEntry
	(*DescriptiveStats)(nil), // 2: model.DescriptiveStats
	(*Character)(nil),        // 3: model.Character
}
var file_protos_model_db_proto_depIdxs = []int32{
	2, // 0: model.DBEntry.sim_duration:type_name -> model.DescriptiveStats
	2, // 1: model.DBEntry.total_damage:type_name -> model.DescriptiveStats
	3, // 2: model.DBEntry.team:type_name -> model.Character
	1, // 3: model.DBEntry.dps_by_target:type_name -> model.DBEntry.DpsByTargetEntry
	2, // 4: model.DBEntry.DpsByTargetEntry.value:type_name -> model.DescriptiveStats
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_model_db_proto_init() }
func file_protos_model_db_proto_init() {
	if File_protos_model_db_proto != nil {
		return
	}
	file_protos_model_result_proto_init()
	file_protos_model_sim_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBEntry); i {
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
			RawDescriptor: file_protos_model_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_db_proto_goTypes,
		DependencyIndexes: file_protos_model_db_proto_depIdxs,
		MessageInfos:      file_protos_model_db_proto_msgTypes,
	}.Build()
	File_protos_model_db_proto = out.File
	file_protos_model_db_proto_rawDesc = nil
	file_protos_model_db_proto_goTypes = nil
	file_protos_model_db_proto_depIdxs = nil
}
