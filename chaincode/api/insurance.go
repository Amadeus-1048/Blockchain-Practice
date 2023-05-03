package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateInsuranceCover 创建保险报销订单
func CreateInsuranceCover(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	prescriptionID := args[0] // 处方id
	patientID := args[1]      // 患者id
	status := args[2]         // 订单状态

	if prescriptionID == "" || patientID == "" || status == "" {
		return shim.Error("参数存在空值")
	}

	// 判断是否为患者操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, []string{patientID})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}

	record := &model.InsuranceCover{
		ID:           stub.GetTxID()[:16],
		Prescription: prescriptionID,
		Patient:      patientID,
		Status:       "处理中",
	}

	// 写入账本
	if err := utils.WriteLedger(record, stub, model.InsuranceKey, []string{record.Patient, record.ID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	recordByte, err := json.Marshal(record)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(recordByte)
}

// QueryInsuranceCover 查询保险报销订单(可查询所有，也可根据所有人查询名下处方)
func QueryInsuranceCover(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var insuranceCoverList []model.InsuranceCover
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.InsuranceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var p model.InsuranceCover
			err := json.Unmarshal(v, &p)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryInsuranceCover-反序列化出错: %s", err))
			}
			insuranceCoverList = append(insuranceCoverList, p)
		}
	}
	insuranceCoverByte, err := json.Marshal(insuranceCoverList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryInsuranceCover-序列化出错: %s", err))
	}
	return shim.Success(insuranceCoverByte)
}
