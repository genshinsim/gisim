// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/model/db.proto

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

type ComputeWorkSource int32

const (
	ComputeWorkSource_InvalidWork    ComputeWorkSource = 0
	ComputeWorkSource_DBWork         ComputeWorkSource = 1
	ComputeWorkSource_SubmissionWork ComputeWorkSource = 2
)

// Enum value maps for ComputeWorkSource.
var (
	ComputeWorkSource_name = map[int32]string{
		0: "InvalidWork",
		1: "DBWork",
		2: "SubmissionWork",
	}
	ComputeWorkSource_value = map[string]int32{
		"InvalidWork":    0,
		"DBWork":         1,
		"SubmissionWork": 2,
	}
)

func (x ComputeWorkSource) Enum() *ComputeWorkSource {
	p := new(ComputeWorkSource)
	*p = x
	return p
}

func (x ComputeWorkSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComputeWorkSource) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_model_db_proto_enumTypes[0].Descriptor()
}

func (ComputeWorkSource) Type() protoreflect.EnumType {
	return &file_protos_model_db_proto_enumTypes[0]
}

func (x ComputeWorkSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComputeWorkSource.Descriptor instead.
func (ComputeWorkSource) EnumDescriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{0}
}

type DBTag int32

const (
	DBTag_DB_TAG_INVALID DBTag = 0
	DBTag_DB_TAG_GCSIM   DBTag = 1
	DBTag_DB_TAG_TESTING DBTag = 2
)

// Enum value maps for DBTag.
var (
	DBTag_name = map[int32]string{
		0: "DB_TAG_INVALID",
		1: "DB_TAG_GCSIM",
		2: "DB_TAG_TESTING",
	}
	DBTag_value = map[string]int32{
		"DB_TAG_INVALID": 0,
		"DB_TAG_GCSIM":   1,
		"DB_TAG_TESTING": 2,
	}
)

func (x DBTag) Enum() *DBTag {
	p := new(DBTag)
	*p = x
	return p
}

func (x DBTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBTag) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_model_db_proto_enumTypes[1].Descriptor()
}

func (DBTag) Type() protoreflect.EnumType {
	return &file_protos_model_db_proto_enumTypes[1]
}

func (x DBTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBTag.Descriptor instead.
func (DBTag) EnumDescriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{1}
}

type DBEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// basic info
	Id          string            `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	ShareKey    string            `protobuf:"bytes,2,opt,name=share_key,proto3" json:"share_key,omitempty" bson:"share_key,omitempty"`
	CreateDate  uint64            `protobuf:"varint,3,opt,name=create_date,proto3" json:"create_date,omitempty" bson:"create_date,omitempty"`
	RunDate     uint64            `protobuf:"varint,4,opt,name=run_date,proto3" json:"run_date,omitempty" bson:"run_date,omitempty"`
	SimDuration *DescriptiveStats `protobuf:"bytes,5,opt,name=sim_duration,proto3" json:"sim_duration,omitempty" bson:"sim_duration,omitempty"`
	Config      string            `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	Hash        string            `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty" bson:"hash,omitempty"`
	Mode        SimMode           `protobuf:"varint,8,opt,name=mode,proto3,enum=model.SimMode" json:"mode,omitempty" bson:"mode,omitempty"`
	// indexing data
	TotalDamage      *DescriptiveStats `protobuf:"bytes,9,opt,name=total_damage,proto3" json:"total_damage,omitempty" bson:"total_damage,omitempty"`
	CharNames        []string          `protobuf:"bytes,10,rep,name=char_names,proto3" json:"char_names,omitempty" bson:"char_names,omitempty"`
	TargetCount      int32             `protobuf:"varint,11,opt,name=target_count,proto3" json:"target_count,omitempty" bson:"target_count,omitempty"`
	MeanDpsPerTarget float64           `protobuf:"fixed64,12,opt,name=mean_dps_per_target,proto3" json:"mean_dps_per_target,omitempty" bson:"mean_dps_per_target,omitempty"`
	// detailed results
	Team        []*Character                 `protobuf:"bytes,13,rep,name=team,proto3" json:"team,omitempty" bson:"team,omitempty"`
	DpsByTarget map[string]*DescriptiveStats `protobuf:"bytes,14,rep,name=dps_by_target,proto3" json:"dps_by_target,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"dps_by_target,omitempty"`
	// other stuff
	Description  string  `protobuf:"bytes,15,opt,name=description,proto3" json:"description,omitempty" bson:"description,omitempty"`
	AcceptedTags []DBTag `protobuf:"varint,16,rep,packed,name=accepted_tags,proto3,enum=model.DBTag" json:"accepted_tags,omitempty" bson:"accepted_tags,omitempty"`
	RejectedTags []DBTag `protobuf:"varint,17,rep,packed,name=rejected_tags,proto3,enum=model.DBTag" json:"rejected_tags,omitempty" bson:"rejected_tags,omitempty"`
	IsDbValid    bool    `protobuf:"varint,18,opt,name=is_db_valid,proto3" json:"is_db_valid,omitempty" bson:"is_db_valid,omitempty"`
	Submitter    string  `protobuf:"bytes,19,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
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

func (x *DBEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DBEntry) GetShareKey() string {
	if x != nil {
		return x.ShareKey
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

func (x *DBEntry) GetMode() SimMode {
	if x != nil {
		return x.Mode
	}
	return SimMode_DURATION_MODE
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

func (x *DBEntry) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DBEntry) GetAcceptedTags() []DBTag {
	if x != nil {
		return x.AcceptedTags
	}
	return nil
}

func (x *DBEntry) GetRejectedTags() []DBTag {
	if x != nil {
		return x.RejectedTags
	}
	return nil
}

func (x *DBEntry) GetIsDbValid() bool {
	if x != nil {
		return x.IsDbValid
	}
	return false
}

func (x *DBEntry) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type DBEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DBEntry `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
}

func (x *DBEntries) Reset() {
	*x = DBEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBEntries) ProtoMessage() {}

func (x *DBEntries) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBEntries.ProtoReflect.Descriptor instead.
func (*DBEntries) Descriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{1}
}

func (x *DBEntries) GetData() []*DBEntry {
	if x != nil {
		return x.Data
	}
	return nil
}

type DBQueryOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   *structpb.Struct `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty" bson:"query,omitempty"`
	Sort    *structpb.Struct `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty" bson:"sort,omitempty"`
	Project *structpb.Struct `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty" bson:"project,omitempty"`
	Skip    int64            `protobuf:"varint,4,opt,name=skip,proto3" json:"skip,omitempty" bson:"skip,omitempty"`
	Limit   int64            `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty" bson:"limit,omitempty"`
}

func (x *DBQueryOpt) Reset() {
	*x = DBQueryOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBQueryOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBQueryOpt) ProtoMessage() {}

func (x *DBQueryOpt) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBQueryOpt.ProtoReflect.Descriptor instead.
func (*DBQueryOpt) Descriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{2}
}

func (x *DBQueryOpt) GetQuery() *structpb.Struct {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DBQueryOpt) GetSort() *structpb.Struct {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *DBQueryOpt) GetProject() *structpb.Struct {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *DBQueryOpt) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *DBQueryOpt) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // auto generated id for this submission
	Config      string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	Submitter   string `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"` //submitter discord id
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" bson:"description,omitempty"`
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{3}
}

func (x *Submission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Submission) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *Submission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ComputeWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	Config     string            `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	Source     ComputeWorkSource `protobuf:"varint,3,opt,name=source,proto3,enum=model.ComputeWorkSource" json:"source,omitempty" bson:"source,omitempty"`
	Iterations int32             `protobuf:"varint,4,opt,name=iterations,proto3" json:"iterations,omitempty" bson:"iterations,omitempty"`
}

func (x *ComputeWork) Reset() {
	*x = ComputeWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeWork) ProtoMessage() {}

func (x *ComputeWork) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeWork.ProtoReflect.Descriptor instead.
func (*ComputeWork) Descriptor() ([]byte, []int) {
	return file_protos_model_db_proto_rawDescGZIP(), []int{4}
}

func (x *ComputeWork) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComputeWork) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *ComputeWork) GetSource() ComputeWorkSource {
	if x != nil {
		return x.Source
	}
	return ComputeWorkSource_InvalidWork
}

func (x *ComputeWork) GetIterations() int32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

var File_protos_model_db_proto protoreflect.FileDescriptor

var file_protos_model_db_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x06, 0x0a, 0x07, 0x44, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x75, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x72, 0x75, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x64,
	0x70, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x70, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x45,
	0x0a, 0x0d, 0x64, 0x70, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44, 0x70, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x64, 0x70, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x67, 0x52, 0x0d, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x0d, 0x72,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x11, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x67,
	0x52, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x62, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x62, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x1a,
	0x57, 0x0a, 0x10, 0x44, 0x70, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x09, 0x44, 0x42, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x0a, 0x44, 0x42,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x75, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2a, 0x44, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x42, 0x57,
	0x6f, 0x72, 0x6b, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x10, 0x02, 0x2a, 0x41, 0x0a, 0x05, 0x44, 0x42, 0x54,
	0x61, 0x67, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x42, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x42, 0x5f, 0x54, 0x41, 0x47,
	0x5f, 0x47, 0x43, 0x53, 0x49, 0x4d, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x42, 0x5f, 0x54,
	0x41, 0x47, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x68,
	0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2f, 0x67, 0x63, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_model_db_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_model_db_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_model_db_proto_goTypes = []interface{}{
	(ComputeWorkSource)(0),   // 0: model.ComputeWorkSource
	(DBTag)(0),               // 1: model.DBTag
	(*DBEntry)(nil),          // 2: model.DBEntry
	(*DBEntries)(nil),        // 3: model.DBEntries
	(*DBQueryOpt)(nil),       // 4: model.DBQueryOpt
	(*Submission)(nil),       // 5: model.Submission
	(*ComputeWork)(nil),      // 6: model.ComputeWork
	nil,                      // 7: model.DBEntry.DpsByTargetEntry
	(*DescriptiveStats)(nil), // 8: model.DescriptiveStats
	(SimMode)(0),             // 9: model.SimMode
	(*Character)(nil),        // 10: model.Character
	(*structpb.Struct)(nil),  // 11: google.protobuf.Struct
}
var file_protos_model_db_proto_depIdxs = []int32{
	8,  // 0: model.DBEntry.sim_duration:type_name -> model.DescriptiveStats
	9,  // 1: model.DBEntry.mode:type_name -> model.SimMode
	8,  // 2: model.DBEntry.total_damage:type_name -> model.DescriptiveStats
	10, // 3: model.DBEntry.team:type_name -> model.Character
	7,  // 4: model.DBEntry.dps_by_target:type_name -> model.DBEntry.DpsByTargetEntry
	1,  // 5: model.DBEntry.accepted_tags:type_name -> model.DBTag
	1,  // 6: model.DBEntry.rejected_tags:type_name -> model.DBTag
	2,  // 7: model.DBEntries.data:type_name -> model.DBEntry
	11, // 8: model.DBQueryOpt.query:type_name -> google.protobuf.Struct
	11, // 9: model.DBQueryOpt.sort:type_name -> google.protobuf.Struct
	11, // 10: model.DBQueryOpt.project:type_name -> google.protobuf.Struct
	0,  // 11: model.ComputeWork.source:type_name -> model.ComputeWorkSource
	8,  // 12: model.DBEntry.DpsByTargetEntry.value:type_name -> model.DescriptiveStats
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
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
		file_protos_model_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBEntries); i {
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
		file_protos_model_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBQueryOpt); i {
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
		file_protos_model_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_protos_model_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeWork); i {
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
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_db_proto_goTypes,
		DependencyIndexes: file_protos_model_db_proto_depIdxs,
		EnumInfos:         file_protos_model_db_proto_enumTypes,
		MessageInfos:      file_protos_model_db_proto_msgTypes,
	}.Build()
	File_protos_model_db_proto = out.File
	file_protos_model_db_proto_rawDesc = nil
	file_protos_model_db_proto_goTypes = nil
	file_protos_model_db_proto_depIdxs = nil
}
