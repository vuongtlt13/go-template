// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.1
// source: admin/permission_service.proto

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

type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Service       string                 `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Method        string                 `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_admin_permission_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Permission) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Permission) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Permission) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CreatePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Service       string                 `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	Method        string                 `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	mi := &file_admin_permission_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePermissionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreatePermissionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePermissionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePermissionRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *CreatePermissionRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type GetPermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPermissionRequest) Reset() {
	*x = GetPermissionRequest{}
	mi := &file_admin_permission_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionRequest) ProtoMessage() {}

func (x *GetPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionRequest) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPermissionRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Service       string                 `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Method        string                 `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	mi := &file_admin_permission_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePermissionRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePermissionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdatePermissionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePermissionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdatePermissionRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *UpdatePermissionRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type DeletePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePermissionRequest) Reset() {
	*x = DeletePermissionRequest{}
	mi := &file_admin_permission_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionRequest) ProtoMessage() {}

func (x *DeletePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionRequest.ProtoReflect.Descriptor instead.
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePermissionRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePermissionResponse) Reset() {
	*x = DeletePermissionResponse{}
	mi := &file_admin_permission_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionResponse) ProtoMessage() {}

func (x *DeletePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionResponse.ProtoReflect.Descriptor instead.
func (*DeletePermissionResponse) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePermissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeletePermissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListPermissionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          uint32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      uint32                 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsRequest) Reset() {
	*x = ListPermissionsRequest{}
	mi := &file_admin_permission_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsRequest) ProtoMessage() {}

func (x *ListPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsRequest.ProtoReflect.Descriptor instead.
func (*ListPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListPermissionsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPermissionsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPermissionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   []*Permission          `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Total         uint32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsResponse) Reset() {
	*x = ListPermissionsResponse{}
	mi := &file_admin_permission_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsResponse) ProtoMessage() {}

func (x *ListPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListPermissionsResponse) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *ListPermissionsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPermissionsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PermissionResponse) Reset() {
	*x = PermissionResponse{}
	mi := &file_admin_permission_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionResponse) ProtoMessage() {}

func (x *PermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_permission_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionResponse.ProtoReflect.Descriptor instead.
func (*PermissionResponse) Descriptor() ([]byte, []int) {
	return file_admin_permission_service_proto_rawDescGZIP(), []int{8}
}

func (x *PermissionResponse) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *PermissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_admin_permission_service_proto protoreflect.FileDescriptor

const file_admin_permission_service_proto_rawDesc = "" +
	"\n" +
	"\x1eadmin/permission_service.proto\x12\ryourapp.admin\x1a\x1cgoogle/api/annotations.proto\x1a\"envoyproxy/validate/validate.proto\"\xbc\x01\n" +
	"\n" +
	"Permission\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x1b\n" +
	"\x04code\x18\x02 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04code\x12\x1b\n" +
	"\x04name\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12!\n" +
	"\aservice\x18\x05 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\aservice\x12\x1f\n" +
	"\x06method\x18\x06 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x06method\"\xb9\x01\n" +
	"\x17CreatePermissionRequest\x12\x1b\n" +
	"\x04code\x18\x01 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04code\x12\x1b\n" +
	"\x04name\x18\x02 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12!\n" +
	"\aservice\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\aservice\x12\x1f\n" +
	"\x06method\x18\x05 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x06method\"/\n" +
	"\x14GetPermissionRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\"\xd2\x01\n" +
	"\x17UpdatePermissionRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\x12\x1b\n" +
	"\x04code\x18\x02 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04code\x12\x1b\n" +
	"\x04name\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12!\n" +
	"\aservice\x18\x05 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\aservice\x12\x1f\n" +
	"\x06method\x18\x06 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x06method\"2\n" +
	"\x17DeletePermissionRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x04B\a\xfaB\x042\x02 \x00R\x02id\"N\n" +
	"\x18DeletePermissionResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"[\n" +
	"\x16ListPermissionsRequest\x12\x1b\n" +
	"\x04page\x18\x01 \x01(\rB\a\xfaB\x04*\x02 \x00R\x04page\x12$\n" +
	"\tpage_size\x18\x02 \x01(\rB\a\xfaB\x04*\x02 \x00R\bpageSize\"\x86\x01\n" +
	"\x17ListPermissionsResponse\x12;\n" +
	"\vpermissions\x18\x01 \x03(\v2\x19.yourapp.admin.PermissionR\vpermissions\x12\x14\n" +
	"\x05total\x18\x02 \x01(\rR\x05total\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"i\n" +
	"\x12PermissionResponse\x129\n" +
	"\n" +
	"permission\x18\x01 \x01(\v2\x19.yourapp.admin.PermissionR\n" +
	"permission\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\x94\x05\n" +
	"\x11PermissionService\x12|\n" +
	"\x10CreatePermission\x12&.yourapp.admin.CreatePermissionRequest\x1a!.yourapp.admin.PermissionResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/admin/permissions\x12x\n" +
	"\rGetPermission\x12#.yourapp.admin.GetPermissionRequest\x1a!.yourapp.admin.PermissionResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/admin/permissions/{id}\x12\x81\x01\n" +
	"\x10UpdatePermission\x12&.yourapp.admin.UpdatePermissionRequest\x1a!.yourapp.admin.PermissionResponse\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\x1a\x17/admin/permissions/{id}\x12\x84\x01\n" +
	"\x10DeletePermission\x12&.yourapp.admin.DeletePermissionRequest\x1a'.yourapp.admin.DeletePermissionResponse\"\x1f\x82\xd3\xe4\x93\x02\x19*\x17/admin/permissions/{id}\x12|\n" +
	"\x0fListPermissions\x12%.yourapp.admin.ListPermissionsRequest\x1a&.yourapp.admin.ListPermissionsResponse\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/admin/permissionsB\x18Z\x16yourapp/pb/admin;adminb\x06proto3"

var (
	file_admin_permission_service_proto_rawDescOnce sync.Once
	file_admin_permission_service_proto_rawDescData []byte
)

func file_admin_permission_service_proto_rawDescGZIP() []byte {
	file_admin_permission_service_proto_rawDescOnce.Do(func() {
		file_admin_permission_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_admin_permission_service_proto_rawDesc), len(file_admin_permission_service_proto_rawDesc)))
	})
	return file_admin_permission_service_proto_rawDescData
}

var file_admin_permission_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_admin_permission_service_proto_goTypes = []any{
	(*Permission)(nil),               // 0: yourapp.admin.Permission
	(*CreatePermissionRequest)(nil),  // 1: yourapp.admin.CreatePermissionRequest
	(*GetPermissionRequest)(nil),     // 2: yourapp.admin.GetPermissionRequest
	(*UpdatePermissionRequest)(nil),  // 3: yourapp.admin.UpdatePermissionRequest
	(*DeletePermissionRequest)(nil),  // 4: yourapp.admin.DeletePermissionRequest
	(*DeletePermissionResponse)(nil), // 5: yourapp.admin.DeletePermissionResponse
	(*ListPermissionsRequest)(nil),   // 6: yourapp.admin.ListPermissionsRequest
	(*ListPermissionsResponse)(nil),  // 7: yourapp.admin.ListPermissionsResponse
	(*PermissionResponse)(nil),       // 8: yourapp.admin.PermissionResponse
}
var file_admin_permission_service_proto_depIdxs = []int32{
	0, // 0: yourapp.admin.ListPermissionsResponse.permissions:type_name -> yourapp.admin.Permission
	0, // 1: yourapp.admin.PermissionResponse.permission:type_name -> yourapp.admin.Permission
	1, // 2: yourapp.admin.PermissionService.CreatePermission:input_type -> yourapp.admin.CreatePermissionRequest
	2, // 3: yourapp.admin.PermissionService.GetPermission:input_type -> yourapp.admin.GetPermissionRequest
	3, // 4: yourapp.admin.PermissionService.UpdatePermission:input_type -> yourapp.admin.UpdatePermissionRequest
	4, // 5: yourapp.admin.PermissionService.DeletePermission:input_type -> yourapp.admin.DeletePermissionRequest
	6, // 6: yourapp.admin.PermissionService.ListPermissions:input_type -> yourapp.admin.ListPermissionsRequest
	8, // 7: yourapp.admin.PermissionService.CreatePermission:output_type -> yourapp.admin.PermissionResponse
	8, // 8: yourapp.admin.PermissionService.GetPermission:output_type -> yourapp.admin.PermissionResponse
	8, // 9: yourapp.admin.PermissionService.UpdatePermission:output_type -> yourapp.admin.PermissionResponse
	5, // 10: yourapp.admin.PermissionService.DeletePermission:output_type -> yourapp.admin.DeletePermissionResponse
	7, // 11: yourapp.admin.PermissionService.ListPermissions:output_type -> yourapp.admin.ListPermissionsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_permission_service_proto_init() }
func file_admin_permission_service_proto_init() {
	if File_admin_permission_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_permission_service_proto_rawDesc), len(file_admin_permission_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_permission_service_proto_goTypes,
		DependencyIndexes: file_admin_permission_service_proto_depIdxs,
		MessageInfos:      file_admin_permission_service_proto_msgTypes,
	}.Build()
	File_admin_permission_service_proto = out.File
	file_admin_permission_service_proto_goTypes = nil
	file_admin_permission_service_proto_depIdxs = nil
}
