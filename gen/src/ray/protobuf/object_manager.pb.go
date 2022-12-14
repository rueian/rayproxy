// Copyright 2017 The Ray Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: src/ray/protobuf/object_manager.proto

package protobuf

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

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The push ID to allow the receiver to differentiate different push attempts
	// from the same sender.
	PushId []byte `protobuf:"bytes,1,opt,name=push_id,json=pushId,proto3" json:"push_id,omitempty"`
	// The object ID being transferred.
	ObjectId []byte `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// The node ID of client sending this object
	NodeId []byte `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// The owner address
	OwnerAddress *Address `protobuf:"bytes,4,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// The index of the chunk being transferred.
	ChunkIndex uint32 `protobuf:"varint,5,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
	// The data_size include object_size and metadata_size
	DataSize uint64 `protobuf:"varint,6,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	// The metadata size.
	MetadataSize uint64 `protobuf:"varint,7,opt,name=metadata_size,json=metadataSize,proto3" json:"metadata_size,omitempty"`
	// The chunk data
	Data []byte `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetPushId() []byte {
	if x != nil {
		return x.PushId
	}
	return nil
}

func (x *PushRequest) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *PushRequest) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *PushRequest) GetOwnerAddress() *Address {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *PushRequest) GetChunkIndex() uint32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

func (x *PushRequest) GetDataSize() uint64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *PushRequest) GetMetadataSize() uint64 {
	if x != nil {
		return x.MetadataSize
	}
	return 0
}

func (x *PushRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node ID of the requesting client.
	NodeId []byte `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Requested ObjectID.
	ObjectId []byte `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *PullRequest) Reset() {
	*x = PullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRequest) ProtoMessage() {}

func (x *PullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRequest.ProtoReflect.Descriptor instead.
func (*PullRequest) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{1}
}

func (x *PullRequest) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *PullRequest) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

type FreeObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectIds [][]byte `protobuf:"bytes,1,rep,name=object_ids,json=objectIds,proto3" json:"object_ids,omitempty"`
}

func (x *FreeObjectsRequest) Reset() {
	*x = FreeObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeObjectsRequest) ProtoMessage() {}

func (x *FreeObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeObjectsRequest.ProtoReflect.Descriptor instead.
func (*FreeObjectsRequest) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{2}
}

func (x *FreeObjectsRequest) GetObjectIds() [][]byte {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

// Reply for request
type PushReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushReply) Reset() {
	*x = PushReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushReply) ProtoMessage() {}

func (x *PushReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushReply.ProtoReflect.Descriptor instead.
func (*PushReply) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{3}
}

type PullReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullReply) Reset() {
	*x = PullReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullReply) ProtoMessage() {}

func (x *PullReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullReply.ProtoReflect.Descriptor instead.
func (*PullReply) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{4}
}

type FreeObjectsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FreeObjectsReply) Reset() {
	*x = FreeObjectsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeObjectsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeObjectsReply) ProtoMessage() {}

func (x *FreeObjectsReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_object_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeObjectsReply.ProtoReflect.Descriptor instead.
func (*FreeObjectsReply) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_object_manager_proto_rawDescGZIP(), []int{5}
}

var File_src_ray_protobuf_object_manager_proto protoreflect.FileDescriptor

var file_src_ray_protobuf_object_manager_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63,
	0x1a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8a, 0x02, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x75, 0x73, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x0b,
	0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x33, 0x0a, 0x12, 0x46, 0x72, 0x65, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x12, 0x0a, 0x10, 0x46, 0x72, 0x65, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0xc1, 0x01, 0x0a, 0x14, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x61,
	0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x30, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x45, 0x0a, 0x0b, 0x46, 0x72, 0x65, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x1b, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x8e, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x12, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x65, 0x69, 0x61,
	0x6e, 0x2f, 0x72, 0x61, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x72,
	0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x52, 0x52, 0x58, 0xaa, 0x02, 0x07, 0x52, 0x61, 0x79, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02,
	0x07, 0x52, 0x61, 0x79, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x13, 0x52, 0x61, 0x79, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x52, 0x61, 0x79, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_src_ray_protobuf_object_manager_proto_rawDescOnce sync.Once
	file_src_ray_protobuf_object_manager_proto_rawDescData = file_src_ray_protobuf_object_manager_proto_rawDesc
)

func file_src_ray_protobuf_object_manager_proto_rawDescGZIP() []byte {
	file_src_ray_protobuf_object_manager_proto_rawDescOnce.Do(func() {
		file_src_ray_protobuf_object_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_ray_protobuf_object_manager_proto_rawDescData)
	})
	return file_src_ray_protobuf_object_manager_proto_rawDescData
}

var file_src_ray_protobuf_object_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_ray_protobuf_object_manager_proto_goTypes = []interface{}{
	(*PushRequest)(nil),        // 0: ray.rpc.PushRequest
	(*PullRequest)(nil),        // 1: ray.rpc.PullRequest
	(*FreeObjectsRequest)(nil), // 2: ray.rpc.FreeObjectsRequest
	(*PushReply)(nil),          // 3: ray.rpc.PushReply
	(*PullReply)(nil),          // 4: ray.rpc.PullReply
	(*FreeObjectsReply)(nil),   // 5: ray.rpc.FreeObjectsReply
	(*Address)(nil),            // 6: ray.rpc.Address
}
var file_src_ray_protobuf_object_manager_proto_depIdxs = []int32{
	6, // 0: ray.rpc.PushRequest.owner_address:type_name -> ray.rpc.Address
	0, // 1: ray.rpc.ObjectManagerService.Push:input_type -> ray.rpc.PushRequest
	1, // 2: ray.rpc.ObjectManagerService.Pull:input_type -> ray.rpc.PullRequest
	2, // 3: ray.rpc.ObjectManagerService.FreeObjects:input_type -> ray.rpc.FreeObjectsRequest
	3, // 4: ray.rpc.ObjectManagerService.Push:output_type -> ray.rpc.PushReply
	4, // 5: ray.rpc.ObjectManagerService.Pull:output_type -> ray.rpc.PullReply
	5, // 6: ray.rpc.ObjectManagerService.FreeObjects:output_type -> ray.rpc.FreeObjectsReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_ray_protobuf_object_manager_proto_init() }
func file_src_ray_protobuf_object_manager_proto_init() {
	if File_src_ray_protobuf_object_manager_proto != nil {
		return
	}
	file_src_ray_protobuf_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_ray_protobuf_object_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_src_ray_protobuf_object_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullRequest); i {
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
		file_src_ray_protobuf_object_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeObjectsRequest); i {
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
		file_src_ray_protobuf_object_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushReply); i {
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
		file_src_ray_protobuf_object_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullReply); i {
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
		file_src_ray_protobuf_object_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeObjectsReply); i {
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
			RawDescriptor: file_src_ray_protobuf_object_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_ray_protobuf_object_manager_proto_goTypes,
		DependencyIndexes: file_src_ray_protobuf_object_manager_proto_depIdxs,
		MessageInfos:      file_src_ray_protobuf_object_manager_proto_msgTypes,
	}.Build()
	File_src_ray_protobuf_object_manager_proto = out.File
	file_src_ray_protobuf_object_manager_proto_rawDesc = nil
	file_src_ray_protobuf_object_manager_proto_goTypes = nil
	file_src_ray_protobuf_object_manager_proto_depIdxs = nil
}
