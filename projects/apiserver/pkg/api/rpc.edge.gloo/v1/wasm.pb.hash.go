// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/wasm.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *WasmFilter) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.WasmFilter")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSource())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetConfig())); err != nil {
		return 0, err
	}

	for _, v := range m.GetLocations() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListWasmFiltersRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.ListWasmFiltersRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListWasmFiltersResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.ListWasmFiltersResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetWasmFilters() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DescribeWasmFilterRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.DescribeWasmFilterRequest")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootId())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGatewayRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GatewayRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGatewayRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GatewayRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DescribeWasmFilterResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.DescribeWasmFilterResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetWasmFilter()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("WasmFilter")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWasmFilter(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("WasmFilter")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmFilter_Location) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rpc.edge.gloo.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1.WasmFilter_Location")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGatewayRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GatewayRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGatewayRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GatewayRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGatewayStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GatewayStatus")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGatewayStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GatewayStatus")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GlooInstanceRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGlooInstanceRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GlooInstanceRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
