console.log("mcisserverprint.js");


function mcisLifeCycle(){

}


// message를 표현할 alert 창
function commonAlert(alertMessage) {
    console.log(alertMessage);
    // $('#alertText').text(alertMessage);
    $('#alertText').html(alertMessage);
    $("#alertArea").modal();
}
// alert창 닫기
function commonAlertClose() {
    $("#alertArea").modal("hide");
}


// message를 표현할 alert 창 : 생성 결과를 표시하는 Alert로 commonAlert와 동일하나 닫힐 때 event처리를 할 수 있게
// 사용할 화면에서 $('#alertResultArea').on('hidden.bs.modal', function () {  // 수행할 일 또는 함수 호출  // })
function commonResultAlert(alertMessage) {
    console.log(alertMessage);
    // $('#alertText').text(alertMessage);
    $('#alertResultText').html(alertMessage);
    $("#alertResultArea").modal();
}
// alert창 닫기
function commonResultAlertClose() {
    $("#alertResultArea").modal("hide");
}

// 에러 메세지 alert 통일 용
function commonErrorAlert(statusCode, message) {
    commonAlert("Error(" + statusCode + ") : " + message);
    
}

// confirm modal창 보이기 modal창이 열릴 때 해당 창의 text 지정, close될 때 action 지정
export function commonConfirmOpen(targetAction, caller) {
    console.log("commonConfirmOpen : " + targetAction)

    //  [ id , 문구]
    let confirmModalTextMap = new Map(
        [
            ["CreateSnapshot", "Would you like to Create Snapshot?"],
            ["DeleteDataDisk", "Would you like to Delete Disk?"],
            ["DeleteMyImage", "Would you like to Delete MyImage?"],
            ["Logout", "Would you like to logout?"],
            ["Config", "Would you like to set Cloud config ?"],
            ["SDK", "Would you like to set Cloud Driver SDK ?"],
            ["Credential", "Would you like to set Credential ?"],
            ["Region", "Would you like to set Region ?"],
            ["Provider", "Would you like to set Cloud Provider ?"],

            ["MoveToConnection", "Would you like to set Cloud config ?"],
            ["ChangeConnection", "Would you like to change Cloud connection ?"],
            ["DeleteCloudConnection", "Would you like to delete <br /> the Cloud connection? "],

            ["DeleteCredential", "Would you like to delete <br /> the Credential? "],
            ["DeleteDriver", "Would you like to delete <br /> the Driver? "],
            ["DeleteRegion", "Would you like to delete <br /> the Region? "],


            // ["IdPassRequired", "ID/Password required !"],    --. 이거는 confirm이 아니잖아
            ["idpwLost", "Illegal account / password 다시 입력 하시겠습니까?"],
            ["ManageNS", "Would you like to manage <br />Name Space?"],
            ["NewNS", "Would you like to add a new Name Space?"],
            ["AddNewNameSpace", "Would you like to register NameSpace <br />Resource ?"],
            ["NameSpace", "Would you like to move <br />selected NameSpace?"],
            ["ChangeNameSpace", "Would you like to move <br />selected NameSpace?"],
            ["DeleteNameSpace", "Would you like to delete <br />selected NameSpace?"],

            ["AddNewVpc", "Would you like to create a new Network <br />Resource ?"],
            ["DeleteVpc", "Are you sure to delete this Network <br />Resource ?"],

            ["AddNewSecurityGroup", "Would you like to create a new Security <br />Resource ?"],
            ["DeleteSecurityGroup", "Would you like to delete Security <br />Resource ?"],

            ["AddNewSshKey", "Would you like to create a new SSH key <br />Resource ?"],
            ["DeleteSshKey", "Would you like to delete SSH key <br />Resource ?"],

            ["AddNewVirtualMachineImage", "Would you like to register Image <br />Resource ?"],
            ["DeleteVirtualMachineImage", "Would you like to un-register Image <br />Resource ?"],
            ["FetchImages", "Would you like to fetch images <br /> to this NameSpace ?"],

            ["AddNewVmSpec", "Would you like to register Spec <br />Resource ?"],
            ["DeleteVmSpec", "Would you like to un-register Spec <br />Resource ?"],
            ["FetchSpecs", "Would you like to fetch Spec <br /> to this NameSpace ?"],

            ["GotoMonitoringPerformance", "Would you like to view performance <br />for MCIS ?"],
            ["GotoMonitoringFault", "Would you like to view fault <br />for MCIS ?"],
            ["GotoMonitoringCost", "Would you like to view cost <br />for MCIS ?"],
            ["GotoMonitoringUtilize", "Would you like to view utilize <br />for MCIS ?"],

            ["McisLifeCycleReboot", "Would you like to reboot MCIS ?"],// mcis_life_cycle('reboot')
            ["McisLifeCycleSuspend", "Would you like to suspend MCIS ?"],//onclick="mcis_life_cycle('suspend')
            ["McisLifeCycleResume", "Would you like to resume MCIS ?"],//onclick="mcis_life_cycle('resume')"
            ["McisLifeCycleTerminate", "Would you like to terminate MCIS ?"],//onclick="mcis_life_cycle('terminate')
            ["McisManagement", "Would you like to manage MCIS ?"],// 해당 function 없음...
            ["MoveToMcisManagement", "Would you like to manage MCIS ?"],
            ["MoveToMcisManagementFromDashboard", "Would you like to manage MCIS ?"],

            ["AddNewMcis", "Would you like to create MCIS ?"],
            ["AddNewMcisDynamic", "Would you like to create MCIS ?"],
            ["DeleteMcis", "Are you sure to delete this MCIS? "],
            ["ImportScriptOfMcis", "Would you like to import MCIS script? "],
            ["ExportScriptOfMcis", "Would you like to export MCIS script? "],
            ["ShowMonitoring", "Would you like to go to the Monitoring page?"],

            ["AddNewVmOfMcis", "Would you like to add a new VM to this MCIS ?"],
            ["DeployServer", "Would you like to deploy?"],

            ["VmLifeCycle", "Would you like to view Server ?"],
            ["VmLifeCycleReboot", "Would you like to reboot VM ?"], //onclick="vm_life_cycle('reboot')"
            ["VmLifeCycleSuspend", "Would you like to suspend VM ?"], // onclick="vm_life_cycle('suspend')"
            ["VmLifeCycleResume", "Would you like to resume VM ?"], // onclick="vm_life_cycle('resume')"
            ["VmLifeCycleTerminate", "Would you like to terminate VM ?"], // onclick="vm_life_cycle('terminate')"
            ["VmManagement", "Would you like to manage VM ?"], // 해당 function 없음
            ["AddNewVm", "Would you like to add VM ?"], //onclick="vm_add()"
            ["ExportVmScriptOfMcis", "Would you like to export VM script ?"], //onclick="vm_add()"


            ["DifferentConnection", "Do you want to set different connectionName?"],
            ["DifferentConnectionAtSecurityGroup", "Do you want to set different connectionName?"],
            ["DifferentConnectionAtAssistPopup", "Do you want to set different connectionName?"],

            ["AddMonitoringAlertPolicy", "Would you like to register Threshold ?"],
            ["DeleteMonitoringAlertPolicy", "Are you sure to delete this Threshold ?"],
            ["AddNewMcks", "Would you like to create MCKS ?"],
            ["DeleteMcks", "Are you sure to delete this MCKS? "],
            ["AddNewNodeOfMcks", "Would you like to add a new Node to this MCKS ?"],
            ["DeleteNodeOfMcks", "Would you like to delete a Node of this MCKS ?"],


            ["AddMonitoringAlertEventHandler", "Would you like to add<br />Monitoring Alert Event-Handler ?"],
            ["deleteMonitoringAlertEventHandler", "Are you sure to delete<br />this Monitoring Alert Event-Handler?"],

            ["RegisterRecommendSpec", "현재 해당 connection에서 사용가능한 spec 이 없습니다. 등록 하시겠습니까?"],

            ["DeleteNlb", "Would you like to delete NLB ?"],

            ["AddNewPmks", "Would you like to create PMKS ?"],
            ["DeletePmks", "Are you sure to delete this PMKS? "],
            ["AddNewNodeGroupOfPmks", "Would you like to add a new NodeGroup to this PMKS ?"],
            ["DeleteNodeGroupOfPmks", "Would you like to delete a NodeGroup of this PMKS ?"],
        ]
    );
    console.log(confirmModalTextMap.get(targetAction));
    try {
        // $('#modalText').text(targetText);// text아니면 html로 해볼까? 태그있는 문구가 있어서
        //$('#modalText').text(confirmModalTextMap.get(targetAction));
        $('#confirmText').html(confirmModalTextMap.get(targetAction));
        $('#confirmOkAction').val(targetAction);
        console.log("caller : ", caller);
        $('#confirmCaller').val(caller);

        if (targetAction == "Region") {
            // button에 target 지정
            // data-target="#Add_Region_Register"
            // TODO : confirm 으로 물어본 뒤 OK버튼 클릭 시 targetDIV 지정하도록
        }
        $('#confirmArea').modal();
    } catch (e) {
        console.log(e);
        alert(e);
    }
}

// confirm modal창 보이기 modal창이 열릴 때 해당 창의 text 지정, close될 때 action 지정, text 내용 전송. caller : 구분자
function commonConfirmMsgOpen(targetAction, message, caller) {
    console.log("commonConfirmMsgOpen : " + targetAction)

    try {
        $('#confirmText').html(message);
        $('#confirmOkAction').val(targetAction);
        $('#confirmCaller').val(caller);

        $('#confirmArea').modal();
    } catch (e) {
        console.log(e);
        alert(e);
    }
}

// confirm modal창에서 ok버튼 클릭시 수행할 method 지정
function commonConfirmOk() {
    //modalArea
    var targetAction = $('#confirmOkAction').val();
    var caller = $('#confirmCaller').val();
    if (targetAction == "Logout") {
        // Logout처리하고 index화면으로 간다. Logout ==> cookie expire
        // location.href="/logout"
        var targetUrl = "/logout"
        changePage(targetUrl)

    } else if (targetAction == "MoveToConnection") {
        var targetUrl = "/setting/connections/cloudconnectionconfig/mngform"
        changePage(targetUrl)
    } else if (targetAction == "ChangeConnection") { // recommendvm에서 다른 connection 선택 시
        changeCloudConnection()
    } else if (targetAction == "DeleteCloudConnection") {
        deleteCloudConnection();
    } else if (targetAction == "Config") {
        //id="Config"
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "SDK") {
        //id="SDK"
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "DeleteCredential") {
        deleteCredential();
    } else if (targetAction == "DeleteDriver") {
        deleteDriver();
    } else if (targetAction == "DeleteRegion") {
        deleteRegion();

    } else if (targetAction == "Credential") {
        //id="Credential"
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "Region") {
        //id="Region"
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "Provider") {
        //id="Provider"
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "required") {//-- IdPassRequired
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "idpwLost") {//-- 
        console.log("commonConfirmOk " + targetAction);
    } else if (targetAction == "ManageNS") {//-- ManageNS
        var targetUrl = "/setting/namespaces/namespace/mngform"
        changePage(targetUrl)
    } else if (targetAction == "NewNS") {//-- NewNS
        var targetUrl = "/setting/namespaces/namespace/mngform"
        changePage(targetUrl)
    } else if (targetAction == "ChangeNameSpace") {//-- ChangeNameSpace
        var changeNameSpaceID = $("#tempSelectedNameSpaceID").val();
        setDefaultNameSpace(changeNameSpaceID)
    } else if (targetAction == "AddNewNameSpace") {//-- AddNewNameSpace
        displayNameSpaceInfo("REG")
        goFocus('ns_reg');// 해당 영역으로 scroll
    } else if (targetAction == "DeleteNameSpace") {
        deleteNameSpace()
    } else if (targetAction == "AddNewVpc") {
        displayVNetInfo("REG")
        goFocus('vnetCreateBox');
    } else if (targetAction == "DeleteVpc") {
        deleteVPC()
    } else if (targetAction == "AddNewSecurityGroup") {
        displaySecurityGroupInfo("REG")
        goFocus('securityGroupCreateBox');
    } else if (targetAction == "DeleteSecurityGroup") {
        deleteSecurityGroup()
    } else if (targetAction == "AddNewSshKey") {
        displaySshKeyInfo("REG")
        goFocus('sshKeyCreateBox');
    } else if (targetAction == "DeleteSshKey") {
        deleteSshKey()
    } else if (targetAction == "AddNewVirtualMachineImage") {
        displayVirtualMachineImageInfo("REG")
        goFocus('virtualMachineImageCreateBox');
    } else if (targetAction == "DeleteVirtualMachineImage") {
        deleteVirtualMachineImage()
    } else if (targetAction == "FetchImages") {
        getCommonFetchImages();
    } else if (targetAction == "AddNewVmSpec") {
        displayVmSpecInfo("REG")
        goFocus('vmSpecCreateBox');
    } else if (targetAction == "ExportVmScriptOfMcis") {
        vmScriptExport();
    } else if (targetAction == "DeleteVmSpec") {
        deleteVmSpec();
    } else if (targetAction == "FetchSpecs") {
        var connectionName = $("#regConnectionName").val();
        putFetchSpecs(connectionName);
    } else if (targetAction == "GotoMonitoringPerformance") {
        // alert("모니터링으로 이동 GotoMonitoringPerformance")
        // location.href ="";//../operation/Monitoring_Mcis.html
        var targetUrl = "/operation/monitorings/mcismng/mngform"
        changePage(targetUrl)
    } else if (targetAction == "GotoMonitoringFault") {
        // alert("모니터링으로 이동 GotoMonitoringFault")
        // location.href ="";//../operation/Monitoring_Mcis.html
        var targetUrl = "/operation/monitorings/mcismng/mngform"
        changePage(targetUrl)
    } else if (targetAction == "GotoMonitoringCost") {
        // alert("모니터링으로 이동 GotoMonitoringCost")
        // location.href ="";//../operation/Monitoring_Mcis.html
        var targetUrl = "/operation/monitorings/mcismng/mngform"
        changePage(targetUrl)
    } else if (targetAction == "GotoMonitoringUtilize") {
        // alert("모니터링으로 이동 GotoMonitoringUtilize")
        // location.href ="";//../operation/Monitoring_Mcis.html    
        var targetUrl = "/operation/monitorings/mcismng/mngform"
        changePage(targetUrl)
    } else if (targetAction == "McisLifeCycleReboot") {
        callMcisLifeCycle('reboot')
    } else if (targetAction == "McisLifeCycleSuspend") {
        callMcisLifeCycle('suspend')
    } else if (targetAction == "McisLifeCycleResume") {
        callMcisLifeCycle('resume')
    } else if (targetAction == "McisLifeCycleTerminate") {
        callMcisLifeCycle('terminate')
    } else if (targetAction == "McisManagement") {
        alert("수행할 function 정의되지 않음");
    } else if (targetAction == "MoveToMcisManagementFromDashboard") {
        var mcisID = $("#mcis_id").val();
        var targetUrl = "/operation/manages/mcismng/mngform?mcisid=" + mcisID;
        changePage(targetUrl)
    } else if (targetAction == "MoveToMcisManagement") {
        var targetUrl = "/operation/manages/mcismng/mngform";
        changePage(targetUrl)
    } else if (targetAction == "AddNewMcis") {
        // $('#loadingContainer').show();
        // location.href ="/operation/manages/mcis/regform/";
        var targetUrl = "/operation/manages/mcismng/regform";
        changePage(targetUrl)
    } else if (targetAction == "DeleteMcis") {
        deleteMCIS();
    } else if (targetAction == "DeployServer") {
        btn_deploy();
    } else if (targetAction == "ImportScriptOfMcis") {
        mcisScriptImport();
    } else if (targetAction == "ExportScriptOfMcis") {
        mcisScriptExport();
    } else if (targetAction == "ShowMonitoring") {
        var mcisID = $("#mcis_id").val();
        var targetUrl = "/operation/monitorings/mcismonitoring/mngform?mcisId=" + mcisID;
        changePage(targetUrl)
    } else if (targetAction == "VmLifeCycle") {
        alert("수행할 function 정의되지 않음");
    } else if (targetAction == "VmLifeCycleReboot") {
        vmLifeCycle('reboot')
    } else if (targetAction == "VmLifeCycleSuspend") {
        vmLifeCycle('suspend')
    } else if (targetAction == "VmLifeCycleResume") {
        vmLifeCycle('resume')
    } else if (targetAction == "VmLifeCycleTerminate") {
        vmLifeCycle('terminate')
    } else if (targetAction == "VmManagement") {
        alert("수행할 function 정의되지 않음");
    } else if (targetAction == "AddNewVm") {
        addNewVirtualMachine()
    } else if (targetAction == "AddNewVmOfMcis") {
        addNewVirtualMachine()
    } else if (targetAction == "ExportVmScriptOfMcis") {
        vmScriptExport();
    } else if (targetAction == "--") {
        addNewVirtualMachine()
    } else if (targetAction == "monitoringConfigPolicyConfig") {
        regMonitoringConfigPolicy()
    } else if (targetAction == "DifferentConnection") {
        setAndClearByDifferentConnectionName(caller);
    } else if (targetAction == "DifferentConnectionAtSecurityGroup") {
        uncheckDifferentConnectionAtSecurityGroup();
    } else if (targetAction == "DifferentConnectionAtAssistPopup") {
        // connection이 다른데도 set 한다고 하면 이전에 설정한 값들을 초기화 한 후 set한다.
        applyAssistValues(caller);
    } else if (targetAction == "AddMonitoringAlertPolicy") {
        addMonitoringAlertPolicy();
    } else if (targetAction == "DeleteMonitoringAlertPolicy") {
        deleteMonitoringAlertPolicy();
    } else if (targetAction == "AddNewMcks") {
        var targetUrl = "/operation/manages/mcksmng/regform";
        changePage(targetUrl)
    } else if (targetAction == "AddNewNodeOfMcks") {
        addNewNode();
    } else if (targetAction == "DeleteNodeOfMcks") {
        deleteNodeOfMcks();
    } else if (targetAction == "AddMonitoringAlertEventHandler") {
        addMonitoringAlertEventHandler();
    } else if (targetAction == "deleteMonitoringAlertEventHandler") {
        deleteMonitoringAlertEventHandler();
    } else if (targetAction == "DeleteMcks") {
        deleteMCKS();
    } else if (targetAction == "RegisterRecommendSpec") {
        commonPromptOpen("RegisterRecommendSpec")
    } else if (targetAction == "AddNewMcisDynamic") {
        createMcisDynamic()
    } else if (targetAction == "DeleteDataDisk") {
        deleteDataDisk();

    } else if (targetAction == "DeleteMyImage") {
        deleteMyImageDisk();

    } else if (targetAction == "CreateSnapshot") {
        commonPromptOk
        createSnapshot();

    } else if (targetAction == "DeleteNlb") {
        deleteNlb();
    } else if (targetAction == "AddNewPmks") {
        changePage("PmksClusterRegForm");
    } else if (targetAction == "DeletePmks") {
        deleteCluster();
    } else if (targetAction == "AddNewNodeGroupOfPmks") {
        changePage("PmksNodeGroupRegForm");
    } else if (targetAction == "DeleteNodeGroupOfPmks") {
        deleteNodeGroupOfPmks();
    } else {
        alert("수행할 function 정의되지 않음 " + targetAction);
    }
    console.log("commonConfirmOk " + targetAction);
    commonConfirmClose();
}

// confirm modal창 닫기. setting값 초기화
function commonConfirmClose() {
    $('#confirmText').text('');
    $('#confirmOkAction').val('');
    // $('#modalArea').hide(); 
    $("#confirmArea").modal("hide");
}