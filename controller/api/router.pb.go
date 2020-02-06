// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RouteChange_Action int32

const (
	// Unused/invalid default value.
	RouteChange_ACTION_UNSPECIFIED RouteChange_Action = 0
	// ACTION_CREATE represents a route that was created.
	RouteChange_ACTION_CREATE RouteChange_Action = 1
	// ACTION_UPDATE represents a route that was updated.
	RouteChange_ACTION_UPDATE RouteChange_Action = 2
	// ACTION_DELETE represents a route that was deleted.
	RouteChange_ACTION_DELETE RouteChange_Action = 3
)

var RouteChange_Action_name = map[int32]string{
	0: "ACTION_UNSPECIFIED",
	1: "ACTION_CREATE",
	2: "ACTION_UPDATE",
	3: "ACTION_DELETE",
}

var RouteChange_Action_value = map[string]int32{
	"ACTION_UNSPECIFIED": 0,
	"ACTION_CREATE":      1,
	"ACTION_UPDATE":      2,
	"ACTION_DELETE":      3,
}

func (x RouteChange_Action) String() string {
	return proto.EnumName(RouteChange_Action_name, int32(x))
}

func (RouteChange_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5, 0}
}

// SetRoutesRequest is a request to set a list of app routes.
type SetRoutesRequest struct {
	// app_routes are the app routes to set.
	AppRoutes []*AppRoutes `protobuf:"bytes,1,rep,name=app_routes,json=appRoutes,proto3" json:"app_routes,omitempty"`
	// dry_run indicates whether to just generate the changes that would be
	// applied to existing routes by this request (true) or to also atomically
	// apply the route changes (false).
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// expected_state is the expected state of existing routes to apply this
	// request to, with the request failing if the actual state differs.
	ExpectedState        []byte   `protobuf:"bytes,3,opt,name=expected_state,json=expectedState,proto3" json:"expected_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRoutesRequest) Reset()         { *m = SetRoutesRequest{} }
func (m *SetRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*SetRoutesRequest) ProtoMessage()    {}
func (*SetRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *SetRoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRoutesRequest.Unmarshal(m, b)
}
func (m *SetRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRoutesRequest.Marshal(b, m, deterministic)
}
func (m *SetRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRoutesRequest.Merge(m, src)
}
func (m *SetRoutesRequest) XXX_Size() int {
	return xxx_messageInfo_SetRoutesRequest.Size(m)
}
func (m *SetRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRoutesRequest proto.InternalMessageInfo

func (m *SetRoutesRequest) GetAppRoutes() []*AppRoutes {
	if m != nil {
		return m.AppRoutes
	}
	return nil
}

func (m *SetRoutesRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *SetRoutesRequest) GetExpectedState() []byte {
	if m != nil {
		return m.ExpectedState
	}
	return nil
}

// SetRoutesResponse is a response to a SetRoutes request.
type SetRoutesResponse struct {
	// route_changes is the list of route changes that were either applied if
	// dry_run=false or that would have been applied if dry_run=true.
	RouteChanges []*RouteChange `protobuf:"bytes,1,rep,name=route_changes,json=routeChanges,proto3" json:"route_changes,omitempty"`
	// dry_run indicates whether the request was a dry run or not.
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// applied_to_state is the state of the existing routes that this request was
	// applied to, and can be used as the expected_state in a subsequent request
	// to confirm the application of a dry run.
	AppliedToState       []byte   `protobuf:"bytes,3,opt,name=applied_to_state,json=appliedToState,proto3" json:"applied_to_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRoutesResponse) Reset()         { *m = SetRoutesResponse{} }
func (m *SetRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*SetRoutesResponse) ProtoMessage()    {}
func (*SetRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *SetRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRoutesResponse.Unmarshal(m, b)
}
func (m *SetRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRoutesResponse.Marshal(b, m, deterministic)
}
func (m *SetRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRoutesResponse.Merge(m, src)
}
func (m *SetRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_SetRoutesResponse.Size(m)
}
func (m *SetRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetRoutesResponse proto.InternalMessageInfo

func (m *SetRoutesResponse) GetRouteChanges() []*RouteChange {
	if m != nil {
		return m.RouteChanges
	}
	return nil
}

func (m *SetRoutesResponse) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *SetRoutesResponse) GetAppliedToState() []byte {
	if m != nil {
		return m.AppliedToState
	}
	return nil
}

// ListAppRoutesRequest is a request to list routes for a set of apps.
type ListAppRoutesRequest struct {
	// apps is the set of apps to list routes for.
	Apps                 []string `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppRoutesRequest) Reset()         { *m = ListAppRoutesRequest{} }
func (m *ListAppRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*ListAppRoutesRequest) ProtoMessage()    {}
func (*ListAppRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *ListAppRoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppRoutesRequest.Unmarshal(m, b)
}
func (m *ListAppRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppRoutesRequest.Marshal(b, m, deterministic)
}
func (m *ListAppRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppRoutesRequest.Merge(m, src)
}
func (m *ListAppRoutesRequest) XXX_Size() int {
	return xxx_messageInfo_ListAppRoutesRequest.Size(m)
}
func (m *ListAppRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppRoutesRequest proto.InternalMessageInfo

func (m *ListAppRoutesRequest) GetApps() []string {
	if m != nil {
		return m.Apps
	}
	return nil
}

// ListAppRoutesResponse is a response to list routes for a set of apps.
type ListAppRoutesResponse struct {
	// app_routes are the requested routes.
	AppRoutes            []*AppRoutes `protobuf:"bytes,1,rep,name=app_routes,json=appRoutes,proto3" json:"app_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListAppRoutesResponse) Reset()         { *m = ListAppRoutesResponse{} }
func (m *ListAppRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*ListAppRoutesResponse) ProtoMessage()    {}
func (*ListAppRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *ListAppRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppRoutesResponse.Unmarshal(m, b)
}
func (m *ListAppRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppRoutesResponse.Marshal(b, m, deterministic)
}
func (m *ListAppRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppRoutesResponse.Merge(m, src)
}
func (m *ListAppRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_ListAppRoutesResponse.Size(m)
}
func (m *ListAppRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppRoutesResponse proto.InternalMessageInfo

func (m *ListAppRoutesResponse) GetAppRoutes() []*AppRoutes {
	if m != nil {
		return m.AppRoutes
	}
	return nil
}

// AppRoutes is a list of desired routes for an app.
type AppRoutes struct {
	// app is an identifier for the app.
	App string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	// routes are the desired routes for the app.
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppRoutes) Reset()         { *m = AppRoutes{} }
func (m *AppRoutes) String() string { return proto.CompactTextString(m) }
func (*AppRoutes) ProtoMessage()    {}
func (*AppRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *AppRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppRoutes.Unmarshal(m, b)
}
func (m *AppRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppRoutes.Marshal(b, m, deterministic)
}
func (m *AppRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppRoutes.Merge(m, src)
}
func (m *AppRoutes) XXX_Size() int {
	return xxx_messageInfo_AppRoutes.Size(m)
}
func (m *AppRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_AppRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_AppRoutes proto.InternalMessageInfo

func (m *AppRoutes) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *AppRoutes) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// RouteChange is a change made to a route in response to a SetRoutes request.
type RouteChange struct {
	// action is the action that was performed.
	Action RouteChange_Action `protobuf:"varint,1,opt,name=action,proto3,enum=flynn.api.v1.RouteChange_Action" json:"action,omitempty"`
	// before is the route before the action was applied.
	Before *Route `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	// after is the route after the action was applied.
	After                *Route   `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteChange) Reset()         { *m = RouteChange{} }
func (m *RouteChange) String() string { return proto.CompactTextString(m) }
func (*RouteChange) ProtoMessage()    {}
func (*RouteChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *RouteChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteChange.Unmarshal(m, b)
}
func (m *RouteChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteChange.Marshal(b, m, deterministic)
}
func (m *RouteChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteChange.Merge(m, src)
}
func (m *RouteChange) XXX_Size() int {
	return xxx_messageInfo_RouteChange.Size(m)
}
func (m *RouteChange) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteChange.DiscardUnknown(m)
}

var xxx_messageInfo_RouteChange proto.InternalMessageInfo

func (m *RouteChange) GetAction() RouteChange_Action {
	if m != nil {
		return m.Action
	}
	return RouteChange_ACTION_UNSPECIFIED
}

func (m *RouteChange) GetBefore() *Route {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *RouteChange) GetAfter() *Route {
	if m != nil {
		return m.After
	}
	return nil
}

// Route is a HTTP or TCP route.
type Route struct {
	// Required. The parent resource name, in the format `apps/{APP_ID}` or
	// `apps/{APP_DISPLAY_NAME}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The discoverd service target for this route.
	ServiceTarget *Route_ServiceTarget `protobuf:"bytes,2,opt,name=service_target,json=serviceTarget,proto3" json:"service_target,omitempty"`
	// The protocol-specific configuration.
	//
	// Types that are valid to be assigned to Config:
	//	*Route_Http
	//	*Route_Tcp
	Config               isRoute_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Route) GetServiceTarget() *Route_ServiceTarget {
	if m != nil {
		return m.ServiceTarget
	}
	return nil
}

type isRoute_Config interface {
	isRoute_Config()
}

type Route_Http struct {
	Http *Route_HTTP `protobuf:"bytes,3,opt,name=http,proto3,oneof"`
}

type Route_Tcp struct {
	Tcp *Route_TCP `protobuf:"bytes,4,opt,name=tcp,proto3,oneof"`
}

func (*Route_Http) isRoute_Config() {}

func (*Route_Tcp) isRoute_Config() {}

func (m *Route) GetConfig() isRoute_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Route) GetHttp() *Route_HTTP {
	if x, ok := m.GetConfig().(*Route_Http); ok {
		return x.Http
	}
	return nil
}

func (m *Route) GetTcp() *Route_TCP {
	if x, ok := m.GetConfig().(*Route_Tcp); ok {
		return x.Tcp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Route) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Route_Http)(nil),
		(*Route_Tcp)(nil),
	}
}

// A discoverd service target.
type Route_ServiceTarget struct {
	// Required. The discoverd service name to route requests to.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Wait for in-flight requests to this target to finish before backends are terminated.
	DrainBackends        bool     `protobuf:"varint,2,opt,name=drain_backends,json=drainBackends,proto3" json:"drain_backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route_ServiceTarget) Reset()         { *m = Route_ServiceTarget{} }
func (m *Route_ServiceTarget) String() string { return proto.CompactTextString(m) }
func (*Route_ServiceTarget) ProtoMessage()    {}
func (*Route_ServiceTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6, 0}
}

func (m *Route_ServiceTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_ServiceTarget.Unmarshal(m, b)
}
func (m *Route_ServiceTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_ServiceTarget.Marshal(b, m, deterministic)
}
func (m *Route_ServiceTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_ServiceTarget.Merge(m, src)
}
func (m *Route_ServiceTarget) XXX_Size() int {
	return xxx_messageInfo_Route_ServiceTarget.Size(m)
}
func (m *Route_ServiceTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_ServiceTarget.DiscardUnknown(m)
}

var xxx_messageInfo_Route_ServiceTarget proto.InternalMessageInfo

func (m *Route_ServiceTarget) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Route_ServiceTarget) GetDrainBackends() bool {
	if m != nil {
		return m.DrainBackends
	}
	return false
}

type Route_HTTP struct {
	// Required. The name of the server that this route matches. May contain up
	// to 10 wildcard labels for plaintext HTTP routes or a single wildcard
	// label for TLS routes, followed by one or more non-wildcard labels. This
	// is matched against SNI to choose the TLS configuration and the Host
	// header to select the route.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// The HTTP path prefix to match against. Defaults to the root path `/`. To
	// create a route with a non-root path prefix, a root path route must already
	// exist for the same listener and domain, which will be used for the TLS and
	// HSTS configuration. Trailing slashes are stripped and only full path
	// segments are matched. The full unstripped path is sent in requests to the
	// target.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route_HTTP) Reset()         { *m = Route_HTTP{} }
func (m *Route_HTTP) String() string { return proto.CompactTextString(m) }
func (*Route_HTTP) ProtoMessage()    {}
func (*Route_HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6, 1}
}

func (m *Route_HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_HTTP.Unmarshal(m, b)
}
func (m *Route_HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_HTTP.Marshal(b, m, deterministic)
}
func (m *Route_HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_HTTP.Merge(m, src)
}
func (m *Route_HTTP) XXX_Size() int {
	return xxx_messageInfo_Route_HTTP.Size(m)
}
func (m *Route_HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_Route_HTTP proto.InternalMessageInfo

func (m *Route_HTTP) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Route_HTTP) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Route_TCP struct {
	// The TCP port configuration for the route. Required and only valid for TCP
	// listeners.
	Port                 *Route_TCPPort `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Route_TCP) Reset()         { *m = Route_TCP{} }
func (m *Route_TCP) String() string { return proto.CompactTextString(m) }
func (*Route_TCP) ProtoMessage()    {}
func (*Route_TCP) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6, 2}
}

func (m *Route_TCP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_TCP.Unmarshal(m, b)
}
func (m *Route_TCP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_TCP.Marshal(b, m, deterministic)
}
func (m *Route_TCP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_TCP.Merge(m, src)
}
func (m *Route_TCP) XXX_Size() int {
	return xxx_messageInfo_Route_TCP.Size(m)
}
func (m *Route_TCP) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_TCP.DiscardUnknown(m)
}

var xxx_messageInfo_Route_TCP proto.InternalMessageInfo

func (m *Route_TCP) GetPort() *Route_TCPPort {
	if m != nil {
		return m.Port
	}
	return nil
}

type Route_TCPPort struct {
	// The TCP port to bind to. If unspecified, a port will be automatically chosen
	// during route creation and provided in the response.
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route_TCPPort) Reset()         { *m = Route_TCPPort{} }
func (m *Route_TCPPort) String() string { return proto.CompactTextString(m) }
func (*Route_TCPPort) ProtoMessage()    {}
func (*Route_TCPPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6, 3}
}

func (m *Route_TCPPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_TCPPort.Unmarshal(m, b)
}
func (m *Route_TCPPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_TCPPort.Marshal(b, m, deterministic)
}
func (m *Route_TCPPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_TCPPort.Merge(m, src)
}
func (m *Route_TCPPort) XXX_Size() int {
	return xxx_messageInfo_Route_TCPPort.Size(m)
}
func (m *Route_TCPPort) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_TCPPort.DiscardUnknown(m)
}

var xxx_messageInfo_Route_TCPPort proto.InternalMessageInfo

func (m *Route_TCPPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("flynn.api.v1.RouteChange_Action", RouteChange_Action_name, RouteChange_Action_value)
	proto.RegisterType((*SetRoutesRequest)(nil), "flynn.api.v1.SetRoutesRequest")
	proto.RegisterType((*SetRoutesResponse)(nil), "flynn.api.v1.SetRoutesResponse")
	proto.RegisterType((*ListAppRoutesRequest)(nil), "flynn.api.v1.ListAppRoutesRequest")
	proto.RegisterType((*ListAppRoutesResponse)(nil), "flynn.api.v1.ListAppRoutesResponse")
	proto.RegisterType((*AppRoutes)(nil), "flynn.api.v1.AppRoutes")
	proto.RegisterType((*RouteChange)(nil), "flynn.api.v1.RouteChange")
	proto.RegisterType((*Route)(nil), "flynn.api.v1.Route")
	proto.RegisterType((*Route_ServiceTarget)(nil), "flynn.api.v1.Route.ServiceTarget")
	proto.RegisterType((*Route_HTTP)(nil), "flynn.api.v1.Route.HTTP")
	proto.RegisterType((*Route_TCP)(nil), "flynn.api.v1.Route.TCP")
	proto.RegisterType((*Route_TCPPort)(nil), "flynn.api.v1.Route.TCPPort")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x96, 0x2e, 0x5b, 0x6f, 0x9b, 0xaa, 0x33, 0xb0, 0x95, 0x22, 0xa0, 0x0b, 0x42, 0x2a,
	0x4c, 0x0a, 0xa2, 0x48, 0x13, 0x4f, 0x48, 0x5d, 0x17, 0xb4, 0xa1, 0x69, 0xab, 0xdc, 0x4c, 0x62,
	0xbc, 0x44, 0x5e, 0xe2, 0x6e, 0x11, 0x9b, 0x63, 0x1c, 0x77, 0x62, 0xbf, 0xc0, 0x07, 0xf0, 0x1d,
	0x3c, 0xf1, 0x71, 0x3c, 0xa1, 0xd8, 0xee, 0x68, 0xa6, 0x76, 0x0f, 0xbc, 0xd9, 0x27, 0xe7, 0xdc,
	0x7b, 0xee, 0x51, 0xae, 0xa1, 0x2e, 0xb2, 0x89, 0xa4, 0xc2, 0xe7, 0x22, 0x93, 0x19, 0xaa, 0x8f,
	0x2f, 0x6f, 0x18, 0xf3, 0x09, 0x4f, 0xfd, 0xeb, 0xb7, 0xde, 0x0f, 0x0b, 0x9a, 0x23, 0x2a, 0x71,
	0xc1, 0xc8, 0x31, 0xfd, 0x36, 0xa1, 0xb9, 0x44, 0x3b, 0x00, 0x84, 0xf3, 0x48, 0xc9, 0xf2, 0x96,
	0xd5, 0xb1, 0xbb, 0xb5, 0xde, 0xa6, 0x3f, 0xab, 0xf3, 0xfb, 0x9c, 0x1b, 0x4d, 0x95, 0x4c, 0x8f,
	0x68, 0x13, 0x56, 0x13, 0x71, 0x13, 0x89, 0x09, 0x6b, 0x2d, 0x77, 0xac, 0xee, 0x1a, 0x76, 0x12,
	0x71, 0x83, 0x27, 0x0c, 0xbd, 0x84, 0x06, 0xfd, 0xce, 0x69, 0x2c, 0x69, 0x12, 0xe5, 0x92, 0x48,
	0xda, 0xb2, 0x3b, 0x56, 0xb7, 0x8e, 0xdd, 0x29, 0x3a, 0x2a, 0x40, 0xef, 0xa7, 0x05, 0xeb, 0x33,
	0x66, 0x72, 0x9e, 0xb1, 0x9c, 0xa2, 0x0f, 0xe0, 0x2a, 0x27, 0x51, 0x7c, 0x41, 0xd8, 0xf9, 0xad,
	0xa1, 0xc7, 0x65, 0x43, 0x4a, 0x34, 0x50, 0x0c, 0xac, 0x07, 0xd6, 0x97, 0x7b, 0x5c, 0x75, 0xa1,
	0x49, 0x38, 0xbf, 0x4c, 0x69, 0x12, 0xc9, 0xac, 0xe4, 0xab, 0x61, 0xf0, 0x30, 0xd3, 0xc6, 0x5e,
	0xc3, 0xc3, 0xc3, 0x34, 0x97, 0xff, 0x86, 0x36, 0x41, 0x21, 0xa8, 0x10, 0xce, 0xb5, 0xa3, 0x2a,
	0x56, 0x67, 0xef, 0x18, 0x1e, 0xdd, 0xe1, 0x9a, 0x39, 0xfe, 0x33, 0x55, 0xef, 0x13, 0x54, 0x6f,
	0x71, 0xd4, 0x04, 0x9b, 0x70, 0xde, 0xb2, 0x3a, 0x56, 0xb7, 0x8a, 0x8b, 0x23, 0xda, 0x06, 0xc7,
	0x94, 0x5c, 0x56, 0x25, 0x1f, 0xcc, 0xc9, 0x05, 0x1b, 0x8a, 0xf7, 0xc7, 0x82, 0xda, 0x4c, 0x52,
	0xe8, 0x3d, 0x38, 0x24, 0x96, 0x69, 0xc6, 0x54, 0xc5, 0x46, 0xaf, 0xb3, 0x30, 0x54, 0xbf, 0xaf,
	0x78, 0xd8, 0xf0, 0x8b, 0xb6, 0x67, 0x74, 0x9c, 0x09, 0xaa, 0x42, 0x5d, 0xd4, 0x56, 0x53, 0xd0,
	0x2b, 0x58, 0x21, 0x63, 0x49, 0x85, 0x8a, 0x77, 0x01, 0x57, 0x33, 0xbc, 0x53, 0x70, 0x74, 0x27,
	0xb4, 0x01, 0xa8, 0x3f, 0x08, 0x0f, 0x8e, 0x8f, 0xa2, 0x93, 0xa3, 0xd1, 0x30, 0x18, 0x1c, 0x7c,
	0x3c, 0x08, 0xf6, 0x9a, 0x4b, 0x68, 0x1d, 0x5c, 0x83, 0x0f, 0x70, 0xd0, 0x0f, 0x83, 0xa6, 0x35,
	0x03, 0x9d, 0x0c, 0xf7, 0x0a, 0x68, 0x79, 0x06, 0xda, 0x0b, 0x0e, 0x83, 0x30, 0x68, 0xda, 0xde,
	0x6f, 0x1b, 0x56, 0x54, 0x2f, 0xb4, 0x01, 0x0e, 0x27, 0x82, 0x32, 0x69, 0x82, 0x34, 0x37, 0xb4,
	0x0f, 0x8d, 0x9c, 0x8a, 0xeb, 0x34, 0xa6, 0x91, 0x24, 0xe2, 0x9c, 0x4a, 0x33, 0xdc, 0xd6, 0x1c,
	0xc3, 0xfe, 0x48, 0x33, 0x43, 0x45, 0xc4, 0x6e, 0x3e, 0x7b, 0x45, 0x3e, 0x54, 0x2e, 0xa4, 0xe4,
	0x66, 0xe0, 0xd6, 0x3c, 0xfd, 0x7e, 0x18, 0x0e, 0xf7, 0x97, 0xb0, 0xe2, 0xa1, 0x6d, 0xb0, 0x65,
	0xcc, 0x5b, 0x15, 0x45, 0xdf, 0x9c, 0x47, 0x0f, 0x07, 0x05, 0xbb, 0x60, 0xb5, 0x4f, 0xc1, 0x2d,
	0x35, 0x47, 0x5b, 0x50, 0x9f, 0xfa, 0x66, 0xe4, 0x8a, 0x9a, 0xa9, 0x6a, 0x06, 0x3b, 0x22, 0x57,
	0xb4, 0x58, 0xc1, 0x44, 0x90, 0x94, 0x45, 0x67, 0x24, 0xfe, 0x4a, 0x59, 0x92, 0x9b, 0x65, 0x70,
	0x15, 0xba, 0x6b, 0xc0, 0x76, 0x0f, 0x2a, 0x85, 0xaf, 0x22, 0xa1, 0x24, 0xbb, 0x22, 0x29, 0x9b,
	0x26, 0xa4, 0x6f, 0xc5, 0x1f, 0xcf, 0x89, 0xbc, 0x50, 0xe2, 0x2a, 0x56, 0xe7, 0xf6, 0x0e, 0xd8,
	0xe1, 0x60, 0x88, 0xde, 0x40, 0x85, 0x67, 0x42, 0x47, 0x5a, 0xeb, 0x3d, 0x59, 0x30, 0xc3, 0x30,
	0x13, 0x12, 0x2b, 0x62, 0xfb, 0x29, 0xac, 0x1a, 0x40, 0x95, 0x9d, 0x6a, 0x5d, 0xfd, 0x79, 0x77,
	0x0d, 0x9c, 0x38, 0x63, 0xe3, 0xf4, 0xbc, 0xf7, 0xcb, 0x02, 0x47, 0x15, 0x10, 0xe8, 0x10, 0xaa,
	0xb7, 0x2f, 0x04, 0x7a, 0x56, 0xee, 0x71, 0xf7, 0x1d, 0x6b, 0x3f, 0x5f, 0xf8, 0xdd, 0xac, 0xe4,
	0x67, 0x70, 0x4b, 0xbb, 0x8a, 0xbc, 0xb2, 0x62, 0xde, 0xd2, 0xb7, 0x5f, 0xdc, 0xcb, 0xd1, 0x95,
	0x77, 0x57, 0xbe, 0xd8, 0x84, 0xa7, 0x67, 0x8e, 0x7a, 0x73, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0x61, 0xa5, 0xa6, 0x83, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	// SetRoutes takes a desired list of routes for a set of apps, calculates the
	// changes that are needed to the existing routes to realise that list, and
	// then either atomically applies those changes or returns them for user
	// confirmation (otherwise known as a dry run).
	//
	// The given list of routes for each app is expected to contain the desired
	// configuration for all of the app's routes, and so if any existing routes
	// are not contained in the list, or they match ones in the list but have
	// different configuration, then they will be either deleted or updated.
	SetRoutes(ctx context.Context, in *SetRoutesRequest, opts ...grpc.CallOption) (*SetRoutesResponse, error)
	// ListAppRoutes list routes for a set of apps.
	ListAppRoutes(ctx context.Context, in *ListAppRoutesRequest, opts ...grpc.CallOption) (*ListAppRoutesResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) SetRoutes(ctx context.Context, in *SetRoutesRequest, opts ...grpc.CallOption) (*SetRoutesResponse, error) {
	out := new(SetRoutesResponse)
	err := c.cc.Invoke(ctx, "/flynn.api.v1.Router/SetRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) ListAppRoutes(ctx context.Context, in *ListAppRoutesRequest, opts ...grpc.CallOption) (*ListAppRoutesResponse, error) {
	out := new(ListAppRoutesResponse)
	err := c.cc.Invoke(ctx, "/flynn.api.v1.Router/ListAppRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	// SetRoutes takes a desired list of routes for a set of apps, calculates the
	// changes that are needed to the existing routes to realise that list, and
	// then either atomically applies those changes or returns them for user
	// confirmation (otherwise known as a dry run).
	//
	// The given list of routes for each app is expected to contain the desired
	// configuration for all of the app's routes, and so if any existing routes
	// are not contained in the list, or they match ones in the list but have
	// different configuration, then they will be either deleted or updated.
	SetRoutes(context.Context, *SetRoutesRequest) (*SetRoutesResponse, error)
	// ListAppRoutes list routes for a set of apps.
	ListAppRoutes(context.Context, *ListAppRoutesRequest) (*ListAppRoutesResponse, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) SetRoutes(ctx context.Context, req *SetRoutesRequest) (*SetRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoutes not implemented")
}
func (*UnimplementedRouterServer) ListAppRoutes(ctx context.Context, req *ListAppRoutesRequest) (*ListAppRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppRoutes not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_SetRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SetRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flynn.api.v1.Router/SetRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SetRoutes(ctx, req.(*SetRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_ListAppRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ListAppRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flynn.api.v1.Router/ListAppRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ListAppRoutes(ctx, req.(*ListAppRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flynn.api.v1.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRoutes",
			Handler:    _Router_SetRoutes_Handler,
		},
		{
			MethodName: "ListAppRoutes",
			Handler:    _Router_ListAppRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
