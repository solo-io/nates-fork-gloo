// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/virtualservice.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *VirtualServiceDetails) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceDetails)
	if !ok {
		that2, ok := that.(VirtualServiceDetails)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualService(), target.GetVirtualService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPlugins()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPlugins()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPlugins(), target.GetPlugins()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRaw()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRaw()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRaw(), target.GetRaw()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Plugins) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Plugins)
	if !ok {
		that2, ok := that.(Plugins)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtAuth(), target.GetExtAuth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRateLimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimit(), target.GetRateLimit()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthPlugin) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthPlugin)
	if !ok {
		that2, ok := that.(ExtAuthPlugin)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	if strings.Compare(m.GetError(), target.GetError()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RateLimitPlugin) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitPlugin)
	if !ok {
		that2, ok := that.(RateLimitPlugin)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	if strings.Compare(m.GetError(), target.GetError()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RepeatedStrings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RepeatedStrings)
	if !ok {
		that2, ok := that.(RepeatedStrings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if strings.Compare(v, target.GetValues()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *RepeatedRoutes) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RepeatedRoutes)
	if !ok {
		that2, ok := that.(RepeatedRoutes)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetValues()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetValues()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *SslConfigValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslConfigValue)
	if !ok {
		that2, ok := that.(SslConfigValue)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *IngressRateLimitValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IngressRateLimitValue)
	if !ok {
		that2, ok := that.(IngressRateLimitValue)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthInput) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthInput)
	if !ok {
		that2, ok := that.(ExtAuthInput)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetVirtualServiceRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceRequest)
	if !ok {
		that2, ok := that.(GetVirtualServiceRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetVirtualServiceResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceResponse)
	if !ok {
		that2, ok := that.(GetVirtualServiceResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListVirtualServicesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListVirtualServicesRequest)
	if !ok {
		that2, ok := that.(ListVirtualServicesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *ListVirtualServicesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListVirtualServicesResponse)
	if !ok {
		that2, ok := that.(ListVirtualServicesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetVirtualServiceDetails()) != len(target.GetVirtualServiceDetails()) {
		return false
	}
	for idx, v := range m.GetVirtualServiceDetails() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualServiceDetails()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualServiceDetails()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualServiceInputV2) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceInputV2)
	if !ok {
		that2, ok := that.(VirtualServiceInputV2)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisplayName()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisplayName()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisplayName(), target.GetDisplayName()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDomains()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDomains()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDomains(), target.GetDomains()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRoutes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRoutes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRoutes(), target.GetRoutes()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRateLimitConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimitConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimitConfig(), target.GetRateLimitConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtAuthConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtAuthConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtAuthConfig(), target.GetExtAuthConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CreateVirtualServiceRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CreateVirtualServiceRequest)
	if !ok {
		that2, ok := that.(CreateVirtualServiceRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetInputV2()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInputV2()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInputV2(), target.GetInputV2()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CreateVirtualServiceResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CreateVirtualServiceResponse)
	if !ok {
		that2, ok := that.(CreateVirtualServiceResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpdateVirtualServiceRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpdateVirtualServiceRequest)
	if !ok {
		that2, ok := that.(UpdateVirtualServiceRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetInputV2()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInputV2()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInputV2(), target.GetInputV2()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpdateVirtualServiceYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpdateVirtualServiceYamlRequest)
	if !ok {
		that2, ok := that.(UpdateVirtualServiceYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetEditedYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEditedYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEditedYamlData(), target.GetEditedYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpdateVirtualServiceResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpdateVirtualServiceResponse)
	if !ok {
		that2, ok := that.(UpdateVirtualServiceResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DeleteVirtualServiceRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DeleteVirtualServiceRequest)
	if !ok {
		that2, ok := that.(DeleteVirtualServiceRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DeleteVirtualServiceResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DeleteVirtualServiceResponse)
	if !ok {
		that2, ok := that.(DeleteVirtualServiceResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *RouteInput) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteInput)
	if !ok {
		that2, ok := that.(RouteInput)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	if m.GetIndex() != target.GetIndex() {
		return false
	}

	if h, ok := interface{}(m.GetRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRoute(), target.GetRoute()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CreateRouteRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CreateRouteRequest)
	if !ok {
		that2, ok := that.(CreateRouteRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetInput()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInput()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInput(), target.GetInput()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CreateRouteResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CreateRouteResponse)
	if !ok {
		that2, ok := that.(CreateRouteResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpdateRouteRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpdateRouteRequest)
	if !ok {
		that2, ok := that.(UpdateRouteRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetInput()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInput()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInput(), target.GetInput()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpdateRouteResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpdateRouteResponse)
	if !ok {
		that2, ok := that.(UpdateRouteResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DeleteRouteRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DeleteRouteRequest)
	if !ok {
		that2, ok := that.(DeleteRouteRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	if m.GetIndex() != target.GetIndex() {
		return false
	}

	return true
}

// Equal function
func (m *DeleteRouteResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DeleteRouteResponse)
	if !ok {
		that2, ok := that.(DeleteRouteResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SwapRoutesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SwapRoutesRequest)
	if !ok {
		that2, ok := that.(SwapRoutesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	if m.GetIndex1() != target.GetIndex1() {
		return false
	}

	if m.GetIndex2() != target.GetIndex2() {
		return false
	}

	return true
}

// Equal function
func (m *SwapRoutesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SwapRoutesResponse)
	if !ok {
		that2, ok := that.(SwapRoutesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ShiftRoutesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ShiftRoutesRequest)
	if !ok {
		that2, ok := that.(ShiftRoutesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	if m.GetFromIndex() != target.GetFromIndex() {
		return false
	}

	if m.GetToIndex() != target.GetToIndex() {
		return false
	}

	return true
}

// Equal function
func (m *ShiftRoutesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ShiftRoutesResponse)
	if !ok {
		that2, ok := that.(ShiftRoutesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceDetails(), target.GetVirtualServiceDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthInput_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthInput_Config)
	if !ok {
		that2, ok := that.(ExtAuthInput_Config)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Value.(type) {

	case *ExtAuthInput_Config_Oauth:
		if _, ok := target.Value.(*ExtAuthInput_Config_Oauth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOauth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOauth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOauth(), target.GetOauth()) {
				return false
			}
		}

	case *ExtAuthInput_Config_CustomAuth:
		if _, ok := target.Value.(*ExtAuthInput_Config_CustomAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCustomAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCustomAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCustomAuth(), target.GetCustomAuth()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Value != target.Value {
			return false
		}
	}

	return true
}
