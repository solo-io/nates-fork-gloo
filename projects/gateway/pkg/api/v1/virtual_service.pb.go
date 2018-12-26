// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/gateway/api/v1/virtual_service.proto

package v1 // import "github.com/solo-io/solo-projects/projects/gateway/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
import v1 "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=vs
// @solo-kit:resource.plural_name=virtual_services
// @solo-kit:resource.resource_groups=api.gateway.solo.io
//
// A virtual service describes the set of routes to match for a set of domains.
//
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *v1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_service_d2ce03de8d2b6452, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (dst *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(dst, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *v1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}
func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_virtual_service_d2ce03de8d2b6452)
}

var fileDescriptor_virtual_service_d2ce03de8d2b6452 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0x3f, 0xc8, 0x17, 0xd4, 0xc5, 0x68, 0x6c, 0x88, 0x16, 0x0e, 0x62, 0x38, 0x79, 0x71,
	0x37, 0x68, 0x62, 0x8c, 0xe1, 0x62, 0x3d, 0xe8, 0x45, 0x0f, 0x25, 0xf1, 0xe0, 0x85, 0x2c, 0x65,
	0x59, 0x56, 0x0a, 0xd3, 0xec, 0x0e, 0x55, 0x1e, 0xc1, 0x37, 0xf1, 0x51, 0x7c, 0x0a, 0x0e, 0x3e,
	0x82, 0x4f, 0x60, 0xba, 0xdd, 0x42, 0x38, 0x98, 0xc8, 0xa9, 0xd3, 0xcc, 0xfc, 0x7e, 0x93, 0xff,
	0x64, 0xc9, 0xa3, 0x54, 0x38, 0x9a, 0xf5, 0x69, 0x04, 0x13, 0x66, 0x20, 0x86, 0x33, 0x05, 0xf9,
	0x37, 0xd1, 0xf0, 0x22, 0x22, 0x34, 0x6c, 0x59, 0x48, 0x8e, 0xe2, 0x95, 0xcf, 0x19, 0x4f, 0x14,
	0x4b, 0xdb, 0x2c, 0x55, 0x1a, 0x67, 0x3c, 0xee, 0x19, 0xa1, 0x53, 0x15, 0x09, 0x9a, 0x68, 0x40,
	0xf0, 0xf6, 0xdd, 0x14, 0xcd, 0x24, 0x54, 0x41, 0xa3, 0x26, 0x41, 0x82, 0xed, 0xb1, 0xac, 0xca,
	0xc7, 0x1a, 0xed, 0xdf, 0xd6, 0x8e, 0x15, 0x16, 0x0b, 0x26, 0x02, 0xf9, 0x80, 0x23, 0x77, 0x08,
	0xfb, 0x03, 0x62, 0x90, 0xe3, 0xcc, 0x38, 0xe0, 0x66, 0x83, 0x68, 0x31, 0x40, 0xe1, 0x48, 0x34,
	0xbc, 0xcd, 0x73, 0x45, 0xeb, 0xbd, 0x4c, 0xf6, 0x9e, 0xf2, 0x9c, 0xdd, 0x3c, 0xa6, 0xd7, 0x21,
	0xbb, 0x45, 0xf2, 0x11, 0x18, 0xf4, 0x4b, 0x27, 0xa5, 0xd3, 0xea, 0x79, 0x9d, 0x66, 0x8a, 0x22,
	0x34, 0x75, 0xcc, 0x3d, 0x18, 0x0c, 0xab, 0xe9, 0xea, 0xc7, 0xbb, 0x24, 0xc4, 0x98, 0xb8, 0x17,
	0xc1, 0x74, 0xa8, 0xa4, 0x5f, 0xb6, 0xec, 0xd1, 0x3a, 0xdb, 0x35, 0xf1, 0xad, 0x6d, 0x87, 0x3b,
	0xa6, 0x28, 0xbd, 0x3b, 0x52, 0xc9, 0xb3, 0xf9, 0x15, 0xcb, 0xd4, 0x68, 0x04, 0x5a, 0xac, 0x18,
	0xdb, 0x0b, 0xea, 0x9f, 0x8b, 0xe6, 0xbf, 0xef, 0x45, 0xf3, 0x00, 0x85, 0xc1, 0x81, 0x1a, 0x0e,
	0xaf, 0x5b, 0x4a, 0x4e, 0x41, 0x8b, 0x56, 0xe8, 0x70, 0xef, 0x8a, 0x6c, 0x17, 0x77, 0xf5, 0xb7,
	0xac, 0xea, 0x70, 0x5d, 0xf5, 0xe0, 0xba, 0xc1, 0xff, 0x4c, 0x16, 0x2e, 0xa7, 0x83, 0xe0, 0xe3,
	0xeb, 0xb8, 0xf4, 0xdc, 0xd9, 0xfc, 0xbd, 0x24, 0x63, 0xe9, 0x6e, 0xdb, 0xaf, 0xd8, 0xb3, 0x5e,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xf9, 0x7a, 0x95, 0x76, 0x02, 0x00, 0x00,
}
