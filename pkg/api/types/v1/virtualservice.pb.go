// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtualservice.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// Virtual Services represent a collection of routes for a set of domains.
// Gloo's Virtual Services can be compared to
// [virtual hosts](https://www.envoyproxy.io/docs/envoy/latest/api-v1/route_config/vService.html?highlight=virtual%20host) in Envoy terminology.
// A virtual service can be used to define "apps"; a collection of APIs that belong to a particular domain.
// The Virtual Service concept allows configuration of per-virtualservice SSL certificates
type VirtualService struct {
	// Name of the virtual service. Names must be unique and follow the following syntax rules:
	// One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Domains represent the list of domains (host/authority header) that will match for all routes on this virtual service.
	// As in Envoy: wildcard hosts are supported in the form of “*.foo.com” or “*-bar.foo.com”.
	// If domains is empty, gloo will set the domain to "*", making that virtual service the "default" virtualservice.
	// The default virtualservice will be the fallback virtual service for all requests that do not match a domain on an existing virtual service.
	// Only one default virtual service can be defined (either with an empty domain list, or a domain list that includes "*")
	Domains []string `protobuf:"bytes,2,rep,name=domains" json:"domains,omitempty"`
	// Routes define the list of [routes](../) that live on this virtual service.
	Routes []*Route `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	// SSL Config is optional for the virtual service. If provided, the virtual service will listen on the envoy HTTPS listener port (default :8443)
	// If left empty, the virtual service will listen on the HTTP listener port (default :8080)
	SslConfig *SSLConfig `protobuf:"bytes,4,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// indicates whether or not this virtual service should be assigned to gateway roles automatically
	// TODO: eventually this flag will be deprecated; gateway roles will have to explicitly state the virtual services
	// they have access to.
	DisableForGateways bool `protobuf:"varint,7,opt,name=disable_for_gateways,json=disableForGateways,proto3" json:"disable_for_gateways,omitempty"`
	// Status indicates the validation status of the virtual service resource. Status is read-only by clients, and set by gloo during validation
	Status *Status `protobuf:"bytes,5,opt,name=status" json:"status,omitempty" testdiff:"ignore"`
	// Metadata contains the resource metadata for the virtual service
	Metadata *Metadata `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *VirtualService) Reset()                    { *m = VirtualService{} }
func (m *VirtualService) String() string            { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()               {}
func (*VirtualService) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{0} }

func (m *VirtualService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualService) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *VirtualService) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *SSLConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetDisableForGateways() bool {
	if m != nil {
		return m.DisableForGateways
	}
	return false
}

func (m *VirtualService) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualService) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// *
// Routes declare the entrypoints on virtual services and the upstreams or functions they route requests to
type Route struct {
	// Matcher defines what properties of a request to match on.
	// Routes will route all requests they match.
	// If a request matches more than one route, the first route on the virtual service's route list will be selected.
	//
	// Types that are valid to be assigned to Matcher:
	//	*Route_RequestMatcher
	//	*Route_EventMatcher
	Matcher isRoute_Matcher `protobuf_oneof:"matcher"`
	// A route is only allowed to specify one of multiple_destinations or single_destination. Setting both will result in an error
	// Multiple Destinations is used when a user wants a route to balance requests between multiple destinations
	// Balancing is done by probability, where weights are specified for each destination
	MultipleDestinations []*WeightedDestination `protobuf:"bytes,3,rep,name=multiple_destinations,json=multipleDestinations" json:"multiple_destinations,omitempty"`
	// A single destination is specified when a route only routes to a single destination.
	SingleDestination *Destination `protobuf:"bytes,4,opt,name=single_destination,json=singleDestination" json:"single_destination,omitempty"`
	// PrefixRewrite can be specified to rewrite the matched path of the request path to a new prefix
	PrefixRewrite string `protobuf:"bytes,5,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	// Extensions provides a way to extend the behavior of a route. In addition to the core route extensions<!--(TODO)-->,
	// gloo provides the means for route plugins<!--(TODO)--> to be added to gloo which add new types of route extensions.
	// <!--See the route extensions section for a more detailed explanation-->
	Extensions *google_protobuf.Struct `protobuf:"bytes,6,opt,name=extensions" json:"extensions,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{1} }

type isRoute_Matcher interface {
	isRoute_Matcher()
	Equal(interface{}) bool
}

type Route_RequestMatcher struct {
	RequestMatcher *RequestMatcher `protobuf:"bytes,1,opt,name=request_matcher,json=requestMatcher,oneof"`
}
type Route_EventMatcher struct {
	EventMatcher *EventMatcher `protobuf:"bytes,2,opt,name=event_matcher,json=eventMatcher,oneof"`
}

func (*Route_RequestMatcher) isRoute_Matcher() {}
func (*Route_EventMatcher) isRoute_Matcher()   {}

func (m *Route) GetMatcher() isRoute_Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (m *Route) GetRequestMatcher() *RequestMatcher {
	if x, ok := m.GetMatcher().(*Route_RequestMatcher); ok {
		return x.RequestMatcher
	}
	return nil
}

func (m *Route) GetEventMatcher() *EventMatcher {
	if x, ok := m.GetMatcher().(*Route_EventMatcher); ok {
		return x.EventMatcher
	}
	return nil
}

func (m *Route) GetMultipleDestinations() []*WeightedDestination {
	if m != nil {
		return m.MultipleDestinations
	}
	return nil
}

func (m *Route) GetSingleDestination() *Destination {
	if m != nil {
		return m.SingleDestination
	}
	return nil
}

func (m *Route) GetPrefixRewrite() string {
	if m != nil {
		return m.PrefixRewrite
	}
	return ""
}

func (m *Route) GetExtensions() *google_protobuf.Struct {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Route) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Route_OneofMarshaler, _Route_OneofUnmarshaler, _Route_OneofSizer, []interface{}{
		(*Route_RequestMatcher)(nil),
		(*Route_EventMatcher)(nil),
	}
}

func _Route_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Route)
	// matcher
	switch x := m.Matcher.(type) {
	case *Route_RequestMatcher:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestMatcher); err != nil {
			return err
		}
	case *Route_EventMatcher:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EventMatcher); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Route.Matcher has unexpected type %T", x)
	}
	return nil
}

func _Route_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Route)
	switch tag {
	case 1: // matcher.request_matcher
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestMatcher)
		err := b.DecodeMessage(msg)
		m.Matcher = &Route_RequestMatcher{msg}
		return true, err
	case 2: // matcher.event_matcher
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventMatcher)
		err := b.DecodeMessage(msg)
		m.Matcher = &Route_EventMatcher{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Route_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Route)
	// matcher
	switch x := m.Matcher.(type) {
	case *Route_RequestMatcher:
		s := proto.Size(x.RequestMatcher)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Route_EventMatcher:
		s := proto.Size(x.EventMatcher)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request Matcher is a route matcher for traditional http requests
// Request Matchers stand in juxtoposition to Event Matchers, which match "events" rather than HTTP Requests
type RequestMatcher struct {
	// Path specifies the :path header in HTTP2, or the request URL path in HTTP 1
	//
	// Types that are valid to be assigned to Path:
	//	*RequestMatcher_PathPrefix
	//	*RequestMatcher_PathRegex
	//	*RequestMatcher_PathExact
	Path isRequestMatcher_Path `protobuf_oneof:"path"`
	// Headers specify a list of request headers and their values the request must contain to match this route
	// If a value is not specified (empty string) for a header, all values will match so long as the header is present on the request
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Query params work the same way as headers, but for query string parameters
	QueryParams map[string]string `protobuf:"bytes,5,rep,name=query_params,json=queryParams" json:"query_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP Verb(s) to match on. If none specified, the matcher will match all verbs
	Verbs []string `protobuf:"bytes,6,rep,name=verbs" json:"verbs,omitempty"`
}

func (m *RequestMatcher) Reset()                    { *m = RequestMatcher{} }
func (m *RequestMatcher) String() string            { return proto.CompactTextString(m) }
func (*RequestMatcher) ProtoMessage()               {}
func (*RequestMatcher) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{2} }

type isRequestMatcher_Path interface {
	isRequestMatcher_Path()
	Equal(interface{}) bool
}

type RequestMatcher_PathPrefix struct {
	PathPrefix string `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3,oneof"`
}
type RequestMatcher_PathRegex struct {
	PathRegex string `protobuf:"bytes,2,opt,name=path_regex,json=pathRegex,proto3,oneof"`
}
type RequestMatcher_PathExact struct {
	PathExact string `protobuf:"bytes,3,opt,name=path_exact,json=pathExact,proto3,oneof"`
}

func (*RequestMatcher_PathPrefix) isRequestMatcher_Path() {}
func (*RequestMatcher_PathRegex) isRequestMatcher_Path()  {}
func (*RequestMatcher_PathExact) isRequestMatcher_Path()  {}

func (m *RequestMatcher) GetPath() isRequestMatcher_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *RequestMatcher) GetPathPrefix() string {
	if x, ok := m.GetPath().(*RequestMatcher_PathPrefix); ok {
		return x.PathPrefix
	}
	return ""
}

func (m *RequestMatcher) GetPathRegex() string {
	if x, ok := m.GetPath().(*RequestMatcher_PathRegex); ok {
		return x.PathRegex
	}
	return ""
}

func (m *RequestMatcher) GetPathExact() string {
	if x, ok := m.GetPath().(*RequestMatcher_PathExact); ok {
		return x.PathExact
	}
	return ""
}

func (m *RequestMatcher) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RequestMatcher) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

func (m *RequestMatcher) GetVerbs() []string {
	if m != nil {
		return m.Verbs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RequestMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RequestMatcher_OneofMarshaler, _RequestMatcher_OneofUnmarshaler, _RequestMatcher_OneofSizer, []interface{}{
		(*RequestMatcher_PathPrefix)(nil),
		(*RequestMatcher_PathRegex)(nil),
		(*RequestMatcher_PathExact)(nil),
	}
}

func _RequestMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RequestMatcher)
	// path
	switch x := m.Path.(type) {
	case *RequestMatcher_PathPrefix:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathPrefix)
	case *RequestMatcher_PathRegex:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathRegex)
	case *RequestMatcher_PathExact:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathExact)
	case nil:
	default:
		return fmt.Errorf("RequestMatcher.Path has unexpected type %T", x)
	}
	return nil
}

func _RequestMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RequestMatcher)
	switch tag {
	case 1: // path.path_prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &RequestMatcher_PathPrefix{x}
		return true, err
	case 2: // path.path_regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &RequestMatcher_PathRegex{x}
		return true, err
	case 3: // path.path_exact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &RequestMatcher_PathExact{x}
		return true, err
	default:
		return false, nil
	}
}

func _RequestMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RequestMatcher)
	// path
	switch x := m.Path.(type) {
	case *RequestMatcher_PathPrefix:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathPrefix)))
		n += len(x.PathPrefix)
	case *RequestMatcher_PathRegex:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathRegex)))
		n += len(x.PathRegex)
	case *RequestMatcher_PathExact:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathExact)))
		n += len(x.PathExact)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Event matcher is a special kind of matcher for CloudEvents
// The CloudEvents API is described here: https://github.com/cloudevents/spec/blob/master/spec.md
type EventMatcher struct {
	// Event Type indicates the event type or topic to match
	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
}

func (m *EventMatcher) Reset()                    { *m = EventMatcher{} }
func (m *EventMatcher) String() string            { return proto.CompactTextString(m) }
func (*EventMatcher) ProtoMessage()               {}
func (*EventMatcher) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{3} }

func (m *EventMatcher) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

// WeightedDestination attaches a weight to a destination
// For use in routes with multiple destinations
type WeightedDestination struct {
	*Destination `protobuf:"bytes,1,opt,name=destination,embedded=destination" json:"destination,omitempty"`
	// Weight must be greater than zero
	// Routing to each destination will be balanced by the ratio of the destination's weight to the total weight on a route
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *WeightedDestination) Reset()         { *m = WeightedDestination{} }
func (m *WeightedDestination) String() string { return proto.CompactTextString(m) }
func (*WeightedDestination) ProtoMessage()    {}
func (*WeightedDestination) Descriptor() ([]byte, []int) {
	return fileDescriptorVirtualservice, []int{4}
}

func (m *WeightedDestination) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// Destination is a destination that requests can be routed to.
type Destination struct {
	// Types that are valid to be assigned to DestinationType:
	//	*Destination_Function
	//	*Destination_Upstream
	DestinationType isDestination_DestinationType `protobuf_oneof:"destination_type"`
}

func (m *Destination) Reset()                    { *m = Destination{} }
func (m *Destination) String() string            { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()               {}
func (*Destination) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{5} }

type isDestination_DestinationType interface {
	isDestination_DestinationType()
	Equal(interface{}) bool
}

type Destination_Function struct {
	Function *FunctionDestination `protobuf:"bytes,1,opt,name=function,oneof"`
}
type Destination_Upstream struct {
	Upstream *UpstreamDestination `protobuf:"bytes,2,opt,name=upstream,oneof"`
}

func (*Destination_Function) isDestination_DestinationType() {}
func (*Destination_Upstream) isDestination_DestinationType() {}

func (m *Destination) GetDestinationType() isDestination_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (m *Destination) GetFunction() *FunctionDestination {
	if x, ok := m.GetDestinationType().(*Destination_Function); ok {
		return x.Function
	}
	return nil
}

func (m *Destination) GetUpstream() *UpstreamDestination {
	if x, ok := m.GetDestinationType().(*Destination_Upstream); ok {
		return x.Upstream
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Destination) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Destination_OneofMarshaler, _Destination_OneofUnmarshaler, _Destination_OneofSizer, []interface{}{
		(*Destination_Function)(nil),
		(*Destination_Upstream)(nil),
	}
}

func _Destination_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Destination)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *Destination_Function:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Function); err != nil {
			return err
		}
	case *Destination_Upstream:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Upstream); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Destination.DestinationType has unexpected type %T", x)
	}
	return nil
}

func _Destination_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Destination)
	switch tag {
	case 1: // destination_type.function
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FunctionDestination)
		err := b.DecodeMessage(msg)
		m.DestinationType = &Destination_Function{msg}
		return true, err
	case 2: // destination_type.upstream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpstreamDestination)
		err := b.DecodeMessage(msg)
		m.DestinationType = &Destination_Upstream{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Destination_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Destination)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *Destination_Function:
		s := proto.Size(x.Function)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Destination_Upstream:
		s := proto.Size(x.Upstream)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// FunctionDestination will route a request to a specific function defined for an upstream
type FunctionDestination struct {
	// Upstream Name is the name of the upstream the function belongs to
	UpstreamName string `protobuf:"bytes,1,opt,name=upstream_name,json=upstreamName,proto3" json:"upstream_name,omitempty"`
	// Function Name is the name of the function as defined on the upstream
	FunctionName string `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (m *FunctionDestination) Reset()         { *m = FunctionDestination{} }
func (m *FunctionDestination) String() string { return proto.CompactTextString(m) }
func (*FunctionDestination) ProtoMessage()    {}
func (*FunctionDestination) Descriptor() ([]byte, []int) {
	return fileDescriptorVirtualservice, []int{6}
}

func (m *FunctionDestination) GetUpstreamName() string {
	if m != nil {
		return m.UpstreamName
	}
	return ""
}

func (m *FunctionDestination) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

// Upstream Destination routes a request to an upstream
type UpstreamDestination struct {
	// Name of the upstream
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *UpstreamDestination) Reset()         { *m = UpstreamDestination{} }
func (m *UpstreamDestination) String() string { return proto.CompactTextString(m) }
func (*UpstreamDestination) ProtoMessage()    {}
func (*UpstreamDestination) Descriptor() ([]byte, []int) {
	return fileDescriptorVirtualservice, []int{7}
}

func (m *UpstreamDestination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SSLConfig contains the options necessary to configure a virtualservice or listener to use TLS
type SSLConfig struct {
	// * SecretRef contains the secret ref to a gloo secret containing the following structure:
	// {
	// "tls.crt": <ca chain data...>,
	// "tls.key": <private key data...>
	// }
	SecretRef string `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// optional. the SNI domains that should be considered for TLS connections
	SniDomains []string `protobuf:"bytes,2,rep,name=sni_domains,json=sniDomains" json:"sni_domains,omitempty"`
}

func (m *SSLConfig) Reset()                    { *m = SSLConfig{} }
func (m *SSLConfig) String() string            { return proto.CompactTextString(m) }
func (*SSLConfig) ProtoMessage()               {}
func (*SSLConfig) Descriptor() ([]byte, []int) { return fileDescriptorVirtualservice, []int{8} }

func (m *SSLConfig) GetSecretRef() string {
	if m != nil {
		return m.SecretRef
	}
	return ""
}

func (m *SSLConfig) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gloo.api.v1.VirtualService")
	proto.RegisterType((*Route)(nil), "gloo.api.v1.Route")
	proto.RegisterType((*RequestMatcher)(nil), "gloo.api.v1.RequestMatcher")
	proto.RegisterType((*EventMatcher)(nil), "gloo.api.v1.EventMatcher")
	proto.RegisterType((*WeightedDestination)(nil), "gloo.api.v1.WeightedDestination")
	proto.RegisterType((*Destination)(nil), "gloo.api.v1.Destination")
	proto.RegisterType((*FunctionDestination)(nil), "gloo.api.v1.FunctionDestination")
	proto.RegisterType((*UpstreamDestination)(nil), "gloo.api.v1.UpstreamDestination")
	proto.RegisterType((*SSLConfig)(nil), "gloo.api.v1.SSLConfig")
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
	if this.Name != that1.Name {
		return false
	}
	if len(this.Domains) != len(that1.Domains) {
		return false
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return false
		}
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if this.DisableForGateways != that1.DisableForGateways {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
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
	if that1.Matcher == nil {
		if this.Matcher != nil {
			return false
		}
	} else if this.Matcher == nil {
		return false
	} else if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	if len(this.MultipleDestinations) != len(that1.MultipleDestinations) {
		return false
	}
	for i := range this.MultipleDestinations {
		if !this.MultipleDestinations[i].Equal(that1.MultipleDestinations[i]) {
			return false
		}
	}
	if !this.SingleDestination.Equal(that1.SingleDestination) {
		return false
	}
	if this.PrefixRewrite != that1.PrefixRewrite {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	return true
}
func (this *Route_RequestMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_RequestMatcher)
	if !ok {
		that2, ok := that.(Route_RequestMatcher)
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
	if !this.RequestMatcher.Equal(that1.RequestMatcher) {
		return false
	}
	return true
}
func (this *Route_EventMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_EventMatcher)
	if !ok {
		that2, ok := that.(Route_EventMatcher)
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
	if !this.EventMatcher.Equal(that1.EventMatcher) {
		return false
	}
	return true
}
func (this *RequestMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestMatcher)
	if !ok {
		that2, ok := that.(RequestMatcher)
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
	if that1.Path == nil {
		if this.Path != nil {
			return false
		}
	} else if this.Path == nil {
		return false
	} else if !this.Path.Equal(that1.Path) {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if len(this.QueryParams) != len(that1.QueryParams) {
		return false
	}
	for i := range this.QueryParams {
		if this.QueryParams[i] != that1.QueryParams[i] {
			return false
		}
	}
	if len(this.Verbs) != len(that1.Verbs) {
		return false
	}
	for i := range this.Verbs {
		if this.Verbs[i] != that1.Verbs[i] {
			return false
		}
	}
	return true
}
func (this *RequestMatcher_PathPrefix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestMatcher_PathPrefix)
	if !ok {
		that2, ok := that.(RequestMatcher_PathPrefix)
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
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	return true
}
func (this *RequestMatcher_PathRegex) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestMatcher_PathRegex)
	if !ok {
		that2, ok := that.(RequestMatcher_PathRegex)
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
	if this.PathRegex != that1.PathRegex {
		return false
	}
	return true
}
func (this *RequestMatcher_PathExact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestMatcher_PathExact)
	if !ok {
		that2, ok := that.(RequestMatcher_PathExact)
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
	if this.PathExact != that1.PathExact {
		return false
	}
	return true
}
func (this *EventMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EventMatcher)
	if !ok {
		that2, ok := that.(EventMatcher)
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
	if this.EventType != that1.EventType {
		return false
	}
	return true
}
func (this *WeightedDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WeightedDestination)
	if !ok {
		that2, ok := that.(WeightedDestination)
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
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if this.Weight != that1.Weight {
		return false
	}
	return true
}
func (this *Destination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination)
	if !ok {
		that2, ok := that.(Destination)
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
	if that1.DestinationType == nil {
		if this.DestinationType != nil {
			return false
		}
	} else if this.DestinationType == nil {
		return false
	} else if !this.DestinationType.Equal(that1.DestinationType) {
		return false
	}
	return true
}
func (this *Destination_Function) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination_Function)
	if !ok {
		that2, ok := that.(Destination_Function)
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
	if !this.Function.Equal(that1.Function) {
		return false
	}
	return true
}
func (this *Destination_Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination_Upstream)
	if !ok {
		that2, ok := that.(Destination_Upstream)
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
	if !this.Upstream.Equal(that1.Upstream) {
		return false
	}
	return true
}
func (this *FunctionDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionDestination)
	if !ok {
		that2, ok := that.(FunctionDestination)
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
	if this.UpstreamName != that1.UpstreamName {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	return true
}
func (this *UpstreamDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamDestination)
	if !ok {
		that2, ok := that.(UpstreamDestination)
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
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *SSLConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SSLConfig)
	if !ok {
		that2, ok := that.(SSLConfig)
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
	if this.SecretRef != that1.SecretRef {
		return false
	}
	if len(this.SniDomains) != len(that1.SniDomains) {
		return false
	}
	for i := range this.SniDomains {
		if this.SniDomains[i] != that1.SniDomains[i] {
			return false
		}
	}
	return true
}

func init() { proto.RegisterFile("virtualservice.proto", fileDescriptorVirtualservice) }

var fileDescriptorVirtualservice = []byte{
	// 891 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x6e, 0x23, 0xb5,
	0x17, 0x6e, 0x9a, 0x34, 0x6d, 0x4e, 0xfe, 0xfc, 0xba, 0x6e, 0xba, 0xbf, 0xa1, 0xfc, 0x69, 0x98,
	0x15, 0x52, 0x40, 0xec, 0x84, 0x16, 0x21, 0x50, 0x2f, 0x56, 0xab, 0xb0, 0xed, 0x56, 0x82, 0x85,
	0xc5, 0x65, 0x41, 0xe2, 0x66, 0xe4, 0x24, 0x27, 0x13, 0x6b, 0x27, 0xe3, 0xa9, 0xed, 0x49, 0x9b,
	0x77, 0xe1, 0x06, 0x89, 0x0b, 0x9e, 0x85, 0x2b, 0x9e, 0x60, 0x2f, 0x78, 0x04, 0x9e, 0x00, 0xd9,
	0x9e, 0x69, 0x67, 0x56, 0x55, 0x25, 0xee, 0x7c, 0xbe, 0xf3, 0x7d, 0xdf, 0xf1, 0x1c, 0xfb, 0x78,
	0xa0, 0xbf, 0xe2, 0x52, 0x67, 0x2c, 0x56, 0x28, 0x57, 0x7c, 0x8a, 0x41, 0x2a, 0x85, 0x16, 0xa4,
	0x1d, 0xc5, 0x42, 0x04, 0x2c, 0xe5, 0xc1, 0xea, 0xe8, 0xe0, 0xbd, 0x48, 0x88, 0x28, 0xc6, 0x91,
	0x4d, 0x4d, 0xb2, 0xf9, 0x48, 0x69, 0x99, 0x4d, 0xb5, 0xa3, 0x1e, 0xf4, 0x23, 0x11, 0x09, 0xbb,
	0x1c, 0x99, 0x55, 0x8e, 0x76, 0x94, 0x66, 0x3a, 0x53, 0x79, 0xd4, 0x5b, 0xa2, 0x66, 0x33, 0xa6,
	0x99, 0x8b, 0xfd, 0x3f, 0x37, 0xa1, 0xf7, 0x93, 0xab, 0x7b, 0xe1, 0xea, 0x12, 0x02, 0x8d, 0x84,
	0x2d, 0xd1, 0xab, 0x0d, 0x6a, 0xc3, 0x16, 0xb5, 0x6b, 0xe2, 0xc1, 0xf6, 0x4c, 0x2c, 0x19, 0x4f,
	0x94, 0xb7, 0x39, 0xa8, 0x0f, 0x5b, 0xb4, 0x08, 0xc9, 0x27, 0xd0, 0x94, 0x22, 0xd3, 0xa8, 0xbc,
	0xfa, 0xa0, 0x3e, 0x6c, 0x1f, 0x93, 0xa0, 0xb4, 0xe1, 0x80, 0x9a, 0x14, 0xcd, 0x19, 0xe4, 0x0b,
	0x00, 0xa5, 0xe2, 0x70, 0x2a, 0x92, 0x39, 0x8f, 0xbc, 0xc6, 0xa0, 0x36, 0x6c, 0x1f, 0x3f, 0xac,
	0xf0, 0x2f, 0x2e, 0xbe, 0xfd, 0xda, 0x66, 0x69, 0x4b, 0xa9, 0xd8, 0x2d, 0xc9, 0x67, 0xd0, 0x9f,
	0x71, 0xc5, 0x26, 0x31, 0x86, 0x73, 0x21, 0xc3, 0x88, 0x69, 0xbc, 0x62, 0x6b, 0xe5, 0x6d, 0x0f,
	0x6a, 0xc3, 0x1d, 0x4a, 0xf2, 0xdc, 0x99, 0x90, 0xcf, 0xf3, 0x0c, 0x19, 0x43, 0xd3, 0x7d, 0xb5,
	0xb7, 0x65, 0x8b, 0xec, 0x55, 0x8b, 0xd8, 0xd4, 0x78, 0xff, 0x9f, 0x37, 0x87, 0x0f, 0x34, 0x2a,
	0x3d, 0xe3, 0xf3, 0xf9, 0x89, 0xcf, 0xa3, 0x44, 0x48, 0xf4, 0x69, 0xae, 0x24, 0x47, 0xb0, 0x53,
	0xf4, 0xca, 0x6b, 0x5a, 0x97, 0xfd, 0x8a, 0xcb, 0x8b, 0x3c, 0x49, 0x6f, 0x68, 0xfe, 0xaf, 0x75,
	0xd8, 0xb2, 0x5f, 0x4c, 0xce, 0xe0, 0x7f, 0x12, 0x2f, 0x33, 0x54, 0x3a, 0x5c, 0x32, 0x3d, 0x5d,
	0xa0, 0xb4, 0xed, 0x6c, 0x1f, 0xbf, 0x5b, 0x6d, 0x8f, 0xe3, 0xbc, 0x70, 0x94, 0xf3, 0x0d, 0xda,
	0x93, 0x15, 0x84, 0x3c, 0x85, 0x2e, 0xae, 0x30, 0xb9, 0x75, 0xd9, 0xb4, 0x2e, 0xef, 0x54, 0x5c,
	0x4e, 0x0d, 0xe3, 0xd6, 0xa3, 0x83, 0xa5, 0x98, 0xbc, 0x82, 0xfd, 0x65, 0x16, 0x6b, 0x9e, 0xc6,
	0x18, 0xce, 0x50, 0x69, 0x9e, 0x30, 0xcd, 0x45, 0x52, 0x1c, 0xd7, 0xa0, 0xe2, 0xf4, 0x33, 0xf2,
	0x68, 0xa1, 0x71, 0xf6, 0xec, 0x96, 0x48, 0xfb, 0x85, 0xbc, 0x04, 0x2a, 0xf2, 0x1c, 0x88, 0xe2,
	0x49, 0x54, 0x35, 0xcd, 0x8f, 0xd4, 0xab, 0x78, 0x96, 0xbd, 0x1e, 0x38, 0x4d, 0x09, 0x22, 0x1f,
	0x41, 0x2f, 0x95, 0x38, 0xe7, 0xd7, 0xa1, 0xc4, 0x2b, 0xc9, 0x35, 0xda, 0x23, 0x6b, 0xd1, 0xae,
	0x43, 0xa9, 0x03, 0xc9, 0x97, 0x00, 0x78, 0xad, 0x31, 0x51, 0x76, 0xef, 0xee, 0x3c, 0xfe, 0x1f,
	0xb8, 0x71, 0x08, 0x8a, 0x71, 0x08, 0x2e, 0xec, 0x38, 0xd0, 0x12, 0x75, 0xdc, 0x82, 0xed, 0xbc,
	0x77, 0xfe, 0xef, 0x75, 0xe8, 0x55, 0x3b, 0x4e, 0x3e, 0x84, 0x76, 0xca, 0xf4, 0x22, 0x74, 0xc5,
	0xdc, 0x95, 0x3f, 0xdf, 0xa0, 0x60, 0xc0, 0x97, 0x16, 0x23, 0x87, 0x60, 0xa3, 0x50, 0x62, 0x84,
	0xd7, 0xb6, 0xff, 0x86, 0xd1, 0x32, 0x18, 0x35, 0xd0, 0x0d, 0x01, 0xaf, 0xd9, 0x54, 0x7b, 0xf5,
	0x32, 0xe1, 0xd4, 0x40, 0x64, 0x0c, 0xdb, 0x0b, 0x64, 0x33, 0x94, 0xca, 0x6b, 0xd8, 0xa6, 0x0f,
	0xef, 0xb9, 0x04, 0xc1, 0xb9, 0xa3, 0x9e, 0x26, 0x5a, 0xae, 0x69, 0x21, 0x24, 0xdf, 0x43, 0xe7,
	0x32, 0x43, 0xb9, 0x0e, 0x53, 0x26, 0xd9, 0xd2, 0xdc, 0x6b, 0x63, 0xf4, 0xe9, 0x7d, 0x46, 0x3f,
	0x18, 0xfe, 0x4b, 0x4b, 0x77, 0x66, 0xed, 0xcb, 0x5b, 0x84, 0xf4, 0x61, 0x6b, 0x85, 0x72, 0x62,
	0x7a, 0x69, 0xe6, 0xd9, 0x05, 0x07, 0x27, 0xd0, 0x29, 0xd7, 0x27, 0xbb, 0x50, 0x7f, 0x8d, 0xeb,
	0xfc, 0x29, 0x30, 0x4b, 0xab, 0x63, 0x71, 0x86, 0xae, 0x13, 0xd4, 0x05, 0x27, 0x9b, 0x5f, 0xd5,
	0x0e, 0x9e, 0xc0, 0xee, 0xdb, 0x25, 0xff, 0x8b, 0x7e, 0xdc, 0x84, 0x86, 0xe9, 0x99, 0xff, 0x18,
	0x3a, 0xe5, 0x1b, 0x4d, 0xde, 0x07, 0x70, 0x33, 0xa0, 0xd7, 0x69, 0xf1, 0x2a, 0xb5, 0x2c, 0xf2,
	0xe3, 0x3a, 0x45, 0x5f, 0xc0, 0xde, 0x1d, 0xd7, 0x96, 0x3c, 0x85, 0x76, 0xf9, 0x66, 0xd6, 0xee,
	0xbf, 0x99, 0xe3, 0xc6, 0x5f, 0x6f, 0x0e, 0x6b, 0xb4, 0x2c, 0x21, 0x0f, 0xa1, 0x79, 0x65, 0x8d,
	0xed, 0x56, 0xbb, 0x34, 0x8f, 0xfc, 0xdf, 0x6a, 0xd0, 0x2e, 0x57, 0x7a, 0x02, 0x3b, 0xf3, 0x2c,
	0x99, 0x96, 0xca, 0x54, 0x87, 0xea, 0x2c, 0x4f, 0x96, 0x34, 0xe7, 0x1b, 0xf4, 0x46, 0x63, 0xf4,
	0x59, 0xaa, 0xb4, 0x44, 0xb6, 0xcc, 0xc7, 0xbb, 0xaa, 0x7f, 0x95, 0x27, 0xdf, 0xd2, 0x17, 0x9a,
	0x31, 0x81, 0xdd, 0xd2, 0xb6, 0x6d, 0x97, 0xfc, 0x10, 0xf6, 0xee, 0x28, 0x4b, 0x1e, 0x41, 0xb7,
	0x90, 0x85, 0xa5, 0x37, 0xbe, 0x53, 0x80, 0xdf, 0x99, 0xb7, 0xfe, 0x11, 0x74, 0x8b, 0xbd, 0x39,
	0x92, 0x3b, 0xa9, 0x4e, 0x01, 0x1a, 0x92, 0xff, 0x31, 0xec, 0xdd, 0xb1, 0xaf, 0xbb, 0xfe, 0x1d,
	0xfe, 0x37, 0xd0, 0xba, 0x79, 0xd6, 0xcd, 0x61, 0x2a, 0x9c, 0x4a, 0xd4, 0xa1, 0xc4, 0x79, 0x71,
	0x98, 0x0e, 0xa1, 0x38, 0x27, 0x87, 0xd0, 0x56, 0x09, 0x0f, 0xab, 0xff, 0x1a, 0x50, 0x09, 0x7f,
	0xe6, 0x90, 0x71, 0xf0, 0xc7, 0xdf, 0x1f, 0xd4, 0x7e, 0x19, 0x46, 0x5c, 0x2f, 0xb2, 0x49, 0x30,
	0x15, 0xcb, 0x91, 0x12, 0xb1, 0x78, 0xcc, 0xc5, 0xc8, 0xb4, 0x6c, 0x94, 0xbe, 0x8e, 0x46, 0x2c,
	0xe5, 0x23, 0xd3, 0x04, 0x35, 0x5a, 0x1d, 0x4d, 0x9a, 0xf6, 0x6d, 0xf8, 0xfc, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x27, 0x9c, 0xad, 0xf5, 0x5d, 0x07, 0x00, 0x00,
}
