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
// source: src/ray/protobuf/job_agent.proto

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

type InitializeJobEnvRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The job description in JSON.
	JobDescription string `protobuf:"bytes,1,opt,name=job_description,json=jobDescription,proto3" json:"job_description,omitempty"`
}

func (x *InitializeJobEnvRequest) Reset() {
	*x = InitializeJobEnvRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_job_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeJobEnvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeJobEnvRequest) ProtoMessage() {}

func (x *InitializeJobEnvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_job_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeJobEnvRequest.ProtoReflect.Descriptor instead.
func (*InitializeJobEnvRequest) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_job_agent_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeJobEnvRequest) GetJobDescription() string {
	if x != nil {
		return x.JobDescription
	}
	return ""
}

type InitializeJobEnvReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of job env initialization.
	Status AgentRpcStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ray.rpc.AgentRpcStatus" json:"status,omitempty"`
	// The error message in InitializeJobEnv.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// The driver pid of job.
	DriverPid int32 `protobuf:"varint,3,opt,name=driver_pid,json=driverPid,proto3" json:"driver_pid,omitempty"`
}

func (x *InitializeJobEnvReply) Reset() {
	*x = InitializeJobEnvReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_ray_protobuf_job_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeJobEnvReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeJobEnvReply) ProtoMessage() {}

func (x *InitializeJobEnvReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_ray_protobuf_job_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeJobEnvReply.ProtoReflect.Descriptor instead.
func (*InitializeJobEnvReply) Descriptor() ([]byte, []int) {
	return file_src_ray_protobuf_job_agent_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeJobEnvReply) GetStatus() AgentRpcStatus {
	if x != nil {
		return x.Status
	}
	return AgentRpcStatus_AGENT_RPC_STATUS_OK
}

func (x *InitializeJobEnvReply) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *InitializeJobEnvReply) GetDriverPid() int32 {
	if x != nil {
		return x.DriverPid
	}
	return 0
}

var File_src_ray_protobuf_job_agent_proto protoreflect.FileDescriptor

var file_src_ray_protobuf_job_agent_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x24, 0x73, 0x72, 0x63,
	0x2f, 0x72, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4a,
	0x6f, 0x62, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x4a, 0x6f, 0x62, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x50, 0x69, 0x64, 0x32, 0x67, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4a, 0x6f, 0x62, 0x45, 0x6e, 0x76, 0x12, 0x20, 0x2e, 0x72, 0x61,
	0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x4a, 0x6f, 0x62, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x4a, 0x6f, 0x62, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x89, 0x01,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0d, 0x4a,
	0x6f, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
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
	file_src_ray_protobuf_job_agent_proto_rawDescOnce sync.Once
	file_src_ray_protobuf_job_agent_proto_rawDescData = file_src_ray_protobuf_job_agent_proto_rawDesc
)

func file_src_ray_protobuf_job_agent_proto_rawDescGZIP() []byte {
	file_src_ray_protobuf_job_agent_proto_rawDescOnce.Do(func() {
		file_src_ray_protobuf_job_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_ray_protobuf_job_agent_proto_rawDescData)
	})
	return file_src_ray_protobuf_job_agent_proto_rawDescData
}

var file_src_ray_protobuf_job_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_ray_protobuf_job_agent_proto_goTypes = []interface{}{
	(*InitializeJobEnvRequest)(nil), // 0: ray.rpc.InitializeJobEnvRequest
	(*InitializeJobEnvReply)(nil),   // 1: ray.rpc.InitializeJobEnvReply
	(AgentRpcStatus)(0),             // 2: ray.rpc.AgentRpcStatus
}
var file_src_ray_protobuf_job_agent_proto_depIdxs = []int32{
	2, // 0: ray.rpc.InitializeJobEnvReply.status:type_name -> ray.rpc.AgentRpcStatus
	0, // 1: ray.rpc.JobAgentService.InitializeJobEnv:input_type -> ray.rpc.InitializeJobEnvRequest
	1, // 2: ray.rpc.JobAgentService.InitializeJobEnv:output_type -> ray.rpc.InitializeJobEnvReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_ray_protobuf_job_agent_proto_init() }
func file_src_ray_protobuf_job_agent_proto_init() {
	if File_src_ray_protobuf_job_agent_proto != nil {
		return
	}
	file_src_ray_protobuf_agent_manager_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_ray_protobuf_job_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeJobEnvRequest); i {
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
		file_src_ray_protobuf_job_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeJobEnvReply); i {
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
			RawDescriptor: file_src_ray_protobuf_job_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_ray_protobuf_job_agent_proto_goTypes,
		DependencyIndexes: file_src_ray_protobuf_job_agent_proto_depIdxs,
		MessageInfos:      file_src_ray_protobuf_job_agent_proto_msgTypes,
	}.Build()
	File_src_ray_protobuf_job_agent_proto = out.File
	file_src_ray_protobuf_job_agent_proto_rawDesc = nil
	file_src_ray_protobuf_job_agent_proto_goTypes = nil
	file_src_ray_protobuf_job_agent_proto_depIdxs = nil
}