// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc.proto

package pb

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

type HLC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Counter   int64  `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	NodeID    uint32 `protobuf:"varint,3,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (x *HLC) Reset() {
	*x = HLC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLC) ProtoMessage() {}

func (x *HLC) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLC.ProtoReflect.Descriptor instead.
func (*HLC) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *HLC) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HLC) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *HLC) GetNodeID() uint32 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

type MergeableString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Point *HLC   `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *MergeableString) Reset() {
	*x = MergeableString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeableString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeableString) ProtoMessage() {}

func (x *MergeableString) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeableString.ProtoReflect.Descriptor instead.
func (*MergeableString) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MergeableString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MergeableString) GetPoint() *HLC {
	if x != nil {
		return x.Point
	}
	return nil
}

type MergeableMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point      *HLC                        `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
	Deleted    bool                        `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Map        map[string]*MergeableString `protobuf:"bytes,3,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tombstones map[string]*HLC             `protobuf:"bytes,4,rep,name=tombstones,proto3" json:"tombstones,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MergeableMap) Reset() {
	*x = MergeableMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeableMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeableMap) ProtoMessage() {}

func (x *MergeableMap) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeableMap.ProtoReflect.Descriptor instead.
func (*MergeableMap) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *MergeableMap) GetPoint() *HLC {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *MergeableMap) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *MergeableMap) GetMap() map[string]*MergeableString {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *MergeableMap) GetTombstones() map[string]*HLC {
	if x != nil {
		return x.Tombstones
	}
	return nil
}

type DocumentCRDTState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string        `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	DocId     string        `protobuf:"bytes,2,opt,name=docId,proto3" json:"docId,omitempty"`
	Map       *MergeableMap `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
}

func (x *DocumentCRDTState) Reset() {
	*x = DocumentCRDTState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentCRDTState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentCRDTState) ProtoMessage() {}

func (x *DocumentCRDTState) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentCRDTState.ProtoReflect.Descriptor instead.
func (*DocumentCRDTState) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DocumentCRDTState) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *DocumentCRDTState) GetDocId() string {
	if x != nil {
		return x.DocId
	}
	return ""
}

func (x *DocumentCRDTState) GetMap() *MergeableMap {
	if x != nil {
		return x.Map
	}
	return nil
}

type MergeCRDTStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*DocumentCRDTState `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *MergeCRDTStatesRequest) Reset() {
	*x = MergeCRDTStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeCRDTStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeCRDTStatesRequest) ProtoMessage() {}

func (x *MergeCRDTStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeCRDTStatesRequest.ProtoReflect.Descriptor instead.
func (*MergeCRDTStatesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *MergeCRDTStatesRequest) GetDocuments() []*DocumentCRDTState {
	if x != nil {
		return x.Documents
	}
	return nil
}

type MergeCRDTStatesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MergeCRDTStatesReply) Reset() {
	*x = MergeCRDTStatesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeCRDTStatesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeCRDTStatesReply) ProtoMessage() {}

func (x *MergeCRDTStatesReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeCRDTStatesReply.ProtoReflect.Descriptor instead.
func (*MergeCRDTStatesReply) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{5}
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x61, 0x6e,
	0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0x55, 0x0a, 0x03, 0x48,
	0x4c, 0x43, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x22, 0x52, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e,
	0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x48, 0x4c, 0x43, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x87, 0x03, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64,
	0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x48, 0x4c, 0x43, 0x52, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x03,
	0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x61, 0x6e, 0x63,
	0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x4c, 0x0a, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x61, 0x6e, 0x63,
	0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x54, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x73, 0x1a, 0x57, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f,
	0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x0f,
	0x54, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f,
	0x73, 0x2e, 0x48, 0x4c, 0x43, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x77, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x52, 0x44, 0x54,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x6d, 0x61, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64,
	0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x61, 0x70, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x22, 0x59, 0x0a, 0x16, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x43, 0x52, 0x44, 0x54, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64,
	0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x52, 0x44, 0x54, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x52, 0x44,
	0x54, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x6d, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x43, 0x52, 0x44, 0x54, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x62, 0x61,
	0x6e, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x43, 0x52, 0x44, 0x54, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64,
	0x61, 0x64, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x52, 0x44, 0x54, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x62,
	0x61, 0x6e, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_proto_goTypes = []interface{}{
	(*HLC)(nil),                    // 0: banco_de_dados.HLC
	(*MergeableString)(nil),        // 1: banco_de_dados.MergeableString
	(*MergeableMap)(nil),           // 2: banco_de_dados.MergeableMap
	(*DocumentCRDTState)(nil),      // 3: banco_de_dados.DocumentCRDTState
	(*MergeCRDTStatesRequest)(nil), // 4: banco_de_dados.MergeCRDTStatesRequest
	(*MergeCRDTStatesReply)(nil),   // 5: banco_de_dados.MergeCRDTStatesReply
	nil,                            // 6: banco_de_dados.MergeableMap.MapEntry
	nil,                            // 7: banco_de_dados.MergeableMap.TombstonesEntry
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: banco_de_dados.MergeableString.point:type_name -> banco_de_dados.HLC
	0, // 1: banco_de_dados.MergeableMap.point:type_name -> banco_de_dados.HLC
	6, // 2: banco_de_dados.MergeableMap.map:type_name -> banco_de_dados.MergeableMap.MapEntry
	7, // 3: banco_de_dados.MergeableMap.tombstones:type_name -> banco_de_dados.MergeableMap.TombstonesEntry
	2, // 4: banco_de_dados.DocumentCRDTState.map:type_name -> banco_de_dados.MergeableMap
	3, // 5: banco_de_dados.MergeCRDTStatesRequest.documents:type_name -> banco_de_dados.DocumentCRDTState
	1, // 6: banco_de_dados.MergeableMap.MapEntry.value:type_name -> banco_de_dados.MergeableString
	0, // 7: banco_de_dados.MergeableMap.TombstonesEntry.value:type_name -> banco_de_dados.HLC
	4, // 8: banco_de_dados.Database.MergeCRDTStates:input_type -> banco_de_dados.MergeCRDTStatesRequest
	5, // 9: banco_de_dados.Database.MergeCRDTStates:output_type -> banco_de_dados.MergeCRDTStatesReply
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLC); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeableString); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeableMap); i {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentCRDTState); i {
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
		file_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeCRDTStatesRequest); i {
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
		file_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeCRDTStatesReply); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}