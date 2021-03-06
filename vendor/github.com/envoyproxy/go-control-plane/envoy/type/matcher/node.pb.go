// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/node.proto

package envoy_type_matcher

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeMatcher struct {
	NodeId               *StringMatcher   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeMetadatas        []*StructMatcher `protobuf:"bytes,2,rep,name=node_metadatas,json=nodeMetadatas,proto3" json:"node_metadatas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeMatcher) Reset()         { *m = NodeMatcher{} }
func (m *NodeMatcher) String() string { return proto.CompactTextString(m) }
func (*NodeMatcher) ProtoMessage()    {}
func (*NodeMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20314fb2f725fb2, []int{0}
}

func (m *NodeMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMatcher.Unmarshal(m, b)
}
func (m *NodeMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMatcher.Marshal(b, m, deterministic)
}
func (m *NodeMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMatcher.Merge(m, src)
}
func (m *NodeMatcher) XXX_Size() int {
	return xxx_messageInfo_NodeMatcher.Size(m)
}
func (m *NodeMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMatcher proto.InternalMessageInfo

func (m *NodeMatcher) GetNodeId() *StringMatcher {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeMatcher) GetNodeMetadatas() []*StructMatcher {
	if m != nil {
		return m.NodeMetadatas
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeMatcher)(nil), "envoy.type.matcher.NodeMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/node.proto", fileDescriptor_c20314fb2f725fb2) }

var fileDescriptor_c20314fb2f725fb2 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x4b, 0xeb, 0x81, 0xa4,
	0xf5, 0xa0, 0xd2, 0x52, 0xf2, 0x58, 0xb4, 0x14, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x43, 0x34, 0xe1,
	0x52, 0x50, 0x9a, 0x5c, 0x02, 0x51, 0xa0, 0x34, 0x99, 0x91, 0x8b, 0xdb, 0x2f, 0x3f, 0x25, 0xd5,
	0x17, 0x22, 0x29, 0x64, 0xc5, 0xc5, 0x0e, 0xb2, 0x33, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x51, 0x0f, 0xd3, 0x5e, 0xbd, 0x60, 0xb0, 0x1d, 0x50, 0x3d, 0x41, 0x6c, 0x20,
	0x1d, 0x9e, 0x29, 0x42, 0x1e, 0x5c, 0x7c, 0x60, 0xbd, 0xb9, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25,
	0x89, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x78, 0x8c, 0x28, 0x4d, 0x2e, 0x81, 0x19, 0xc1, 0x0b, 0xd2,
	0xe8, 0x0b, 0xd3, 0xe7, 0xa4, 0xcf, 0xa5, 0x90, 0x99, 0x0f, 0xd1, 0x55, 0x50, 0x94, 0x5f, 0x51,
	0x89, 0xc5, 0x00, 0x27, 0x4e, 0x90, 0xb3, 0x03, 0x40, 0x9e, 0x08, 0x60, 0x4c, 0x62, 0x03, 0xfb,
	0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xd0, 0x05, 0xde, 0x44, 0x01, 0x00, 0x00,
}
