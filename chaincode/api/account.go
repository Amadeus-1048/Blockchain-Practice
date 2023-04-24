package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// QueryAccountList 查询账户列表
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// args：用于创建复合主键的键值列表，一个字符串数组。
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []model.Account
	// 根据 model.AccountKey 和 args 的前缀从账本中获取所有的符合条件的账户信息
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// 遍历账户信息，将其转换为 model.Account 对象添加到 accountList 列表中
	for _, v := range results {
		if v != nil {
			var account model.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	// 将 accountList 转换为字节数组，然后使用 shim.Success 方法将其作为查询结果返回给调用方
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}
