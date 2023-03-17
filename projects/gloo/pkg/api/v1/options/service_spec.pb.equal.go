// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/service_spec.proto

package options

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
func (m *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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

	switch m.PluginType.(type) {

	case *ServiceSpec_Rest:
		if _, ok := target.PluginType.(*ServiceSpec_Rest); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRest()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRest()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRest(), target.GetRest()) {
				return false
			}
		}

	case *ServiceSpec_Grpc:
		if _, ok := target.PluginType.(*ServiceSpec_Grpc); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpc(), target.GetGrpc()) {
				return false
			}
		}

	case *ServiceSpec_GrpcJsonTranscoder:
		if _, ok := target.PluginType.(*ServiceSpec_GrpcJsonTranscoder); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpcJsonTranscoder()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpcJsonTranscoder()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpcJsonTranscoder(), target.GetGrpcJsonTranscoder()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.PluginType != target.PluginType {
			return false
		}
	}

	return true
}
