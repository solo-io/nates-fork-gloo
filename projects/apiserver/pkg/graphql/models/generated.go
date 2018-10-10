// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	fmt "fmt"
	io "io"
	strconv "strconv"

	customtypes "github.com/solo-io/solo-kit/projects/apiserver/pkg/graphql/customtypes"
)

type Artifact struct {
	Data     string   `json:"data"`
	Metadata Metadata `json:"metadata"`
}

func (Artifact) IsResource() {}

type AwsDestinationSpec struct {
	LogicalName            string                   `json:"logicalName"`
	InvocationStyle        AwsLambdaInvocationStyle `json:"invocationStyle"`
	ResponseTransformation bool                     `json:"responseTransformation"`
}

func (AwsDestinationSpec) IsDestinationSpec() {}

type AwsLambdaFunction struct {
	LogicalName  string `json:"logicalName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}

type AwsSecret struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func (AwsSecret) IsSecretKind() {}

type AwsUpstreamSpec struct {
	Region    string              `json:"region"`
	SecretRef ResourceRef         `json:"secretRef"`
	Functions []AwsLambdaFunction `json:"functions"`
}

func (AwsUpstreamSpec) IsUpstreamSpec() {}

type AzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}

func (AzureDestinationSpec) IsDestinationSpec() {}

type AzureFunction struct {
	FunctionName string           `json:"functionName"`
	AuthLevel    AzureFnAuthLevel `json:"authLevel"`
}

type AzureSecret struct {
	APIKeys *MapStringString `json:"apiKeys"`
}

func (AzureSecret) IsSecretKind() {}

type AzureUpstreamSpec struct {
	FunctionAppName string          `json:"functionAppName"`
	SecretRef       ResourceRef     `json:"secretRef"`
	Functions       []AzureFunction `json:"functions"`
}

func (AzureUpstreamSpec) IsUpstreamSpec() {}

type Destination interface {
	IsDestination()
}

type DestinationSpec interface {
	IsDestinationSpec()
}

type FieldResolver struct {
	FieldName string   `json:"fieldName"`
	Resolver  Resolver `json:"resolver"`
}

type GlooResolver struct {
	RequestTemplate  *RequestTemplate  `json:"requestTemplate"`
	ResponseTemplate *ResponseTemplate `json:"responseTemplate"`
	Destination      Destination       `json:"destination"`
}

func (GlooResolver) IsResolver() {}

type GrpcDestinationSpec struct {
	FunctionName string                    `json:"functionName"`
	Parameters   *TransformationParameters `json:"parameters"`
}

func (GrpcDestinationSpec) IsDestinationSpec() {}

type GrpcService struct {
	PackageName   string   `json:"packageName"`
	ServiceName   string   `json:"serviceName"`
	FunctionNames []string `json:"functionNames"`
}

type GrpcServiceSpec struct {
	GrpcServices []*GrpcService `json:"grpcServices"`
}

func (GrpcServiceSpec) IsServiceSpec() {}

type InputArtifact struct {
	Data     string        `json:"data"`
	Metadata InputMetadata `json:"metadata"`
}

type InputAwsDestinationSpec struct {
	LogicalName            string                   `json:"logicalName"`
	InvocationStyle        AwsLambdaInvocationStyle `json:"invocationStyle"`
	ResponseTransformation bool                     `json:"responseTransformation"`
}

type InputAwsLambdaFunction struct {
	LogicalName  string `json:"logicalName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}

type InputAwsSecret struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type InputAwsUpstreamSpec struct {
	Region    string                   `json:"region"`
	SecretRef InputResourceRef         `json:"secretRef"`
	Functions []InputAwsLambdaFunction `json:"functions"`
}

type InputAzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}

type InputAzureFunction struct {
	FunctionName string `json:"functionName"`
	AuthLevel    string `json:"authLevel"`
}

type InputAzureSecret struct {
	APIKeys InputMapStringString `json:"apiKeys"`
}

type InputAzureUpstreamSpec struct {
	FunctionAppName string               `json:"functionAppName"`
	SecretRef       *InputResourceRef    `json:"secretRef"`
	Functions       []InputAzureFunction `json:"functions"`
}

type InputDestination struct {
	SingleDestination *InputSingleDestination `json:"singleDestination"`
	MultiDestination  *InputMultiDestination  `json:"multiDestination"`
}

type InputDestinationSpec struct {
	Aws   *InputAwsDestinationSpec   `json:"aws"`
	Azure *InputAzureDestinationSpec `json:"azure"`
	Rest  *InputRestDestinationSpec  `json:"rest"`
}

type InputFieldResolver struct {
	FieldName string        `json:"fieldName"`
	Resolver  InputResolver `json:"resolver"`
}

type InputGlooResolver struct {
	RequestTemplate  *InputRequestTemplate  `json:"requestTemplate"`
	ResponseTemplate *InputResponseTemplate `json:"responseTemplate"`
	Destination      InputDestination       `json:"destination"`
}

type InputGrpcServiceSpec struct {
	Empty *string `json:"empty"`
}

type InputKeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}

type InputKubeUpstreamSpec struct {
	ServiceName      string                `json:"serviceName"`
	ServiceNamespace string                `json:"serviceNamespace"`
	ServicePort      int                   `json:"servicePort"`
	Selector         *InputMapStringString `json:"selector"`
	ServiceSpec      *InputServiceSpec     `json:"serviceSpec"`
}

type InputMapStringString struct {
	Values []InputValue `json:"values"`
}

type InputMatcher struct {
	PathMatch       string                 `json:"pathMatch"`
	PathMatchType   PathMatchType          `json:"pathMatchType"`
	Headers         []InputKeyValueMatcher `json:"headers"`
	QueryParameters []InputKeyValueMatcher `json:"queryParameters"`
	Methods         []string               `json:"methods"`
}

type InputMetadata struct {
	Name            string                `json:"name"`
	Namespace       string                `json:"namespace"`
	ResourceVersion string                `json:"resourceVersion"`
	Labels          *InputMapStringString `json:"labels"`
	Annotations     *InputMapStringString `json:"annotations"`
}

type InputMultiDestination struct {
	Destinations []InputWeightedDestination `json:"destinations"`
}

type InputNodeJSResolver struct {
	Empty *string `json:"empty"`
}

type InputRequestTemplate struct {
	Verb    string                `json:"verb"`
	Path    string                `json:"path"`
	Body    string                `json:"body"`
	Headers *InputMapStringString `json:"headers"`
}

type InputResolver struct {
	GlooResolver     *InputGlooResolver     `json:"glooResolver"`
	TemplateResolver *InputTemplateResolver `json:"templateResolver"`
	NodeResolver     *InputNodeJSResolver   `json:"nodeResolver"`
}

type InputResolverMap struct {
	Types    []InputTypeResolver `json:"types"`
	Metadata InputMetadata       `json:"metadata"`
}

type InputResourceRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type InputResponseTemplate struct {
	Body    string                `json:"body"`
	Headers *InputMapStringString `json:"headers"`
}

type InputRestDestinationSpec struct {
	FunctionName string                         `json:"functionName"`
	Parameters   *InputTransformationParameters `json:"parameters"`
}

type InputRestServiceSpec struct {
	Functions        []InputTransformation `json:"functions"`
	InlineSwaggerDoc *string               `json:"inlineSwaggerDoc"`
}

type InputRoute struct {
	Matcher     InputMatcher       `json:"matcher"`
	Destination InputDestination   `json:"destination"`
	Plugins     *InputRoutePlugins `json:"plugins"`
}

type InputRoutePlugins struct {
	Empty *string `json:"empty"`
}

type InputSchema struct {
	InlineSchema string        `json:"inlineSchema"`
	Metadata     InputMetadata `json:"metadata"`
}

type InputSecret struct {
	Kind     InputSecretKind `json:"kind"`
	Metadata InputMetadata   `json:"metadata"`
}

type InputSecretKind struct {
	Aws   *InputAwsSecret   `json:"aws"`
	Azure *InputAzureSecret `json:"azure"`
	TLS   *InputTlsSecret   `json:"tls"`
}

type InputServiceSpec struct {
	Rest *InputRestServiceSpec `json:"rest"`
	Grpc *InputGrpcServiceSpec `json:"grpc"`
}

type InputSettings struct {
	WatchNamespaces []string              `json:"watchNamespaces"`
	RefreshRate     *customtypes.Duration `json:"refreshRate"`
	Metadata        InputMetadata         `json:"metadata"`
}

type InputSingleDestination struct {
	Upstream        InputResourceRef      `json:"upstream"`
	DestinationSpec *InputDestinationSpec `json:"destinationSpec"`
}

type InputSslConfig struct {
	SecretRef InputResourceRef `json:"secretRef"`
}

type InputStaticHost struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type InputStaticUpstreamSpec struct {
	Hosts       []InputStaticHost `json:"hosts"`
	ServiceSpec *InputServiceSpec `json:"serviceSpec"`
	UseTLS      bool              `json:"useTls"`
}

type InputStatus struct {
	State  State  `json:"state"`
	Reason string `json:"reason"`
}

type InputTemplateResolver struct {
	Empty *string `json:"empty"`
}

type InputTlsSecret struct {
	CertChain  string `json:"certChain"`
	PrivateKey string `json:"privateKey"`
	RootCa     string `json:"rootCa"`
}

type InputTransformation struct {
	FunctionName string                `json:"functionName"`
	Body         *string               `json:"body"`
	Headers      *InputMapStringString `json:"headers"`
}

type InputTransformationParameters struct {
	Headers *InputMapStringString `json:"headers"`
	Path    *string               `json:"path"`
}

type InputTypeResolver struct {
	TypeName string               `json:"typeName"`
	Fields   []InputFieldResolver `json:"fields"`
}

type InputUpdateMetadata struct {
	Name        *string               `json:"name"`
	Labels      *InputMapStringString `json:"labels"`
	Annotations *InputMapStringString `json:"annotations"`
}

type InputUpdateVirtualService struct {
	Domains   []string                    `json:"domains"`
	SslConfig *InputSslConfig             `json:"sslConfig"`
	Plugins   *InputVirtualServicePlugins `json:"plugins"`
	Metadata  *InputUpdateMetadata        `json:"metadata"`
}

type InputUpstream struct {
	Spec     InputUpstreamSpec `json:"spec"`
	Metadata InputMetadata     `json:"metadata"`
}

type InputUpstreamSpec struct {
	Aws    *InputAwsUpstreamSpec    `json:"aws"`
	Azure  *InputAzureUpstreamSpec  `json:"azure"`
	Kube   *InputKubeUpstreamSpec   `json:"kube"`
	Static *InputStaticUpstreamSpec `json:"static"`
}

type InputValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InputVirtualService struct {
	Domains   []string                    `json:"domains"`
	Routes    []InputRoute                `json:"routes"`
	SslConfig *InputSslConfig             `json:"sslConfig"`
	Plugins   *InputVirtualServicePlugins `json:"plugins"`
	Metadata  InputMetadata               `json:"metadata"`
}

type InputVirtualServicePlugins struct {
	Empty *string `json:"empty"`
}

type InputWeightedDestination struct {
	Destination InputSingleDestination `json:"destination"`
	Weight      int                    `json:"weight"`
}

type KeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}

type KubeUpstreamSpec struct {
	ServiceName      string           `json:"serviceName"`
	ServiceNamespace string           `json:"serviceNamespace"`
	ServicePort      int              `json:"servicePort"`
	Selector         *MapStringString `json:"selector"`
	ServiceSpec      ServiceSpec      `json:"serviceSpec"`
}

func (KubeUpstreamSpec) IsUpstreamSpec() {}

type MapStringString struct {
	Values []Value `json:"values"`
}

type Matcher struct {
	PathMatch       string            `json:"pathMatch"`
	PathMatchType   PathMatchType     `json:"pathMatchType"`
	Headers         []KeyValueMatcher `json:"headers"`
	QueryParameters []KeyValueMatcher `json:"queryParameters"`
	Methods         []string          `json:"methods"`
}

type Metadata struct {
	Name            string           `json:"name"`
	Namespace       string           `json:"namespace"`
	ResourceVersion string           `json:"resourceVersion"`
	Labels          *MapStringString `json:"labels"`
	Annotations     *MapStringString `json:"annotations"`
	GUID            string           `json:"guid"`
}

type MultiDestination struct {
	Destinations []WeightedDestination `json:"destinations"`
}

func (MultiDestination) IsDestination() {}

type NodeJSResolver struct {
	Empty *string `json:"empty"`
}

func (NodeJSResolver) IsResolver() {}

type OAuthEndpoint struct {
	URL        string `json:"url"`
	ClientName string `json:"clientName"`
}

type RequestTemplate struct {
	Verb    string           `json:"verb"`
	Path    string           `json:"path"`
	Body    string           `json:"body"`
	Headers *MapStringString `json:"headers"`
}

type Resolver interface {
	IsResolver()
}

type ResolverMap struct {
	Types    []TypeResolver `json:"types"`
	Metadata Metadata       `json:"metadata"`
	Status   Status         `json:"status"`
}

func (ResolverMap) IsResource() {}

type Resource interface {
	IsResource()
}

type ResourceRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ResponseTemplate struct {
	Body    string           `json:"body"`
	Headers *MapStringString `json:"headers"`
}

type RestDestinationSpec struct {
	FunctionName string                    `json:"functionName"`
	Parameters   *TransformationParameters `json:"parameters"`
}

func (RestDestinationSpec) IsDestinationSpec() {}

type RestServiceSpec struct {
	Functions []Transformation `json:"functions"`
}

func (RestServiceSpec) IsServiceSpec() {}

type Route struct {
	Matcher     Matcher       `json:"matcher"`
	Destination Destination   `json:"destination"`
	Plugins     *RoutePlugins `json:"plugins"`
}

type RoutePlugins struct {
	Empty *string `json:"empty"`
}

type Schema struct {
	InlineSchema string   `json:"inlineSchema"`
	Metadata     Metadata `json:"metadata"`
	Status       Status   `json:"status"`
}

func (Schema) IsResource() {}

type Secret struct {
	Kind     SecretKind `json:"kind"`
	Metadata Metadata   `json:"metadata"`
}

func (Secret) IsResource() {}

type SecretKind interface {
	IsSecretKind()
}

type ServiceSpec interface {
	IsServiceSpec()
}

type Settings struct {
	WatchNamespaces []string              `json:"watchNamespaces"`
	RefreshRate     *customtypes.Duration `json:"refreshRate"`
	Metadata        Metadata              `json:"metadata"`
}

func (Settings) IsResource() {}

type SingleDestination struct {
	Upstream        Upstream        `json:"upstream"`
	DestinationSpec DestinationSpec `json:"destinationSpec"`
}

func (SingleDestination) IsDestination() {}

type SslConfig struct {
	SecretRef ResourceRef `json:"secretRef"`
}

type StaticHost struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type StaticUpstreamSpec struct {
	Hosts       []StaticHost `json:"hosts"`
	ServiceSpec ServiceSpec  `json:"serviceSpec"`
	UseTLS      bool         `json:"useTls"`
}

func (StaticUpstreamSpec) IsUpstreamSpec() {}

type Status struct {
	State  State   `json:"state"`
	Reason *string `json:"reason"`
}

type TemplateResolver struct {
	InlineTemplate *string `json:"inlineTemplate"`
}

func (TemplateResolver) IsResolver() {}

type TlsSecret struct {
	CertChain  string `json:"certChain"`
	PrivateKey string `json:"privateKey"`
	RootCa     string `json:"rootCa"`
}

func (TlsSecret) IsSecretKind() {}

type Transformation struct {
	FunctionName string           `json:"functionName"`
	Body         *string          `json:"body"`
	Headers      *MapStringString `json:"headers"`
}

type TransformationParameters struct {
	Headers *MapStringString `json:"headers"`
	Path    *string          `json:"path"`
}

type TypeResolver struct {
	TypeName string          `json:"typeName"`
	Fields   []FieldResolver `json:"fields"`
}

type Upstream struct {
	Spec     UpstreamSpec `json:"spec"`
	Metadata Metadata     `json:"metadata"`
	Status   Status       `json:"status"`
}

func (Upstream) IsResource() {}

type UpstreamSpec interface {
	IsUpstreamSpec()
}

type Value struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VirtualService struct {
	Domains   []string               `json:"domains"`
	Routes    []Route                `json:"routes"`
	SslConfig *SslConfig             `json:"sslConfig"`
	Plugins   *VirtualServicePlugins `json:"plugins"`
	Metadata  Metadata               `json:"metadata"`
	Status    Status                 `json:"status"`
}

func (VirtualService) IsResource() {}

type VirtualServicePlugins struct {
	Empty *string `json:"empty"`
}

type WeightedDestination struct {
	Destination SingleDestination `json:"destination"`
	Weight      int               `json:"weight"`
}

type AwsLambdaInvocationStyle string

const (
	AwsLambdaInvocationStyleSync  AwsLambdaInvocationStyle = "SYNC"
	AwsLambdaInvocationStyleAsync AwsLambdaInvocationStyle = "ASYNC"
)

func (e AwsLambdaInvocationStyle) IsValid() bool {
	switch e {
	case AwsLambdaInvocationStyleSync, AwsLambdaInvocationStyleAsync:
		return true
	}
	return false
}

func (e AwsLambdaInvocationStyle) String() string {
	return string(e)
}

func (e *AwsLambdaInvocationStyle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AwsLambdaInvocationStyle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AwsLambdaInvocationStyle", str)
	}
	return nil
}

func (e AwsLambdaInvocationStyle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AzureFnAuthLevel string

const (
	AzureFnAuthLevelAnonymous AzureFnAuthLevel = "ANONYMOUS"
	AzureFnAuthLevelFunction  AzureFnAuthLevel = "FUNCTION"
	AzureFnAuthLevelAdmin     AzureFnAuthLevel = "ADMIN"
)

func (e AzureFnAuthLevel) IsValid() bool {
	switch e {
	case AzureFnAuthLevelAnonymous, AzureFnAuthLevelFunction, AzureFnAuthLevelAdmin:
		return true
	}
	return false
}

func (e AzureFnAuthLevel) String() string {
	return string(e)
}

func (e *AzureFnAuthLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AzureFnAuthLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AzureFnAuthLevel", str)
	}
	return nil
}

func (e AzureFnAuthLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PathMatchType string

const (
	PathMatchTypePrefix PathMatchType = "PREFIX"
	PathMatchTypeRegex  PathMatchType = "REGEX"
	PathMatchTypeExact  PathMatchType = "EXACT"
)

func (e PathMatchType) IsValid() bool {
	switch e {
	case PathMatchTypePrefix, PathMatchTypeRegex, PathMatchTypeExact:
		return true
	}
	return false
}

func (e PathMatchType) String() string {
	return string(e)
}

func (e *PathMatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PathMatchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PathMatchType", str)
	}
	return nil
}

func (e PathMatchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type State string

const (
	StatePending  State = "PENDING"
	StateAccepted State = "ACCEPTED"
	StateRejected State = "REJECTED"
)

func (e State) IsValid() bool {
	switch e {
	case StatePending, StateAccepted, StateRejected:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
