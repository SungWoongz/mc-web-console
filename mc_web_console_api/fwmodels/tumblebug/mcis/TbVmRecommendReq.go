package mcis

import (
	tbcommon "mc_web_console_api/fwmodels/tumblebug/common"
)

type TbVmRecommendReq struct {
	DiskSize       string                `json:"diskSize"`
	MaxResultNum   string                `json:"maxResultNum"`
	MemorySize     string                `json:"memorySize"`
	PlacementAlgo  string                `json:"placementAlgo"`
	PlacementParam []tbcommon.TbKeyValue `json:"placementParam"`

	RequestName string `json:"requestName"`
	VcpuSize    string `json:"vcpuSize"`
}
