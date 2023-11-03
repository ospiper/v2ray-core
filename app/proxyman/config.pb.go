package proxyman

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
	internet "github.com/v2fly/v2ray-core/v5/transport/internet"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KnownProtocols int32

const (
	KnownProtocols_HTTP KnownProtocols = 0
	KnownProtocols_TLS  KnownProtocols = 1
)

// Enum value maps for KnownProtocols.
var (
	KnownProtocols_name = map[int32]string{
		0: "HTTP",
		1: "TLS",
	}
	KnownProtocols_value = map[string]int32{
		"HTTP": 0,
		"TLS":  1,
	}
)

func (x KnownProtocols) Enum() *KnownProtocols {
	p := new(KnownProtocols)
	*p = x
	return p
}

func (x KnownProtocols) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KnownProtocols) Descriptor() protoreflect.EnumDescriptor {
	return file_app_proxyman_config_proto_enumTypes[0].Descriptor()
}

func (KnownProtocols) Type() protoreflect.EnumType {
	return &file_app_proxyman_config_proto_enumTypes[0]
}

func (x KnownProtocols) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KnownProtocols.Descriptor instead.
func (KnownProtocols) EnumDescriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{0}
}

type AllocationStrategy_Type int32

const (
	// Always allocate all connection handlers.
	AllocationStrategy_Always AllocationStrategy_Type = 0
	// Randomly allocate specific range of handlers.
	AllocationStrategy_Random AllocationStrategy_Type = 1
	// External. Not supported yet.
	AllocationStrategy_External AllocationStrategy_Type = 2
)

// Enum value maps for AllocationStrategy_Type.
var (
	AllocationStrategy_Type_name = map[int32]string{
		0: "Always",
		1: "Random",
		2: "External",
	}
	AllocationStrategy_Type_value = map[string]int32{
		"Always":   0,
		"Random":   1,
		"External": 2,
	}
)

func (x AllocationStrategy_Type) Enum() *AllocationStrategy_Type {
	p := new(AllocationStrategy_Type)
	*p = x
	return p
}

func (x AllocationStrategy_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AllocationStrategy_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_app_proxyman_config_proto_enumTypes[1].Descriptor()
}

func (AllocationStrategy_Type) Type() protoreflect.EnumType {
	return &file_app_proxyman_config_proto_enumTypes[1]
}

func (x AllocationStrategy_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AllocationStrategy_Type.Descriptor instead.
func (AllocationStrategy_Type) EnumDescriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{1, 0}
}

type SenderConfig_DomainStrategy int32

const (
	SenderConfig_AS_IS   SenderConfig_DomainStrategy = 0
	SenderConfig_USE_IP  SenderConfig_DomainStrategy = 1
	SenderConfig_USE_IP4 SenderConfig_DomainStrategy = 2
	SenderConfig_USE_IP6 SenderConfig_DomainStrategy = 3
)

// Enum value maps for SenderConfig_DomainStrategy.
var (
	SenderConfig_DomainStrategy_name = map[int32]string{
		0: "AS_IS",
		1: "USE_IP",
		2: "USE_IP4",
		3: "USE_IP6",
	}
	SenderConfig_DomainStrategy_value = map[string]int32{
		"AS_IS":   0,
		"USE_IP":  1,
		"USE_IP4": 2,
		"USE_IP6": 3,
	}
)

func (x SenderConfig_DomainStrategy) Enum() *SenderConfig_DomainStrategy {
	p := new(SenderConfig_DomainStrategy)
	*p = x
	return p
}

func (x SenderConfig_DomainStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SenderConfig_DomainStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_app_proxyman_config_proto_enumTypes[2].Descriptor()
}

func (SenderConfig_DomainStrategy) Type() protoreflect.EnumType {
	return &file_app_proxyman_config_proto_enumTypes[2]
}

func (x SenderConfig_DomainStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SenderConfig_DomainStrategy.Descriptor instead.
func (SenderConfig_DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{6, 0}
}

type InboundConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InboundConfig) Reset() {
	*x = InboundConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboundConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboundConfig) ProtoMessage() {}

func (x *InboundConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboundConfig.ProtoReflect.Descriptor instead.
func (*InboundConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{0}
}

type AllocationStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type AllocationStrategy_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.proxyman.AllocationStrategy_Type" json:"type,omitempty"`
	// Number of handlers (ports) running in parallel.
	// Default value is 3 if unset.
	Concurrency *AllocationStrategy_AllocationStrategyConcurrency `protobuf:"bytes,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// Number of minutes before a handler is regenerated.
	// Default value is 5 if unset.
	Refresh *AllocationStrategy_AllocationStrategyRefresh `protobuf:"bytes,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *AllocationStrategy) Reset() {
	*x = AllocationStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationStrategy) ProtoMessage() {}

func (x *AllocationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationStrategy.ProtoReflect.Descriptor instead.
func (*AllocationStrategy) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{1}
}

func (x *AllocationStrategy) GetType() AllocationStrategy_Type {
	if x != nil {
		return x.Type
	}
	return AllocationStrategy_Always
}

func (x *AllocationStrategy) GetConcurrency() *AllocationStrategy_AllocationStrategyConcurrency {
	if x != nil {
		return x.Concurrency
	}
	return nil
}

func (x *AllocationStrategy) GetRefresh() *AllocationStrategy_AllocationStrategyRefresh {
	if x != nil {
		return x.Refresh
	}
	return nil
}

type SniffingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to enable content sniffing on an inbound connection.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Override target destination if sniff'ed protocol is in the given list.
	// Supported values are "http", "tls", "fakedns".
	DestinationOverride []string `protobuf:"bytes,2,rep,name=destination_override,json=destinationOverride,proto3" json:"destination_override,omitempty"`
	// Whether should only try to sniff metadata without waiting for client input.
	// Can be used to support SMTP like protocol where server send the first message.
	MetadataOnly bool `protobuf:"varint,3,opt,name=metadata_only,json=metadataOnly,proto3" json:"metadata_only,omitempty"`
}

func (x *SniffingConfig) Reset() {
	*x = SniffingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SniffingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SniffingConfig) ProtoMessage() {}

func (x *SniffingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SniffingConfig.ProtoReflect.Descriptor instead.
func (*SniffingConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{2}
}

func (x *SniffingConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SniffingConfig) GetDestinationOverride() []string {
	if x != nil {
		return x.DestinationOverride
	}
	return nil
}

func (x *SniffingConfig) GetMetadataOnly() bool {
	if x != nil {
		return x.MetadataOnly
	}
	return false
}

type ReceiverConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PortRange specifies the ports which the Receiver should listen on.
	PortRange *net.PortRange `protobuf:"bytes,1,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"`
	// Listen specifies the IP address that the Receiver should listen on.
	Listen                     *net.IPOrDomain        `protobuf:"bytes,2,opt,name=listen,proto3" json:"listen,omitempty"`
	AllocationStrategy         *AllocationStrategy    `protobuf:"bytes,3,opt,name=allocation_strategy,json=allocationStrategy,proto3" json:"allocation_strategy,omitempty"`
	StreamSettings             *internet.StreamConfig `protobuf:"bytes,4,opt,name=stream_settings,json=streamSettings,proto3" json:"stream_settings,omitempty"`
	ReceiveOriginalDestination bool                   `protobuf:"varint,5,opt,name=receive_original_destination,json=receiveOriginalDestination,proto3" json:"receive_original_destination,omitempty"`
	// Override domains for the given protocol.
	// Deprecated. Use sniffing_settings.
	//
	// Deprecated: Marked as deprecated in app/proxyman/config.proto.
	DomainOverride   []KnownProtocols `protobuf:"varint,7,rep,packed,name=domain_override,json=domainOverride,proto3,enum=v2ray.core.app.proxyman.KnownProtocols" json:"domain_override,omitempty"`
	SniffingSettings *SniffingConfig  `protobuf:"bytes,8,opt,name=sniffing_settings,json=sniffingSettings,proto3" json:"sniffing_settings,omitempty"`
}

func (x *ReceiverConfig) Reset() {
	*x = ReceiverConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiverConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiverConfig) ProtoMessage() {}

func (x *ReceiverConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiverConfig.ProtoReflect.Descriptor instead.
func (*ReceiverConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiverConfig) GetPortRange() *net.PortRange {
	if x != nil {
		return x.PortRange
	}
	return nil
}

func (x *ReceiverConfig) GetListen() *net.IPOrDomain {
	if x != nil {
		return x.Listen
	}
	return nil
}

func (x *ReceiverConfig) GetAllocationStrategy() *AllocationStrategy {
	if x != nil {
		return x.AllocationStrategy
	}
	return nil
}

func (x *ReceiverConfig) GetStreamSettings() *internet.StreamConfig {
	if x != nil {
		return x.StreamSettings
	}
	return nil
}

func (x *ReceiverConfig) GetReceiveOriginalDestination() bool {
	if x != nil {
		return x.ReceiveOriginalDestination
	}
	return false
}

// Deprecated: Marked as deprecated in app/proxyman/config.proto.
func (x *ReceiverConfig) GetDomainOverride() []KnownProtocols {
	if x != nil {
		return x.DomainOverride
	}
	return nil
}

func (x *ReceiverConfig) GetSniffingSettings() *SniffingConfig {
	if x != nil {
		return x.SniffingSettings
	}
	return nil
}

type InboundHandlerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag              string     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	ReceiverSettings *anypb.Any `protobuf:"bytes,2,opt,name=receiver_settings,json=receiverSettings,proto3" json:"receiver_settings,omitempty"`
	ProxySettings    *anypb.Any `protobuf:"bytes,3,opt,name=proxy_settings,json=proxySettings,proto3" json:"proxy_settings,omitempty"`
}

func (x *InboundHandlerConfig) Reset() {
	*x = InboundHandlerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboundHandlerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboundHandlerConfig) ProtoMessage() {}

func (x *InboundHandlerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboundHandlerConfig.ProtoReflect.Descriptor instead.
func (*InboundHandlerConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{4}
}

func (x *InboundHandlerConfig) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *InboundHandlerConfig) GetReceiverSettings() *anypb.Any {
	if x != nil {
		return x.ReceiverSettings
	}
	return nil
}

func (x *InboundHandlerConfig) GetProxySettings() *anypb.Any {
	if x != nil {
		return x.ProxySettings
	}
	return nil
}

type OutboundConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OutboundConfig) Reset() {
	*x = OutboundConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutboundConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutboundConfig) ProtoMessage() {}

func (x *OutboundConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutboundConfig.ProtoReflect.Descriptor instead.
func (*OutboundConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{5}
}

type SenderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Send traffic through the given IP. Only IP is allowed.
	Via               *net.IPOrDomain             `protobuf:"bytes,1,opt,name=via,proto3" json:"via,omitempty"`
	StreamSettings    *internet.StreamConfig      `protobuf:"bytes,2,opt,name=stream_settings,json=streamSettings,proto3" json:"stream_settings,omitempty"`
	ProxySettings     *internet.ProxyConfig       `protobuf:"bytes,3,opt,name=proxy_settings,json=proxySettings,proto3" json:"proxy_settings,omitempty"`
	MultiplexSettings *MultiplexingConfig         `protobuf:"bytes,4,opt,name=multiplex_settings,json=multiplexSettings,proto3" json:"multiplex_settings,omitempty"`
	DomainStrategy    SenderConfig_DomainStrategy `protobuf:"varint,5,opt,name=domain_strategy,json=domainStrategy,proto3,enum=v2ray.core.app.proxyman.SenderConfig_DomainStrategy" json:"domain_strategy,omitempty"`
}

func (x *SenderConfig) Reset() {
	*x = SenderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderConfig) ProtoMessage() {}

func (x *SenderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderConfig.ProtoReflect.Descriptor instead.
func (*SenderConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{6}
}

func (x *SenderConfig) GetVia() *net.IPOrDomain {
	if x != nil {
		return x.Via
	}
	return nil
}

func (x *SenderConfig) GetStreamSettings() *internet.StreamConfig {
	if x != nil {
		return x.StreamSettings
	}
	return nil
}

func (x *SenderConfig) GetProxySettings() *internet.ProxyConfig {
	if x != nil {
		return x.ProxySettings
	}
	return nil
}

func (x *SenderConfig) GetMultiplexSettings() *MultiplexingConfig {
	if x != nil {
		return x.MultiplexSettings
	}
	return nil
}

func (x *SenderConfig) GetDomainStrategy() SenderConfig_DomainStrategy {
	if x != nil {
		return x.DomainStrategy
	}
	return SenderConfig_AS_IS
}

type MultiplexingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not Mux is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Max number of concurrent connections that one Mux connection can handle.
	Concurrency uint32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
}

func (x *MultiplexingConfig) Reset() {
	*x = MultiplexingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplexingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplexingConfig) ProtoMessage() {}

func (x *MultiplexingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplexingConfig.ProtoReflect.Descriptor instead.
func (*MultiplexingConfig) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{7}
}

func (x *MultiplexingConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *MultiplexingConfig) GetConcurrency() uint32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

type AllocationStrategy_AllocationStrategyConcurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AllocationStrategy_AllocationStrategyConcurrency) Reset() {
	*x = AllocationStrategy_AllocationStrategyConcurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationStrategy_AllocationStrategyConcurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationStrategy_AllocationStrategyConcurrency) ProtoMessage() {}

func (x *AllocationStrategy_AllocationStrategyConcurrency) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationStrategy_AllocationStrategyConcurrency.ProtoReflect.Descriptor instead.
func (*AllocationStrategy_AllocationStrategyConcurrency) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AllocationStrategy_AllocationStrategyConcurrency) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AllocationStrategy_AllocationStrategyRefresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AllocationStrategy_AllocationStrategyRefresh) Reset() {
	*x = AllocationStrategy_AllocationStrategyRefresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationStrategy_AllocationStrategyRefresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationStrategy_AllocationStrategyRefresh) ProtoMessage() {}

func (x *AllocationStrategy_AllocationStrategyRefresh) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationStrategy_AllocationStrategyRefresh.ProtoReflect.Descriptor instead.
func (*AllocationStrategy_AllocationStrategyRefresh) Descriptor() ([]byte, []int) {
	return file_app_proxyman_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AllocationStrategy_AllocationStrategyRefresh) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_app_proxyman_config_proto protoreflect.FileDescriptor

var file_app_proxyman_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x6d, 0x61, 0x6e, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xc0, 0x03, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61,
	0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x6b, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x5f, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x1a, 0x35, 0x0a,
	0x1d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x31, 0x0a, 0x19, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x10, 0x02, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x53, 0x6e, 0x69, 0x66, 0x66, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xb4, 0x04, 0x0a, 0x0e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a,
	0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x5c, 0x0a, 0x13, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x54, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x40, 0x0a,
	0x1c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x54, 0x0a, 0x0f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d,
	0x61, 0x6e, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x53, 0x6e, 0x69, 0x66, 0x66,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x73, 0x6e, 0x69, 0x66, 0x66,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4a, 0x04, 0x08, 0x06, 0x10,
	0x07, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x41, 0x0a, 0x11,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x10, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x3b, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x10, 0x0a, 0x0e,
	0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xea,
	0x03, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x33, 0x0a, 0x03, 0x76, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x03, 0x76, 0x69, 0x61, 0x12, 0x54, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5a, 0x0a,
	0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6d, 0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x78, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5d, 0x0a, 0x0f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x41, 0x0a, 0x0e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53,
	0x5f, 0x49, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x03, 0x22, 0x50, 0x0a, 0x12, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2a, 0x23, 0x0a,
	0x0e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4c, 0x53,
	0x10, 0x01, 0x42, 0x66, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61,
	0x6e, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e,
	0xaa, 0x02, 0x17, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70,
	0x70, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_proxyman_config_proto_rawDescOnce sync.Once
	file_app_proxyman_config_proto_rawDescData = file_app_proxyman_config_proto_rawDesc
)

func file_app_proxyman_config_proto_rawDescGZIP() []byte {
	file_app_proxyman_config_proto_rawDescOnce.Do(func() {
		file_app_proxyman_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proxyman_config_proto_rawDescData)
	})
	return file_app_proxyman_config_proto_rawDescData
}

var file_app_proxyman_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_app_proxyman_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_app_proxyman_config_proto_goTypes = []interface{}{
	(KnownProtocols)(0),                                      // 0: v2ray.core.app.proxyman.KnownProtocols
	(AllocationStrategy_Type)(0),                             // 1: v2ray.core.app.proxyman.AllocationStrategy.Type
	(SenderConfig_DomainStrategy)(0),                         // 2: v2ray.core.app.proxyman.SenderConfig.DomainStrategy
	(*InboundConfig)(nil),                                    // 3: v2ray.core.app.proxyman.InboundConfig
	(*AllocationStrategy)(nil),                               // 4: v2ray.core.app.proxyman.AllocationStrategy
	(*SniffingConfig)(nil),                                   // 5: v2ray.core.app.proxyman.SniffingConfig
	(*ReceiverConfig)(nil),                                   // 6: v2ray.core.app.proxyman.ReceiverConfig
	(*InboundHandlerConfig)(nil),                             // 7: v2ray.core.app.proxyman.InboundHandlerConfig
	(*OutboundConfig)(nil),                                   // 8: v2ray.core.app.proxyman.OutboundConfig
	(*SenderConfig)(nil),                                     // 9: v2ray.core.app.proxyman.SenderConfig
	(*MultiplexingConfig)(nil),                               // 10: v2ray.core.app.proxyman.MultiplexingConfig
	(*AllocationStrategy_AllocationStrategyConcurrency)(nil), // 11: v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyConcurrency
	(*AllocationStrategy_AllocationStrategyRefresh)(nil),     // 12: v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyRefresh
	(*net.PortRange)(nil),                                    // 13: v2ray.core.common.net.PortRange
	(*net.IPOrDomain)(nil),                                   // 14: v2ray.core.common.net.IPOrDomain
	(*internet.StreamConfig)(nil),                            // 15: v2ray.core.transport.internet.StreamConfig
	(*anypb.Any)(nil),                                        // 16: google.protobuf.Any
	(*internet.ProxyConfig)(nil),                             // 17: v2ray.core.transport.internet.ProxyConfig
}
var file_app_proxyman_config_proto_depIdxs = []int32{
	1,  // 0: v2ray.core.app.proxyman.AllocationStrategy.type:type_name -> v2ray.core.app.proxyman.AllocationStrategy.Type
	11, // 1: v2ray.core.app.proxyman.AllocationStrategy.concurrency:type_name -> v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyConcurrency
	12, // 2: v2ray.core.app.proxyman.AllocationStrategy.refresh:type_name -> v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyRefresh
	13, // 3: v2ray.core.app.proxyman.ReceiverConfig.port_range:type_name -> v2ray.core.common.net.PortRange
	14, // 4: v2ray.core.app.proxyman.ReceiverConfig.listen:type_name -> v2ray.core.common.net.IPOrDomain
	4,  // 5: v2ray.core.app.proxyman.ReceiverConfig.allocation_strategy:type_name -> v2ray.core.app.proxyman.AllocationStrategy
	15, // 6: v2ray.core.app.proxyman.ReceiverConfig.stream_settings:type_name -> v2ray.core.transport.internet.StreamConfig
	0,  // 7: v2ray.core.app.proxyman.ReceiverConfig.domain_override:type_name -> v2ray.core.app.proxyman.KnownProtocols
	5,  // 8: v2ray.core.app.proxyman.ReceiverConfig.sniffing_settings:type_name -> v2ray.core.app.proxyman.SniffingConfig
	16, // 9: v2ray.core.app.proxyman.InboundHandlerConfig.receiver_settings:type_name -> google.protobuf.Any
	16, // 10: v2ray.core.app.proxyman.InboundHandlerConfig.proxy_settings:type_name -> google.protobuf.Any
	14, // 11: v2ray.core.app.proxyman.SenderConfig.via:type_name -> v2ray.core.common.net.IPOrDomain
	15, // 12: v2ray.core.app.proxyman.SenderConfig.stream_settings:type_name -> v2ray.core.transport.internet.StreamConfig
	17, // 13: v2ray.core.app.proxyman.SenderConfig.proxy_settings:type_name -> v2ray.core.transport.internet.ProxyConfig
	10, // 14: v2ray.core.app.proxyman.SenderConfig.multiplex_settings:type_name -> v2ray.core.app.proxyman.MultiplexingConfig
	2,  // 15: v2ray.core.app.proxyman.SenderConfig.domain_strategy:type_name -> v2ray.core.app.proxyman.SenderConfig.DomainStrategy
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_app_proxyman_config_proto_init() }
func file_app_proxyman_config_proto_init() {
	if File_app_proxyman_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proxyman_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InboundConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationStrategy); i {
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
		file_app_proxyman_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SniffingConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiverConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InboundHandlerConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutboundConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplexingConfig); i {
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
		file_app_proxyman_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationStrategy_AllocationStrategyConcurrency); i {
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
		file_app_proxyman_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationStrategy_AllocationStrategyRefresh); i {
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
			RawDescriptor: file_app_proxyman_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_proxyman_config_proto_goTypes,
		DependencyIndexes: file_app_proxyman_config_proto_depIdxs,
		EnumInfos:         file_app_proxyman_config_proto_enumTypes,
		MessageInfos:      file_app_proxyman_config_proto_msgTypes,
	}.Build()
	File_app_proxyman_config_proto = out.File
	file_app_proxyman_config_proto_rawDesc = nil
	file_app_proxyman_config_proto_goTypes = nil
	file_app_proxyman_config_proto_depIdxs = nil
}
