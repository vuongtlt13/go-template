// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.1
// source: admin/user_service.proto

package admin

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_admin_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[0]
	if x != nil {
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
	return file_admin_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_admin_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_admin_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_admin_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateUserRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_admin_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	mi := &file_admin_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          uint32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      uint32                 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	mi := &file_admin_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListUsersRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUsersRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Total         uint32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_admin_user_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_admin_user_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_admin_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_admin_user_service_proto protoreflect.FileDescriptor

const file_admin_user_service_proto_rawDesc = "" +
	"\n" +
	"\x18admin/user_service.proto\x12\ryourapp.admin\x1a\x1cgoogle/api/annotations.proto\x1a\"envoyproxy/validate/validate.proto\"\xa0\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x1d\n" +
	"\x05email\x18\x02 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x12&\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\tfirstName\x12$\n" +
	"\tlast_name\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\blastName\x12\x1b\n" +
	"\tis_active\x18\x05 \x01(\bR\bisActive\"\xa5\x01\n" +
	"\x11CreateUserRequest\x12\x1d\n" +
	"\x05email\x18\x01 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x12#\n" +
	"\bpassword\x18\x02 \x01(\tB\a\xfaB\x04r\x02\x10\bR\bpassword\x12&\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\tfirstName\x12$\n" +
	"\tlast_name\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\blastName\")\n" +
	"\x0eGetUserRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\"\xb6\x01\n" +
	"\x11UpdateUserRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\x12\x1d\n" +
	"\x05email\x18\x02 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x12&\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\tfirstName\x12$\n" +
	"\tlast_name\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\blastName\x12\x1b\n" +
	"\tis_active\x18\x05 \x01(\bR\bisActive\",\n" +
	"\x11DeleteUserRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\".\n" +
	"\x12DeleteUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"U\n" +
	"\x10ListUsersRequest\x12\x1b\n" +
	"\x04page\x18\x01 \x01(\rB\a\xfaB\x04*\x02 \x00R\x04page\x12$\n" +
	"\tpage_size\x18\x02 \x01(\rB\a\xfaB\x04*\x02 \x00R\bpageSize\"T\n" +
	"\x11ListUsersResponse\x12)\n" +
	"\x05users\x18\x01 \x03(\v2\x13.yourapp.admin.UserR\x05users\x12\x14\n" +
	"\x05total\x18\x02 \x01(\rR\x05total\"7\n" +
	"\fUserResponse\x12'\n" +
	"\x04user\x18\x01 \x01(\v2\x13.yourapp.admin.UserR\x04user2\x94\x04\n" +
	"\vUserService\x12d\n" +
	"\n" +
	"CreateUser\x12 .yourapp.admin.CreateUserRequest\x1a\x1b.yourapp.admin.UserResponse\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/admin/users\x12`\n" +
	"\aGetUser\x12\x1d.yourapp.admin.GetUserRequest\x1a\x1b.yourapp.admin.UserResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/admin/users/{id}\x12i\n" +
	"\n" +
	"UpdateUser\x12 .yourapp.admin.UpdateUserRequest\x1a\x1b.yourapp.admin.UserResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\x1a\x11/admin/users/{id}\x12l\n" +
	"\n" +
	"DeleteUser\x12 .yourapp.admin.DeleteUserRequest\x1a!.yourapp.admin.DeleteUserResponse\"\x19\x82\xd3\xe4\x93\x02\x13*\x11/admin/users/{id}\x12d\n" +
	"\tListUsers\x12\x1f.yourapp.admin.ListUsersRequest\x1a .yourapp.admin.ListUsersResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/admin/usersB\x18Z\x16yourapp/pb/admin;adminb\x06proto3"

var (
	file_admin_user_service_proto_rawDescOnce sync.Once
	file_admin_user_service_proto_rawDescData []byte
)

func file_admin_user_service_proto_rawDescGZIP() []byte {
	file_admin_user_service_proto_rawDescOnce.Do(func() {
		file_admin_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_admin_user_service_proto_rawDesc), len(file_admin_user_service_proto_rawDesc)))
	})
	return file_admin_user_service_proto_rawDescData
}

var file_admin_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_admin_user_service_proto_goTypes = []any{
	(*User)(nil),               // 0: yourapp.admin.User
	(*CreateUserRequest)(nil),  // 1: yourapp.admin.CreateUserRequest
	(*GetUserRequest)(nil),     // 2: yourapp.admin.GetUserRequest
	(*UpdateUserRequest)(nil),  // 3: yourapp.admin.UpdateUserRequest
	(*DeleteUserRequest)(nil),  // 4: yourapp.admin.DeleteUserRequest
	(*DeleteUserResponse)(nil), // 5: yourapp.admin.DeleteUserResponse
	(*ListUsersRequest)(nil),   // 6: yourapp.admin.ListUsersRequest
	(*ListUsersResponse)(nil),  // 7: yourapp.admin.ListUsersResponse
	(*UserResponse)(nil),       // 8: yourapp.admin.UserResponse
}
var file_admin_user_service_proto_depIdxs = []int32{
	0, // 0: yourapp.admin.ListUsersResponse.users:type_name -> yourapp.admin.User
	0, // 1: yourapp.admin.UserResponse.user:type_name -> yourapp.admin.User
	1, // 2: yourapp.admin.UserService.CreateUser:input_type -> yourapp.admin.CreateUserRequest
	2, // 3: yourapp.admin.UserService.GetUser:input_type -> yourapp.admin.GetUserRequest
	3, // 4: yourapp.admin.UserService.UpdateUser:input_type -> yourapp.admin.UpdateUserRequest
	4, // 5: yourapp.admin.UserService.DeleteUser:input_type -> yourapp.admin.DeleteUserRequest
	6, // 6: yourapp.admin.UserService.ListUsers:input_type -> yourapp.admin.ListUsersRequest
	8, // 7: yourapp.admin.UserService.CreateUser:output_type -> yourapp.admin.UserResponse
	8, // 8: yourapp.admin.UserService.GetUser:output_type -> yourapp.admin.UserResponse
	8, // 9: yourapp.admin.UserService.UpdateUser:output_type -> yourapp.admin.UserResponse
	5, // 10: yourapp.admin.UserService.DeleteUser:output_type -> yourapp.admin.DeleteUserResponse
	7, // 11: yourapp.admin.UserService.ListUsers:output_type -> yourapp.admin.ListUsersResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_user_service_proto_init() }
func file_admin_user_service_proto_init() {
	if File_admin_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_user_service_proto_rawDesc), len(file_admin_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_user_service_proto_goTypes,
		DependencyIndexes: file_admin_user_service_proto_depIdxs,
		MessageInfos:      file_admin_user_service_proto_msgTypes,
	}.Build()
	File_admin_user_service_proto = out.File
	file_admin_user_service_proto_goTypes = nil
	file_admin_user_service_proto_depIdxs = nil
}
