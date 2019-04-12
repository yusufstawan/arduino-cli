// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compile.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CompileReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketchPath,proto3" json:"sketchPath,omitempty"`
	ShowProperties       bool      `protobuf:"varint,4,opt,name=showProperties,proto3" json:"showProperties,omitempty"`
	Preprocess           bool      `protobuf:"varint,5,opt,name=preprocess,proto3" json:"preprocess,omitempty"`
	BuildCachePath       string    `protobuf:"bytes,6,opt,name=buildCachePath,proto3" json:"buildCachePath,omitempty"`
	BuildPath            string    `protobuf:"bytes,7,opt,name=buildPath,proto3" json:"buildPath,omitempty"`
	BuildProperties      []string  `protobuf:"bytes,8,rep,name=buildProperties,proto3" json:"buildProperties,omitempty"`
	Warnings             string    `protobuf:"bytes,9,opt,name=warnings,proto3" json:"warnings,omitempty"`
	Verbose              bool      `protobuf:"varint,10,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Quiet                bool      `protobuf:"varint,11,opt,name=quiet,proto3" json:"quiet,omitempty"`
	VidPid               string    `protobuf:"bytes,12,opt,name=vidPid,proto3" json:"vidPid,omitempty"`
	ExportFile           string    `protobuf:"bytes,13,opt,name=exportFile,proto3" json:"exportFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CompileReq) Reset()         { *m = CompileReq{} }
func (m *CompileReq) String() string { return proto.CompactTextString(m) }
func (*CompileReq) ProtoMessage()    {}
func (*CompileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_b05aa4b7d1cad046, []int{0}
}
func (m *CompileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileReq.Unmarshal(m, b)
}
func (m *CompileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileReq.Marshal(b, m, deterministic)
}
func (dst *CompileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileReq.Merge(dst, src)
}
func (m *CompileReq) XXX_Size() int {
	return xxx_messageInfo_CompileReq.Size(m)
}
func (m *CompileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompileReq proto.InternalMessageInfo

func (m *CompileReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CompileReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *CompileReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *CompileReq) GetShowProperties() bool {
	if m != nil {
		return m.ShowProperties
	}
	return false
}

func (m *CompileReq) GetPreprocess() bool {
	if m != nil {
		return m.Preprocess
	}
	return false
}

func (m *CompileReq) GetBuildCachePath() string {
	if m != nil {
		return m.BuildCachePath
	}
	return ""
}

func (m *CompileReq) GetBuildPath() string {
	if m != nil {
		return m.BuildPath
	}
	return ""
}

func (m *CompileReq) GetBuildProperties() []string {
	if m != nil {
		return m.BuildProperties
	}
	return nil
}

func (m *CompileReq) GetWarnings() string {
	if m != nil {
		return m.Warnings
	}
	return ""
}

func (m *CompileReq) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *CompileReq) GetQuiet() bool {
	if m != nil {
		return m.Quiet
	}
	return false
}

func (m *CompileReq) GetVidPid() string {
	if m != nil {
		return m.VidPid
	}
	return ""
}

func (m *CompileReq) GetExportFile() string {
	if m != nil {
		return m.ExportFile
	}
	return ""
}

type CompileResp struct {
	OutStream            []byte            `protobuf:"bytes,1,opt,name=out_stream,json=outStream,proto3" json:"out_stream,omitempty"`
	ErrStream            []byte            `protobuf:"bytes,2,opt,name=err_stream,json=errStream,proto3" json:"err_stream,omitempty"`
	DownloadProgress     *DownloadProgress `protobuf:"bytes,3,opt,name=download_progress,json=downloadProgress,proto3" json:"download_progress,omitempty"`
	TaskProgress         *TaskProgress     `protobuf:"bytes,4,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CompileResp) Reset()         { *m = CompileResp{} }
func (m *CompileResp) String() string { return proto.CompactTextString(m) }
func (*CompileResp) ProtoMessage()    {}
func (*CompileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_b05aa4b7d1cad046, []int{1}
}
func (m *CompileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileResp.Unmarshal(m, b)
}
func (m *CompileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileResp.Marshal(b, m, deterministic)
}
func (dst *CompileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileResp.Merge(dst, src)
}
func (m *CompileResp) XXX_Size() int {
	return xxx_messageInfo_CompileResp.Size(m)
}
func (m *CompileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileResp.DiscardUnknown(m)
}

var xxx_messageInfo_CompileResp proto.InternalMessageInfo

func (m *CompileResp) GetOutStream() []byte {
	if m != nil {
		return m.OutStream
	}
	return nil
}

func (m *CompileResp) GetErrStream() []byte {
	if m != nil {
		return m.ErrStream
	}
	return nil
}

func (m *CompileResp) GetDownloadProgress() *DownloadProgress {
	if m != nil {
		return m.DownloadProgress
	}
	return nil
}

func (m *CompileResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileReq)(nil), "arduino.CompileReq")
	proto.RegisterType((*CompileResp)(nil), "arduino.CompileResp")
}

func init() { proto.RegisterFile("compile.proto", fileDescriptor_compile_b05aa4b7d1cad046) }

var fileDescriptor_compile_b05aa4b7d1cad046 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x18, 0x84, 0xd5, 0x4d, 0xb7, 0x4d, 0xfe, 0xa6, 0xc0, 0x5a, 0x80, 0xcc, 0x0a, 0xd0, 0xaa, 0x42,
	0x68, 0x2f, 0x9b, 0x95, 0xe0, 0xc6, 0x91, 0x45, 0x95, 0xb8, 0x45, 0x81, 0x13, 0x97, 0xca, 0x49,
	0x4c, 0x63, 0x35, 0xb5, 0x53, 0xdb, 0x69, 0x79, 0x27, 0x1e, 0x86, 0x57, 0xc2, 0xf9, 0x93, 0x26,
	0x55, 0x4e, 0xed, 0x7c, 0x33, 0x1e, 0x3b, 0xfe, 0x0d, 0xcb, 0x4c, 0xed, 0x2b, 0x51, 0xf2, 0xa8,
	0xd2, 0xca, 0x2a, 0x32, 0x67, 0x3a, 0xaf, 0x85, 0x54, 0xb7, 0xa1, 0xe3, 0x7b, 0x25, 0x5b, 0xbc,
	0xfa, 0xeb, 0x01, 0x3c, 0xb5, 0xc1, 0x84, 0x1f, 0xc8, 0x03, 0xf8, 0x42, 0x1a, 0xcb, 0x64, 0xc6,
	0xe9, 0xe4, 0x6e, 0x72, 0xbf, 0xf8, 0x74, 0x13, 0x75, 0x0b, 0xa3, 0xef, 0x9d, 0x91, 0xf4, 0x11,
	0x42, 0x60, 0xfa, 0xfb, 0x90, 0x4a, 0x7a, 0xe5, 0xa2, 0x41, 0x82, 0xff, 0xc9, 0x7b, 0x00, 0xb3,
	0xe3, 0x36, 0x2b, 0x62, 0x66, 0x0b, 0xea, 0xa1, 0x73, 0x41, 0xc8, 0x47, 0x78, 0x66, 0x0a, 0x75,
	0x8a, 0xb5, 0xaa, 0xb8, 0xb6, 0x82, 0x1b, 0x3a, 0x75, 0x19, 0x3f, 0x19, 0xd1, 0xa6, 0xa7, 0xd2,
	0xdc, 0x9d, 0x32, 0xe3, 0xc6, 0xd0, 0x6b, 0xcc, 0x5c, 0x90, 0xa6, 0x27, 0xad, 0x45, 0x99, 0x3f,
	0xb1, 0xac, 0xe0, 0xb8, 0xd7, 0x0c, 0xf7, 0x1a, 0x51, 0xf2, 0x16, 0x02, 0x24, 0x18, 0x99, 0x63,
	0x64, 0x00, 0xe4, 0x1e, 0x9e, 0xb7, 0x62, 0x38, 0x8e, 0x7f, 0xe7, 0xb9, 0xcc, 0x18, 0x93, 0x5b,
	0xf0, 0x4f, 0x4c, 0x4b, 0x21, 0xb7, 0x86, 0x06, 0x58, 0xd3, 0x6b, 0x42, 0x61, 0x7e, 0xe4, 0x3a,
	0x55, 0x86, 0x53, 0xc0, 0x83, 0x9e, 0x25, 0x79, 0x09, 0xd7, 0x87, 0x5a, 0x70, 0x4b, 0x17, 0xc8,
	0x5b, 0x41, 0x5e, 0xc3, 0xec, 0x28, 0xf2, 0x58, 0xe4, 0x34, 0xc4, 0xa6, 0x4e, 0x35, 0xdf, 0xcc,
	0xff, 0x54, 0x4a, 0xdb, 0xb5, 0x9b, 0x07, 0x5d, 0xb6, 0x77, 0x37, 0x90, 0xd5, 0xbf, 0x09, 0x2c,
	0xfa, 0x69, 0x99, 0x8a, 0xbc, 0x03, 0x50, 0xb5, 0xdd, 0x18, 0xab, 0x39, 0xdb, 0xe3, 0xc0, 0xc2,
	0x24, 0x70, 0xe4, 0x07, 0x82, 0xc6, 0xe6, 0x5a, 0x9f, 0xed, 0xab, 0xd6, 0x76, 0xa4, 0xb3, 0xd7,
	0x70, 0x93, 0xab, 0x93, 0x2c, 0x15, 0xcb, 0x37, 0xee, 0x56, 0xb7, 0xba, 0xb9, 0x68, 0x0f, 0xa7,
	0xfe, 0xa6, 0x9f, 0xfa, 0xb7, 0x2e, 0x11, 0x77, 0x81, 0xe4, 0x45, 0x3e, 0x22, 0xe4, 0x0b, 0x2c,
	0x2d, 0x33, 0xbb, 0xa1, 0x63, 0x8a, 0x1d, 0xaf, 0xfa, 0x8e, 0x9f, 0xce, 0xed, 0xd7, 0x87, 0xf6,
	0x42, 0x7d, 0xfd, 0xf0, 0x6b, 0xb5, 0x15, 0xb6, 0xa8, 0xd3, 0xc8, 0x3d, 0xcb, 0xc7, 0x6e, 0xc1,
	0xf9, 0xf7, 0x21, 0x2b, 0xc5, 0xa3, 0xae, 0xb2, 0x74, 0x86, 0x8f, 0xf5, 0xf3, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0x3a, 0xba, 0x6b, 0xd4, 0x02, 0x00, 0x00,
}
