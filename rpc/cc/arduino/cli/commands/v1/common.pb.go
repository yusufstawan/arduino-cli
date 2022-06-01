// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cc/arduino/cli/commands/v1/common.proto

package commands

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the instance.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Instance) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DownloadProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL of the download.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The file being downloaded.
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	// Total size of the file being downloaded.
	TotalSize int64 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// Size of the downloaded portion of the file.
	Downloaded int64 `protobuf:"varint,4,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
	// Whether the download is complete.
	Completed bool `protobuf:"varint,5,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *DownloadProgress) Reset() {
	*x = DownloadProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadProgress) ProtoMessage() {}

func (x *DownloadProgress) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadProgress.ProtoReflect.Descriptor instead.
func (*DownloadProgress) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadProgress) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DownloadProgress) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *DownloadProgress) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *DownloadProgress) GetDownloaded() int64 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

func (x *DownloadProgress) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TaskProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description of the task.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional information about the task.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Whether the task is complete.
	Completed bool `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	// Amount in percent of the task completion (optional)
	Percent float32 `protobuf:"fixed32,4,opt,name=percent,proto3" json:"percent,omitempty"`
}

func (x *TaskProgress) Reset() {
	*x = TaskProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskProgress) ProtoMessage() {}

func (x *TaskProgress) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskProgress.ProtoReflect.Descriptor instead.
func (*TaskProgress) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *TaskProgress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskProgress) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TaskProgress) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *TaskProgress) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

type Programmer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Programmer) Reset() {
	*x = Programmer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Programmer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Programmer) ProtoMessage() {}

func (x *Programmer) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Programmer.ProtoReflect.Descriptor instead.
func (*Programmer) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *Programmer) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Programmer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Programmer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Platform ID (e.g., `arduino:avr`).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Version of the platform.
	Installed string `protobuf:"bytes,2,opt,name=installed,proto3" json:"installed,omitempty"`
	// Newest available version of the platform.
	Latest string `protobuf:"bytes,3,opt,name=latest,proto3" json:"latest,omitempty"`
	// Name used to identify the platform to humans (e.g., "Arduino AVR Boards").
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Maintainer of the platform's package.
	Maintainer string `protobuf:"bytes,5,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	// A URL provided by the author of the platform's package, intended to point
	// to their website.
	Website string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	// Email of the maintainer of the platform's package.
	Email string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	// List of boards provided by the platform. If the platform is installed,
	// this is the boards listed in the platform's boards.txt. If the platform is
	// not installed, this is an arbitrary list of board names provided by the
	// platform author for display and may not match boards.txt.
	Boards []*Board `protobuf:"bytes,8,rep,name=boards,proto3" json:"boards,omitempty"`
	// If true this Platform has been installed manually in the user' sketchbook
	// hardware folder
	ManuallyInstalled bool `protobuf:"varint,9,opt,name=manually_installed,json=manuallyInstalled,proto3" json:"manually_installed,omitempty"`
	// If true this Platform has been deprecated
	Deprecated bool `protobuf:"varint,10,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Platform) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Platform) GetInstalled() string {
	if x != nil {
		return x.Installed
	}
	return ""
}

func (x *Platform) GetLatest() string {
	if x != nil {
		return x.Latest
	}
	return ""
}

func (x *Platform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

func (x *Platform) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Platform) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Platform) GetBoards() []*Board {
	if x != nil {
		return x.Boards
	}
	return nil
}

func (x *Platform) GetManuallyInstalled() bool {
	if x != nil {
		return x.ManuallyInstalled
	}
	return false
}

func (x *Platform) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

type InstalledPlatformReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Platform ID (e.g., `arduino:avr`).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Version of the platform.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Installation directory of the platform
	InstallDir string `protobuf:"bytes,3,opt,name=install_dir,json=installDir,proto3" json:"install_dir,omitempty"`
	// 3rd party platform URL
	PackageUrl string `protobuf:"bytes,4,opt,name=package_url,json=packageUrl,proto3" json:"package_url,omitempty"`
}

func (x *InstalledPlatformReference) Reset() {
	*x = InstalledPlatformReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstalledPlatformReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstalledPlatformReference) ProtoMessage() {}

func (x *InstalledPlatformReference) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstalledPlatformReference.ProtoReflect.Descriptor instead.
func (*InstalledPlatformReference) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *InstalledPlatformReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstalledPlatformReference) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InstalledPlatformReference) GetInstallDir() string {
	if x != nil {
		return x.InstallDir
	}
	return ""
}

func (x *InstalledPlatformReference) GetPackageUrl() string {
	if x != nil {
		return x.PackageUrl
	}
	return ""
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name used to identify the board to humans.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fully qualified board name used to identify the board to machines. The FQBN
	// is only available for installed boards.
	Fqbn string `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *Board) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Board) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name used to identify the profile within the sketch.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// FQBN specified in the profile.
	Fqbn string `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP(), []int{7}
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

var File_cc_arduino_cli_commands_v1_common_proto protoreflect.FileDescriptor

var file_cc_arduino_cli_commands_v1_common_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x63, 0x2e, 0x61, 0x72,
	0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x74, 0x0a, 0x0c, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22,
	0x4c, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbe, 0x02,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e,
	0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x6c, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x88,
	0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x05, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x22, 0x31, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x62,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x42, 0x48, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x64, 0x75,
	0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63,
	0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cc_arduino_cli_commands_v1_common_proto_rawDescOnce sync.Once
	file_cc_arduino_cli_commands_v1_common_proto_rawDescData = file_cc_arduino_cli_commands_v1_common_proto_rawDesc
)

func file_cc_arduino_cli_commands_v1_common_proto_rawDescGZIP() []byte {
	file_cc_arduino_cli_commands_v1_common_proto_rawDescOnce.Do(func() {
		file_cc_arduino_cli_commands_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_arduino_cli_commands_v1_common_proto_rawDescData)
	})
	return file_cc_arduino_cli_commands_v1_common_proto_rawDescData
}

var file_cc_arduino_cli_commands_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cc_arduino_cli_commands_v1_common_proto_goTypes = []interface{}{
	(*Instance)(nil),                   // 0: cc.arduino.cli.commands.v1.Instance
	(*DownloadProgress)(nil),           // 1: cc.arduino.cli.commands.v1.DownloadProgress
	(*TaskProgress)(nil),               // 2: cc.arduino.cli.commands.v1.TaskProgress
	(*Programmer)(nil),                 // 3: cc.arduino.cli.commands.v1.Programmer
	(*Platform)(nil),                   // 4: cc.arduino.cli.commands.v1.Platform
	(*InstalledPlatformReference)(nil), // 5: cc.arduino.cli.commands.v1.InstalledPlatformReference
	(*Board)(nil),                      // 6: cc.arduino.cli.commands.v1.Board
	(*Profile)(nil),                    // 7: cc.arduino.cli.commands.v1.Profile
}
var file_cc_arduino_cli_commands_v1_common_proto_depIdxs = []int32{
	6, // 0: cc.arduino.cli.commands.v1.Platform.boards:type_name -> cc.arduino.cli.commands.v1.Board
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cc_arduino_cli_commands_v1_common_proto_init() }
func file_cc_arduino_cli_commands_v1_common_proto_init() {
	if File_cc_arduino_cli_commands_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadProgress); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskProgress); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Programmer); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstalledPlatformReference); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_cc_arduino_cli_commands_v1_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_cc_arduino_cli_commands_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cc_arduino_cli_commands_v1_common_proto_goTypes,
		DependencyIndexes: file_cc_arduino_cli_commands_v1_common_proto_depIdxs,
		MessageInfos:      file_cc_arduino_cli_commands_v1_common_proto_msgTypes,
	}.Build()
	File_cc_arduino_cli_commands_v1_common_proto = out.File
	file_cc_arduino_cli_commands_v1_common_proto_rawDesc = nil
	file_cc_arduino_cli_commands_v1_common_proto_goTypes = nil
	file_cc_arduino_cli_commands_v1_common_proto_depIdxs = nil
}
