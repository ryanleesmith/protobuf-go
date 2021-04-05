// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/messageset/messagesetpb/message_set.proto

package messagesetpb

import (
	protoreflect "github.com/ryanleesmith/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/ryanleesmith/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type MessageSet struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields
}

func (x *MessageSet) Reset() {
	*x = MessageSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSet) ProtoMessage() {}

func (x *MessageSet) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSet.ProtoReflect.Descriptor instead.
func (*MessageSet) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescGZIP(), []int{0}
}

type MessageSetContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSet *MessageSet `protobuf:"bytes,1,opt,name=message_set,json=messageSet" json:"message_set,omitempty"`
}

func (x *MessageSetContainer) Reset() {
	*x = MessageSetContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSetContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSetContainer) ProtoMessage() {}

func (x *MessageSetContainer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSetContainer.ProtoReflect.Descriptor instead.
func (*MessageSetContainer) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescGZIP(), []int{1}
}

func (x *MessageSetContainer) GetMessageSet() *MessageSet {
	if x != nil {
		return x.MessageSet
	}
	return nil
}

var File_internal_testprotos_messageset_messagesetpb_message_set_proto protoreflect.FileDescriptor

var file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x22, 0x1a, 0x0a, 0x0a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x2a, 0x08, 0x08, 0x04, 0x10, 0xff, 0xff, 0xff, 0xff,
	0x07, 0x3a, 0x02, 0x08, 0x01, 0x22, 0x5c, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62,
}

var (
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescOnce sync.Once
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescData = file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDesc
)

func file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescGZIP() []byte {
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescData)
	})
	return file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDescData
}

var file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_testprotos_messageset_messagesetpb_message_set_proto_goTypes = []interface{}{
	(*MessageSet)(nil),          // 0: goproto.proto.messageset.MessageSet
	(*MessageSetContainer)(nil), // 1: goproto.proto.messageset.MessageSetContainer
}
var file_internal_testprotos_messageset_messagesetpb_message_set_proto_depIdxs = []int32{
	0, // 0: goproto.proto.messageset.MessageSetContainer.message_set:type_name -> goproto.proto.messageset.MessageSet
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_testprotos_messageset_messagesetpb_message_set_proto_init() }
func file_internal_testprotos_messageset_messagesetpb_message_set_proto_init() {
	if File_internal_testprotos_messageset_messagesetpb_message_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSetContainer); i {
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
			RawDescriptor: file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_messageset_messagesetpb_message_set_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_messageset_messagesetpb_message_set_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_messageset_messagesetpb_message_set_proto_msgTypes,
	}.Build()
	File_internal_testprotos_messageset_messagesetpb_message_set_proto = out.File
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_rawDesc = nil
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_goTypes = nil
	file_internal_testprotos_messageset_messagesetpb_message_set_proto_depIdxs = nil
}
