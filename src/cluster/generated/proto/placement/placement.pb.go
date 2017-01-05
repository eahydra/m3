// Code generated by protoc-gen-go.
// source: placement.proto
// DO NOT EDIT!

/*
Package placement is a generated protocol buffer package.

It is generated from these files:
	placement.proto

It has these top-level messages:
	Placement
	Instance
	Shard
*/
package placement

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShardState int32

const (
	ShardState_Initializing ShardState = 0
	ShardState_Available    ShardState = 1
	ShardState_Leaving      ShardState = 2
)

var ShardState_name = map[int32]string{
	0: "Initializing",
	1: "Available",
	2: "Leaving",
}
var ShardState_value = map[string]int32{
	"Initializing": 0,
	"Available":    1,
	"Leaving":      2,
}

func (x ShardState) String() string {
	return proto.EnumName(ShardState_name, int32(x))
}
func (ShardState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Placement struct {
	Instances     map[string]*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReplicaFactor uint32               `protobuf:"varint,2,opt,name=replica_factor,json=replicaFactor" json:"replica_factor,omitempty"`
	NumShards     uint32               `protobuf:"varint,3,opt,name=num_shards,json=numShards" json:"num_shards,omitempty"`
}

func (m *Placement) Reset()                    { *m = Placement{} }
func (m *Placement) String() string            { return proto.CompactTextString(m) }
func (*Placement) ProtoMessage()               {}
func (*Placement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Placement) GetInstances() map[string]*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type Instance struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Rack     string   `protobuf:"bytes,2,opt,name=rack" json:"rack,omitempty"`
	Zone     string   `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	Weight   uint32   `protobuf:"varint,4,opt,name=weight" json:"weight,omitempty"`
	Endpoint string   `protobuf:"bytes,5,opt,name=endpoint" json:"endpoint,omitempty"`
	Shards   []*Shard `protobuf:"bytes,6,rep,name=shards" json:"shards,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Instance) GetShards() []*Shard {
	if m != nil {
		return m.Shards
	}
	return nil
}

type Shard struct {
	Id       uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	State    ShardState `protobuf:"varint,2,opt,name=state,enum=placement.ShardState" json:"state,omitempty"`
	SourceId string     `protobuf:"bytes,3,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Placement)(nil), "placement.placement")
	proto.RegisterType((*Instance)(nil), "placement.instance")
	proto.RegisterType((*Shard)(nil), "placement.shard")
	proto.RegisterEnum("placement.ShardState", ShardState_name, ShardState_value)
}

func init() { proto.RegisterFile("placement.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x0a, 0xd3, 0x40,
	0x10, 0x86, 0xdd, 0xb4, 0x89, 0xdd, 0x89, 0x89, 0x61, 0x84, 0x12, 0x2a, 0x42, 0xa9, 0x08, 0x51,
	0xa4, 0x87, 0x7a, 0x11, 0x3d, 0xf5, 0xa0, 0x50, 0xf0, 0xe2, 0xfa, 0x00, 0x61, 0x9b, 0xac, 0xed,
	0xd2, 0x74, 0x13, 0x92, 0x4d, 0xa5, 0x7d, 0x1a, 0xdf, 0xce, 0xd7, 0x90, 0x4e, 0xd2, 0x36, 0x7a,
	0xfb, 0xe7, 0x9f, 0x7f, 0x66, 0xbf, 0x09, 0x81, 0xe7, 0x55, 0x21, 0x33, 0x75, 0x54, 0xc6, 0x2e,
	0xab, 0xba, 0xb4, 0x25, 0xf2, 0xbb, 0xb1, 0xf8, 0xc3, 0xe0, 0x51, 0xe1, 0x1a, 0xb8, 0x36, 0x8d,
	0x95, 0x26, 0x53, 0x4d, 0xcc, 0xe6, 0xa3, 0xc4, 0x5f, 0xbd, 0x5e, 0x0e, 0xa6, 0xef, 0x6a, 0x73,
	0x4b, 0x7d, 0x31, 0xb6, 0x3e, 0x8b, 0xc7, 0x14, 0xbe, 0x81, 0xb0, 0x56, 0x55, 0xa1, 0x33, 0x99,
	0xfe, 0x94, 0x99, 0x2d, 0xeb, 0xd8, 0x99, 0xb3, 0x24, 0x10, 0x41, 0xef, 0x7e, 0x25, 0x13, 0x5f,
	0x01, 0x98, 0xf6, 0x98, 0x36, 0x7b, 0x59, 0xe7, 0x4d, 0x3c, 0xa2, 0x08, 0x37, 0xed, 0xf1, 0x07,
	0x19, 0xb3, 0xef, 0x10, 0xfe, 0xfb, 0x04, 0x46, 0x30, 0x3a, 0xa8, 0x73, 0xcc, 0xe6, 0x2c, 0xe1,
	0xe2, 0x2a, 0xf1, 0x2d, 0xb8, 0x27, 0x59, 0xb4, 0x8a, 0x1e, 0xf0, 0x57, 0x2f, 0x06, 0x78, 0x37,
	0x1c, 0xd1, 0x25, 0x3e, 0x39, 0x1f, 0xd9, 0xe2, 0x37, 0x83, 0xc9, 0xcd, 0xc7, 0x10, 0x1c, 0x9d,
	0xf7, 0xcb, 0x1c, 0x9d, 0x23, 0xc2, 0xb8, 0x96, 0xd9, 0x81, 0x56, 0x71, 0x41, 0xfa, 0xea, 0x5d,
	0x4a, 0xa3, 0x08, 0x8e, 0x0b, 0xd2, 0x38, 0x05, 0xef, 0x97, 0xd2, 0xbb, 0xbd, 0x8d, 0xc7, 0x84,
	0xdc, 0x57, 0x38, 0x83, 0x89, 0x32, 0x79, 0x55, 0x6a, 0x63, 0x63, 0x97, 0xf2, 0xf7, 0x1a, 0x13,
	0xf0, 0xfa, 0x33, 0x3d, 0xfa, 0xa2, 0xd1, 0x00, 0x94, 0x1a, 0xa2, 0xef, 0x2f, 0xb6, 0xe0, 0x92,
	0x1a, 0xe0, 0x05, 0x84, 0xf7, 0x1e, 0xdc, 0xc6, 0x4a, 0xdb, 0x9d, 0x1a, 0xae, 0xa6, 0xff, 0x6f,
	0x48, 0xa9, 0x2b, 0xba, 0x10, 0xbe, 0x04, 0xde, 0x94, 0x6d, 0x9d, 0xa9, 0x54, 0xe7, 0x3d, 0xfd,
	0xa4, 0x33, 0x36, 0xf9, 0xbb, 0xcf, 0xe0, 0x0f, 0x46, 0x30, 0x82, 0x67, 0x1b, 0xa3, 0xad, 0x96,
	0x85, 0xbe, 0x68, 0xb3, 0x8b, 0x9e, 0x60, 0x00, 0x7c, 0x7d, 0x92, 0xba, 0x90, 0xdb, 0x42, 0x45,
	0x0c, 0x7d, 0x78, 0xfa, 0x4d, 0xc9, 0xd3, 0xb5, 0xe7, 0x6c, 0x3d, 0xfa, 0x7f, 0x3e, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0x52, 0xa0, 0x4c, 0x52, 0x02, 0x00, 0x00,
}