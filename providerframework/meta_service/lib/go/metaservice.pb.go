// Copyright 2022 The SODA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: proto/metaservice.proto

package metaservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Resource specifies the name of a resource and whether it is namespaced.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the plural name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespaced indicates if a resource is namespaced or not.
	Namespaced bool `protobuf:"varint,2,opt,name=namespaced,proto3" json:"namespaced,omitempty"`
	// kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// group is the group of the resource.  Empty implies the group of the core resource list.
	// For subresources, this may have a different value, for example: Scale".
	Group string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	// version is the version of the resource.
	// For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)".
	Version string `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetNamespaced() bool {
	if x != nil {
		return x.Namespaced
	}
	return false
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Resource) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type BackupResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Data     []byte    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BackupResource) Reset() {
	*x = BackupResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupResource) ProtoMessage() {}

func (x *BackupResource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupResource.ProtoReflect.Descriptor instead.
func (*BackupResource) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{1}
}

func (x *BackupResource) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *BackupResource) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BackupIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackupHandle   string            `protobuf:"bytes,1,opt,name=backup_handle,json=backupHandle,proto3" json:"backup_handle,omitempty"`
	BackupLocation string            `protobuf:"bytes,2,opt,name=backup_location,json=backupLocation,proto3" json:"backup_location,omitempty"`
	Parameters     map[string]string `protobuf:"bytes,9,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BackupIdentifier) Reset() {
	*x = BackupIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupIdentifier) ProtoMessage() {}

func (x *BackupIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupIdentifier.ProtoReflect.Descriptor instead.
func (*BackupIdentifier) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{2}
}

func (x *BackupIdentifier) GetBackupHandle() string {
	if x != nil {
		return x.BackupHandle
	}
	return ""
}

func (x *BackupIdentifier) GetBackupLocation() string {
	if x != nil {
		return x.BackupLocation
	}
	return ""
}

func (x *BackupIdentifier) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type BackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Backup:
	//	*BackupRequest_Identifier
	//	*BackupRequest_BackupResource
	Backup isBackupRequest_Backup `protobuf_oneof:"Backup"`
}

func (x *BackupRequest) Reset() {
	*x = BackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRequest) ProtoMessage() {}

func (x *BackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRequest.ProtoReflect.Descriptor instead.
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{3}
}

func (m *BackupRequest) GetBackup() isBackupRequest_Backup {
	if m != nil {
		return m.Backup
	}
	return nil
}

func (x *BackupRequest) GetIdentifier() *BackupIdentifier {
	if x, ok := x.GetBackup().(*BackupRequest_Identifier); ok {
		return x.Identifier
	}
	return nil
}

func (x *BackupRequest) GetBackupResource() *BackupResource {
	if x, ok := x.GetBackup().(*BackupRequest_BackupResource); ok {
		return x.BackupResource
	}
	return nil
}

type isBackupRequest_Backup interface {
	isBackupRequest_Backup()
}

type BackupRequest_Identifier struct {
	Identifier *BackupIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3,oneof"`
}

type BackupRequest_BackupResource struct {
	BackupResource *BackupResource `protobuf:"bytes,2,opt,name=backup_resource,json=backupResource,proto3,oneof"`
}

func (*BackupRequest_Identifier) isBackupRequest_Backup() {}

func (*BackupRequest_BackupResource) isBackupRequest_Backup() {}

type RestoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *BackupIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RestoreRequest) Reset() {
	*x = RestoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreRequest) ProtoMessage() {}

func (x *RestoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreRequest.ProtoReflect.Descriptor instead.
func (*RestoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{4}
}

func (x *RestoreRequest) GetId() *BackupIdentifier {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Restore:
	//	*GetResponse_Identifier
	//	*GetResponse_BackupResource
	Restore isGetResponse_Restore `protobuf_oneof:"Restore"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_metaservice_proto_rawDescGZIP(), []int{5}
}

func (m *GetResponse) GetRestore() isGetResponse_Restore {
	if m != nil {
		return m.Restore
	}
	return nil
}

func (x *GetResponse) GetIdentifier() *BackupIdentifier {
	if x, ok := x.GetRestore().(*GetResponse_Identifier); ok {
		return x.Identifier
	}
	return nil
}

func (x *GetResponse) GetBackupResource() *BackupResource {
	if x, ok := x.GetRestore().(*GetResponse_BackupResource); ok {
		return x.BackupResource
	}
	return nil
}

type isGetResponse_Restore interface {
	isGetResponse_Restore()
}

type GetResponse_Identifier struct {
	Identifier *BackupIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3,oneof"`
}

type GetResponse_BackupResource struct {
	BackupResource *BackupResource `protobuf:"bytes,2,opt,name=backup_resource,json=backupResource,proto3,oneof"`
}

func (*GetResponse_Identifier) isGetResponse_Restore() {}

func (*GetResponse_BackupResource) isGetResponse_Restore() {}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metaservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metaservice_proto_msgTypes[6]
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
	return file_proto_metaservice_proto_rawDescGZIP(), []int{6}
}

var File_proto_metaservice_proto protoreflect.FileDescriptor

var file_proto_metaservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6b, 0x61, 0x68, 0x75, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x5c, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf3,
	0x01, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x52, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x61, 0x68,
	0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x4b, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xa5, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x68,
	0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x61,
	0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4e, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6b, 0x61, 0x68, 0x75, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x6c, 0x69, 0x62, 0x2f,
	0x67, 0x6f, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_metaservice_proto_rawDescOnce sync.Once
	file_proto_metaservice_proto_rawDescData = file_proto_metaservice_proto_rawDesc
)

func file_proto_metaservice_proto_rawDescGZIP() []byte {
	file_proto_metaservice_proto_rawDescOnce.Do(func() {
		file_proto_metaservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metaservice_proto_rawDescData)
	})
	return file_proto_metaservice_proto_rawDescData
}

var file_proto_metaservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_metaservice_proto_goTypes = []interface{}{
	(*Resource)(nil),         // 0: kahu.metaservice.Resource
	(*BackupResource)(nil),   // 1: kahu.metaservice.BackupResource
	(*BackupIdentifier)(nil), // 2: kahu.metaservice.BackupIdentifier
	(*BackupRequest)(nil),    // 3: kahu.metaservice.BackupRequest
	(*RestoreRequest)(nil),   // 4: kahu.metaservice.RestoreRequest
	(*GetResponse)(nil),      // 5: kahu.metaservice.GetResponse
	(*Empty)(nil),            // 6: kahu.metaservice.Empty
	nil,                      // 7: kahu.metaservice.BackupIdentifier.ParametersEntry
}
var file_proto_metaservice_proto_depIdxs = []int32{
	0, // 0: kahu.metaservice.BackupResource.resource:type_name -> kahu.metaservice.Resource
	7, // 1: kahu.metaservice.BackupIdentifier.parameters:type_name -> kahu.metaservice.BackupIdentifier.ParametersEntry
	2, // 2: kahu.metaservice.BackupRequest.identifier:type_name -> kahu.metaservice.BackupIdentifier
	1, // 3: kahu.metaservice.BackupRequest.backup_resource:type_name -> kahu.metaservice.BackupResource
	2, // 4: kahu.metaservice.RestoreRequest.id:type_name -> kahu.metaservice.BackupIdentifier
	2, // 5: kahu.metaservice.GetResponse.identifier:type_name -> kahu.metaservice.BackupIdentifier
	1, // 6: kahu.metaservice.GetResponse.backup_resource:type_name -> kahu.metaservice.BackupResource
	3, // 7: kahu.metaservice.MetaService.Backup:input_type -> kahu.metaservice.BackupRequest
	4, // 8: kahu.metaservice.MetaService.Restore:input_type -> kahu.metaservice.RestoreRequest
	6, // 9: kahu.metaservice.MetaService.Backup:output_type -> kahu.metaservice.Empty
	5, // 10: kahu.metaservice.MetaService.Restore:output_type -> kahu.metaservice.GetResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_metaservice_proto_init() }
func file_proto_metaservice_proto_init() {
	if File_proto_metaservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metaservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_proto_metaservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupResource); i {
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
		file_proto_metaservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupIdentifier); i {
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
		file_proto_metaservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupRequest); i {
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
		file_proto_metaservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreRequest); i {
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
		file_proto_metaservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_metaservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	file_proto_metaservice_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*BackupRequest_Identifier)(nil),
		(*BackupRequest_BackupResource)(nil),
	}
	file_proto_metaservice_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetResponse_Identifier)(nil),
		(*GetResponse_BackupResource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_metaservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_metaservice_proto_goTypes,
		DependencyIndexes: file_proto_metaservice_proto_depIdxs,
		MessageInfos:      file_proto_metaservice_proto_msgTypes,
	}.Build()
	File_proto_metaservice_proto = out.File
	file_proto_metaservice_proto_rawDesc = nil
	file_proto_metaservice_proto_goTypes = nil
	file_proto_metaservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaServiceClient interface {
	Backup(ctx context.Context, opts ...grpc.CallOption) (MetaService_BackupClient, error)
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (MetaService_RestoreClient, error)
}

type metaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaServiceClient(cc grpc.ClientConnInterface) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) Backup(ctx context.Context, opts ...grpc.CallOption) (MetaService_BackupClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetaService_serviceDesc.Streams[0], "/kahu.metaservice.MetaService/Backup", opts...)
	if err != nil {
		return nil, err
	}
	x := &metaServiceBackupClient{stream}
	return x, nil
}

type MetaService_BackupClient interface {
	Send(*BackupRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type metaServiceBackupClient struct {
	grpc.ClientStream
}

func (x *metaServiceBackupClient) Send(m *BackupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metaServiceBackupClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metaServiceClient) Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (MetaService_RestoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetaService_serviceDesc.Streams[1], "/kahu.metaservice.MetaService/Restore", opts...)
	if err != nil {
		return nil, err
	}
	x := &metaServiceRestoreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetaService_RestoreClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type metaServiceRestoreClient struct {
	grpc.ClientStream
}

func (x *metaServiceRestoreClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetaServiceServer is the server API for MetaService service.
type MetaServiceServer interface {
	Backup(MetaService_BackupServer) error
	Restore(*RestoreRequest, MetaService_RestoreServer) error
}

// UnimplementedMetaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (*UnimplementedMetaServiceServer) Backup(MetaService_BackupServer) error {
	return status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (*UnimplementedMetaServiceServer) Restore(*RestoreRequest, MetaService_RestoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

func RegisterMetaServiceServer(s *grpc.Server, srv MetaServiceServer) {
	s.RegisterService(&_MetaService_serviceDesc, srv)
}

func _MetaService_Backup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetaServiceServer).Backup(&metaServiceBackupServer{stream})
}

type MetaService_BackupServer interface {
	SendAndClose(*Empty) error
	Recv() (*BackupRequest, error)
	grpc.ServerStream
}

type metaServiceBackupServer struct {
	grpc.ServerStream
}

func (x *metaServiceBackupServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metaServiceBackupServer) Recv() (*BackupRequest, error) {
	m := new(BackupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MetaService_Restore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RestoreRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetaServiceServer).Restore(m, &metaServiceRestoreServer{stream})
}

type MetaService_RestoreServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type metaServiceRestoreServer struct {
	grpc.ServerStream
}

func (x *metaServiceRestoreServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kahu.metaservice.MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Backup",
			Handler:       _MetaService_Backup_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Restore",
			Handler:       _MetaService_Restore_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/metaservice.proto",
}
