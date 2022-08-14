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
// source: src/ray/protobuf/agent_manager.proto

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

type AgentRpcStatus int32

const (
	// OK.
	AgentRpcStatus_AGENT_RPC_STATUS_OK AgentRpcStatus = 0
	// Failed.
	AgentRpcStatus_AGENT_RPC_STATUS_FAILED AgentRpcStatus = 1
)

// Enum value maps for AgentRpcStatus.
var (
	AgentRpcStatus_name = map[int32]string{
		0: "AGENT_RPC_STATUS_OK",
		1: "AGENT_RPC_STATUS_FAILED",
	}
	AgentRpcStatus_value = map[string]int32{
		"AGENT_RPC_STATUS_OK":     0,
		"AGENT_RPC_STATUS_FAILED": 1,
	}
)

func (x AgentRpcStatus) Enum() *AgentRpcStatus {
	p := new(AgentRpcStatus)
	*p = x
	return p
}

func (x AgentRpcStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentRpcStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_src_ray_protobuf_agent_manager_proto_enumTypes[0].Descriptor()
}

func (AgentRpcStatus) Type() protoreflect.EnumType {
	return &file_src_ray_protobuf_agent_manager_proto_enumTypes[0]
}

func (x AgentRpcStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentRpcStatus.Descriptor instead.
func (AgentRpcStatus) EnumDescriptor() ([]byte, []int) {
	return file_src_ray_protobuf_agent_manager_proto_rawDescGZIP(), []int{0}
}

type RegisterAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentPid       int32  `protobuf:"varint,1,opt,name=agent_pid,json=agentPid,proto3" json:"agent_pid,omitempty"`
	AgentPort      int32  `protobuf:"varint,2,opt,name=agent_port,json=agentPort,proto3" json:"agent_port,omitempty"`
	AgentIpAddress string `protobuf:"bytes,3,opt,name=agent_ip_address,json=agentIpAddress,proto3" json:"agent_ip_address,omitempty"`
}

func (x *RegisterAgentRequest) Reset() {
	*x = RegisterAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_agent_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentRequest) ProtoMessage() {}

func (x *RegisterAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_agent_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentRequest.ProtoReflect.Descriptor instead.
func (*RegisterAgentRequest) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_agent_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAgentRequest) GetAgentPid() int32 {
	if x != nil {
		return x.AgentPid
	}
	return 0
}

func (x *RegisterAgentRequest) GetAgentPort() int32 {
	if x != nil {
		return x.AgentPort
	}
	return 0
}

func (x *RegisterAgentRequest) GetAgentIpAddress() string {
	if x != nil {
		return x.AgentIpAddress
	}
	return ""
}

type RegisterAgentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status AgentRpcStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ray.rpc.AgentRpcStatus" json:"status,omitempty"`
}

func (x *RegisterAgentReply) Reset() {
	*x = RegisterAgentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_agent_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentReply) ProtoMessage() {}

func (x *RegisterAgentReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_agent_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentReply.ProtoReflect.Descriptor instead.
func (*RegisterAgentReply) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_agent_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterAgentReply) GetStatus() AgentRpcStatus {
	if x != nil {
		return x.Status
	}
	return AgentRpcStatus_AGENT_RPC_STATUS_OK
}

var File_src_ray_protobuf_agent_manager_proto protoreflect.FileDescriptor

var file_src_ray_protobuf_agent_manager_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x22,
	0x7c, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a,
	0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x46, 0x0a, 0x0e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x52, 0x50, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0x62, 0x0a, 0x13,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x8d, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63,
	0x42, 0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x75, 0x65, 0x69, 0x61, 0x6e, 0x2f, 0x72, 0x61, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x52, 0x52, 0x58, 0xaa, 0x02, 0x07, 0x52, 0x61,
	0x79, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x07, 0x52, 0x61, 0x79, 0x5c, 0x52, 0x70, 0x63, 0xe2,
	0x02, 0x13, 0x52, 0x61, 0x79, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x61, 0x79, 0x3a, 0x3a, 0x52, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_ray_protobuf_agent_manager_proto_rawDescOnce sync.Once
	file_src_ray_protobuf_agent_manager_proto_rawDescData = file_src_ray_protobuf_agent_manager_proto_rawDesc
)

func file_src_ray_protobuf_agent_manager_proto_rawDescGZIP() []byte {
	file_src_ray_protobuf_agent_manager_proto_rawDescOnce.Do(func() {
		file_src_ray_protobuf_agent_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_ray_protobuf_agent_manager_proto_rawDescData)
	})
	return file_src_ray_protobuf_agent_manager_proto_rawDescData
}

var file_src_ray_protobuf_agent_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_ray_protobuf_agent_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_ray_protobuf_agent_manager_proto_goTypes = []interface{}{
	(AgentRpcStatus)(0),          // 0: ray.rpc.AgentRpcStatus
	(*RegisterAgentRequest)(nil), // 1: ray.rpc.RegisterAgentRequest
	(*RegisterAgentReply)(nil),   // 2: ray.rpc.RegisterAgentReply
}
var file_src_ray_protobuf_agent_manager_proto_depIdxs = []int32{
	0, // 0: ray.rpc.RegisterAgentReply.status:type_name -> ray.rpc.AgentRpcStatus
	1, // 1: ray.rpc.AgentManagerService.RegisterAgent:input_type -> ray.rpc.RegisterAgentRequest
	2, // 2: ray.rpc.AgentManagerService.RegisterAgent:output_type -> ray.rpc.RegisterAgentReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_ray_protobuf_agent_manager_proto_init() }
func file_src_ray_protobuf_agent_manager_proto_init() {
	if File_src_ray_protobuf_agent_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_ray_protobuf_agent_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentRequest); i {
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
		file_src_ray_protobuf_agent_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentReply); i {
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
			RawDescriptor: file_src_ray_protobuf_agent_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_ray_protobuf_agent_manager_proto_goTypes,
		DependencyIndexes: file_src_ray_protobuf_agent_manager_proto_depIdxs,
		EnumInfos:         file_src_ray_protobuf_agent_manager_proto_enumTypes,
		MessageInfos:      file_src_ray_protobuf_agent_manager_proto_msgTypes,
	}.Build()
	File_src_ray_protobuf_agent_manager_proto = out.File
	file_src_ray_protobuf_agent_manager_proto_rawDesc = nil
	file_src_ray_protobuf_agent_manager_proto_goTypes = nil
	file_src_ray_protobuf_agent_manager_proto_depIdxs = nil
}