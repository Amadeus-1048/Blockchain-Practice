package v2

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PrescriptionRequestBody struct {
	Doctor    string `json:"doctor"`    // 医生ID
	Patient   string `json:"patient"`   // 患者Id
	Diagnosis string `json:"diagnosis"` // 诊断结果
	//Drug      []Drug `json:"drug"`      // 药品列表及用量
	DrugName   string `json:"drug_name"`   // 药品名
	DrugAmount string `json:"drug_amount"` // 药品用量
	Hospital   string `json:"hospital"`    // 医院 ID
	Comment    string `json:"comment"`     // 备注
}

type PrescriptionQueryRequestBody struct {
	Patient string `json:"patient"` // 患者AccountId
}

func CreatePrescription(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(PrescriptionRequestBody)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Doctor == "" || body.Patient == "" || body.Diagnosis == "" || body.Hospital == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数存在空值")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Doctor))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	bodyBytes = append(bodyBytes, []byte(body.Diagnosis))
	bodyBytes = append(bodyBytes, []byte(body.DrugName))
	bodyBytes = append(bodyBytes, []byte(body.DrugAmount))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	bodyBytes = append(bodyBytes, []byte(body.Comment))
	//bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.TotalArea, 'E', -1, 64)))

	// 调用智能合约
	resp, err := bc.ChannelExecute("createPrescription", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryPrescriptionList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(PrescriptionQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Patient != "" {
		bodyBytes = append(bodyBytes, []byte(body.Patient))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryPrescription", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
