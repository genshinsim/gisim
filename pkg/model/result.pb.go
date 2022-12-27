// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/model/result.proto

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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int64 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int64 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() int64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int64 {
	if x != nil {
		return x.Minor
	}
	return 0
}

type SimulationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//metadata
	SchemaVersion *Version `protobuf:"bytes,1,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	SimVersion    string   `protobuf:"bytes,2,opt,name=sim_version,json=simVersion,proto3" json:"sim_version,omitempty"`
	BuildDate     string   `protobuf:"bytes,3,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	Modified      bool     `protobuf:"varint,4,opt,name=modified,proto3" json:"modified,omitempty"`
	//charater and enemey data
	InitialCharacter  string                `protobuf:"bytes,5,opt,name=initial_character,json=initialCharacter,proto3" json:"initial_character,omitempty"`
	CharacterDetails  []*Character          `protobuf:"bytes,6,rep,name=character_details,json=characterDetails,proto3" json:"character_details,omitempty"`
	TargetDetails     []*Enemy              `protobuf:"bytes,7,rep,name=target_details,json=targetDetails,proto3" json:"target_details,omitempty"`
	SimulatorSettings *SimulatorSettings    `protobuf:"bytes,8,opt,name=simulator_settings,json=simulatorSettings,proto3" json:"simulator_settings,omitempty"`
	EnergySettings    *EnergySettings       `protobuf:"bytes,9,opt,name=energy_settings,json=energySettings,proto3" json:"energy_settings,omitempty"`
	Config            string                `protobuf:"bytes,10,opt,name=config,proto3" json:"config,omitempty"`
	SampleSeed        string                `protobuf:"bytes,11,opt,name=sample_seed,json=sampleSeed,proto3" json:"sample_seed,omitempty"`
	Statistics        *SimulationStatistics `protobuf:"bytes,12,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *SimulationResult) Reset() {
	*x = SimulationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationResult) ProtoMessage() {}

func (x *SimulationResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationResult.ProtoReflect.Descriptor instead.
func (*SimulationResult) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{1}
}

func (x *SimulationResult) GetSchemaVersion() *Version {
	if x != nil {
		return x.SchemaVersion
	}
	return nil
}

func (x *SimulationResult) GetSimVersion() string {
	if x != nil {
		return x.SimVersion
	}
	return ""
}

func (x *SimulationResult) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

func (x *SimulationResult) GetModified() bool {
	if x != nil {
		return x.Modified
	}
	return false
}

func (x *SimulationResult) GetInitialCharacter() string {
	if x != nil {
		return x.InitialCharacter
	}
	return ""
}

func (x *SimulationResult) GetCharacterDetails() []*Character {
	if x != nil {
		return x.CharacterDetails
	}
	return nil
}

func (x *SimulationResult) GetTargetDetails() []*Enemy {
	if x != nil {
		return x.TargetDetails
	}
	return nil
}

func (x *SimulationResult) GetSimulatorSettings() *SimulatorSettings {
	if x != nil {
		return x.SimulatorSettings
	}
	return nil
}

func (x *SimulationResult) GetEnergySettings() *EnergySettings {
	if x != nil {
		return x.EnergySettings
	}
	return nil
}

func (x *SimulationResult) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *SimulationResult) GetSampleSeed() string {
	if x != nil {
		return x.SampleSeed
	}
	return ""
}

func (x *SimulationResult) GetStatistics() *SimulationStatistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

type SimulationStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//metadata
	MinSeed    string  `protobuf:"bytes,1,opt,name=min_seed,json=minSeed,proto3" json:"min_seed,omitempty"`
	MaxSeed    string  `protobuf:"bytes,2,opt,name=max_seed,json=maxSeed,proto3" json:"max_seed,omitempty"`
	P25Seed    string  `protobuf:"bytes,3,opt,name=p25_seed,json=p25Seed,proto3" json:"p25_seed,omitempty"`
	P50Seed    string  `protobuf:"bytes,4,opt,name=p50_seed,json=p50Seed,proto3" json:"p50_seed,omitempty"`
	P75Seed    string  `protobuf:"bytes,5,opt,name=p75_seed,json=p75Seed,proto3" json:"p75_seed,omitempty"`
	Runtime    float64 `protobuf:"fixed64,6,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Iterations int64   `protobuf:"varint,7,opt,name=iterations,proto3" json:"iterations,omitempty"`
	// global overview (global/no group by)
	Duration      *OverviewStats   `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
	DPS           *OverviewStats   `protobuf:"bytes,9,opt,name=DPS,json=dps,proto3" json:"DPS,omitempty"`
	RPS           *OverviewStats   `protobuf:"bytes,10,opt,name=RPS,json=rps,proto3" json:"RPS,omitempty"`
	EPS           *OverviewStats   `protobuf:"bytes,11,opt,name=EPS,json=eps,proto3" json:"EPS,omitempty"`
	HPS           *OverviewStats   `protobuf:"bytes,12,opt,name=HPS,json=hps,proto3" json:"HPS,omitempty"`
	SPS           *OverviewStats   `protobuf:"bytes,13,opt,name=SPS,json=sps,proto3" json:"SPS,omitempty"`
	TotalDamage   *OverviewStats   `protobuf:"bytes,14,opt,name=total_damage,json=totalDamage,proto3" json:"total_damage,omitempty"`
	Warnings      *Warnings        `protobuf:"bytes,15,opt,name=warnings,proto3" json:"warnings,omitempty"`
	FailedActions []*FailedActions `protobuf:"bytes,16,rep,name=failed_actions,json=failedActions,proto3" json:"failed_actions,omitempty"`
}

func (x *SimulationStatistics) Reset() {
	*x = SimulationStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationStatistics) ProtoMessage() {}

func (x *SimulationStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationStatistics.ProtoReflect.Descriptor instead.
func (*SimulationStatistics) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{2}
}

func (x *SimulationStatistics) GetMinSeed() string {
	if x != nil {
		return x.MinSeed
	}
	return ""
}

func (x *SimulationStatistics) GetMaxSeed() string {
	if x != nil {
		return x.MaxSeed
	}
	return ""
}

func (x *SimulationStatistics) GetP25Seed() string {
	if x != nil {
		return x.P25Seed
	}
	return ""
}

func (x *SimulationStatistics) GetP50Seed() string {
	if x != nil {
		return x.P50Seed
	}
	return ""
}

func (x *SimulationStatistics) GetP75Seed() string {
	if x != nil {
		return x.P75Seed
	}
	return ""
}

func (x *SimulationStatistics) GetRuntime() float64 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *SimulationStatistics) GetIterations() int64 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *SimulationStatistics) GetDuration() *OverviewStats {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *SimulationStatistics) GetDPS() *OverviewStats {
	if x != nil {
		return x.DPS
	}
	return nil
}

func (x *SimulationStatistics) GetRPS() *OverviewStats {
	if x != nil {
		return x.RPS
	}
	return nil
}

func (x *SimulationStatistics) GetEPS() *OverviewStats {
	if x != nil {
		return x.EPS
	}
	return nil
}

func (x *SimulationStatistics) GetHPS() *OverviewStats {
	if x != nil {
		return x.HPS
	}
	return nil
}

func (x *SimulationStatistics) GetSPS() *OverviewStats {
	if x != nil {
		return x.SPS
	}
	return nil
}

func (x *SimulationStatistics) GetTotalDamage() *OverviewStats {
	if x != nil {
		return x.TotalDamage
	}
	return nil
}

func (x *SimulationStatistics) GetWarnings() *Warnings {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *SimulationStatistics) GetFailedActions() []*FailedActions {
	if x != nil {
		return x.FailedActions
	}
	return nil
}

type OverviewStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min  float64  `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max  float64  `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Mean float64  `protobuf:"fixed64,3,opt,name=mean,proto3" json:"mean,omitempty"`
	SD   float64  `protobuf:"fixed64,4,opt,name=SD,json=sd,proto3" json:"SD,omitempty"`
	Q1   float64  `protobuf:"fixed64,5,opt,name=q1,proto3" json:"q1,omitempty"`
	Q2   float64  `protobuf:"fixed64,6,opt,name=q2,proto3" json:"q2,omitempty"`
	Q3   float64  `protobuf:"fixed64,7,opt,name=q3,proto3" json:"q3,omitempty"`
	Hist []uint64 `protobuf:"varint,8,rep,packed,name=hist,json=histogram,proto3" json:"hist,omitempty"`
}

func (x *OverviewStats) Reset() {
	*x = OverviewStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverviewStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverviewStats) ProtoMessage() {}

func (x *OverviewStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverviewStats.ProtoReflect.Descriptor instead.
func (*OverviewStats) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{3}
}

func (x *OverviewStats) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *OverviewStats) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *OverviewStats) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *OverviewStats) GetSD() float64 {
	if x != nil {
		return x.SD
	}
	return 0
}

func (x *OverviewStats) GetQ1() float64 {
	if x != nil {
		return x.Q1
	}
	return 0
}

func (x *OverviewStats) GetQ2() float64 {
	if x != nil {
		return x.Q2
	}
	return 0
}

func (x *OverviewStats) GetQ3() float64 {
	if x != nil {
		return x.Q3
	}
	return 0
}

func (x *OverviewStats) GetHist() []uint64 {
	if x != nil {
		return x.Hist
	}
	return nil
}

type DescriptiveStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min  float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max  float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Mean float64 `protobuf:"fixed64,3,opt,name=mean,proto3" json:"mean,omitempty"`
	SD   float64 `protobuf:"fixed64,4,opt,name=SD,json=sd,proto3" json:"SD,omitempty"`
}

func (x *DescriptiveStats) Reset() {
	*x = DescriptiveStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptiveStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptiveStats) ProtoMessage() {}

func (x *DescriptiveStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptiveStats.ProtoReflect.Descriptor instead.
func (*DescriptiveStats) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{4}
}

func (x *DescriptiveStats) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DescriptiveStats) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *DescriptiveStats) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *DescriptiveStats) GetSD() float64 {
	if x != nil {
		return x.SD
	}
	return 0
}

type Warnings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetOverlap       bool `protobuf:"varint,1,opt,name=target_overlap,json=targetOverlap,proto3" json:"target_overlap,omitempty"`
	InsufficientEnergy  bool `protobuf:"varint,2,opt,name=insufficient_energy,json=insufficientEnergy,proto3" json:"insufficient_energy,omitempty"`
	InsufficientStamina bool `protobuf:"varint,3,opt,name=insufficient_stamina,json=insufficientStamina,proto3" json:"insufficient_stamina,omitempty"`
	SwapCd              bool `protobuf:"varint,4,opt,name=swap_cd,json=swapCd,proto3" json:"swap_cd,omitempty"`
	SkillCd             bool `protobuf:"varint,5,opt,name=skill_cd,json=skillCd,proto3" json:"skill_cd,omitempty"`
}

func (x *Warnings) Reset() {
	*x = Warnings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warnings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warnings) ProtoMessage() {}

func (x *Warnings) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warnings.ProtoReflect.Descriptor instead.
func (*Warnings) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{5}
}

func (x *Warnings) GetTargetOverlap() bool {
	if x != nil {
		return x.TargetOverlap
	}
	return false
}

func (x *Warnings) GetInsufficientEnergy() bool {
	if x != nil {
		return x.InsufficientEnergy
	}
	return false
}

func (x *Warnings) GetInsufficientStamina() bool {
	if x != nil {
		return x.InsufficientStamina
	}
	return false
}

func (x *Warnings) GetSwapCd() bool {
	if x != nil {
		return x.SwapCd
	}
	return false
}

func (x *Warnings) GetSkillCd() bool {
	if x != nil {
		return x.SkillCd
	}
	return false
}

type FailedActions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsufficientEnergy  *DescriptiveStats `protobuf:"bytes,1,opt,name=insufficient_energy,json=insufficientEnergy,proto3" json:"insufficient_energy,omitempty"`
	InsufficientStamina *DescriptiveStats `protobuf:"bytes,2,opt,name=insufficient_stamina,json=insufficientStamina,proto3" json:"insufficient_stamina,omitempty"`
	SwapCd              *DescriptiveStats `protobuf:"bytes,3,opt,name=swap_cd,json=swapCd,proto3" json:"swap_cd,omitempty"`
	SkillCd             *DescriptiveStats `protobuf:"bytes,4,opt,name=skill_cd,json=skillCd,proto3" json:"skill_cd,omitempty"`
}

func (x *FailedActions) Reset() {
	*x = FailedActions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedActions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedActions) ProtoMessage() {}

func (x *FailedActions) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_result_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedActions.ProtoReflect.Descriptor instead.
func (*FailedActions) Descriptor() ([]byte, []int) {
	return file_protos_model_result_proto_rawDescGZIP(), []int{6}
}

func (x *FailedActions) GetInsufficientEnergy() *DescriptiveStats {
	if x != nil {
		return x.InsufficientEnergy
	}
	return nil
}

func (x *FailedActions) GetInsufficientStamina() *DescriptiveStats {
	if x != nil {
		return x.InsufficientStamina
	}
	return nil
}

func (x *FailedActions) GetSwapCd() *DescriptiveStats {
	if x != nil {
		return x.SwapCd
	}
	return nil
}

func (x *FailedActions) GetSkillCd() *DescriptiveStats {
	if x != nil {
		return x.SkillCd
	}
	return nil
}

var File_protos_model_result_proto protoreflect.FileDescriptor

var file_protos_model_result_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x22, 0xc5, 0x04, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x69, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x10, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x0d, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x11, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0xf4, 0x04, 0x0a, 0x14, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x32, 0x35, 0x5f,
	0x73, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x32, 0x35, 0x53,
	0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x35, 0x30, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x35, 0x30, 0x53, 0x65, 0x65, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x37, 0x35, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x37, 0x35, 0x53, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x76,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x03, 0x44, 0x50, 0x53, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x64, 0x70, 0x73, 0x12, 0x26, 0x0a,
	0x03, 0x52, 0x50, 0x53, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x03, 0x72, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x03, 0x45, 0x50, 0x53, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x65, 0x70, 0x73, 0x12, 0x26, 0x0a,
	0x03, 0x48, 0x50, 0x53, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x03, 0x68, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x03, 0x53, 0x50, 0x53, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x73, 0x70, 0x73, 0x12, 0x37, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x73, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x31,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x71, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x32,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x71, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x33,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x71, 0x33, 0x12, 0x17, 0x0a, 0x04, 0x68, 0x69,
	0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x22, 0x5a, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x53, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x73, 0x64, 0x22,
	0xc9, 0x01, 0x0a, 0x08, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x70, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x77, 0x61, 0x70, 0x5f,
	0x63, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x77, 0x61, 0x70, 0x43, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x64, 0x22, 0x8b, 0x02, 0x0a, 0x0d,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a,
	0x13, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x4a, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x13,
	0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x6d,
	0x69, 0x6e, 0x61, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x73,
	0x77, 0x61, 0x70, 0x43, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x63,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_result_proto_rawDescOnce sync.Once
	file_protos_model_result_proto_rawDescData = file_protos_model_result_proto_rawDesc
)

func file_protos_model_result_proto_rawDescGZIP() []byte {
	file_protos_model_result_proto_rawDescOnce.Do(func() {
		file_protos_model_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_result_proto_rawDescData)
	})
	return file_protos_model_result_proto_rawDescData
}

var file_protos_model_result_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_model_result_proto_goTypes = []interface{}{
	(*Version)(nil),              // 0: model.Version
	(*SimulationResult)(nil),     // 1: model.SimulationResult
	(*SimulationStatistics)(nil), // 2: model.SimulationStatistics
	(*OverviewStats)(nil),        // 3: model.OverviewStats
	(*DescriptiveStats)(nil),     // 4: model.DescriptiveStats
	(*Warnings)(nil),             // 5: model.Warnings
	(*FailedActions)(nil),        // 6: model.FailedActions
	(*Character)(nil),            // 7: model.Character
	(*Enemy)(nil),                // 8: model.Enemy
	(*SimulatorSettings)(nil),    // 9: model.SimulatorSettings
	(*EnergySettings)(nil),       // 10: model.EnergySettings
}
var file_protos_model_result_proto_depIdxs = []int32{
	0,  // 0: model.SimulationResult.schema_version:type_name -> model.Version
	7,  // 1: model.SimulationResult.character_details:type_name -> model.Character
	8,  // 2: model.SimulationResult.target_details:type_name -> model.Enemy
	9,  // 3: model.SimulationResult.simulator_settings:type_name -> model.SimulatorSettings
	10, // 4: model.SimulationResult.energy_settings:type_name -> model.EnergySettings
	2,  // 5: model.SimulationResult.statistics:type_name -> model.SimulationStatistics
	3,  // 6: model.SimulationStatistics.duration:type_name -> model.OverviewStats
	3,  // 7: model.SimulationStatistics.DPS:type_name -> model.OverviewStats
	3,  // 8: model.SimulationStatistics.RPS:type_name -> model.OverviewStats
	3,  // 9: model.SimulationStatistics.EPS:type_name -> model.OverviewStats
	3,  // 10: model.SimulationStatistics.HPS:type_name -> model.OverviewStats
	3,  // 11: model.SimulationStatistics.SPS:type_name -> model.OverviewStats
	3,  // 12: model.SimulationStatistics.total_damage:type_name -> model.OverviewStats
	5,  // 13: model.SimulationStatistics.warnings:type_name -> model.Warnings
	6,  // 14: model.SimulationStatistics.failed_actions:type_name -> model.FailedActions
	4,  // 15: model.FailedActions.insufficient_energy:type_name -> model.DescriptiveStats
	4,  // 16: model.FailedActions.insufficient_stamina:type_name -> model.DescriptiveStats
	4,  // 17: model.FailedActions.swap_cd:type_name -> model.DescriptiveStats
	4,  // 18: model.FailedActions.skill_cd:type_name -> model.DescriptiveStats
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_protos_model_result_proto_init() }
func file_protos_model_result_proto_init() {
	if File_protos_model_result_proto != nil {
		return
	}
	file_protos_model_sim_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_protos_model_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationResult); i {
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
		file_protos_model_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationStatistics); i {
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
		file_protos_model_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverviewStats); i {
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
		file_protos_model_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptiveStats); i {
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
		file_protos_model_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warnings); i {
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
		file_protos_model_result_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedActions); i {
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
			RawDescriptor: file_protos_model_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_result_proto_goTypes,
		DependencyIndexes: file_protos_model_result_proto_depIdxs,
		MessageInfos:      file_protos_model_result_proto_msgTypes,
	}.Build()
	File_protos_model_result_proto = out.File
	file_protos_model_result_proto_rawDesc = nil
	file_protos_model_result_proto_goTypes = nil
	file_protos_model_result_proto_depIdxs = nil
}
