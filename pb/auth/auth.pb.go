// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: auth.proto

package Auth

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

type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Err     string `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthData) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *AuthData) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

func (x *AuthData) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type UserCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserCode) Reset() {
	*x = UserCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCode) ProtoMessage() {}

func (x *UserCode) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCode.ProtoReflect.Descriptor instead.
func (*UserCode) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserCode) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UserCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x22, 0x3d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x4e, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2e, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xe6, 0x01, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x0e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x0e, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x09,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x1a, 0x0c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x42, 0x02, 0x5a, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_auth_proto_goTypes = []interface{}{
	(*LoginData)(nil), // 0: Auth.LoginData
	(*AuthData)(nil),  // 1: Auth.AuthData
	(*User)(nil),      // 2: Auth.User
	(*Status)(nil),    // 3: Auth.Status
	(*UserCode)(nil),  // 4: Auth.UserCode
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: Auth.Auth.Register:input_type -> Auth.LoginData
	0, // 1: Auth.Auth.Login:input_type -> Auth.LoginData
	1, // 2: Auth.Auth.Refresh:input_type -> Auth.AuthData
	2, // 3: Auth.Auth.SendEmail:input_type -> Auth.User
	4, // 4: Auth.Auth.VerifyEmail:input_type -> Auth.UserCode
	1, // 5: Auth.Auth.Register:output_type -> Auth.AuthData
	1, // 6: Auth.Auth.Login:output_type -> Auth.AuthData
	1, // 7: Auth.Auth.Refresh:output_type -> Auth.AuthData
	3, // 8: Auth.Auth.SendEmail:output_type -> Auth.Status
	3, // 9: Auth.Auth.VerifyEmail:output_type -> Auth.Status
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCode); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
