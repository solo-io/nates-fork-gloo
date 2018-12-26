// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/gateway/api/v1/gateway.proto

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
// @solo-kit:resource.short_name=gw
// @solo-kit:resource.plural_name=gateways
// @solo-kit:resource.resource_groups=api.gateway.solo.io
//
// A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// names of the the virtual services, which contain the actual routes for the gateway
	// if the list is empty, the gateway will apply all virtual services to this gateway
	VirtualServices []core.ResourceRef `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways in a namespace
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level plugin configuration for all routes on the gateway
	Plugins *v1.ListenerPlugins `protobuf:"bytes,5,opt,name=plugins" json:"plugins,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_5740f57dcb851702, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (dst *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(dst, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetVirtualServices() []core.ResourceRef {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Gateway) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *Gateway) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *Gateway) GetPlugins() *v1.ListenerPlugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *Gateway) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Gateway) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
}
func (this *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
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
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(&that1.VirtualServices[i]) {
			return false
		}
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.Plugins.Equal(that1.Plugins) {
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
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/gateway/api/v1/gateway.proto", fileDescriptor_gateway_5740f57dcb851702)
}

var fileDescriptor_gateway_5740f57dcb851702 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4f, 0xe2, 0x40,
	0x18, 0xc6, 0xb7, 0xc0, 0xf2, 0x67, 0xd8, 0x0d, 0xbb, 0x0d, 0xd9, 0x14, 0x36, 0xbb, 0x54, 0x4e,
	0x3d, 0x68, 0x1b, 0xf0, 0xa0, 0x31, 0x5e, 0xa8, 0x89, 0x24, 0x46, 0x13, 0x52, 0x6e, 0x5e, 0xc8,
	0xd0, 0x4e, 0xeb, 0x48, 0xe9, 0xdb, 0xcc, 0x4c, 0x51, 0xbf, 0x91, 0x1f, 0xc5, 0xab, 0x5f, 0x80,
	0x83, 0x1f, 0xc1, 0x4f, 0x60, 0xda, 0x4e, 0x49, 0x38, 0x98, 0x80, 0xa7, 0xce, 0xcc, 0xfb, 0xfc,
	0x9e, 0x3e, 0xf3, 0xbe, 0x83, 0x2e, 0x03, 0x2a, 0xee, 0x92, 0xb9, 0xe9, 0xc2, 0xd2, 0xe2, 0x10,
	0xc2, 0x11, 0x85, 0xfc, 0x1b, 0x33, 0xb8, 0x27, 0xae, 0xe0, 0xd6, 0x66, 0x11, 0x60, 0x41, 0x1e,
	0xf0, 0x93, 0x85, 0x63, 0x6a, 0xad, 0x06, 0xc5, 0xd6, 0x8c, 0x19, 0x08, 0x50, 0x5b, 0xc5, 0x36,
	0x85, 0x4d, 0x0a, 0xdd, 0x76, 0x00, 0x01, 0x64, 0x35, 0x2b, 0x5d, 0xe5, 0xb2, 0xee, 0xe0, 0xb3,
	0xdf, 0x2d, 0xa8, 0x28, 0x8c, 0x97, 0x44, 0x60, 0x0f, 0x0b, 0x2c, 0x11, 0x6b, 0x07, 0x84, 0x0b,
	0x2c, 0x12, 0x2e, 0x81, 0xc3, 0x1d, 0x00, 0x46, 0x7c, 0xa9, 0x1e, 0xed, 0xd1, 0x80, 0x10, 0xa0,
	0x30, 0x88, 0x19, 0x3c, 0xca, 0xbb, 0x77, 0x2f, 0xbe, 0x68, 0x11, 0x26, 0x01, 0x8d, 0x64, 0xea,
	0xfe, 0x6b, 0x09, 0xd5, 0xc6, 0x79, 0x0f, 0xd5, 0x2b, 0xf4, 0x6b, 0x45, 0x99, 0x48, 0x70, 0x38,
	0xe3, 0x84, 0xad, 0xa8, 0x4b, 0xb8, 0x56, 0xd2, 0xcb, 0x46, 0x73, 0xd8, 0x31, 0x5d, 0x60, 0xa4,
	0x68, 0xb2, 0xe9, 0x10, 0x0e, 0x09, 0x73, 0x89, 0x43, 0x7c, 0xbb, 0xf2, 0xb2, 0xee, 0x7d, 0x73,
	0x5a, 0x12, 0x9c, 0x4a, 0x4e, 0x3d, 0x40, 0x3f, 0xe6, 0x34, 0xf2, 0x66, 0xd8, 0xf3, 0x18, 0xe1,
	0x5c, 0x2b, 0xeb, 0x8a, 0xd1, 0x70, 0x9a, 0xe9, 0xd9, 0x28, 0x3f, 0x52, 0xff, 0xa2, 0x46, 0x26,
	0x89, 0x81, 0x09, 0xad, 0xa2, 0x2b, 0xc6, 0x4f, 0xa7, 0x9e, 0x1e, 0x4c, 0x80, 0x09, 0xf5, 0x04,
	0xd5, 0x64, 0x50, 0xed, 0xbb, 0xae, 0x18, 0xcd, 0xe1, 0x3f, 0x33, 0xbd, 0xc4, 0x26, 0xc2, 0x35,
	0xe5, 0x82, 0x44, 0x84, 0x4d, 0x72, 0x91, 0x53, 0xa8, 0xd5, 0x31, 0xaa, 0xe6, 0x63, 0xd1, 0xaa,
	0x19, 0xd7, 0xde, 0x8e, 0x3e, 0xcd, 0x6a, 0x76, 0x27, 0x4d, 0xfd, 0xbe, 0xee, 0xfd, 0x16, 0x84,
	0x0b, 0x8f, 0xfa, 0xfe, 0x59, 0x9f, 0x06, 0x11, 0x30, 0xd2, 0x77, 0x24, 0xae, 0x9e, 0xa2, 0x7a,
	0xf1, 0x24, 0xb4, 0x5a, 0x66, 0xf5, 0x67, 0xdb, 0xea, 0x46, 0x56, 0x65, 0x0b, 0x36, 0x6a, 0xdb,
	0x7e, 0x7e, 0xfb, 0xaf, 0xdc, 0x9e, 0xef, 0xff, 0xc4, 0xe3, 0x45, 0x20, 0xa7, 0x34, 0xaf, 0x66,
	0xe3, 0x39, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x71, 0xfd, 0xd1, 0x9c, 0x29, 0x03, 0x00, 0x00,
}
