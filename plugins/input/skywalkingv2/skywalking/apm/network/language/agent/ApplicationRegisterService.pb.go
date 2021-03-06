//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: language-agent/ApplicationRegisterService.proto

package agent

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

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationCode string `protobuf:"bytes,1,opt,name=applicationCode,proto3" json:"applicationCode,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_ApplicationRegisterService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_ApplicationRegisterService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_language_agent_ApplicationRegisterService_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetApplicationCode() string {
	if x != nil {
		return x.ApplicationCode
	}
	return ""
}

type ApplicationMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *KeyWithIntegerValue `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *ApplicationMapping) Reset() {
	*x = ApplicationMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_ApplicationRegisterService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationMapping) ProtoMessage() {}

func (x *ApplicationMapping) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_ApplicationRegisterService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationMapping.ProtoReflect.Descriptor instead.
func (*ApplicationMapping) Descriptor() ([]byte, []int) {
	return file_language_agent_ApplicationRegisterService_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationMapping) GetApplication() *KeyWithIntegerValue {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_language_agent_ApplicationRegisterService_proto protoreflect.FileDescriptor

var file_language_agent_ApplicationRegisterService_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x42, 0xa1, 0x01, 0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x4e, 0x6c, 0x6f, 0x67, 0x74, 0x61, 0x69, 0x6c,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x76,
	0x32, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x6d,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x1a, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_language_agent_ApplicationRegisterService_proto_rawDescOnce sync.Once
	file_language_agent_ApplicationRegisterService_proto_rawDescData = file_language_agent_ApplicationRegisterService_proto_rawDesc
)

func file_language_agent_ApplicationRegisterService_proto_rawDescGZIP() []byte {
	file_language_agent_ApplicationRegisterService_proto_rawDescOnce.Do(func() {
		file_language_agent_ApplicationRegisterService_proto_rawDescData = protoimpl.X.CompressGZIP(file_language_agent_ApplicationRegisterService_proto_rawDescData)
	})
	return file_language_agent_ApplicationRegisterService_proto_rawDescData
}

var file_language_agent_ApplicationRegisterService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_language_agent_ApplicationRegisterService_proto_goTypes = []interface{}{
	(*Application)(nil),         // 0: Application
	(*ApplicationMapping)(nil),  // 1: ApplicationMapping
	(*KeyWithIntegerValue)(nil), // 2: KeyWithIntegerValue
}
var file_language_agent_ApplicationRegisterService_proto_depIdxs = []int32{
	2, // 0: ApplicationMapping.application:type_name -> KeyWithIntegerValue
	0, // 1: ApplicationRegisterService.applicationCodeRegister:input_type -> Application
	1, // 2: ApplicationRegisterService.applicationCodeRegister:output_type -> ApplicationMapping
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_language_agent_ApplicationRegisterService_proto_init() }
func file_language_agent_ApplicationRegisterService_proto_init() {
	if File_language_agent_ApplicationRegisterService_proto != nil {
		return
	}
	file_language_agent_KeyWithIntegerValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_language_agent_ApplicationRegisterService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_language_agent_ApplicationRegisterService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationMapping); i {
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
			RawDescriptor: file_language_agent_ApplicationRegisterService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_ApplicationRegisterService_proto_goTypes,
		DependencyIndexes: file_language_agent_ApplicationRegisterService_proto_depIdxs,
		MessageInfos:      file_language_agent_ApplicationRegisterService_proto_msgTypes,
	}.Build()
	File_language_agent_ApplicationRegisterService_proto = out.File
	file_language_agent_ApplicationRegisterService_proto_rawDesc = nil
	file_language_agent_ApplicationRegisterService_proto_goTypes = nil
	file_language_agent_ApplicationRegisterService_proto_depIdxs = nil
}
