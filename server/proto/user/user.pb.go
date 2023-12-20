// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user.proto

package __

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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DBCreateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Birthday    string `protobuf:"bytes,4,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Icon        string `protobuf:"bytes,7,opt,name=Icon,proto3" json:"Icon,omitempty"`
}

func (x *DBCreateUser) Reset() {
	*x = DBCreateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCreateUser) ProtoMessage() {}

func (x *DBCreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCreateUser.ProtoReflect.Descriptor instead.
func (*DBCreateUser) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *DBCreateUser) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DBCreateUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DBCreateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DBCreateUser) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *DBCreateUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DBCreateUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DBCreateUser) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type DBUpdateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Birthday    string `protobuf:"bytes,4,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Icon        string `protobuf:"bytes,7,opt,name=Icon,proto3" json:"Icon,omitempty"`
}

func (x *DBUpdateUser) Reset() {
	*x = DBUpdateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBUpdateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBUpdateUser) ProtoMessage() {}

func (x *DBUpdateUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBUpdateUser.ProtoReflect.Descriptor instead.
func (*DBUpdateUser) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *DBUpdateUser) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DBUpdateUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DBUpdateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DBUpdateUser) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *DBUpdateUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DBUpdateUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DBUpdateUser) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type DBGetUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Birthday    string `protobuf:"bytes,4,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Icon        string `protobuf:"bytes,7,opt,name=Icon,proto3" json:"Icon,omitempty"`
}

func (x *DBGetUser) Reset() {
	*x = DBGetUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBGetUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBGetUser) ProtoMessage() {}

func (x *DBGetUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBGetUser.ProtoReflect.Descriptor instead.
func (*DBGetUser) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *DBGetUser) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DBGetUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DBGetUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DBGetUser) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *DBGetUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DBGetUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DBGetUser) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x44, 0x42, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x44, 0x42, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x44, 0x42,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0xa7, 0x02, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x50, 0x43, 0x12, 0x29, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x1a, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x2f, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(*ID)(nil),           // 0: user.ID
	(*Value)(nil),        // 1: user.Value
	(*DBCreateUser)(nil), // 2: user.DBCreateUser
	(*DBUpdateUser)(nil), // 3: user.DBUpdateUser
	(*DBGetUser)(nil),    // 4: user.DBGetUser
	(*Nothing)(nil),      // 5: user.Nothing
}
var file_user_proto_depIdxs = []int32{
	0, // 0: user.UserRPC.FindUserByID:input_type -> user.ID
	2, // 1: user.UserRPC.CreateUser:input_type -> user.DBCreateUser
	3, // 2: user.UserRPC.UpdateUser:input_type -> user.DBUpdateUser
	1, // 3: user.UserRPC.FindUserByUsername:input_type -> user.Value
	1, // 4: user.UserRPC.FindUserByEmail:input_type -> user.Value
	1, // 5: user.UserRPC.FindUserByPhone:input_type -> user.Value
	4, // 6: user.UserRPC.FindUserByID:output_type -> user.DBGetUser
	0, // 7: user.UserRPC.CreateUser:output_type -> user.ID
	5, // 8: user.UserRPC.UpdateUser:output_type -> user.Nothing
	4, // 9: user.UserRPC.FindUserByUsername:output_type -> user.DBGetUser
	4, // 10: user.UserRPC.FindUserByEmail:output_type -> user.DBGetUser
	4, // 11: user.UserRPC.FindUserByPhone:output_type -> user.DBGetUser
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCreateUser); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBUpdateUser); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBGetUser); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
