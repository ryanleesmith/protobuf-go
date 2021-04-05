// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: benchmarks.proto

package benchmarks

import (
	protoreflect "github.com/ryanleesmith/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/ryanleesmith/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type BenchmarkDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the benchmark dataset.  This should be unique across all datasets.
	// Should only contain word characters: [a-zA-Z0-9_]
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fully-qualified name of the protobuf message for this dataset.
	// It will be one of the messages defined benchmark_messages_proto2.proto
	// or benchmark_messages_proto3.proto.
	//
	// Implementations that do not support reflection can implement this with
	// an explicit "if/else" chain that lists every known message defined
	// in those files.
	MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	// The payload(s) for this dataset.  They should be parsed or serialized
	// in sequence, in a loop, ie.
	//
	//  while (!benchmarkDone) {  // Benchmark runner decides when to exit.
	//    for (i = 0; i < benchmark.payload.length; i++) {
	//      parse(benchmark.payload[i])
	//    }
	//  }
	//
	// This is intended to let datasets include a variety of data to provide
	// potentially more realistic results than just parsing the same message
	// over and over.  A single message parsed repeatedly could yield unusually
	// good branch prediction performance.
	Payload [][]byte `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload,omitempty"`
}

func (x *BenchmarkDataset) Reset() {
	*x = BenchmarkDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmarks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkDataset) ProtoMessage() {}

func (x *BenchmarkDataset) ProtoReflect() protoreflect.Message {
	mi := &file_benchmarks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkDataset.ProtoReflect.Descriptor instead.
func (*BenchmarkDataset) Descriptor() ([]byte, []int) {
	return file_benchmarks_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkDataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenchmarkDataset) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

func (x *BenchmarkDataset) GetPayload() [][]byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_benchmarks_proto protoreflect.FileDescriptor

var file_benchmarks_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x63,
	0x0a, 0x10, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x20, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmarks_proto_rawDescOnce sync.Once
	file_benchmarks_proto_rawDescData = file_benchmarks_proto_rawDesc
)

func file_benchmarks_proto_rawDescGZIP() []byte {
	file_benchmarks_proto_rawDescOnce.Do(func() {
		file_benchmarks_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmarks_proto_rawDescData)
	})
	return file_benchmarks_proto_rawDescData
}

var file_benchmarks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_benchmarks_proto_goTypes = []interface{}{
	(*BenchmarkDataset)(nil), // 0: benchmarks.BenchmarkDataset
}
var file_benchmarks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_benchmarks_proto_init() }
func file_benchmarks_proto_init() {
	if File_benchmarks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmarks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkDataset); i {
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
			RawDescriptor: file_benchmarks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmarks_proto_goTypes,
		DependencyIndexes: file_benchmarks_proto_depIdxs,
		MessageInfos:      file_benchmarks_proto_msgTypes,
	}.Build()
	File_benchmarks_proto = out.File
	file_benchmarks_proto_rawDesc = nil
	file_benchmarks_proto_goTypes = nil
	file_benchmarks_proto_depIdxs = nil
}
