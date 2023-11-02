package mcir

type TbSecurityGroupReq struct {
	ConnectionName     string               `json:"connectionName"`
	Description        string               `json:"description"`
	FirewallRules      []TbFirewallRuleInfo `json:"firewallRules"`
	Name               string               `json:"name"`
	VNetID             string               `json:"vNetId"`
	CspSecurityGroupId string               `json:"cspSecurityGroupId"`

	ProviderID   string `json:"providerId"`
	ProviderName string `json:"providerName"`
	RegionName   string `json:"regionName"`
	ZoneName     string `json:"zoneName"`
}
