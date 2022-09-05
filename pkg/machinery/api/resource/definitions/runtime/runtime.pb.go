// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: resource/definitions/runtime/runtime.proto

package runtime

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	enums "github.com/talos-systems/talos/pkg/machinery/api/resource/definitions/enums"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// KernelModuleSpecSpec describes Linux kernel module to load.
type KernelModuleSpecSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *KernelModuleSpecSpec) Reset() {
	*x = KernelModuleSpecSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelModuleSpecSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelModuleSpecSpec) ProtoMessage() {}

func (x *KernelModuleSpecSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelModuleSpecSpec.ProtoReflect.Descriptor instead.
func (*KernelModuleSpecSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *KernelModuleSpecSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// KernelParamSpecSpec describes status of the defined sysctls.
type KernelParamSpecSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	IgnoreErrors bool   `protobuf:"varint,2,opt,name=ignore_errors,json=ignoreErrors,proto3" json:"ignore_errors,omitempty"`
}

func (x *KernelParamSpecSpec) Reset() {
	*x = KernelParamSpecSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelParamSpecSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelParamSpecSpec) ProtoMessage() {}

func (x *KernelParamSpecSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelParamSpecSpec.ProtoReflect.Descriptor instead.
func (*KernelParamSpecSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{1}
}

func (x *KernelParamSpecSpec) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *KernelParamSpecSpec) GetIgnoreErrors() bool {
	if x != nil {
		return x.IgnoreErrors
	}
	return false
}

// KernelParamStatusSpec describes status of the defined sysctls.
type KernelParamStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current     string `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Default     string `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Unsupported bool   `protobuf:"varint,3,opt,name=unsupported,proto3" json:"unsupported,omitempty"`
}

func (x *KernelParamStatusSpec) Reset() {
	*x = KernelParamStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelParamStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelParamStatusSpec) ProtoMessage() {}

func (x *KernelParamStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelParamStatusSpec.ProtoReflect.Descriptor instead.
func (*KernelParamStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{2}
}

func (x *KernelParamStatusSpec) GetCurrent() string {
	if x != nil {
		return x.Current
	}
	return ""
}

func (x *KernelParamStatusSpec) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *KernelParamStatusSpec) GetUnsupported() bool {
	if x != nil {
		return x.Unsupported
	}
	return false
}

// MachineStatusSpec describes status of the defined sysctls.
type MachineStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage  enums.RuntimeMachineStage `protobuf:"varint,1,opt,name=stage,proto3,enum=talos.resource.definitions.enums.RuntimeMachineStage" json:"stage,omitempty"`
	Status *MachineStatusStatus      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MachineStatusSpec) Reset() {
	*x = MachineStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineStatusSpec) ProtoMessage() {}

func (x *MachineStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineStatusSpec.ProtoReflect.Descriptor instead.
func (*MachineStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{3}
}

func (x *MachineStatusSpec) GetStage() enums.RuntimeMachineStage {
	if x != nil {
		return x.Stage
	}
	return enums.RuntimeMachineStage(0)
}

func (x *MachineStatusSpec) GetStatus() *MachineStatusStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// MachineStatusStatus describes machine current status at the stage.
type MachineStatusStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready           bool              `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	UnmetConditions []*UnmetCondition `protobuf:"bytes,2,rep,name=unmet_conditions,json=unmetConditions,proto3" json:"unmet_conditions,omitempty"`
}

func (x *MachineStatusStatus) Reset() {
	*x = MachineStatusStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineStatusStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineStatusStatus) ProtoMessage() {}

func (x *MachineStatusStatus) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineStatusStatus.ProtoReflect.Descriptor instead.
func (*MachineStatusStatus) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{4}
}

func (x *MachineStatusStatus) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *MachineStatusStatus) GetUnmetConditions() []*UnmetCondition {
	if x != nil {
		return x.UnmetConditions
	}
	return nil
}

// MountStatusSpec describes status of the defined sysctls.
type MountStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source         string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target         string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	FilesystemType string   `protobuf:"bytes,3,opt,name=filesystem_type,json=filesystemType,proto3" json:"filesystem_type,omitempty"`
	Options        []string `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *MountStatusSpec) Reset() {
	*x = MountStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountStatusSpec) ProtoMessage() {}

func (x *MountStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountStatusSpec.ProtoReflect.Descriptor instead.
func (*MountStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{5}
}

func (x *MountStatusSpec) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MountStatusSpec) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *MountStatusSpec) GetFilesystemType() string {
	if x != nil {
		return x.FilesystemType
	}
	return ""
}

func (x *MountStatusSpec) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

// UnmetCondition is a failure which prevents machine from being ready at the stage.
type UnmetCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UnmetCondition) Reset() {
	*x = UnmetCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnmetCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnmetCondition) ProtoMessage() {}

func (x *UnmetCondition) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_runtime_runtime_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnmetCondition.ProtoReflect.Descriptor instead.
func (*UnmetCondition) Descriptor() ([]byte, []int) {
	return file_resource_definitions_runtime_runtime_proto_rawDescGZIP(), []int{6}
}

func (x *UnmetCondition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnmetCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_resource_definitions_runtime_runtime_proto protoreflect.FileDescriptor

var file_resource_definitions_runtime_runtime_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x61,
	0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x1a, 0x26, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x14, 0x4b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x13, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x6d, 0x0a, 0x15, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x6e, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x5d, 0x0a, 0x10, 0x75, 0x6e, 0x6d, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x6d, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x75, 0x6e, 0x6d, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a,
	0x0e, 0x55, 0x6e, 0x6d, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x4f, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_runtime_runtime_proto_rawDescOnce sync.Once
	file_resource_definitions_runtime_runtime_proto_rawDescData = file_resource_definitions_runtime_runtime_proto_rawDesc
)

func file_resource_definitions_runtime_runtime_proto_rawDescGZIP() []byte {
	file_resource_definitions_runtime_runtime_proto_rawDescOnce.Do(func() {
		file_resource_definitions_runtime_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_runtime_runtime_proto_rawDescData)
	})
	return file_resource_definitions_runtime_runtime_proto_rawDescData
}

var file_resource_definitions_runtime_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_resource_definitions_runtime_runtime_proto_goTypes = []interface{}{
	(*KernelModuleSpecSpec)(nil),   // 0: talos.resource.definitions.runtime.KernelModuleSpecSpec
	(*KernelParamSpecSpec)(nil),    // 1: talos.resource.definitions.runtime.KernelParamSpecSpec
	(*KernelParamStatusSpec)(nil),  // 2: talos.resource.definitions.runtime.KernelParamStatusSpec
	(*MachineStatusSpec)(nil),      // 3: talos.resource.definitions.runtime.MachineStatusSpec
	(*MachineStatusStatus)(nil),    // 4: talos.resource.definitions.runtime.MachineStatusStatus
	(*MountStatusSpec)(nil),        // 5: talos.resource.definitions.runtime.MountStatusSpec
	(*UnmetCondition)(nil),         // 6: talos.resource.definitions.runtime.UnmetCondition
	(enums.RuntimeMachineStage)(0), // 7: talos.resource.definitions.enums.RuntimeMachineStage
}
var file_resource_definitions_runtime_runtime_proto_depIdxs = []int32{
	7, // 0: talos.resource.definitions.runtime.MachineStatusSpec.stage:type_name -> talos.resource.definitions.enums.RuntimeMachineStage
	4, // 1: talos.resource.definitions.runtime.MachineStatusSpec.status:type_name -> talos.resource.definitions.runtime.MachineStatusStatus
	6, // 2: talos.resource.definitions.runtime.MachineStatusStatus.unmet_conditions:type_name -> talos.resource.definitions.runtime.UnmetCondition
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resource_definitions_runtime_runtime_proto_init() }
func file_resource_definitions_runtime_runtime_proto_init() {
	if File_resource_definitions_runtime_runtime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_definitions_runtime_runtime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelModuleSpecSpec); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelParamSpecSpec); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelParamStatusSpec); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineStatusSpec); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineStatusStatus); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountStatusSpec); i {
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
		file_resource_definitions_runtime_runtime_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnmetCondition); i {
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
			RawDescriptor: file_resource_definitions_runtime_runtime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_runtime_runtime_proto_goTypes,
		DependencyIndexes: file_resource_definitions_runtime_runtime_proto_depIdxs,
		MessageInfos:      file_resource_definitions_runtime_runtime_proto_msgTypes,
	}.Build()
	File_resource_definitions_runtime_runtime_proto = out.File
	file_resource_definitions_runtime_runtime_proto_rawDesc = nil
	file_resource_definitions_runtime_runtime_proto_goTypes = nil
	file_resource_definitions_runtime_runtime_proto_depIdxs = nil
}
