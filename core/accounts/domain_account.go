package accounts

import (
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/tietang/dbx"
	"resk/infra/base"
	"resk/services"
)

//有状态的，每次使用都要实例化
type accountDomain struct {
	account    Account
	accountLog AccountLog
}

//创建流水的记录
func (domain *accountDomain) createAccountLogNo() {
	//暂时采用ksuid的ID生成策略来创建No ，以后可以优化成可读性比较好的分布式ID
	domain.accountLog.LogNo = ksuid.New().Next().String()

}

//创建流水No logNo
func (domain *accountDomain) createAccountNo() {
	domain.accountLog.AccountNo = ksuid.New().Next().String()
}

//生成账户No accountNo
func (domain *accountDomain) createAccountLog() {
	//通过account来创建流水，创建账户逻辑在前
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo
	//流水中的交易主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String
	//交易对象信息
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String
	//交易金额
	domain.accountLog.Amount = domain.account.Balance
	domain.accountLog.Balance = domain.account.Balance
	//交易变化的属性
	domain.accountLog.Decs = "账户创建"
	domain.accountLog.ChangeType = services.AccountCreated
	domain.accountLog.ChangeFlag = services.FlagAccountCreated
}

func (domain *accountDomain) Create(dto services.AccountDTO) (*services.AccountDTO, error) {
	//创建账户的持久化对象
	domain.account = Account{}
	domain.account.FromDTO(&dto)
	domain.createAccountNo()
	domain.account.Username.Valid = true
	//创建账户流水的持久化对象
	domain.createAccountLog()
	accountDao := AccountDao{}
	accountLogDao := AccountLogDao{}
	var rdto *services.AccountDTO
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		accountLogDao.runner = runner
		//插入账户数据
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}
		//如果插入成功，就插入流水数据
		id, err = accountLogDao.Insert(&domain.accountLog)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户流水失败")
		}
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	rdto = domain.account.ToDTO()
	return rdto, err
}
