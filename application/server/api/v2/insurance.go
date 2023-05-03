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

type InsuranceCoverRequestBody struct {
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者Id
	Status       string `json:"status"`       // 订单状态
}

type InsuranceCoverQueryRequestBody struct {
	Patient string `json:"patient"` // 患者AccountId
}

func CreateInsuranceCover(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(InsuranceCoverRequestBody)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Prescription == "" || body.Patient == "" || body.Status == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数存在空值")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Prescription))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	//bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.TotalArea, 'E', -1, 64)))

	// 调用智能合约
	resp, err := bc.ChannelExecute("createInsuranceCover", bodyBytes)
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

func QueryInsuranceCoverList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(InsuranceCoverQueryRequestBody)
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
	resp, err := bc.ChannelQuery("queryInsuranceCover", bodyBytes)
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
