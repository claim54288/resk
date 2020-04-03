package infra

//资源启动器上下文，
// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

//基础资源启动器接口
type Starter interface {
	//1.系统启动，初始化的一些基础资源
	Init(StarterContext)
	//2.系统基础资源的安装
	Setup(StarterContext)
	//3.启动基础资源
	Start(StarterContext)
	//4.资源停止和销毁
	Stop(StarterContext)
}

var _ Starter = new(BaseStarter)

//基础空启动器实现，为了方便资源启动器的代码实现
type BaseStarter struct{}
