package infra

import "github.com/tietang/props/kvs"

const (
	KeyProps = "props"
)

//资源启动器上下文，
// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(kvs.ConfigSource)
}

//基础资源启动器接口
type Starter interface {
	//1.系统启动，初始化的一些基础资源
	Init(StarterContext)
	//2.系统基础资源的安装
	Setup(StarterContext)
	//3.启动基础资源
	Start(StarterContext)
	//启动器是否可阻塞
	StartBlocking() bool
	//4.资源停止和销毁
	Stop(StarterContext)
}

var _ Starter = new(BaseStarter) //这一行是验证一下BaseStarter有没有实现Starter接口的，没实际作用

//基础空启动器实现，为了方便资源启动器的代码实现
type BaseStarter struct{}

func (*BaseStarter) Init(StarterContext)  {}
func (*BaseStarter) Setup(StarterContext) {}
func (*BaseStarter) Start(StarterContext) {}
func (*BaseStarter) StartBlocking() bool {
	return false
}
func (*BaseStarter) Stop(StarterContext) {}

//启动器注册器
type startRegister struct {
	starters []Starter
}

//启动器注册
func (r *startRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *startRegister) AllStarters() []Starter {
	return r.starters
}

var StarterRegister = new(startRegister)

//注册函数，被放在包的init函数里
func Register(s Starter) {
	StarterRegister.Register(s)
}

/**
系统基础资源的启动管理，循环执行StarterRegister里的所有的资源
这个StarterRegister里的资源是通过Register函数注册进去的，Register函数被放在包的初始化函数init里
*/
func SystemRun() {
	//1.初始化
	ctx := StarterContext{}
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(ctx)
	}
	//2.安装
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(ctx)
	}
	//3.启动
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(ctx)
	}
}
