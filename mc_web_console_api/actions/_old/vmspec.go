package actions

import (
	"log"
	"mc_web_console_api/fwmodels"
	"mc_web_console_api/fwmodels/tumblebug/common"

	tbmcir "mc_web_console_api/fwmodels/tumblebug/mcir"
	tbmcis "mc_web_console_api/fwmodels/tumblebug/mcis"
	"mc_web_console_api/handler"
	"mc_web_console_api/models"
	"mc_web_console_api/models/views"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

// Get "/lookupvmspecs"
func VmSpecLookupList(c buffalo.Context) error {
	connectionName := &common.TbConnectionName{}

	paramProviderId := c.Params().Get("providerId")
	paramRegionName := c.Params().Get("regionName")
	log.Println("************provider ID : ", paramProviderId)
	log.Println("************regionName : ", paramRegionName)
	paramViewConnection := views.ViewCloudConnection{}
	paramViewConnection.ProviderID = paramProviderId
	paramViewConnection.RegionName = paramRegionName
	viewConnection, err := handler.GetAvailableConnection(paramViewConnection, c)
	if err != nil {
		return c.Render(500, r.JSON(map[string]interface{}{
			"error":  "there is no available connection",
			"status": "500",
		}))
	}
	connectionName.ConnectionName = viewConnection.ConnectionName
	cspVmSpecInfoList, respStatus := handler.LookupVmSpecInfoList(connectionName)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":       "success",
		"status":        respStatus.StatusCode,
		"CspVmSpecList": cspVmSpecInfoList,
	}))
}

// VmspecList default implementation.
func VmSpecList(c buffalo.Context) error {
	log.Println("GetVmSpecList : ")
	namespaceID := c.Session().Get("current_namespace_id").(string)
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	log.Println("namespaceID at vmspecList : ", namespaceID)
	optionParam := c.Params().Get("option")
	filterKeyParam := c.Params().Get("filterKey")
	filterValParam := c.Params().Get("filterVal")

	paramConnectionName := c.Params().Get("connectionName")

	if optionParam == "id" {
		vmSpecInfoList, respStatus := handler.GetVmSpecInfoListByID(namespaceID, filterKeyParam, filterValParam)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			}))
		}

		return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": namespaceID,
			"VmSpecList":         vmSpecInfoList,
		}))
	} else {
		log.Println("option 값이 없어서 여기로 간다.")

		if !strings.EqualFold(paramConnectionName, "") && strings.EqualFold(filterKeyParam, "") {
			filterKeyParam = "connectionName"
			filterValParam = paramConnectionName
		}

		vmSpecInfoList, respStatus := handler.GetVmSpecInfoListByOption(namespaceID, optionParam, filterKeyParam, filterValParam)
		log.Println(vmSpecInfoList)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			}))
		}

		return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": namespaceID,
			"VmSpecList":         vmSpecInfoList,
		}))
	}
}

func VmSpecGet(c buffalo.Context) error {
	namespaceID := c.Session().Get("current_namespace_id").(string)

	paramVmSpecId := c.Param("vmSpecId")
	vmSpecInfo, respStatus := handler.GetVmSpecInfoData(namespaceID, paramVmSpecId)

	paramViewConnection := views.ViewCloudConnection{}
	paramViewConnection.ConnectionName = vmSpecInfo.ConnectionName
	viewConnection, err := handler.GetViewConnection(paramViewConnection)
	if err != nil {
		// cb에서 정보는 가져왔으니 오류로 뱉지는 않기.
	} else {
		vmSpecInfo.ProviderID = viewConnection.ProviderID
		vmSpecInfo.ProviderName = viewConnection.ProviderID
		vmSpecInfo.RegionName = viewConnection.RegionName
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":            "success",
		"status":             respStatus.StatusCode,
		"DefaultNameSpaceID": namespaceID,
		"VmSpec":             vmSpecInfo,
	}))
}

func VmSpecReg(c buffalo.Context) error {
	namespaceID := c.Session().Get("current_namespace_id").(string)
	namespaceName := c.Session().Get("current_namespace").(string)

	vmSpecReq := &tbmcir.TbSpecReq{}
	vmSpecRegInfo := new(tbmcir.TbSpecInfo)
	resultVmSpecInfo := new(tbmcir.TbSpecInfo)

	paramVMSpecregisteringMethod := c.Param("specregisteringMethod")

	respStatus := fwmodels.WebStatus{}
	paramViewConnection := views.ViewCloudConnection{}

	if strings.EqualFold(paramVMSpecregisteringMethod, "registerWithInfo") {
		if err := c.Bind(vmSpecRegInfo); err != nil {
			log.Println(err)
			return c.Render(http.StatusBadRequest, r.JSON(err))
		}
		paramViewConnection.ProviderID = vmSpecRegInfo.ProviderName
		paramViewConnection.RegionName = vmSpecRegInfo.RegionName

	} else {
		if err := c.Bind(vmSpecReq); err != nil {
			log.Println(err)
			return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
				"error":  "fail",
				"status": "fail",
			}))
		}
		paramViewConnection.ProviderID = vmSpecReq.ProviderID
		paramViewConnection.RegionName = vmSpecReq.RegionName
		paramViewConnection.ZoneName = vmSpecReq.ZoneName
	}

	viewConnection, err := handler.GetAvailableConnection(paramViewConnection, c)

	log.Println("*************vmspeceq : ", vmSpecReq, "|*************namespaceID : ", namespaceID, "|**********viewconnection :", viewConnection, "|*************error : ", err)
	if err != nil {
		return c.Render(500, r.JSON(map[string]interface{}{
			"error":  "there is no available connection",
			"status": "500",
		}))
	}

	if strings.EqualFold(paramVMSpecregisteringMethod, "registerWithInfo") {
		vmSpecRegInfo.ConnectionName = viewConnection.ConnectionName
		resultVmSpecInfo, respStatus = handler.RegVmSpecWithInfo(namespaceID, paramVMSpecregisteringMethod, vmSpecRegInfo)
	} else {
		vmSpecReq.ConnectionName = viewConnection.ConnectionName
		resultVmSpecInfo, respStatus = handler.RegVmSpec(namespaceID, paramVMSpecregisteringMethod, vmSpecReq)
	}

	connectionMapping := &models.CloudConnectionMapping{}
	connectionMapping.Status = "C"
	connectionMapping.ResourceType = "vmspec"
	connectionMapping.ResourceID = resultVmSpecInfo.ID
	connectionMapping.ResourceName = resultVmSpecInfo.Name
	connectionMapping.CloudConnectionID = viewConnection.ID
	connectionMapping.CredentialID = viewConnection.CredentialID
	connectionMapping.NamespaceID = namespaceID
	connectionMapping.NamespaceName = namespaceName
	err = handler.SaveConnectionMapping(connectionMapping, c)
	if err != nil {
		// 실패시 생성한 vpc Resource 제거
		_, respStatus := handler.DelVMSpec(namespaceID, vmSpecReq.Name)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			}))
		}
		return c.Render(500, r.JSON(map[string]interface{}{
			"error":  err.Error(),
			"status": "500",
		}))
	}

	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
		"VMSpec":  resultVmSpecInfo,
	}))
}

func VmSpecDel(c buffalo.Context) error {
	namespaceID := c.Session().Get("current_namespace_id").(string)
	namespaceName := c.Session().Get("current_namespace").(string)

	paramVmSpecId := c.Param("vmSpecId")
	log.Println("***********paramVMSpecID : ")
	log.Println(paramVmSpecId)
	respMessage, respStatus := handler.DelVMSpec(namespaceID, paramVmSpecId)

	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}

	connectionMapping := &models.CloudConnectionMapping{}
	connectionMapping.NamespaceID = namespaceID
	connectionMapping.NamespaceName = namespaceName
	connectionMapping.ResourceType = "vmspec"
	connectionMapping.ResourceID = paramVmSpecId

	usedConnectionMapping, err := handler.GetUsedConnection(connectionMapping)

	if err != nil {
		// TODO : mapping table에 없을 때, 에러로 처리할 것인가?
		// return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
		// 	"error":  respStatus.Message,
		// 	"status": respStatus.StatusCode,
		// }))
	} else {
		// 삭제 Row 추가
		usedConnectionMapping.ID = uuid.UUID{}
		usedConnectionMapping.Status = "D"

		err = handler.SaveConnectionMapping(usedConnectionMapping, c)

		if err != nil {
			return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			}))
		}
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message": respMessage.Message,
		"status":  respMessage.StatusCode,
	}))
}

func FetchVmSpecList(c buffalo.Context) error {
	log.Println("FetchVMSpecList : ")

	namespaceID := c.Session().Get("current_namespace_id").(string)

	// vmSpecInfoList, respStatus := service.FetchVmSpecInfoList(defaultNameSpaceID)
	go handler.FetchVmSpecInfoListByAsync(namespaceID, c)
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message": "success",
		"status":  http.StatusOK,
	}))
}

func FilterVmSpecList(c buffalo.Context) error {
	log.Println("FilterVmspecList : ")

	namespaceID := c.Session().Get("current_namespace_id").(string)

	vmSpecRange := &tbmcir.FilterSpecsByRangeRequest{}
	if err := c.Bind(vmSpecRange); err != nil {
		log.Println(err)
		return c.Render(http.StatusBadRequest, r.JSON(err))
	}

	vmSpecInfoList, respStatus := handler.FilterVmSpecInfoListByRange(namespaceID, vmSpecRange)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}

	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":    "success",
		"status":     respStatus.StatusCode,
		"VmSpecList": vmSpecInfoList,
	}))
}

// 추천 VMSpec 목록 조회
func RecommendVmSpecList(c buffalo.Context) error {
	mcisDeploymentPlan := &tbmcis.DeploymentPlan{}
	if err := c.Bind(mcisDeploymentPlan); err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		}))
	}
	vmSpecList, _ := handler.GetRecommendVmSpecList(mcisDeploymentPlan)

	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":    "success",
		"status":     200,
		"VmSpecList": vmSpecList,
	}))
}
