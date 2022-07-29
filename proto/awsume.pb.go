// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: awsume.proto

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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleArn            string `protobuf:"bytes,1,opt,name=RoleArn,proto3" json:"RoleArn,omitempty"`
	SourceProfile      string `protobuf:"bytes,2,opt,name=SourceProfile,proto3" json:"SourceProfile,omitempty"`
	AwsAccessKeyId     string `protobuf:"bytes,3,opt,name=AwsAccessKeyId,proto3" json:"AwsAccessKeyId,omitempty"`
	AwsAccessKeySecret string `protobuf:"bytes,4,opt,name=AwsAccessKeySecret,proto3" json:"AwsAccessKeySecret,omitempty"`
	MfaSerial          string `protobuf:"bytes,5,opt,name=MfaSerial,proto3" json:"MfaSerial,omitempty"`
	Region             string `protobuf:"bytes,6,opt,name=Region,proto3" json:"Region,omitempty"`
	Output             string `protobuf:"bytes,7,opt,name=Output,proto3" json:"Output,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awsume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_awsume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_awsume_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *Profile) GetSourceProfile() string {
	if x != nil {
		return x.SourceProfile
	}
	return ""
}

func (x *Profile) GetAwsAccessKeyId() string {
	if x != nil {
		return x.AwsAccessKeyId
	}
	return ""
}

func (x *Profile) GetAwsAccessKeySecret() string {
	if x != nil {
		return x.AwsAccessKeySecret
	}
	return ""
}

func (x *Profile) GetMfaSerial() string {
	if x != nil {
		return x.MfaSerial
	}
	return ""
}

func (x *Profile) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Profile) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type ProfilesMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *ProfilesMsg) Reset() {
	*x = ProfilesMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awsume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfilesMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfilesMsg) ProtoMessage() {}

func (x *ProfilesMsg) ProtoReflect() protoreflect.Message {
	mi := &file_awsume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfilesMsg.ProtoReflect.Descriptor instead.
func (*ProfilesMsg) Descriptor() ([]byte, []int) {
	return file_awsume_proto_rawDescGZIP(), []int{1}
}

func (x *ProfilesMsg) GetProfiles() []*Profile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type Argument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Flag  string `protobuf:"bytes,4,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *Argument) Reset() {
	*x = Argument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awsume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Argument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Argument) ProtoMessage() {}

func (x *Argument) ProtoReflect() protoreflect.Message {
	mi := &file_awsume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Argument.ProtoReflect.Descriptor instead.
func (*Argument) Descriptor() ([]byte, []int) {
	return file_awsume_proto_rawDescGZIP(), []int{2}
}

func (x *Argument) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Argument) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Argument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Argument) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

type ArgumentsMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arguments []*Argument `protobuf:"bytes,1,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *ArgumentsMsg) Reset() {
	*x = ArgumentsMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awsume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgumentsMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgumentsMsg) ProtoMessage() {}

func (x *ArgumentsMsg) ProtoReflect() protoreflect.Message {
	mi := &file_awsume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgumentsMsg.ProtoReflect.Descriptor instead.
func (*ArgumentsMsg) Descriptor() ([]byte, []int) {
	return file_awsume_proto_rawDescGZIP(), []int{3}
}

func (x *ArgumentsMsg) GetArguments() []*Argument {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awsume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_awsume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_awsume_proto_rawDescGZIP(), []int{4}
}

var File_awsume_proto protoreflect.FileDescriptor

var file_awsume_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x77, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x77, 0x73, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x77, 0x73,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x66, 0x61,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x66,
	0x61, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x39, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x22, 0x5c, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x22, 0x3d, 0x0a, 0x0c, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x73, 0x67,
	0x12, 0x2d, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x83, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x50, 0x72, 0x65, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x73, 0x67,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x80,
	0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x50,
	0x72, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x4d, 0x73, 0x67, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0x77, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x21, 0x0a, 0x03, 0x50, 0x72, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x78, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x50, 0x72,
	0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x21, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x22, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_awsume_proto_rawDescOnce sync.Once
	file_awsume_proto_rawDescData = file_awsume_proto_rawDesc
)

func file_awsume_proto_rawDescGZIP() []byte {
	file_awsume_proto_rawDescOnce.Do(func() {
		file_awsume_proto_rawDescData = protoimpl.X.CompressGZIP(file_awsume_proto_rawDescData)
	})
	return file_awsume_proto_rawDescData
}

var file_awsume_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_awsume_proto_goTypes = []interface{}{
	(*Profile)(nil),      // 0: proto.Profile
	(*ProfilesMsg)(nil),  // 1: proto.ProfilesMsg
	(*Argument)(nil),     // 2: proto.Argument
	(*ArgumentsMsg)(nil), // 3: proto.ArgumentsMsg
	(*Empty)(nil),        // 4: proto.Empty
}
var file_awsume_proto_depIdxs = []int32{
	0,  // 0: proto.ProfilesMsg.profiles:type_name -> proto.Profile
	2,  // 1: proto.ArgumentsMsg.arguments:type_name -> proto.Argument
	4,  // 2: proto.Arguments.Pre:input_type -> proto.Empty
	4,  // 3: proto.Arguments.Get:input_type -> proto.Empty
	3,  // 4: proto.Arguments.Post:input_type -> proto.ArgumentsMsg
	4,  // 5: proto.Profiles.Pre:input_type -> proto.Empty
	4,  // 6: proto.Profiles.Get:input_type -> proto.Empty
	1,  // 7: proto.Profiles.Post:input_type -> proto.ProfilesMsg
	4,  // 8: proto.Credentials.Pre:input_type -> proto.Empty
	4,  // 9: proto.Credentials.Get:input_type -> proto.Empty
	4,  // 10: proto.Credentials.Post:input_type -> proto.Empty
	4,  // 11: proto.ProfileNames.Pre:input_type -> proto.Empty
	4,  // 12: proto.ProfileNames.Get:input_type -> proto.Empty
	4,  // 13: proto.ProfileNames.Post:input_type -> proto.Empty
	4,  // 14: proto.Arguments.Pre:output_type -> proto.Empty
	3,  // 15: proto.Arguments.Get:output_type -> proto.ArgumentsMsg
	4,  // 16: proto.Arguments.Post:output_type -> proto.Empty
	4,  // 17: proto.Profiles.Pre:output_type -> proto.Empty
	1,  // 18: proto.Profiles.Get:output_type -> proto.ProfilesMsg
	4,  // 19: proto.Profiles.Post:output_type -> proto.Empty
	4,  // 20: proto.Credentials.Pre:output_type -> proto.Empty
	4,  // 21: proto.Credentials.Get:output_type -> proto.Empty
	4,  // 22: proto.Credentials.Post:output_type -> proto.Empty
	4,  // 23: proto.ProfileNames.Pre:output_type -> proto.Empty
	4,  // 24: proto.ProfileNames.Get:output_type -> proto.Empty
	4,  // 25: proto.ProfileNames.Post:output_type -> proto.Empty
	14, // [14:26] is the sub-list for method output_type
	2,  // [2:14] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_awsume_proto_init() }
func file_awsume_proto_init() {
	if File_awsume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_awsume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_awsume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfilesMsg); i {
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
		file_awsume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Argument); i {
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
		file_awsume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgumentsMsg); i {
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
		file_awsume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_awsume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_awsume_proto_goTypes,
		DependencyIndexes: file_awsume_proto_depIdxs,
		MessageInfos:      file_awsume_proto_msgTypes,
	}.Build()
	File_awsume_proto = out.File
	file_awsume_proto_rawDesc = nil
	file_awsume_proto_goTypes = nil
	file_awsume_proto_depIdxs = nil
}