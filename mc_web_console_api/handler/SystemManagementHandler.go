package handler

import (
	"encoding/json"
	"fmt"
	"log"

	// "math"
	"net/http"
	// "strconv"
	// "sync"

	//"github.com/davecgh/go-spew/spew"

	// "mc_web_console_api/fwmodels/spider"
	// "mc_web_console_api/fwmodels/tumblebug"
	"mc_web_console_api/fwmodels"
	tbcommon "mc_web_console_api/fwmodels/tumblebug/common"

	// tbmcir "mc_web_console_api/fwmodels/tumblebug/mcir"
	// tbmcis "mc_web_console_api/fwmodels/tumblebug/mcis"

	util "mc_web_console_api/util"
	// "github.com/labstack/echo"
)

func GetCheckResourceExistence(nameSpaceID string, resourceType string, resourceId string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	var originalUrl = "/{nsId}/checkResource/{resourceType}/{resourceId}"

	var paramMapper = make(map[string]string)
	paramMapper["{nsId}"] = nameSpaceID
	paramMapper["{resourceType}"] = resourceType
	paramMapper["{resourceId}"] = resourceId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm/" + vmID

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func GetHealth() (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Check Tumblebug alive start")
	var originalUrl = "/health"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func DelObject(optionParam string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Delete object start")
	var originalUrl = "/object"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam + "?option=" + optionParam
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func GetObject(optionParam string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Get object start")
	var originalUrl = "/object"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	//"?key=" + key
	optionParamVal := ""
	// install, init, cpus, cpum, memR, memW, fioR, fioW, dbR, dbW, rtt, mrtt, clean
	if optionParam != "" {
		optionParamVal = "?key=" + optionParam
	}

	url := util.TUMBLEBUG + urlParam + optionParamVal
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func DelObjects(optionParam string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Delete objects start")
	var originalUrl = "/objects"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	//"?key=" + key
	optionParamVal := ""
	// install, init, cpus, cpum, memR, memW, fioR, fioW, dbR, dbW, rtt, mrtt, clean
	if optionParam != "" {
		optionParamVal = "?key=" + optionParam
	}

	url := util.TUMBLEBUG + urlParam + optionParamVal
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func GetObjectList(optionParam string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Get object list start")
	var originalUrl = "/objects"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	//"?key=" + key
	optionParamVal := ""
	// install, init, cpus, cpum, memR, memW, fioR, fioW, dbR, dbW, rtt, mrtt, clean
	if optionParam != "" {
		optionParamVal = "?key=" + optionParam
	}

	url := util.TUMBLEBUG + urlParam + optionParamVal
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}

func GetAPIDocument(optionParam string) (tbcommon.TbSimpleMsg, fwmodels.WebStatus) {
	fmt.Println("Get API document start")
	var originalUrl = "/swaggerActive"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	resultInfo := tbcommon.TbSimpleMsg{}

	if err != nil {
		return resultInfo, fwmodels.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultInfo)
	log.Println(resultInfo)
	log.Println("ResultMessage : " + resultInfo.Message)

	if respStatus != 200 && respStatus != 201 {
		return resultInfo, fwmodels.WebStatus{StatusCode: respStatus, Message: resultInfo.Message}
	}

	return resultInfo, fwmodels.WebStatus{StatusCode: respStatus}
}
