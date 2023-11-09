package mcir

type TbSpecInfo struct {
	AssociatedObjectList	[]string	`json:"associatedObjectList"`
	EbsBwMbps		`json:"ebsBwMbps"`
	EvaluationScore10		`json:"evaluationScore10"`
	MaxTotalStorageTiB		`json:"maxTotalStorageTiB"`
	OsType	string	`json:"osType"`
	StorageGiB		`json:"storageGiB"`
	Description	string	`json:"description"`
	EvaluationScore03		`json:"evaluationScore03"`
	EvaluationScore07		`json:"evaluationScore07"`
	ConnectionName	string	`json:"connectionName"`
	CostPerHour		`json:"costPerHour"`
	EvaluationScore01		`json:"evaluationScore01"`
	EvaluationScore02		`json:"evaluationScore02"`
	EvaluationScore09		`json:"evaluationScore09"`
	GpuMemGiB		`json:"gpuMemGiB"`
	EvaluationScore05		`json:"evaluationScore05"`
	EvaluationScore06		`json:"evaluationScore06"`
	EvaluationScore08		`json:"evaluationScore08"`
	MemGiB		`json:"memGiB"`
	Namespace	string	`json:"namespace"`
	NumGpu		`json:"numGpu"`
	GpuP2p	string	`json:"gpuP2p"`
	Name	string	`json:"name"`
	NetBwGbps		`json:"netBwGbps"`
	NumStorage		`json:"numStorage"`
	EvaluationScore04		`json:"evaluationScore04"`
	EvaluationStatus	string	`json:"evaluationStatus"`
	GpuModel	string	`json:"gpuModel"`
	MaxNumStorage		`json:"maxNumStorage"`
	NumCore		`json:"numCore"`
	NumvCPU		`json:"numvCPU"`
	CspSpecName	string	`json:"cspSpecName"`
	Id	string	`json:"id"`
	IsAutoGenerated		`json:"isAutoGenerated"`
	OrderInFilteredResult		`json:"orderInFilteredResult"`
}