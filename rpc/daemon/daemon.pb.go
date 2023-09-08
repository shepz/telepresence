// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: daemon/daemon.proto

package daemon

import (
	common "github.com/telepresenceio/telepresence/rpc/v2/common"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DaemonStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutboundConfig *OutboundInfo       `protobuf:"bytes,4,opt,name=outbound_config,json=outboundConfig,proto3" json:"outbound_config,omitempty"`
	Version        *common.VersionInfo `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DaemonStatus) Reset() {
	*x = DaemonStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaemonStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaemonStatus) ProtoMessage() {}

func (x *DaemonStatus) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaemonStatus.ProtoReflect.Descriptor instead.
func (*DaemonStatus) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *DaemonStatus) GetOutboundConfig() *OutboundInfo {
	if x != nil {
		return x.OutboundConfig
	}
	return nil
}

func (x *DaemonStatus) GetVersion() *common.VersionInfo {
	if x != nil {
		return x.Version
	}
	return nil
}

type Paths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths      []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *Paths) Reset() {
	*x = Paths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paths) ProtoMessage() {}

func (x *Paths) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paths.ProtoReflect.Descriptor instead.
func (*Paths) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{1}
}

func (x *Paths) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Paths) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type DNSMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AliasFor string `protobuf:"bytes,2,opt,name=alias_for,json=aliasFor,proto3" json:"alias_for,omitempty"`
}

func (x *DNSMapping) Reset() {
	*x = DNSMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSMapping) ProtoMessage() {}

func (x *DNSMapping) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSMapping.ProtoReflect.Descriptor instead.
func (*DNSMapping) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{2}
}

func (x *DNSMapping) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSMapping) GetAliasFor() string {
	if x != nil {
		return x.AliasFor
	}
	return ""
}

// DNS configuration for the local DNS resolver
type DNSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// local_ip is the address of the local DNS server. Only used by Linux systems that have no
	// systemd-resolved configured. Defaults to the first line of /etc/resolv.conf
	LocalIp []byte `protobuf:"bytes,1,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	// remote_ip is the address of the kube-dns.kube-system, dns-default.openshift-dns, or similar service,
	RemoteIp []byte `protobuf:"bytes,2,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// Suffixes to exclude
	ExcludeSuffixes []string `protobuf:"bytes,3,rep,name=exclude_suffixes,json=excludeSuffixes,proto3" json:"exclude_suffixes,omitempty"`
	// Suffixes to include. Has higher prio than the excludes
	IncludeSuffixes []string `protobuf:"bytes,4,rep,name=include_suffixes,json=includeSuffixes,proto3" json:"include_suffixes,omitempty"`
	// Exclude are a list of hostname that the DNS resolver will not resolve even if they exist.
	Excludes []string `protobuf:"bytes,8,rep,name=excludes,proto3" json:"excludes,omitempty"`
	// DNSMapping contains a hostname and its associated alias. When requesting the name, the intended behavior is
	// to resolve the alias instead.
	Mappings []*DNSMapping `protobuf:"bytes,9,rep,name=mappings,proto3" json:"mappings,omitempty"`
	// The maximum time wait for a cluster side host lookup.
	LookupTimeout *durationpb.Duration `protobuf:"bytes,6,opt,name=lookup_timeout,json=lookupTimeout,proto3" json:"lookup_timeout,omitempty"`
	// If set, this error indicates why DNS is not working.
	Error string `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DNSConfig) Reset() {
	*x = DNSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSConfig) ProtoMessage() {}

func (x *DNSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSConfig.ProtoReflect.Descriptor instead.
func (*DNSConfig) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{3}
}

func (x *DNSConfig) GetLocalIp() []byte {
	if x != nil {
		return x.LocalIp
	}
	return nil
}

func (x *DNSConfig) GetRemoteIp() []byte {
	if x != nil {
		return x.RemoteIp
	}
	return nil
}

func (x *DNSConfig) GetExcludeSuffixes() []string {
	if x != nil {
		return x.ExcludeSuffixes
	}
	return nil
}

func (x *DNSConfig) GetIncludeSuffixes() []string {
	if x != nil {
		return x.IncludeSuffixes
	}
	return nil
}

func (x *DNSConfig) GetExcludes() []string {
	if x != nil {
		return x.Excludes
	}
	return nil
}

func (x *DNSConfig) GetMappings() []*DNSMapping {
	if x != nil {
		return x.Mappings
	}
	return nil
}

func (x *DNSConfig) GetLookupTimeout() *durationpb.Duration {
	if x != nil {
		return x.LookupTimeout
	}
	return nil
}

func (x *DNSConfig) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// OutboundInfo contains all information that the root daemon needs in order to
// establish outbound traffic to the cluster.
type OutboundInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// session makes it possible for the root daemon to identify itself as the
	// same client as the user daemon.
	Session *manager.SessionInfo `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	// DNS configuration
	Dns *DNSConfig `protobuf:"bytes,3,opt,name=dns,proto3" json:"dns,omitempty"`
	// also_proxy are user-added subnets.
	AlsoProxySubnets []*manager.IPNet `protobuf:"bytes,5,rep,name=also_proxy_subnets,json=alsoProxySubnets,proto3" json:"also_proxy_subnets,omitempty"`
	// never_proxy_subnets are subnets that the daemon should not proxy but resolve
	// via the underlying network interface.
	NeverProxySubnets []*manager.IPNet `protobuf:"bytes,6,rep,name=never_proxy_subnets,json=neverProxySubnets,proto3" json:"never_proxy_subnets,omitempty"`
	// Users home directory
	HomeDir string `protobuf:"bytes,7,opt,name=home_dir,json=homeDir,proto3" json:"home_dir,omitempty"`
	// Traffic manager namespace
	ManagerNamespace string `protobuf:"bytes,8,opt,name=manager_namespace,json=managerNamespace,proto3" json:"manager_namespace,omitempty"`
}

func (x *OutboundInfo) Reset() {
	*x = OutboundInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutboundInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutboundInfo) ProtoMessage() {}

func (x *OutboundInfo) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutboundInfo.ProtoReflect.Descriptor instead.
func (*OutboundInfo) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{4}
}

func (x *OutboundInfo) GetSession() *manager.SessionInfo {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *OutboundInfo) GetDns() *DNSConfig {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *OutboundInfo) GetAlsoProxySubnets() []*manager.IPNet {
	if x != nil {
		return x.AlsoProxySubnets
	}
	return nil
}

func (x *OutboundInfo) GetNeverProxySubnets() []*manager.IPNet {
	if x != nil {
		return x.NeverProxySubnets
	}
	return nil
}

func (x *OutboundInfo) GetHomeDir() string {
	if x != nil {
		return x.HomeDir
	}
	return ""
}

func (x *OutboundInfo) GetManagerNamespace() string {
	if x != nil {
		return x.ManagerNamespace
	}
	return ""
}

type NetworkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnets      []*manager.IPNet `protobuf:"bytes,1,rep,name=subnets,proto3" json:"subnets,omitempty"`
	OutboundInfo *OutboundInfo    `protobuf:"bytes,2,opt,name=outbound_info,json=outboundInfo,proto3" json:"outbound_info,omitempty"`
}

func (x *NetworkConfig) Reset() {
	*x = NetworkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConfig) ProtoMessage() {}

func (x *NetworkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConfig.ProtoReflect.Descriptor instead.
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkConfig) GetSubnets() []*manager.IPNet {
	if x != nil {
		return x.Subnets
	}
	return nil
}

func (x *NetworkConfig) GetOutboundInfo() *OutboundInfo {
	if x != nil {
		return x.OutboundInfo
	}
	return nil
}

type SetDNSExcludesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Excludes []string `protobuf:"bytes,1,rep,name=excludes,proto3" json:"excludes,omitempty"`
}

func (x *SetDNSExcludesRequest) Reset() {
	*x = SetDNSExcludesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSExcludesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSExcludesRequest) ProtoMessage() {}

func (x *SetDNSExcludesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSExcludesRequest.ProtoReflect.Descriptor instead.
func (*SetDNSExcludesRequest) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{6}
}

func (x *SetDNSExcludesRequest) GetExcludes() []string {
	if x != nil {
		return x.Excludes
	}
	return nil
}

type SetDNSMappingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mappings []*DNSMapping `protobuf:"bytes,1,rep,name=mappings,proto3" json:"mappings,omitempty"`
}

func (x *SetDNSMappingsRequest) Reset() {
	*x = SetDNSMappingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_daemon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSMappingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSMappingsRequest) ProtoMessage() {}

func (x *SetDNSMappingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_daemon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSMappingsRequest.ProtoReflect.Descriptor instead.
func (*SetDNSMappingsRequest) Descriptor() ([]byte, []int) {
	return file_daemon_daemon_proto_rawDescGZIP(), []int{7}
}

func (x *SetDNSMappingsRequest) GetMappings() []*DNSMapping {
	if x != nil {
		return x.Mappings
	}
	return nil
}

var File_daemon_daemon_proto protoreflect.FileDescriptor

var file_daemon_daemon_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4a, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22,
	0x3d, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x3d,
	0x0a, 0x0a, 0x44, 0x4e, 0x53, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x46, 0x6f, 0x72, 0x22, 0xd0, 0x02,
	0x0a, 0x09, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x4e,
	0x53, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06,
	0x22, 0xef, 0x02, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x64, 0x6e, 0x73,
	0x12, 0x49, 0x0a, 0x12, 0x61, 0x6c, 0x73, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x52, 0x10, 0x61, 0x6c, 0x73, 0x6f, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x6e,
	0x65, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x49, 0x50, 0x4e, 0x65, 0x74, 0x52, 0x11, 0x6e, 0x65, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65,
	0x5f, 0x64, 0x69, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d, 0x65,
	0x44, 0x69, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x09,
	0x10, 0x0a, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x50, 0x4e,
	0x65, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x6f,
	0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44,
	0x4e, 0x53, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x4e, 0x53, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xad,
	0x06, 0x0a, 0x06, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0a,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x45, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e,
	0x53, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44,
	0x4e, 0x53, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x0e,
	0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32, 0x2f,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daemon_daemon_proto_rawDescOnce sync.Once
	file_daemon_daemon_proto_rawDescData = file_daemon_daemon_proto_rawDesc
)

func file_daemon_daemon_proto_rawDescGZIP() []byte {
	file_daemon_daemon_proto_rawDescOnce.Do(func() {
		file_daemon_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_daemon_daemon_proto_rawDescData)
	})
	return file_daemon_daemon_proto_rawDescData
}

var file_daemon_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_daemon_daemon_proto_goTypes = []interface{}{
	(*DaemonStatus)(nil),            // 0: telepresence.daemon.DaemonStatus
	(*Paths)(nil),                   // 1: telepresence.daemon.Paths
	(*DNSMapping)(nil),              // 2: telepresence.daemon.DNSMapping
	(*DNSConfig)(nil),               // 3: telepresence.daemon.DNSConfig
	(*OutboundInfo)(nil),            // 4: telepresence.daemon.OutboundInfo
	(*NetworkConfig)(nil),           // 5: telepresence.daemon.NetworkConfig
	(*SetDNSExcludesRequest)(nil),   // 6: telepresence.daemon.SetDNSExcludesRequest
	(*SetDNSMappingsRequest)(nil),   // 7: telepresence.daemon.SetDNSMappingsRequest
	(*common.VersionInfo)(nil),      // 8: telepresence.common.VersionInfo
	(*durationpb.Duration)(nil),     // 9: google.protobuf.Duration
	(*manager.SessionInfo)(nil),     // 10: telepresence.manager.SessionInfo
	(*manager.IPNet)(nil),           // 11: telepresence.manager.IPNet
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
	(*manager.LogLevelRequest)(nil), // 13: telepresence.manager.LogLevelRequest
}
var file_daemon_daemon_proto_depIdxs = []int32{
	4,  // 0: telepresence.daemon.DaemonStatus.outbound_config:type_name -> telepresence.daemon.OutboundInfo
	8,  // 1: telepresence.daemon.DaemonStatus.version:type_name -> telepresence.common.VersionInfo
	2,  // 2: telepresence.daemon.DNSConfig.mappings:type_name -> telepresence.daemon.DNSMapping
	9,  // 3: telepresence.daemon.DNSConfig.lookup_timeout:type_name -> google.protobuf.Duration
	10, // 4: telepresence.daemon.OutboundInfo.session:type_name -> telepresence.manager.SessionInfo
	3,  // 5: telepresence.daemon.OutboundInfo.dns:type_name -> telepresence.daemon.DNSConfig
	11, // 6: telepresence.daemon.OutboundInfo.also_proxy_subnets:type_name -> telepresence.manager.IPNet
	11, // 7: telepresence.daemon.OutboundInfo.never_proxy_subnets:type_name -> telepresence.manager.IPNet
	11, // 8: telepresence.daemon.NetworkConfig.subnets:type_name -> telepresence.manager.IPNet
	4,  // 9: telepresence.daemon.NetworkConfig.outbound_info:type_name -> telepresence.daemon.OutboundInfo
	2,  // 10: telepresence.daemon.SetDNSMappingsRequest.mappings:type_name -> telepresence.daemon.DNSMapping
	12, // 11: telepresence.daemon.Daemon.Version:input_type -> google.protobuf.Empty
	12, // 12: telepresence.daemon.Daemon.Status:input_type -> google.protobuf.Empty
	12, // 13: telepresence.daemon.Daemon.Quit:input_type -> google.protobuf.Empty
	4,  // 14: telepresence.daemon.Daemon.Connect:input_type -> telepresence.daemon.OutboundInfo
	12, // 15: telepresence.daemon.Daemon.Disconnect:input_type -> google.protobuf.Empty
	12, // 16: telepresence.daemon.Daemon.GetNetworkConfig:input_type -> google.protobuf.Empty
	1,  // 17: telepresence.daemon.Daemon.SetDnsSearchPath:input_type -> telepresence.daemon.Paths
	6,  // 18: telepresence.daemon.Daemon.SetDNSExcludes:input_type -> telepresence.daemon.SetDNSExcludesRequest
	7,  // 19: telepresence.daemon.Daemon.SetDNSMappings:input_type -> telepresence.daemon.SetDNSMappingsRequest
	13, // 20: telepresence.daemon.Daemon.SetLogLevel:input_type -> telepresence.manager.LogLevelRequest
	12, // 21: telepresence.daemon.Daemon.WaitForNetwork:input_type -> google.protobuf.Empty
	8,  // 22: telepresence.daemon.Daemon.Version:output_type -> telepresence.common.VersionInfo
	0,  // 23: telepresence.daemon.Daemon.Status:output_type -> telepresence.daemon.DaemonStatus
	12, // 24: telepresence.daemon.Daemon.Quit:output_type -> google.protobuf.Empty
	0,  // 25: telepresence.daemon.Daemon.Connect:output_type -> telepresence.daemon.DaemonStatus
	12, // 26: telepresence.daemon.Daemon.Disconnect:output_type -> google.protobuf.Empty
	5,  // 27: telepresence.daemon.Daemon.GetNetworkConfig:output_type -> telepresence.daemon.NetworkConfig
	12, // 28: telepresence.daemon.Daemon.SetDnsSearchPath:output_type -> google.protobuf.Empty
	12, // 29: telepresence.daemon.Daemon.SetDNSExcludes:output_type -> google.protobuf.Empty
	12, // 30: telepresence.daemon.Daemon.SetDNSMappings:output_type -> google.protobuf.Empty
	12, // 31: telepresence.daemon.Daemon.SetLogLevel:output_type -> google.protobuf.Empty
	12, // 32: telepresence.daemon.Daemon.WaitForNetwork:output_type -> google.protobuf.Empty
	22, // [22:33] is the sub-list for method output_type
	11, // [11:22] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_daemon_daemon_proto_init() }
func file_daemon_daemon_proto_init() {
	if File_daemon_daemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daemon_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DaemonStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paths); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutboundInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSExcludesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_daemon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSMappingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_daemon_daemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daemon_daemon_proto_goTypes,
		DependencyIndexes: file_daemon_daemon_proto_depIdxs,
		MessageInfos:      file_daemon_daemon_proto_msgTypes,
	}.Build()
	File_daemon_daemon_proto = out.File
	file_daemon_daemon_proto_rawDesc = nil
	file_daemon_daemon_proto_goTypes = nil
	file_daemon_daemon_proto_depIdxs = nil
}
