// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Operation int32

const (
	Operation_GET    Operation = 0
	Operation_POST   Operation = 1
	Operation_DELETE Operation = 2
	Operation_UPDATE Operation = 3
)

var Operation_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "DELETE",
	3: "UPDATE",
}

var Operation_value = map[string]int32{
	"GET":    0,
	"POST":   1,
	"DELETE": 2,
	"UPDATE": 3,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{0}
}

type ProvisionState int32

const (
	ProvisionState_UNKNOWN            ProvisionState = 0
	ProvisionState_CREATING           ProvisionState = 1
	ProvisionState_CREATED            ProvisionState = 2
	ProvisionState_CREATE_FAILED      ProvisionState = 3
	ProvisionState_DELETING           ProvisionState = 4
	ProvisionState_DELETE_FAILED      ProvisionState = 5
	ProvisionState_DELETED            ProvisionState = 6
	ProvisionState_UPDATING           ProvisionState = 7
	ProvisionState_UPDATE_FAILED      ProvisionState = 8
	ProvisionState_UPDATED            ProvisionState = 9
	ProvisionState_PROVISIONING       ProvisionState = 10
	ProvisionState_PROVISIONED        ProvisionState = 11
	ProvisionState_PROVISION_FAILED   ProvisionState = 12
	ProvisionState_DEPROVISIONING     ProvisionState = 13
	ProvisionState_DEPROVISIONED      ProvisionState = 14
	ProvisionState_DEPROVISION_FAILED ProvisionState = 15
	ProvisionState_DELETE_PENDING     ProvisionState = 16
)

var ProvisionState_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CREATING",
	2:  "CREATED",
	3:  "CREATE_FAILED",
	4:  "DELETING",
	5:  "DELETE_FAILED",
	6:  "DELETED",
	7:  "UPDATING",
	8:  "UPDATE_FAILED",
	9:  "UPDATED",
	10: "PROVISIONING",
	11: "PROVISIONED",
	12: "PROVISION_FAILED",
	13: "DEPROVISIONING",
	14: "DEPROVISIONED",
	15: "DEPROVISION_FAILED",
	16: "DELETE_PENDING",
}

var ProvisionState_value = map[string]int32{
	"UNKNOWN":            0,
	"CREATING":           1,
	"CREATED":            2,
	"CREATE_FAILED":      3,
	"DELETING":           4,
	"DELETE_FAILED":      5,
	"DELETED":            6,
	"UPDATING":           7,
	"UPDATE_FAILED":      8,
	"UPDATED":            9,
	"PROVISIONING":       10,
	"PROVISIONED":        11,
	"PROVISION_FAILED":   12,
	"DEPROVISIONING":     13,
	"DEPROVISIONED":      14,
	"DEPROVISION_FAILED": 15,
	"DELETE_PENDING":     16,
}

func (x ProvisionState) String() string {
	return proto.EnumName(ProvisionState_name, int32(x))
}

func (ProvisionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{1}
}

type HighAvailabilityState int32

const (
	HighAvailabilityState_UNKNOWN_HA_STATE HighAvailabilityState = 0
	HighAvailabilityState_STABLE           HighAvailabilityState = 1
	HighAvailabilityState_PENDING          HighAvailabilityState = 2
)

var HighAvailabilityState_name = map[int32]string{
	0: "UNKNOWN_HA_STATE",
	1: "STABLE",
	2: "PENDING",
}

var HighAvailabilityState_value = map[string]int32{
	"UNKNOWN_HA_STATE": 0,
	"STABLE":           1,
	"PENDING":          2,
}

func (x HighAvailabilityState) String() string {
	return proto.EnumName(HighAvailabilityState_name, int32(x))
}

func (HighAvailabilityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{2}
}

type HealthState int32

const (
	HealthState_NOTKNOWN HealthState = 0
	HealthState_OK       HealthState = 1
	HealthState_WARNING  HealthState = 2
	HealthState_CRITICAL HealthState = 3
	// The entity went missing from the platform
	HealthState_MISSING  HealthState = 4
	HealthState_DEGRADED HealthState = 5
	// The entity went missing from the agent
	HealthState_NOTFOUND HealthState = 6
)

var HealthState_name = map[int32]string{
	0: "NOTKNOWN",
	1: "OK",
	2: "WARNING",
	3: "CRITICAL",
	4: "MISSING",
	5: "DEGRADED",
	6: "NOTFOUND",
}

var HealthState_value = map[string]int32{
	"NOTKNOWN": 0,
	"OK":       1,
	"WARNING":  2,
	"CRITICAL": 3,
	"MISSING":  4,
	"DEGRADED": 5,
	"NOTFOUND": 6,
}

func (x HealthState) String() string {
	return proto.EnumName(HealthState_name, int32(x))
}

func (HealthState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{3}
}

type ClientType int32

const (
	ClientType_CONTROLPLANE   ClientType = 0
	ClientType_EXTERNALCLIENT ClientType = 1
	ClientType_NODE           ClientType = 2
	ClientType_ADMIN          ClientType = 3
	ClientType_BAREMETAL      ClientType = 4
	ClientType_LOADBALANCER   ClientType = 5
)

var ClientType_name = map[int32]string{
	0: "CONTROLPLANE",
	1: "EXTERNALCLIENT",
	2: "NODE",
	3: "ADMIN",
	4: "BAREMETAL",
	5: "LOADBALANCER",
}

var ClientType_value = map[string]int32{
	"CONTROLPLANE":   0,
	"EXTERNALCLIENT": 1,
	"NODE":           2,
	"ADMIN":          3,
	"BAREMETAL":      4,
	"LOADBALANCER":   5,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{4}
}

type AuthenticationType int32

const (
	AuthenticationType_SELFSIGNED AuthenticationType = 0
	AuthenticationType_CASIGNED   AuthenticationType = 1
)

var AuthenticationType_name = map[int32]string{
	0: "SELFSIGNED",
	1: "CASIGNED",
}

var AuthenticationType_value = map[string]int32{
	"SELFSIGNED": 0,
	"CASIGNED":   1,
}

func (x AuthenticationType) String() string {
	return proto.EnumName(AuthenticationType_name, int32(x))
}

func (AuthenticationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{5}
}

type ProviderType int32

const (
	ProviderType_AnyProvider            ProviderType = 0
	ProviderType_VirtualMachine         ProviderType = 1
	ProviderType_VirtualMachineScaleSet ProviderType = 2
	ProviderType_LoadBalancer           ProviderType = 3
	ProviderType_VirtualNetwork         ProviderType = 4
	ProviderType_VirtualHardDisk        ProviderType = 5
	ProviderType_GalleryImage           ProviderType = 6
	ProviderType_VirtualMachineImage    ProviderType = 7
	ProviderType_NetworkInterface       ProviderType = 8
	ProviderType_Certificate            ProviderType = 9
	ProviderType_Key                    ProviderType = 10
	ProviderType_Secret                 ProviderType = 11
	ProviderType_KeyVault               ProviderType = 12
	ProviderType_Identity               ProviderType = 13
	ProviderType_Role                   ProviderType = 14
	ProviderType_RoleAssignment         ProviderType = 15
	ProviderType_Kubernetes             ProviderType = 16
	ProviderType_Cluster                ProviderType = 17
	ProviderType_ControlPlane           ProviderType = 18
	ProviderType_Group                  ProviderType = 19
	ProviderType_Node                   ProviderType = 20
	ProviderType_Location               ProviderType = 21
	ProviderType_StorageContainer       ProviderType = 22
	ProviderType_StorageFile            ProviderType = 23
	ProviderType_StorageDirectory       ProviderType = 24
	ProviderType_Subscription           ProviderType = 25
	ProviderType_VipPool                ProviderType = 26
	ProviderType_MacPool                ProviderType = 27
	ProviderType_EtcdCluster            ProviderType = 28
	ProviderType_EtcdServer             ProviderType = 29
	ProviderType_BareMetalMachine       ProviderType = 30
	ProviderType_CredentialMonitor      ProviderType = 31
	ProviderType_Logging                ProviderType = 32
	ProviderType_Recovery               ProviderType = 33
	ProviderType_Debug                  ProviderType = 34
	ProviderType_BareMetalHost          ProviderType = 35
	ProviderType_Authentication         ProviderType = 36
)

var ProviderType_name = map[int32]string{
	0:  "AnyProvider",
	1:  "VirtualMachine",
	2:  "VirtualMachineScaleSet",
	3:  "LoadBalancer",
	4:  "VirtualNetwork",
	5:  "VirtualHardDisk",
	6:  "GalleryImage",
	7:  "VirtualMachineImage",
	8:  "NetworkInterface",
	9:  "Certificate",
	10: "Key",
	11: "Secret",
	12: "KeyVault",
	13: "Identity",
	14: "Role",
	15: "RoleAssignment",
	16: "Kubernetes",
	17: "Cluster",
	18: "ControlPlane",
	19: "Group",
	20: "Node",
	21: "Location",
	22: "StorageContainer",
	23: "StorageFile",
	24: "StorageDirectory",
	25: "Subscription",
	26: "VipPool",
	27: "MacPool",
	28: "EtcdCluster",
	29: "EtcdServer",
	30: "BareMetalMachine",
	31: "CredentialMonitor",
	32: "Logging",
	33: "Recovery",
	34: "Debug",
	35: "BareMetalHost",
	36: "Authentication",
}

var ProviderType_value = map[string]int32{
	"AnyProvider":            0,
	"VirtualMachine":         1,
	"VirtualMachineScaleSet": 2,
	"LoadBalancer":           3,
	"VirtualNetwork":         4,
	"VirtualHardDisk":        5,
	"GalleryImage":           6,
	"VirtualMachineImage":    7,
	"NetworkInterface":       8,
	"Certificate":            9,
	"Key":                    10,
	"Secret":                 11,
	"KeyVault":               12,
	"Identity":               13,
	"Role":                   14,
	"RoleAssignment":         15,
	"Kubernetes":             16,
	"Cluster":                17,
	"ControlPlane":           18,
	"Group":                  19,
	"Node":                   20,
	"Location":               21,
	"StorageContainer":       22,
	"StorageFile":            23,
	"StorageDirectory":       24,
	"Subscription":           25,
	"VipPool":                26,
	"MacPool":                27,
	"EtcdCluster":            28,
	"EtcdServer":             29,
	"BareMetalMachine":       30,
	"CredentialMonitor":      31,
	"Logging":                32,
	"Recovery":               33,
	"Debug":                  34,
	"BareMetalHost":          35,
	"Authentication":         36,
}

func (x ProviderType) String() string {
	return proto.EnumName(ProviderType_name, int32(x))
}

func (ProviderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{6}
}

type ImageSource int32

const (
	ImageSource_LOCAL_SOURCE ImageSource = 0
	ImageSource_SFS_SOURCE   ImageSource = 1
	ImageSource_HTTP_SOURCE  ImageSource = 2
	ImageSource_CLONE_SOURCE ImageSource = 3
)

var ImageSource_name = map[int32]string{
	0: "LOCAL_SOURCE",
	1: "SFS_SOURCE",
	2: "HTTP_SOURCE",
	3: "CLONE_SOURCE",
}

var ImageSource_value = map[string]int32{
	"LOCAL_SOURCE": 0,
	"SFS_SOURCE":   1,
	"HTTP_SOURCE":  2,
	"CLONE_SOURCE": 3,
}

func (x ImageSource) String() string {
	return proto.EnumName(ImageSource_name, int32(x))
}

func (ImageSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{7}
}

type CloudInitDataSource int32

const (
	CloudInitDataSource_NoCloud CloudInitDataSource = 0
	CloudInitDataSource_Azure   CloudInitDataSource = 1
)

var CloudInitDataSource_name = map[int32]string{
	0: "NoCloud",
	1: "Azure",
}

var CloudInitDataSource_value = map[string]int32{
	"NoCloud": 0,
	"Azure":   1,
}

func (x CloudInitDataSource) String() string {
	return proto.EnumName(CloudInitDataSource_name, int32(x))
}

func (CloudInitDataSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{8}
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Parameters           string   `protobuf:"bytes,3,opt,name=Parameters,proto3" json:"Parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetParameters() string {
	if m != nil {
		return m.Parameters
	}
	return ""
}

type ProvisionStatus struct {
	CurrentState         ProvisionState `protobuf:"varint,1,opt,name=currentState,proto3,enum=moc.ProvisionState" json:"currentState,omitempty"`
	PreviousState        ProvisionState `protobuf:"varint,2,opt,name=previousState,proto3,enum=moc.ProvisionState" json:"previousState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProvisionStatus) Reset()         { *m = ProvisionStatus{} }
func (m *ProvisionStatus) String() string { return proto.CompactTextString(m) }
func (*ProvisionStatus) ProtoMessage()    {}
func (*ProvisionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{1}
}

func (m *ProvisionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvisionStatus.Unmarshal(m, b)
}
func (m *ProvisionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvisionStatus.Marshal(b, m, deterministic)
}
func (m *ProvisionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionStatus.Merge(m, src)
}
func (m *ProvisionStatus) XXX_Size() int {
	return xxx_messageInfo_ProvisionStatus.Size(m)
}
func (m *ProvisionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionStatus proto.InternalMessageInfo

func (m *ProvisionStatus) GetCurrentState() ProvisionState {
	if m != nil {
		return m.CurrentState
	}
	return ProvisionState_UNKNOWN
}

func (m *ProvisionStatus) GetPreviousState() ProvisionState {
	if m != nil {
		return m.PreviousState
	}
	return ProvisionState_UNKNOWN
}

type DownloadStatus struct {
	Progress             int64    `protobuf:"varint,1,opt,name=progress,proto3" json:"progress,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	ErrorDetails         *Error   `protobuf:"bytes,3,opt,name=errorDetails,proto3" json:"errorDetails,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadStatus) Reset()         { *m = DownloadStatus{} }
func (m *DownloadStatus) String() string { return proto.CompactTextString(m) }
func (*DownloadStatus) ProtoMessage()    {}
func (*DownloadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{2}
}

func (m *DownloadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadStatus.Unmarshal(m, b)
}
func (m *DownloadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadStatus.Marshal(b, m, deterministic)
}
func (m *DownloadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadStatus.Merge(m, src)
}
func (m *DownloadStatus) XXX_Size() int {
	return xxx_messageInfo_DownloadStatus.Size(m)
}
func (m *DownloadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadStatus proto.InternalMessageInfo

func (m *DownloadStatus) GetProgress() int64 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *DownloadStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *DownloadStatus) GetErrorDetails() *Error {
	if m != nil {
		return m.ErrorDetails
	}
	return nil
}

type Health struct {
	CurrentState         HealthState `protobuf:"varint,1,opt,name=currentState,proto3,enum=moc.HealthState" json:"currentState,omitempty"`
	PreviousState        HealthState `protobuf:"varint,2,opt,name=previousState,proto3,enum=moc.HealthState" json:"previousState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{3}
}

func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (m *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(m, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetCurrentState() HealthState {
	if m != nil {
		return m.CurrentState
	}
	return HealthState_NOTKNOWN
}

func (m *Health) GetPreviousState() HealthState {
	if m != nil {
		return m.PreviousState
	}
	return HealthState_NOTKNOWN
}

type Version struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{4}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Status struct {
	Health               *Health          `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	ProvisioningStatus   *ProvisionStatus `protobuf:"bytes,2,opt,name=provisioningStatus,proto3" json:"provisioningStatus,omitempty"`
	DownloadStatus       *DownloadStatus  `protobuf:"bytes,3,opt,name=downloadStatus,proto3" json:"downloadStatus,omitempty"`
	LastError            *Error           `protobuf:"bytes,4,opt,name=lastError,proto3" json:"lastError,omitempty"`
	Version              *Version         `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{5}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetHealth() *Health {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *Status) GetProvisioningStatus() *ProvisionStatus {
	if m != nil {
		return m.ProvisioningStatus
	}
	return nil
}

func (m *Status) GetDownloadStatus() *DownloadStatus {
	if m != nil {
		return m.DownloadStatus
	}
	return nil
}

func (m *Status) GetLastError() *Error {
	if m != nil {
		return m.LastError
	}
	return nil
}

func (m *Status) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Entity struct {
	IsPlaceholder        bool     `protobuf:"varint,1,opt,name=IsPlaceholder,proto3" json:"IsPlaceholder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{6}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetIsPlaceholder() bool {
	if m != nil {
		return m.IsPlaceholder
	}
	return false
}

type Tag struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{7}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Tags struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tags) Reset()         { *m = Tags{} }
func (m *Tags) String() string { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()    {}
func (*Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{8}
}

func (m *Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tags.Unmarshal(m, b)
}
func (m *Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tags.Marshal(b, m, deterministic)
}
func (m *Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tags.Merge(m, src)
}
func (m *Tags) XXX_Size() int {
	return xxx_messageInfo_Tags.Size(m)
}
func (m *Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_Tags proto.InternalMessageInfo

func (m *Tags) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

var E_Sensitive = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50001,
	Name:          "moc.sensitive",
	Tag:           "varint,50001,opt,name=sensitive",
	Filename:      "moc_common_common.proto",
}

func init() {
	proto.RegisterEnum("moc.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("moc.ProvisionState", ProvisionState_name, ProvisionState_value)
	proto.RegisterEnum("moc.HighAvailabilityState", HighAvailabilityState_name, HighAvailabilityState_value)
	proto.RegisterEnum("moc.HealthState", HealthState_name, HealthState_value)
	proto.RegisterEnum("moc.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("moc.AuthenticationType", AuthenticationType_name, AuthenticationType_value)
	proto.RegisterEnum("moc.ProviderType", ProviderType_name, ProviderType_value)
	proto.RegisterEnum("moc.ImageSource", ImageSource_name, ImageSource_value)
	proto.RegisterEnum("moc.CloudInitDataSource", CloudInitDataSource_name, CloudInitDataSource_value)
	proto.RegisterType((*Error)(nil), "moc.Error")
	proto.RegisterType((*ProvisionStatus)(nil), "moc.ProvisionStatus")
	proto.RegisterType((*DownloadStatus)(nil), "moc.DownloadStatus")
	proto.RegisterType((*Health)(nil), "moc.Health")
	proto.RegisterType((*Version)(nil), "moc.Version")
	proto.RegisterType((*Status)(nil), "moc.Status")
	proto.RegisterType((*Entity)(nil), "moc.Entity")
	proto.RegisterType((*Tag)(nil), "moc.Tag")
	proto.RegisterType((*Tags)(nil), "moc.Tags")
	proto.RegisterExtension(E_Sensitive)
}

func init() { proto.RegisterFile("moc_common_common.proto", fileDescriptor_681f78e46755eb93) }

var fileDescriptor_681f78e46755eb93 = []byte{
	// 1379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0x6d, 0x4f, 0xe3, 0xca,
	0x15, 0x26, 0xef, 0xc9, 0x49, 0x08, 0xb3, 0x03, 0xcb, 0xa6, 0xf4, 0xde, 0x5b, 0x6e, 0x76, 0x5b,
	0xa1, 0x48, 0x37, 0x48, 0xb4, 0xba, 0x55, 0x5b, 0xf5, 0x83, 0x89, 0x0d, 0x44, 0x18, 0x27, 0xb2,
	0x0d, 0x5b, 0xf5, 0x0b, 0x9a, 0x38, 0x07, 0xc7, 0x5a, 0xc7, 0x13, 0x8d, 0xc7, 0x59, 0xa5, 0xdf,
	0x5a, 0xa9, 0xfd, 0x43, 0xfd, 0x25, 0xfd, 0x47, 0xd5, 0x8c, 0x1d, 0x96, 0xec, 0x72, 0x3f, 0xe1,
	0xf3, 0xcc, 0x73, 0xde, 0x9e, 0x73, 0x86, 0x0c, 0xbc, 0x5b, 0xf2, 0xe0, 0x31, 0xe0, 0xcb, 0x25,
	0x4f, 0x8a, 0x3f, 0xc3, 0x95, 0xe0, 0x92, 0xd3, 0xca, 0x92, 0x07, 0x27, 0xa7, 0x21, 0xe7, 0x61,
	0x8c, 0xe7, 0x1a, 0x9a, 0x65, 0x4f, 0xe7, 0x73, 0x4c, 0x03, 0x11, 0xad, 0x24, 0x17, 0x39, 0xad,
	0x7f, 0x0f, 0x35, 0x4b, 0x08, 0x2e, 0x68, 0x0f, 0x1a, 0x77, 0x98, 0xa6, 0x2c, 0xc4, 0x5e, 0xe9,
	0xb4, 0x74, 0xd6, 0x72, 0xb7, 0x26, 0xa5, 0x50, 0x1d, 0xf1, 0x39, 0xf6, 0xca, 0xa7, 0xa5, 0xb3,
	0x9a, 0xab, 0xbf, 0xe9, 0x0f, 0x00, 0x53, 0x26, 0xd8, 0x12, 0x25, 0x8a, 0xb4, 0x57, 0xd1, 0x0e,
	0x2f, 0x90, 0xfe, 0xbf, 0x4b, 0x70, 0x30, 0x15, 0x7c, 0x1d, 0xa5, 0x11, 0x4f, 0x3c, 0xc9, 0x64,
	0x96, 0xd2, 0x3f, 0x42, 0x27, 0xc8, 0x84, 0xc0, 0x44, 0x2a, 0x20, 0x4f, 0xd3, 0xbd, 0x38, 0x1c,
	0x2e, 0x79, 0x30, 0xdc, 0xe1, 0xa2, 0xbb, 0x43, 0xa4, 0x7f, 0x82, 0xfd, 0x95, 0xc0, 0x75, 0xc4,
	0xb3, 0x34, 0xf7, 0x2c, 0xff, 0xb2, 0xe7, 0x2e, 0xb3, 0x2f, 0xa0, 0x6b, 0xf2, 0xcf, 0x49, 0xcc,
	0xd9, 0xbc, 0xa8, 0xe2, 0x04, 0x9a, 0x2b, 0xc1, 0x43, 0x81, 0x69, 0xaa, 0x2b, 0xa8, 0xb8, 0xcf,
	0x36, 0x3d, 0x82, 0x5a, 0xfa, 0x9c, 0xa0, 0xe5, 0xe6, 0x06, 0x1d, 0x42, 0x07, 0x95, 0x44, 0x26,
	0x4a, 0x16, 0xc5, 0x79, 0xb7, 0xed, 0x0b, 0xd0, 0xd9, 0xb5, 0x76, 0xee, 0xce, 0x79, 0x7f, 0x0d,
	0xf5, 0x1b, 0x64, 0xb1, 0x5c, 0xd0, 0x3f, 0xbc, 0xda, 0x31, 0xd1, 0x9e, 0x39, 0xe5, 0xb5, 0x76,
	0x7f, 0x7e, 0xbd, 0xdd, 0x6f, 0xdd, 0xbe, 0xea, 0xf5, 0x47, 0x68, 0x3c, 0xa0, 0x50, 0x52, 0xd0,
	0x63, 0xa8, 0x27, 0xd9, 0x72, 0x86, 0xa2, 0x98, 0x65, 0x61, 0xf5, 0xff, 0x55, 0x86, 0x7a, 0xa1,
	0xc3, 0x7b, 0xa8, 0x2f, 0x74, 0x2c, 0x4d, 0x69, 0x5f, 0xb4, 0x5f, 0x84, 0x77, 0x8b, 0x23, 0x6a,
	0x02, 0x5d, 0x6d, 0xf5, 0x8d, 0x92, 0x30, 0x77, 0xd5, 0xf5, 0xb4, 0x2f, 0x8e, 0xbe, 0x95, 0x3f,
	0x4b, 0xdd, 0x57, 0xf8, 0xf4, 0x2f, 0xd0, 0x9d, 0xef, 0x0c, 0xa1, 0x90, 0x30, 0x1f, 0xe0, 0xee,
	0x7c, 0xdc, 0xaf, 0xa8, 0xf4, 0x0c, 0x5a, 0x31, 0x4b, 0xa5, 0x16, 0xba, 0x57, 0xfd, 0x46, 0xfa,
	0x2f, 0x87, 0xf4, 0x77, 0xd0, 0x58, 0xe7, 0xfd, 0xf7, 0x6a, 0x9a, 0xd7, 0xd1, 0xbc, 0x42, 0x13,
	0x77, 0x7b, 0xd8, 0x1f, 0x42, 0xdd, 0x4a, 0x64, 0x24, 0x37, 0xf4, 0x03, 0xec, 0x8f, 0xd3, 0x69,
	0xcc, 0x02, 0x5c, 0xf0, 0x78, 0x5e, 0xa8, 0xd5, 0x74, 0x77, 0xc1, 0xfe, 0x4f, 0x50, 0xf1, 0x59,
	0x48, 0x09, 0x54, 0x3e, 0xe1, 0xa6, 0x10, 0x54, 0x7d, 0xaa, 0x75, 0x59, 0xb3, 0x38, 0x7b, 0x5e,
	0x17, 0x6d, 0xf4, 0x3f, 0x40, 0xd5, 0x67, 0x61, 0x4a, 0xbf, 0x83, 0xaa, 0x64, 0xa1, 0x5a, 0xb2,
	0xca, 0x59, 0xfb, 0xa2, 0xa9, 0x6b, 0xf1, 0x59, 0xe8, 0x6a, 0x74, 0xf0, 0x33, 0xb4, 0x26, 0x2b,
	0x14, 0x4c, 0xaa, 0x71, 0x35, 0xa0, 0x72, 0x6d, 0xf9, 0x64, 0x8f, 0x36, 0xa1, 0x3a, 0x9d, 0x78,
	0x3e, 0x29, 0x51, 0x80, 0xba, 0x69, 0xd9, 0x96, 0x6f, 0x91, 0xb2, 0xfa, 0xbe, 0x9f, 0x9a, 0x86,
	0x6f, 0x91, 0xca, 0xe0, 0xbf, 0x65, 0xe8, 0xee, 0xae, 0x3c, 0x6d, 0x43, 0xe3, 0xde, 0xb9, 0x75,
	0x26, 0x1f, 0x1d, 0xb2, 0x47, 0x3b, 0xd0, 0x1c, 0xb9, 0x96, 0xe1, 0x8f, 0x9d, 0x6b, 0x52, 0x52,
	0x47, 0xda, 0xb2, 0x4c, 0x52, 0xa6, 0x6f, 0x60, 0x3f, 0x37, 0x1e, 0xaf, 0x8c, 0xb1, 0x6d, 0x99,
	0xa4, 0xa2, 0xd8, 0x3a, 0x8b, 0x62, 0x57, 0x15, 0x21, 0xcf, 0xb9, 0x25, 0xd4, 0x54, 0x80, 0x1c,
	0x32, 0x49, 0x5d, 0xb1, 0x75, 0x1d, 0x8a, 0xdd, 0x50, 0xec, 0xbc, 0xaa, 0x2d, 0xbb, 0xa9, 0x2b,
	0xd1, 0x90, 0x49, 0x5a, 0x94, 0x40, 0x67, 0xea, 0x4e, 0x1e, 0xc6, 0xde, 0x78, 0xe2, 0x28, 0x0f,
	0xa0, 0x07, 0xd0, 0x7e, 0x46, 0x2c, 0x93, 0xb4, 0xe9, 0x11, 0x90, 0x67, 0x60, 0x1b, 0xa5, 0x43,
	0x29, 0x74, 0x4d, 0x6b, 0xc7, 0x75, 0x3f, 0x2f, 0xed, 0xa5, 0x73, 0x97, 0x1e, 0x03, 0x7d, 0x01,
	0x6d, 0xdd, 0x0f, 0x72, 0x77, 0xdd, 0xc5, 0xd4, 0x72, 0x4c, 0xe5, 0x4e, 0x06, 0x57, 0xf0, 0xf6,
	0x26, 0x0a, 0x17, 0xc6, 0x9a, 0x45, 0x31, 0x9b, 0x45, 0x71, 0x24, 0x37, 0xb9, 0x76, 0x47, 0x40,
	0x0a, 0xed, 0x1e, 0x6f, 0x8c, 0x47, 0xcf, 0x57, 0x22, 0xef, 0x29, 0xc1, 0x3d, 0xdf, 0xb8, 0xb4,
	0xad, 0x5c, 0xc2, 0x6d, 0x9c, 0xf2, 0x20, 0x84, 0xf6, 0x8b, 0x0b, 0xa8, 0x04, 0x71, 0x26, 0xfe,
	0x56, 0xfa, 0x3a, 0x94, 0x27, 0xb7, 0xb9, 0xc7, 0x47, 0xc3, 0xd5, 0x85, 0x97, 0xf3, 0x79, 0x8c,
	0xfd, 0xf1, 0xc8, 0xb0, 0x49, 0x45, 0x1d, 0xdd, 0x8d, 0x3d, 0x2f, 0x97, 0x5b, 0x8b, 0x7f, 0xed,
	0x1a, 0xa6, 0x56, 0x3a, 0x8f, 0x75, 0x35, 0xb9, 0x77, 0x4c, 0x52, 0x1f, 0x2c, 0x00, 0x46, 0x71,
	0x84, 0x89, 0xf4, 0x37, 0x2b, 0x54, 0x52, 0x8e, 0x26, 0x8e, 0xef, 0x4e, 0xec, 0xa9, 0x6d, 0x38,
	0xaa, 0x42, 0x0a, 0x5d, 0xeb, 0x6f, 0xbe, 0xe5, 0x3a, 0x86, 0x3d, 0xb2, 0xc7, 0x96, 0xa3, 0x56,
	0xa6, 0x09, 0x55, 0x67, 0x62, 0xaa, 0x85, 0x69, 0x41, 0xcd, 0x30, 0xef, 0xc6, 0x0e, 0xa9, 0xd0,
	0x7d, 0x68, 0x5d, 0x1a, 0xae, 0x75, 0x67, 0xf9, 0x86, 0x4d, 0xaa, 0x2a, 0x92, 0x3d, 0x31, 0xcc,
	0x4b, 0xc3, 0x36, 0x9c, 0x91, 0xe5, 0x92, 0xda, 0xe0, 0x02, 0xa8, 0x91, 0xc9, 0x05, 0x26, 0x32,
	0x0a, 0xf4, 0x36, 0xea, 0x8c, 0x5d, 0x00, 0xcf, 0xb2, 0xaf, 0xbc, 0xf1, 0xb5, 0x12, 0x3b, 0x5f,
	0x2b, 0xa3, 0xb0, 0x4a, 0x83, 0x7f, 0xd6, 0xa0, 0xa3, 0x97, 0x70, 0x8e, 0x42, 0xd3, 0x0f, 0xa0,
	0x6d, 0x24, 0x9b, 0x2d, 0x94, 0xd7, 0xf7, 0x10, 0x09, 0x99, 0xb1, 0xf8, 0x8e, 0x05, 0x8b, 0x28,
	0x41, 0x52, 0xa2, 0x27, 0x70, 0xbc, 0x8b, 0x79, 0x01, 0x8b, 0xd1, 0x43, 0x49, 0xca, 0xba, 0x2e,
	0xce, 0xe6, 0x97, 0x2c, 0x66, 0x49, 0x80, 0x82, 0x54, 0x5e, 0x44, 0x70, 0x50, 0x7e, 0xe6, 0xe2,
	0x13, 0xa9, 0xd2, 0x43, 0x38, 0x28, 0xb0, 0x1b, 0x26, 0xe6, 0x66, 0x94, 0x7e, 0x22, 0x35, 0xe5,
	0x7a, 0xcd, 0xe2, 0x18, 0xc5, 0x66, 0xbc, 0x64, 0x21, 0x92, 0x3a, 0x7d, 0x07, 0x87, 0xbb, 0x89,
	0xf2, 0x83, 0x86, 0x9a, 0x76, 0x11, 0x6c, 0x9c, 0x48, 0x14, 0x4f, 0x2c, 0x40, 0xd2, 0x54, 0xc5,
	0x8f, 0x50, 0xc8, 0xe8, 0x49, 0x09, 0x80, 0xa4, 0xa5, 0xae, 0xe3, 0x2d, 0x6e, 0x08, 0xe8, 0x3d,
	0xc0, 0x40, 0xa0, 0x24, 0x6d, 0xa5, 0xc0, 0x2d, 0x6e, 0x1e, 0x58, 0x16, 0x4b, 0xd2, 0x51, 0xd6,
	0x78, 0x8e, 0xfa, 0xbf, 0x08, 0xd9, 0x57, 0xca, 0xbb, 0x3c, 0x46, 0xd2, 0x55, 0x55, 0xab, 0x2f,
	0x23, 0x4d, 0xa3, 0x30, 0x59, 0x62, 0x22, 0xc9, 0x81, 0xd2, 0xf2, 0x36, 0x9b, 0xa1, 0x48, 0x50,
	0x62, 0x4a, 0x88, 0xbe, 0x94, 0x71, 0x96, 0x4a, 0x14, 0xe4, 0x8d, 0x1e, 0x2d, 0x4f, 0xa4, 0xe0,
	0xf1, 0x34, 0x66, 0x09, 0x12, 0xaa, 0x86, 0x77, 0x2d, 0x78, 0xb6, 0x22, 0x87, 0x7a, 0xa2, 0x7c,
	0x8e, 0xe4, 0x48, 0xe5, 0xb3, 0x79, 0x3e, 0x1f, 0xf2, 0x56, 0xf5, 0xe1, 0x49, 0x2e, 0x58, 0x88,
	0xca, 0x97, 0x45, 0x09, 0x0a, 0x72, 0xac, 0xfa, 0x28, 0xd0, 0xab, 0x28, 0x46, 0xf2, 0xee, 0x05,
	0xcd, 0x8c, 0x04, 0x06, 0x92, 0x8b, 0x0d, 0xe9, 0xa9, 0x8c, 0x5e, 0x36, 0xcb, 0x9f, 0x01, 0x2a,
	0xdc, 0xaf, 0x54, 0x41, 0x0f, 0xd1, 0x6a, 0xca, 0x79, 0x4c, 0x4e, 0xf4, 0x8a, 0xb2, 0x40, 0x1b,
	0xbf, 0x56, 0x21, 0x2d, 0x19, 0xcc, 0xb7, 0xe5, 0x7e, 0xa7, 0x7a, 0x51, 0x80, 0x87, 0x62, 0x8d,
	0x82, 0x7c, 0xaf, 0x52, 0x5c, 0x32, 0x81, 0x77, 0x28, 0xbf, 0x4c, 0xfa, 0x07, 0xfa, 0x16, 0xde,
	0x8c, 0x04, 0x6a, 0x7d, 0x58, 0x7c, 0xc7, 0x93, 0x48, 0x72, 0x41, 0x7e, 0xa3, 0x42, 0xdb, 0x3c,
	0x0c, 0xa3, 0x24, 0x24, 0xa7, 0xaa, 0x23, 0x17, 0x03, 0xbe, 0x46, 0xb1, 0x21, 0x3f, 0xaa, 0xa6,
	0x4d, 0x9c, 0x65, 0x21, 0xe9, 0xab, 0xab, 0xfe, 0x1c, 0xf2, 0x86, 0xa7, 0x92, 0xbc, 0x57, 0xaa,
	0xee, 0xee, 0x28, 0xf9, 0x30, 0x70, 0xa1, 0xad, 0xc7, 0xea, 0xf1, 0x4c, 0x04, 0x98, 0x2f, 0xf6,
	0xc8, 0xb0, 0x1f, 0xbd, 0xc9, 0xbd, 0x3b, 0x52, 0x57, 0x44, 0xad, 0xf0, 0x95, 0xb7, 0xb5, 0x4b,
	0xaa, 0x97, 0x1b, 0xdf, 0x9f, 0x6e, 0x01, 0xbd, 0x73, 0x23, 0x7b, 0xe2, 0x58, 0x5b, 0xa4, 0x32,
	0xf8, 0x09, 0x0e, 0x47, 0x31, 0xcf, 0xe6, 0xe3, 0x24, 0x92, 0x26, 0x93, 0xac, 0x88, 0xdd, 0x86,
	0x86, 0xc3, 0xf5, 0x01, 0xd9, 0xd3, 0x77, 0xeb, 0x1f, 0x99, 0x40, 0x52, 0xfa, 0xf3, 0x5f, 0xa1,
	0x95, 0x62, 0x92, 0x46, 0x32, 0x5a, 0x23, 0xfd, 0x7e, 0x98, 0xbf, 0xb5, 0x86, 0xdb, 0xb7, 0xd6,
	0xf0, 0x2a, 0xc2, 0x78, 0x3e, 0xd1, 0x22, 0xa7, 0xbd, 0xff, 0xfd, 0xa7, 0xa2, 0x7f, 0x5b, 0xbe,
	0x78, 0x5c, 0xfe, 0xf6, 0xef, 0xef, 0xc3, 0x48, 0x2e, 0xb2, 0xd9, 0x30, 0xe0, 0xcb, 0xf3, 0x65,
	0x14, 0x08, 0x9e, 0xf2, 0x27, 0x79, 0xbe, 0xe4, 0xc1, 0xb9, 0x58, 0x05, 0xe7, 0xf9, 0x73, 0x6e,
	0x56, 0xd7, 0x01, 0x7f, 0xff, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xde, 0xea, 0x3e, 0xea,
	0x09, 0x00, 0x00,
}
