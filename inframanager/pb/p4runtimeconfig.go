// Copyright (c) 2022 Intel Corporation.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: infraconfig.proto

package pb

type Chip struct {
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChipFamily      string `protobuf:"bytes,2,opt,name=chip_family,json=chipFamily,proto3" json:"chip_family,omitempty"`
	Instance        uint32 `protobuf:"varint,3,opt,name=instance,proto3" json:"instance,omitempty"`
	PcieSysfsPrefix string `protobuf:"bytes,4,opt,name=pcie_sysfs_prefix,json=pcieSysfsPrefix,proto3" json:"pcie_sysfs_prefix,omitempty"`
	PcieBdf         string `protobuf:"bytes,5,opt,name=pcie_bdf,json=pcieBdf,proto3" json:"pcie_bdf,omitempty"`
	PcieIntMode     uint32 `protobuf:"varint,6,opt,name=pcie_int_mode,json=pcieIntMode,proto3" json:"pcie_int_mode,omitempty"`
	IommuGrpNum     uint32 `protobuf:"varint,7,opt,name=iommu_grp_num,json=iommuGrpNum,proto3" json:"iommu_grp_num,omitempty"`
	SdsFwPath       string `protobuf:"bytes,8,opt,name=sds_fw_path,json=sdsFwPath,proto3" json:"sds_fw_path,omitempty"`
}

func (x *Chip) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chip) GetChipFamily() string {
	if x != nil {
		return x.ChipFamily
	}
	return ""
}

func (x *Chip) GetInstance() uint32 {
	if x != nil {
		return x.Instance
	}
	return 0
}

func (x *Chip) GetPcieSysfsPrefix() string {
	if x != nil {
		return x.PcieSysfsPrefix
	}
	return ""
}

func (x *Chip) GetPcieBdf() string {
	if x != nil {
		return x.PcieBdf
	}
	return ""
}

func (x *Chip) GetPcieIntMode() uint32 {
	if x != nil {
		return x.PcieIntMode
	}
	return 0
}

func (x *Chip) GetIommuGrpNum() uint32 {
	if x != nil {
		return x.IommuGrpNum
	}
	return 0
}

func (x *Chip) GetSdsFwPath() string {
	if x != nil {
		return x.SdsFwPath
	}
	return ""
}

type P4Device struct {
	DeviceId   uint32       `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	P4Programs []*P4Program `protobuf:"bytes,2,rep,name=p4_programs,json=p4Programs,proto3" json:"p4_programs,omitempty"`
	Agent0     string       `protobuf:"bytes,3,opt,name=agent0,proto3" json:"agent0,omitempty"`
}

func (x *P4Device) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *P4Device) GetP4Programs() []*P4Program {
	if x != nil {
		return x.P4Programs
	}
	return nil
}

func (x *P4Device) GetAgent0() string {
	if x != nil {
		return x.Agent0
	}
	return ""
}

type P4Program struct {
	ProgramName string        `protobuf:"bytes,1,opt,name=program_name,json=programName,proto3" json:"program_name,omitempty"`
	BfrtConfig  string        `protobuf:"bytes,2,opt,name=bfrt_config,json=bfrtConfig,proto3" json:"bfrt_config,omitempty"`
	P4Pipelines []*P4Pipeline `protobuf:"bytes,3,rep,name=p4_pipelines,json=p4Pipelines,proto3" json:"p4_pipelines,omitempty"`
}

func (x *P4Program) GetProgramName() string {
	if x != nil {
		return x.ProgramName
	}
	return ""
}

func (x *P4Program) GetBfrtConfig() string {
	if x != nil {
		return x.BfrtConfig
	}
	return ""
}

func (x *P4Program) GetP4Pipelines() []*P4Pipeline {
	if x != nil {
		return x.P4Pipelines
	}
	return nil
}

type P4Pipeline struct {
	P4PipelineName string   `protobuf:"bytes,1,opt,name=p4_pipeline_name,json=p4PipelineName,proto3" json:"p4_pipeline_name,omitempty"`
	Context        string   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Config         string   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	PipeScope      []uint32 `protobuf:"varint,4,rep,packed,name=pipe_scope,json=pipeScope,proto3" json:"pipe_scope,omitempty"`
	Path           string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *P4Pipeline) GetP4PipelineName() string {
	if x != nil {
		return x.P4PipelineName
	}
	return ""
}

func (x *P4Pipeline) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *P4Pipeline) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *P4Pipeline) GetPipeScope() []uint32 {
	if x != nil {
		return x.PipeScope
	}
	return nil
}

func (x *P4Pipeline) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type P4RuntimeConfig struct {
	ChipList  []*Chip     `protobuf:"bytes,1,rep,name=chip_list,json=chipList,proto3" json:"chip_list,omitempty"`
	Instance  uint32      `protobuf:"varint,2,opt,name=instance,proto3" json:"instance,omitempty"`
	P4Devices []*P4Device `protobuf:"bytes,3,rep,name=p4_devices,json=p4Devices,proto3" json:"p4_devices,omitempty"`
}

func (x *P4RuntimeConfig) GetChipList() []*Chip {
	if x != nil {
		return x.ChipList
	}
	return nil
}

func (x *P4RuntimeConfig) GetInstance() uint32 {
	if x != nil {
		return x.Instance
	}
	return 0
}

func (x *P4RuntimeConfig) GetP4Devices() []*P4Device {
	if x != nil {
		return x.P4Devices
	}
	return nil
}

var file_infraconfig_proto_goTypes = []interface{}{
	(*Chip)(nil),            // 0: pb.Chip
	(*P4Device)(nil),        // 1: pb.P4Device
	(*P4Program)(nil),       // 2: pb.P4Program
	(*P4Pipeline)(nil),      // 3: pb.P4Pipeline
	(*P4RuntimeConfig)(nil), // 4: pb.P4RuntimeConfig
}
