package services

//转账状态
type TransferedStatus int8

const (
	//转账失败
	TansferedStatusFailure TransferedStatus = -1
	//余额不足
	TransferdStatusSufficientFunds TransferedStatus = 0
	//转账成功
	TransferedStatusSuccess TransferedStatus = 1
)

//转账类型：0=创建账户 >0进账 <0支出
type ChangeType int8

const (
	//账户创建
	AccountCreated ChangeType = 0
	//储值
	AccountStoreValue ChangeType = 1
	//红包资金的支出
	EnvelopOutgoing ChangeType = -2
	//红包资金的收入
	EnvelopeIncoming ChangeType = 2
	//红包过期的退款
	EnvelopExpiredRefund ChangeType = 3
)

//资金交易的变化标识
type ChangeFlag int8

const (
	//创建账户 = 0
	FlagAccountCreated ChangeFlag = 0
	//支出=-1
	FlagTransferOut ChangeFlag = -1
	//收入=1
	FlagTransferIn ChangeFlag = 1
)