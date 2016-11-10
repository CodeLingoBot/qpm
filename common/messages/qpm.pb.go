// Code generated by protoc-gen-go.
// source: qpm.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	qpm.proto

It has these top-level messages:
	DependencyMessage
	Package
	Dependency
	VersionInfo
	SearchResult
	InstallStats
	PingRequest
	PingResponse
	PublishRequest
	PublishResponse
	DependencyRequest
	DependencyResponse
	SearchRequest
	SearchResponse
	ListRequest
	ListResponse
	LoginRequest
	LoginResponse
	InfoRequest
	InfoResponse
	LicenseRequest
	LicenseResponse
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RepoType int32

const (
	RepoType_AUTO      RepoType = 0
	RepoType_GITHUB    RepoType = 1
	RepoType_GIT       RepoType = 2
	RepoType_MERCURIAL RepoType = 3
)

var RepoType_name = map[int32]string{
	0: "AUTO",
	1: "GITHUB",
	2: "GIT",
	3: "MERCURIAL",
}
var RepoType_value = map[string]int32{
	"AUTO":      0,
	"GITHUB":    1,
	"GIT":       2,
	"MERCURIAL": 3,
}

func (x RepoType) String() string {
	return proto.EnumName(RepoType_name, int32(x))
}
func (RepoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The values in this enum should correspond to an SPDX identifier
type LicenseType int32

const (
	LicenseType_NONE         LicenseType = 0
	LicenseType_MIT          LicenseType = 1
	LicenseType_AGPL_3_0     LicenseType = 2
	LicenseType_APACHE_2_0   LicenseType = 3
	LicenseType_ARTISTIC_2_0 LicenseType = 4
	LicenseType_BSD_2_CLAUSE LicenseType = 5
	LicenseType_BSD_3_CLAUSE LicenseType = 6
	LicenseType_CC0_1_0      LicenseType = 7
	LicenseType_EPL_1_0      LicenseType = 8
	LicenseType_GPL_2_0      LicenseType = 9
	LicenseType_GPL_3_0      LicenseType = 10
	LicenseType_ISC          LicenseType = 11
	LicenseType_LGPL_2_1     LicenseType = 12
	LicenseType_LGPL_3_0     LicenseType = 13
	LicenseType_UNLICENSE    LicenseType = 14
	LicenseType_MPL_2_0      LicenseType = 15
)

var LicenseType_name = map[int32]string{
	0:  "NONE",
	1:  "MIT",
	2:  "AGPL_3_0",
	3:  "APACHE_2_0",
	4:  "ARTISTIC_2_0",
	5:  "BSD_2_CLAUSE",
	6:  "BSD_3_CLAUSE",
	7:  "CC0_1_0",
	8:  "EPL_1_0",
	9:  "GPL_2_0",
	10: "GPL_3_0",
	11: "ISC",
	12: "LGPL_2_1",
	13: "LGPL_3_0",
	14: "UNLICENSE",
	15: "MPL_2_0",
}
var LicenseType_value = map[string]int32{
	"NONE":         0,
	"MIT":          1,
	"AGPL_3_0":     2,
	"APACHE_2_0":   3,
	"ARTISTIC_2_0": 4,
	"BSD_2_CLAUSE": 5,
	"BSD_3_CLAUSE": 6,
	"CC0_1_0":      7,
	"EPL_1_0":      8,
	"GPL_2_0":      9,
	"GPL_3_0":      10,
	"ISC":          11,
	"LGPL_2_1":     12,
	"LGPL_3_0":     13,
	"UNLICENSE":    14,
	"MPL_2_0":      15,
}

func (x LicenseType) String() string {
	return proto.EnumName(LicenseType_name, int32(x))
}
func (LicenseType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MessageType int32

const (
	MessageType_INFO    MessageType = 0
	MessageType_WARNING MessageType = 1
	MessageType_ERROR   MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
}
var MessageType_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DependencyMessage struct {
	Type   MessageType `protobuf:"varint,1,opt,name=type,enum=messages.MessageType" json:"type,omitempty"`
	Title  string      `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Body   string      `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Prompt bool        `protobuf:"varint,4,opt,name=prompt" json:"prompt,omitempty"`
}

func (m *DependencyMessage) Reset()                    { *m = DependencyMessage{} }
func (m *DependencyMessage) String() string            { return proto.CompactTextString(m) }
func (*DependencyMessage) ProtoMessage()               {}
func (*DependencyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Package struct {
	Name         string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description  string              `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Author       *Package_Author     `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Repository   *Package_Repository `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Version      *Package_Version    `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	Dependencies []string            `protobuf:"bytes,6,rep,name=dependencies" json:"dependencies,omitempty"`
	License      LicenseType         `protobuf:"varint,7,opt,name=license,enum=messages.LicenseType" json:"license,omitempty"`
	PriFilename  string              `protobuf:"bytes,8,opt,name=pri_filename,json=priFilename" json:"pri_filename,omitempty"`
	Webpage      string              `protobuf:"bytes,10,opt,name=webpage" json:"webpage,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Package) GetAuthor() *Package_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Package) GetRepository() *Package_Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Package) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Package_Repository struct {
	Type RepoType `protobuf:"varint,1,opt,name=type,enum=messages.RepoType" json:"type,omitempty"`
	Url  string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Package_Repository) Reset()                    { *m = Package_Repository{} }
func (m *Package_Repository) String() string            { return proto.CompactTextString(m) }
func (*Package_Repository) ProtoMessage()               {}
func (*Package_Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Package_Version struct {
	Label       string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Revision    string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Fingerprint string `protobuf:"bytes,3,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *Package_Version) Reset()                    { *m = Package_Version{} }
func (m *Package_Version) String() string            { return proto.CompactTextString(m) }
func (*Package_Version) ProtoMessage()               {}
func (*Package_Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type Package_Author struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Package_Author) Reset()                    { *m = Package_Author{} }
func (m *Package_Author) String() string            { return proto.CompactTextString(m) }
func (*Package_Author) ProtoMessage()               {}
func (*Package_Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

type Dependency struct {
	Name       string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Repository *Package_Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Version    *Package_Version    `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *Dependency) Reset()                    { *m = Dependency{} }
func (m *Dependency) String() string            { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()               {}
func (*Dependency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Dependency) GetRepository() *Package_Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Dependency) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type VersionInfo struct {
	Version       *Package_Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	DatePublished string           `protobuf:"bytes,2,opt,name=date_published,json=datePublished" json:"date_published,omitempty"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VersionInfo) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type SearchResult struct {
	Name        string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version     string          `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Author      *Package_Author `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Description string          `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	License     LicenseType     `protobuf:"varint,5,opt,name=license,enum=messages.LicenseType" json:"license,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchResult) GetAuthor() *Package_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

type InstallStats struct {
	Daily   uint32 `protobuf:"varint,1,opt,name=daily" json:"daily,omitempty"`
	Weekly  uint32 `protobuf:"varint,2,opt,name=weekly" json:"weekly,omitempty"`
	Monthly uint32 `protobuf:"varint,3,opt,name=monthly" json:"monthly,omitempty"`
	Yearly  uint32 `protobuf:"varint,4,opt,name=yearly" json:"yearly,omitempty"`
	Total   uint32 `protobuf:"varint,5,opt,name=total" json:"total,omitempty"`
}

func (m *InstallStats) Reset()                    { *m = InstallStats{} }
func (m *InstallStats) String() string            { return proto.CompactTextString(m) }
func (*InstallStats) ProtoMessage()               {}
func (*InstallStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PingResponse struct {
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type PublishRequest struct {
	PackageDescription *Package `protobuf:"bytes,1,opt,name=package_description,json=packageDescription" json:"package_description,omitempty"`
	Token              string   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PublishRequest) GetPackageDescription() *Package {
	if m != nil {
		return m.PackageDescription
	}
	return nil
}

type PublishResponse struct {
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DependencyRequest struct {
	PackageNames  []string    `protobuf:"bytes,1,rep,name=package_names,json=packageNames" json:"package_names,omitempty"`
	CompatLicense LicenseType `protobuf:"varint,4,opt,name=compat_license,json=compatLicense,enum=messages.LicenseType" json:"compat_license,omitempty"`
}

func (m *DependencyRequest) Reset()                    { *m = DependencyRequest{} }
func (m *DependencyRequest) String() string            { return proto.CompactTextString(m) }
func (*DependencyRequest) ProtoMessage()               {}
func (*DependencyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type DependencyResponse struct {
	Dependencies []*Dependency        `protobuf:"bytes,1,rep,name=dependencies" json:"dependencies,omitempty"`
	Messages     []*DependencyMessage `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *DependencyResponse) Reset()                    { *m = DependencyResponse{} }
func (m *DependencyResponse) String() string            { return proto.CompactTextString(m) }
func (*DependencyResponse) ProtoMessage()               {}
func (*DependencyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DependencyResponse) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *DependencyResponse) GetMessages() []*DependencyMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type SearchRequest struct {
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SearchResponse struct {
	Results []*SearchResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SearchResponse) GetResults() []*SearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type ListResponse struct {
	Results []*SearchResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ListResponse) GetResults() []*SearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Create   bool   `protobuf:"varint,3,opt,name=create" json:"create,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type InfoRequest struct {
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type InfoResponse struct {
	Package      *Package       `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Versions     []*VersionInfo `protobuf:"bytes,2,rep,name=versions" json:"versions,omitempty"`
	Dependencies []*Dependency  `protobuf:"bytes,3,rep,name=dependencies" json:"dependencies,omitempty"`
	InstallStats *InstallStats  `protobuf:"bytes,4,opt,name=install_stats,json=installStats" json:"install_stats,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *InfoResponse) GetPackage() *Package {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *InfoResponse) GetVersions() []*VersionInfo {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *InfoResponse) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *InfoResponse) GetInstallStats() *InstallStats {
	if m != nil {
		return m.InstallStats
	}
	return nil
}

type LicenseRequest struct {
	Package *Package `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
}

func (m *LicenseRequest) Reset()                    { *m = LicenseRequest{} }
func (m *LicenseRequest) String() string            { return proto.CompactTextString(m) }
func (*LicenseRequest) ProtoMessage()               {}
func (*LicenseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *LicenseRequest) GetPackage() *Package {
	if m != nil {
		return m.Package
	}
	return nil
}

type LicenseResponse struct {
	Body string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
}

func (m *LicenseResponse) Reset()                    { *m = LicenseResponse{} }
func (m *LicenseResponse) String() string            { return proto.CompactTextString(m) }
func (*LicenseResponse) ProtoMessage()               {}
func (*LicenseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func init() {
	proto.RegisterType((*DependencyMessage)(nil), "messages.DependencyMessage")
	proto.RegisterType((*Package)(nil), "messages.Package")
	proto.RegisterType((*Package_Repository)(nil), "messages.Package.Repository")
	proto.RegisterType((*Package_Version)(nil), "messages.Package.Version")
	proto.RegisterType((*Package_Author)(nil), "messages.Package.Author")
	proto.RegisterType((*Dependency)(nil), "messages.Dependency")
	proto.RegisterType((*VersionInfo)(nil), "messages.VersionInfo")
	proto.RegisterType((*SearchResult)(nil), "messages.SearchResult")
	proto.RegisterType((*InstallStats)(nil), "messages.InstallStats")
	proto.RegisterType((*PingRequest)(nil), "messages.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "messages.PingResponse")
	proto.RegisterType((*PublishRequest)(nil), "messages.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "messages.PublishResponse")
	proto.RegisterType((*DependencyRequest)(nil), "messages.DependencyRequest")
	proto.RegisterType((*DependencyResponse)(nil), "messages.DependencyResponse")
	proto.RegisterType((*SearchRequest)(nil), "messages.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "messages.SearchResponse")
	proto.RegisterType((*ListRequest)(nil), "messages.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "messages.ListResponse")
	proto.RegisterType((*LoginRequest)(nil), "messages.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "messages.LoginResponse")
	proto.RegisterType((*InfoRequest)(nil), "messages.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "messages.InfoResponse")
	proto.RegisterType((*LicenseRequest)(nil), "messages.LicenseRequest")
	proto.RegisterType((*LicenseResponse)(nil), "messages.LicenseResponse")
	proto.RegisterEnum("messages.RepoType", RepoType_name, RepoType_value)
	proto.RegisterEnum("messages.LicenseType", LicenseType_name, LicenseType_value)
	proto.RegisterEnum("messages.MessageType", MessageType_name, MessageType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Qpm service

type QpmClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	GetLicense(ctx context.Context, in *LicenseRequest, opts ...grpc.CallOption) (*LicenseResponse, error)
}

type qpmClient struct {
	cc *grpc.ClientConn
}

func NewQpmClient(cc *grpc.ClientConn) QpmClient {
	return &qpmClient{cc}
}

func (c *qpmClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error) {
	out := new(DependencyResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/GetDependencies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) GetLicense(ctx context.Context, in *LicenseRequest, opts ...grpc.CallOption) (*LicenseResponse, error) {
	out := new(LicenseResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/GetLicense", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Qpm service

type QpmServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	GetDependencies(context.Context, *DependencyRequest) (*DependencyResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	GetLicense(context.Context, *LicenseRequest) (*LicenseResponse, error)
}

func RegisterQpmServer(s *grpc.Server, srv QpmServer) {
	s.RegisterService(&_Qpm_serviceDesc, srv)
}

func _Qpm_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_GetDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).GetDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/GetDependencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).GetDependencies(ctx, req.(*DependencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qpm_GetLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QpmServer).GetLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Qpm/GetLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QpmServer).GetLicense(ctx, req.(*LicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Qpm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Qpm",
	HandlerType: (*QpmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Qpm_Ping_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Qpm_Publish_Handler,
		},
		{
			MethodName: "GetDependencies",
			Handler:    _Qpm_GetDependencies_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Qpm_Search_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Qpm_List_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Qpm_Login_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Qpm_Info_Handler,
		},
		{
			MethodName: "GetLicense",
			Handler:    _Qpm_GetLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("qpm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x57, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x36, 0xad, 0x03, 0xa5, 0xa1, 0x28, 0x33, 0xfb, 0x27, 0x0e, 0xa3, 0x3f, 0x17, 0x2e, 0x8b,
	0x14, 0x6e, 0x0a, 0x38, 0x8a, 0x7c, 0x91, 0xa0, 0x4d, 0x80, 0xc8, 0xb2, 0xe2, 0x10, 0x90, 0x15,
	0x77, 0x25, 0xb7, 0xbd, 0x29, 0x08, 0x5a, 0xda, 0xd8, 0x6c, 0x28, 0x92, 0x21, 0xe9, 0x04, 0xba,
	0x2b, 0x8a, 0x02, 0x7d, 0x80, 0x3e, 0x51, 0x1f, 0xa1, 0x0f, 0xd1, 0x8b, 0xbe, 0x45, 0xb1, 0x27,
	0x6a, 0x75, 0x68, 0xe1, 0xe4, 0x4e, 0x33, 0x3b, 0xa7, 0x9d, 0xf9, 0xe6, 0x5b, 0x0a, 0xea, 0xef,
	0x92, 0xd9, 0x41, 0x92, 0xc6, 0x79, 0x8c, 0x6a, 0x33, 0x92, 0x65, 0xfe, 0x25, 0xc9, 0x9c, 0x9f,
	0x35, 0xb8, 0x75, 0x4c, 0x12, 0x12, 0x4d, 0x49, 0x34, 0x99, 0x9f, 0x72, 0x35, 0xfa, 0x12, 0xca,
	0xf9, 0x3c, 0x21, 0xb6, 0xb6, 0xa7, 0xed, 0x37, 0x3b, 0x77, 0x0e, 0xa4, 0xf9, 0x81, 0x30, 0x18,
	0xcf, 0x13, 0x82, 0x99, 0x09, 0xba, 0x0d, 0x95, 0x3c, 0xc8, 0x43, 0x62, 0x6f, 0xef, 0x69, 0xfb,
	0x75, 0xcc, 0x05, 0x84, 0xa0, 0x7c, 0x11, 0x4f, 0xe7, 0x76, 0x89, 0x29, 0xd9, 0x6f, 0xb4, 0x0b,
	0xd5, 0x24, 0x8d, 0x67, 0x49, 0x6e, 0x97, 0xf7, 0xb4, 0xfd, 0x1a, 0x16, 0x92, 0xf3, 0x67, 0x19,
	0xf4, 0x33, 0x7f, 0xf2, 0x96, 0x26, 0x46, 0x50, 0x8e, 0xfc, 0x19, 0x4f, 0x5c, 0xc7, 0xec, 0x37,
	0xda, 0x03, 0x63, 0x4a, 0xb2, 0x49, 0x1a, 0x24, 0x79, 0x10, 0x47, 0x22, 0x8f, 0xaa, 0x42, 0x6d,
	0xa8, 0xfa, 0xd7, 0xf9, 0x55, 0x9c, 0xb2, 0x7c, 0x46, 0xc7, 0x5e, 0x14, 0x2c, 0x02, 0x1f, 0x74,
	0xd9, 0x39, 0x16, 0x76, 0xe8, 0x19, 0x40, 0x4a, 0x92, 0x38, 0x0b, 0xf2, 0x38, 0x9d, 0xb3, 0x7a,
	0x8c, 0xce, 0xfd, 0x75, 0x2f, 0x5c, 0xd8, 0x60, 0xc5, 0x1e, 0x1d, 0x82, 0xfe, 0x9e, 0xa4, 0x19,
	0xad, 0xa6, 0xc2, 0x5c, 0xef, 0xad, 0xbb, 0x7e, 0xc7, 0x0d, 0xb0, 0xb4, 0x44, 0x0e, 0x34, 0xa6,
	0xb2, 0xd1, 0x01, 0xc9, 0xec, 0xea, 0x5e, 0x69, 0xbf, 0x8e, 0x97, 0x74, 0xe8, 0x11, 0xe8, 0x61,
	0x30, 0x21, 0x51, 0x46, 0x6c, 0x7d, 0xb5, 0xf5, 0x03, 0x7e, 0xc0, 0x5a, 0x2f, 0xad, 0xd0, 0x67,
	0xd0, 0x48, 0xd2, 0xc0, 0x7b, 0x13, 0x84, 0x84, 0xf5, 0xad, 0xc6, 0x9b, 0x93, 0xa4, 0xc1, 0x4b,
	0xa1, 0x42, 0x36, 0xe8, 0x1f, 0xc8, 0x45, 0xe2, 0x5f, 0x12, 0x1b, 0xd8, 0xa9, 0x14, 0x5b, 0x2f,
	0x01, 0x16, 0x17, 0x44, 0x5f, 0x2c, 0xcd, 0x1c, 0x2d, 0x12, 0x53, 0x1b, 0x65, 0xe0, 0x16, 0x94,
	0xae, 0xd3, 0x50, 0x8c, 0x81, 0xfe, 0x6c, 0xfd, 0x08, 0xba, 0xb8, 0x2d, 0x45, 0x43, 0xe8, 0x5f,
	0x90, 0x50, 0x0c, 0x90, 0x0b, 0xa8, 0x05, 0xb5, 0x94, 0xbc, 0x0f, 0xb2, 0xc5, 0xf8, 0x0a, 0x99,
	0x4e, 0xf7, 0x4d, 0x10, 0x5d, 0x92, 0x34, 0x49, 0x83, 0x28, 0x17, 0x80, 0x51, 0x55, 0xad, 0x0e,
	0x54, 0xf9, 0xf4, 0x36, 0xa2, 0xe3, 0x36, 0x54, 0xc8, 0xcc, 0x0f, 0x64, 0x41, 0x5c, 0x70, 0x7e,
	0xd7, 0x00, 0x16, 0xb0, 0xde, 0xe8, 0xb8, 0x0c, 0x81, 0xed, 0x4f, 0x87, 0x40, 0xe9, 0xa6, 0x10,
	0x70, 0x02, 0x30, 0x84, 0xce, 0x8d, 0xde, 0xc4, 0x6a, 0x0c, 0xed, 0xc6, 0x30, 0x7a, 0x00, 0xcd,
	0xa9, 0x9f, 0x13, 0x2f, 0xb9, 0xbe, 0x08, 0x83, 0xec, 0x8a, 0x4c, 0xc5, 0xc5, 0x4d, 0xaa, 0x3d,
	0x93, 0x4a, 0xe7, 0x0f, 0x0d, 0x1a, 0x23, 0xe2, 0xa7, 0x93, 0x2b, 0x4c, 0xb2, 0xeb, 0x30, 0xdf,
	0xd8, 0x02, 0x7b, 0x51, 0x00, 0x0f, 0x52, 0x64, 0xf9, 0xf8, 0x8d, 0x5a, 0xd9, 0xd2, 0xf2, 0xfa,
	0x96, 0x2a, 0xe0, 0xae, 0xdc, 0x04, 0xdc, 0xce, 0xaf, 0x1a, 0x34, 0xdc, 0x28, 0xcb, 0xfd, 0x30,
	0x1c, 0xe5, 0x7e, 0x9e, 0xd1, 0x59, 0x4f, 0xfd, 0x20, 0x9c, 0xb3, 0x4b, 0x98, 0x98, 0x0b, 0x94,
	0x57, 0x3e, 0x10, 0xf2, 0x36, 0xe4, 0x43, 0x34, 0xb1, 0x90, 0xe8, 0xed, 0x66, 0x71, 0x94, 0x5f,
	0x85, 0x9c, 0x86, 0x4c, 0x2c, 0x45, 0xea, 0x31, 0x27, 0x7e, 0x1a, 0xf2, 0xcd, 0x37, 0xb1, 0x90,
	0x18, 0x97, 0xc5, 0xb9, 0x1f, 0xb2, 0xfa, 0x4c, 0xcc, 0x05, 0xc7, 0x04, 0xe3, 0x2c, 0x88, 0x2e,
	0x31, 0x79, 0x77, 0x4d, 0xb2, 0xdc, 0x69, 0x42, 0x83, 0x8b, 0x59, 0x12, 0xd3, 0x2a, 0x7f, 0x82,
	0xa6, 0x68, 0xbb, 0xb0, 0x40, 0x47, 0xf0, 0xbf, 0x84, 0x37, 0xc9, 0x53, 0x5b, 0xc2, 0x67, 0x7c,
	0x6b, 0xad, 0x93, 0x18, 0x09, 0xeb, 0x63, 0xa5, 0x59, 0xac, 0x94, 0xb7, 0x24, 0x2a, 0x68, 0x95,
	0x0a, 0xce, 0x2d, 0xd8, 0x29, 0x72, 0x89, 0xf4, 0xef, 0x55, 0xfe, 0x96, 0x15, 0x7c, 0x0e, 0xa6,
	0xac, 0x80, 0x0e, 0x3a, 0xb3, 0x35, 0x4e, 0x36, 0x42, 0x39, 0xa4, 0x3a, 0xf4, 0x0c, 0x9a, 0x93,
	0x78, 0x96, 0xf8, 0xb9, 0x27, 0xc7, 0x52, 0xfe, 0xaf, 0xb1, 0x98, 0xdc, 0x58, 0xa8, 0x9c, 0xdf,
	0x34, 0x40, 0x6a, 0x62, 0x5e, 0x0e, 0x7a, 0xba, 0xc2, 0x72, 0x34, 0xb1, 0xd1, 0xb9, 0xbd, 0x08,
	0xa9, 0xf8, 0x2c, 0x73, 0xdf, 0x13, 0x28, 0x5e, 0x25, 0x7b, 0x9b, 0x79, 0xfd, 0x7f, 0x93, 0x97,
	0x78, 0x81, 0xf0, 0xe2, 0x09, 0xeb, 0x80, 0x29, 0x91, 0xce, 0x6f, 0x4f, 0x49, 0x51, 0xb9, 0xbd,
	0x80, 0xbc, 0xa1, 0x5c, 0xde, 0x39, 0x82, 0x66, 0xb1, 0x1d, 0xbc, 0xf0, 0x36, 0xe8, 0x29, 0xdb,
	0x14, 0x59, 0xf3, 0xee, 0x22, 0xbb, 0xba, 0x48, 0x58, 0x9a, 0x51, 0x5c, 0x0c, 0x82, 0x2c, 0x97,
	0xb8, 0x78, 0x01, 0x0d, 0x2e, 0x7e, 0x72, 0xc0, 0x1f, 0xa0, 0x31, 0x88, 0x2f, 0x83, 0x48, 0xde,
	0xa3, 0xa0, 0x36, 0x4d, 0xa1, 0x36, 0x4a, 0xa6, 0x89, 0x9f, 0x65, 0x1f, 0xe2, 0x54, 0xae, 0x7e,
	0x21, 0x53, 0x60, 0x4f, 0x52, 0xe2, 0xe7, 0x84, 0x21, 0xbe, 0x86, 0x85, 0xe4, 0x3c, 0x00, 0x53,
	0x44, 0x16, 0xc5, 0x15, 0xf0, 0xd2, 0x54, 0x78, 0xb5, 0xc1, 0xa0, 0xc4, 0xf4, 0x11, 0x7d, 0xfc,
	0x8b, 0xad, 0x28, 0x75, 0x11, 0x81, 0xbf, 0x02, 0x5d, 0x9c, 0xff, 0x3b, 0xde, 0xa5, 0x05, 0x7a,
	0x0c, 0x35, 0x41, 0x38, 0x72, 0xe4, 0x0a, 0xf6, 0x14, 0xa6, 0xc4, 0x85, 0xd9, 0x1a, 0xbe, 0x4a,
	0x37, 0xc6, 0xd7, 0x37, 0x60, 0x06, 0x9c, 0x4c, 0xbc, 0x8c, 0xb2, 0x89, 0x78, 0xf5, 0x95, 0xa9,
	0xa8, 0x5c, 0x83, 0x1b, 0x81, 0x22, 0x39, 0xcf, 0xa1, 0x29, 0x80, 0x2f, 0x9b, 0xf3, 0x31, 0x17,
	0x75, 0x1e, 0xc0, 0x4e, 0xe1, 0x2e, 0x1a, 0x25, 0xbf, 0x90, 0xb4, 0xc5, 0x17, 0xd2, 0xc3, 0xa7,
	0x50, 0x93, 0x8f, 0x2d, 0xaa, 0x41, 0xb9, 0x7b, 0x3e, 0x7e, 0x6d, 0x6d, 0x21, 0x80, 0xea, 0x89,
	0x3b, 0x7e, 0x75, 0x7e, 0x64, 0x69, 0x48, 0x87, 0xd2, 0x89, 0x3b, 0xb6, 0xb6, 0x91, 0x09, 0xf5,
	0xd3, 0x3e, 0xee, 0x9d, 0x63, 0xb7, 0x3b, 0xb0, 0x4a, 0x0f, 0xff, 0xd6, 0x28, 0x18, 0x8b, 0x65,
	0xa5, 0xde, 0xc3, 0xd7, 0xc3, 0xbe, 0xb5, 0x45, 0x3d, 0x4e, 0xdd, 0xb1, 0xa5, 0xa1, 0x06, 0xd4,
	0xba, 0x27, 0x67, 0x03, 0xef, 0xd0, 0x6b, 0x5b, 0xdb, 0xa8, 0x09, 0xd0, 0x3d, 0xeb, 0xf6, 0x5e,
	0xf5, 0xbd, 0x8e, 0xd7, 0xb6, 0x4a, 0xc8, 0x82, 0x46, 0x17, 0x8f, 0xdd, 0xd1, 0xd8, 0xed, 0x31,
	0x4d, 0x99, 0x6a, 0x8e, 0x46, 0xc7, 0x5e, 0xc7, 0xeb, 0x0d, 0xba, 0xe7, 0xa3, 0xbe, 0x55, 0x91,
	0x9a, 0x43, 0xa9, 0xa9, 0x22, 0x03, 0xf4, 0x5e, 0xaf, 0xed, 0x3d, 0xf6, 0xda, 0x96, 0x4e, 0x85,
	0xfe, 0xd9, 0x80, 0x09, 0x35, 0x2a, 0xd0, 0x64, 0x34, 0x54, 0x5d, 0x0a, 0x34, 0x33, 0xd0, 0x82,
	0xdc, 0x51, 0xcf, 0x32, 0x68, 0x41, 0x03, 0x6e, 0xf3, 0xd8, 0x6a, 0x14, 0x12, 0x35, 0x32, 0xe9,
	0xf5, 0xce, 0x87, 0x03, 0xb7, 0xd7, 0x1f, 0x8e, 0xfa, 0x56, 0x93, 0x06, 0x38, 0x15, 0xd1, 0x76,
	0x1e, 0x3e, 0x02, 0x43, 0xf9, 0x0c, 0xa5, 0x57, 0x75, 0x87, 0x2f, 0x69, 0xa3, 0x0c, 0xd0, 0xbf,
	0xef, 0xe2, 0xa1, 0x3b, 0x3c, 0xb1, 0x34, 0x54, 0x87, 0x4a, 0x1f, 0xe3, 0xd7, 0xd8, 0xda, 0xee,
	0xfc, 0x52, 0x86, 0xd2, 0xb7, 0xc9, 0x0c, 0x3d, 0x81, 0x32, 0x65, 0x6e, 0xa4, 0x80, 0x4c, 0x21,
	0xf6, 0xd6, 0xee, 0xaa, 0x5a, 0x30, 0xec, 0x16, 0x7a, 0x01, 0xba, 0xa0, 0x5d, 0xa4, 0x3e, 0x84,
	0x4b, 0xac, 0xdf, 0xba, 0xb7, 0xe1, 0xa4, 0x88, 0x30, 0x84, 0x9d, 0x13, 0x92, 0x1f, 0xab, 0x78,
	0xdc, 0xc8, 0x6e, 0x32, 0xd8, 0xfd, 0xcd, 0x87, 0x45, 0xbc, 0xe7, 0x50, 0xe5, 0x1c, 0x82, 0xee,
	0xae, 0xb3, 0x0a, 0x0f, 0x61, 0x6f, 0xa0, 0x1b, 0xe9, 0xfe, 0x04, 0xca, 0x94, 0xab, 0xd0, 0x12,
	0xd5, 0x17, 0x54, 0xa6, 0x76, 0x42, 0xa5, 0x34, 0x67, 0x0b, 0x7d, 0x0d, 0x15, 0x46, 0x24, 0x48,
	0x35, 0x51, 0x38, 0xab, 0x75, 0x77, 0x4d, 0xaf, 0x26, 0x65, 0x9f, 0x3d, 0x77, 0xd4, 0x8d, 0x2b,
	0xd8, 0xa6, 0xb5, 0xbb, 0xaa, 0x2e, 0x1c, 0x7b, 0x00, 0x27, 0x44, 0x3e, 0x3c, 0xea, 0x04, 0x96,
	0x57, 0x52, 0x9d, 0xc0, 0xca, 0xb6, 0x39, 0x5b, 0x17, 0x55, 0xf6, 0xcf, 0xe7, 0xf0, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0x97, 0x82, 0x94, 0x06, 0x0d, 0x00, 0x00,
}
