package actions

import (
	"log"
	"mc_web_console_api/fwmodels"
	"mc_web_console_api/fwmodels/dragonfly"
	"mc_web_console_api/handler"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func McisMonitoringDurationMetric(c buffalo.Context) error {
	log.Println("McisMonitoringData")
	namespaceID := c.Session().Get("current_namespace_id").(string)

	mcisMonitoring := &dragonfly.McisMonitoringOnDemandInfoReq{}
	if err := c.Bind(mcisMonitoring); err != nil {
		log.Println(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		}))
	}
	mcisMonitoring.NameSpaceID = namespaceID

	returnMcisMonitoringInfo, respStatus := handler.GetMcisMonitoringDurationInfo(mcisMonitoring.NameSpaceID, mcisMonitoring.McisID, mcisMonitoring.MetricName)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":            "success",
		"status":             respStatus.StatusCode,
		"McisMonitoringInfo": returnMcisMonitoringInfo,
	}))
}

// MCIS의 monitoring data 조회 : 가져오는 것은 GET이지만 설정 항목이 많아 POST로 받음
// 모니터링정보는 DF를 호출하지만 해당기능이 DF에 완전하지 않아 TB에 있는 API 호출
func McisMonitoringOndemandMetric(c buffalo.Context) error {
	log.Println("McisMonitoringData")
	namespaceID := c.Session().Get("current_namespace_id").(string)

	mcisMonitoring := &dragonfly.McisMonitoringOnDemandInfoReq{}
	if err := c.Bind(mcisMonitoring); err != nil {
		log.Println(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		}))
	}

	mcisMonitoring.NameSpaceID = namespaceID

	returnVMMonitoringInfo, respStatus := handler.GetMcisOnDemandMonitoringMetricInfo(mcisMonitoring)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":              "success",
		"status":               respStatus.StatusCode,
		"McisVMMonitoringInfo": returnVMMonitoringInfo,
	}))
}

// MCIS내 특정 vm monitoring data 조회 : 가져오는 것은 GET이지만 설정 항목이 많아 POST로 받음
// option 으로 구분 ( 이전에는 3개의 다른 function이었음)
func McisVmMonitoringOndemand(c buffalo.Context) error {
	log.Println("McisVmMonitoringOndemand")
	namespaceID := c.Session().Get("current_namespace_id").(string)

	optionParam := c.Params().Get("option")

	vmMonitoringReq := &dragonfly.VmMonitoringReq{}
	if err := c.Bind(vmMonitoringReq); err != nil {
		log.Println(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		}))
	}

	returnVMMonitoringInfo := &dragonfly.VmMonitoringOnDemandInfo{}
	respStatus := fwmodels.WebStatus{}

	vmMonitoringReq.NameSpaceID = namespaceID

	if optionParam == "NetworkPacket" {
		returnVMMonitoringInfo, respStatus = handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
	} else if optionParam == "Metric" {
		returnVMMonitoringInfo, respStatus = handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
	} else if optionParam == "ProcessUsage" {
		returnVMMonitoringInfo, respStatus = handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
	} else {
		respStatus.Message = "optionParam is invalid " + optionParam
		respStatus.StatusCode = 500
	}

	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":              "success",
		"status":               respStatus.StatusCode,
		"McisVMMonitoringInfo": returnVMMonitoringInfo,
	}))
}

// MCIS내 특정 vm monitoring data 조회 : 가져오는 것은 GET이지만 설정 항목이 많아 POST로 받음
// func  McisVmMonitoringOndemandMetric(c buffalo.Context) error {
// 	log.Println("McisVmMonitoringData")
// 	namespaceID := c.Session().Get("current_namespace_id").(string)

// 	vmMonitoringReq := &dragonfly.VmMonitoringReq{}
// 	if err := c.Bind(vmMonitoringReq); err != nil {
// 		log.Println(err)
// 		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
// 			"message": "fail",
// 			"status":  "fail",
// 		}))
// 	}

// 	vmMonitoringReq.NameSpaceID = namespaceID

// 	returnVMMonitoringInfo, respStatus := handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
// 	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
// 		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
// 			"error":  respStatus.Message,
// 			"status": respStatus.StatusCode,
// 		}))
// 	}
// 	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
// 		"message":              "success",
// 		"status":               respStatus.StatusCode,
// 		"McisVMMonitoringInfo": returnVMMonitoringInfo,
// 	}))
// }

// func  McisVmMonitoringOndemandNetworkPacket(c buffalo.Context) error {
// 	log.Println("McisVmMonitoringOndemandNetworkPacket")
// 	namespaceID := c.Session().Get("current_namespace_id").(string)

// 	vmMonitoringReq := &dragonfly.VmMonitoringReq{}
// 	if err := c.Bind(vmMonitoringReq); err != nil {
// 		log.Println(err)
// 		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
// 			"message": "fail",
// 			"status":  "fail",
// 		}))
// 	}

// 	vmMonitoringReq.NameSpaceID = namespaceID

// 	returnVMMonitoringInfo, respStatus := handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
// 	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
// 		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
// 			"error":  respStatus.Message,
// 			"status": respStatus.StatusCode,
// 		}))
// 	}
// 	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
// 		"message":              "success",
// 		"status":               respStatus.StatusCode,
// 		"McisVMMonitoringInfo": returnVMMonitoringInfo,
// 	}))
// }

// func  McisVmMonitoringOndemandProcessUsage(c buffalo.Context) error {
// 	log.Println("McisVmMonitoringOndemandProcessUsage")
// 	namespaceID := c.Session().Get("current_namespace_id").(string)

// 	vmMonitoringReq := &dragonfly.VmMonitoringReq{}
// 	if err := c.Bind(vmMonitoringReq); err != nil {
// 		log.Println(err)
// 		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
// 			"message": "fail",
// 			"status":  "fail",
// 		}))
// 	}

// 	vmMonitoringReq.NameSpaceID = namespaceID

// 	returnVMMonitoringInfo, respStatus := handler.GetMcisVmOnDemandMonitoringMetricInfo(vmMonitoringReq)
// 	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
// 		return c.Render(respStatus.StatusCode, r.JSON(map[string]interface{}{
// 			"error":  respStatus.Message,
// 			"status": respStatus.StatusCode,
// 		}))
// 	}
// 	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
// 		"message":              "success",
// 		"status":               respStatus.StatusCode,
// 		"McisVMMonitoringInfo": returnVMMonitoringInfo,
// 	}))
// }
