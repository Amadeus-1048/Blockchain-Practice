package model

// Account 账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

// RealEstate 房地产作为担保出售、捐赠或质押时Encumbrance为true，默认状态false。
// 仅当Encumbrance为false时，才可发起出售、捐赠或质押
// Proprietor和RealEstateID一起作为复合键,保证可以通过Proprietor查询到名下所有的房产信息
type RealEstate struct {
	RealEstateID string  `json:"realEstateId"` //房地产ID
	Proprietor   string  `json:"proprietor"`   //所有者(业主)(业主AccountId)
	Encumbrance  bool    `json:"encumbrance"`  //是否作为担保
	TotalArea    float64 `json:"totalArea"`    //总面积
	LivingSpace  float64 `json:"livingSpace"`  //生活空间
}

// Selling 销售要约
// 需要确定ObjectOfSale是否属于Seller
// 买家初始为空
// Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的销售
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //销售对象(正在出售的房地产RealEstateID)
	Seller        string  `json:"seller"`        //发起销售人、卖家(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与销售人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

// SellingStatusConstant 销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// SellingBuy 买家参与销售
// 销售对象不能是买家发起的
// Buyer和CreateTime作为复合键,保证可以通过buyer查询到名下所有参与的销售
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //参与销售人、买家(买家AccountId)
	CreateTime string  `json:"createTime"` //创建时间
	Selling    Selling `json:"selling"`    //销售对象
}

// Donating 捐赠要约
// 需要确定ObjectOfDonating是否属于Donor
// 需要指定受赠人Grantee，并等待受赠人同意接收
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //捐赠对象(正在捐赠的房地产RealEstateID)
	Donor            string `json:"donor"`            //捐赠人(捐赠人AccountId)
	Grantee          string `json:"grantee"`          //受赠人(受赠人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //捐赠状态
}

// DonatingStatusConstant 捐赠状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "捐赠中", //捐赠人发起捐赠合约，等待受赠人确认受赠
		"cancelled":     "已取消", //捐赠人在受赠人确认受赠之前取消捐赠或受赠人取消接收受赠
		"done":          "完成",  //受赠人确认接收，交易完成
	}
}

// DonatingGrantee 供受赠人查询的
type DonatingGrantee struct {
	Grantee    string   `json:"grantee"`    //受赠人(受赠人AccountId)
	CreateTime string   `json:"createTime"` //创建时间
	Donating   Donating `json:"donating"`   //捐赠对象
}

// objectType  对象类型，用于创建复合主键
const (
	AccountKey         = "account-key"
	RealEstateKey      = "real-estate-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
	DonatingKey        = "donating-key"
	DonatingGranteeKey = "donating-grantee-key"

	AccountV2Key    = "account-v2-key"
	PrescriptionKey = "prescription-key"
	PatientKey      = "patient-key"
	InsuranceKey    = "insurance-key"
	DrugKey         = "drug-key"
)

// --------------------------------------------------------------------

// AccountV2 账号
type AccountV2 struct {
	AccountId   string `json:"account_id"`   // 账号ID
	AccountName string `json:"account_name"` // 账号名
}

// Hospital 医院
type Hospital struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Admins  []HospitalAdmin `json:"admins"`
	Doctors []Doctor        `json:"doctors"`
}

// HospitalAdmin 医院管理员
type HospitalAdmin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Doctor 医生
type Doctor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Prescription 医疗处方
type Prescription struct {
	ID        string `json:"id"`        // 医疗处方ID
	Patient   string `json:"patient"`   // 患者ID
	Diagnosis string `json:"diagnosis"` // 诊断结果
	Drug      []Drug `json:"drug"`      // 药品列表及用量
	Doctor    string `json:"doctor"`    // 开方医师 AccountV2Id
	Hospital  string `json:"hospital"`  // 医院 ID
	Comment   string `json:"comment"`   // 备注
}

// Patient 患者
type Patient struct {
	ID     string `json:"id"`     // 患者 AccountV2Id
	Name   string `json:"name"`   // 患者姓名
	Age    int    `json:"age"`    // 患者年龄
	Gender string `json:"gender"` // 患者性别
}

// Drug 药品
type Drug struct {
	//ID      string `json:"id"`
	Name   string `json:"Name"`   // 药品名
	Amount string `json:"amount"` // 药品数量
}

// DrugOrder 药品订单
type DrugOrder struct {
	ID           string `json:"id"`           // 订单ID
	Name         string `json:"Name"`         // 药品名
	Amount       string `json:"amount"`       // 药品数量
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者ID
	DrugStore    string `json:"drug_store"`   // 药店id
}

// DrugStore 药店
type DrugStore struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Insurance 保险机构
type Insurance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InsuranceCover 保险报销订单
type InsuranceCover struct {
	ID           string `json:"id"`           // 订单ID
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者ID
	Status       string `json:"status"`       // 订单状态
}

// InsuranceStatusConstant 保险状态
var InsuranceStatusConstant = func() map[string]string {
	return map[string]string{
		"processing": "处理中", // 患者发起保险报销申请，等待保险公司确认报销
		"cancelled":  "已取消", // 患者在保险公司确认报销之前取消保险报销申请
		"refused":    "已拒绝", // 保险公司拒绝确认报销
		"approved":   "已通过", // 保险公司确认报销，保险报销完成
	}
}

// DrugStatusConstant 药品状态
//var DrugStatusConstant = func() map[string]string {
//	return map[string]string{
//		"processing": "处理中", //
//		"done":       "完成",   //
//	}
//}
