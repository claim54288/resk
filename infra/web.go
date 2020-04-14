package infra

var apiInitializerRegister = new(InitializeRegister)

//注册WEB API初始化对象
func RegisterApi(ai Initializer) {
	apiInitializerRegister.Regidter(ai)
}

//获取注册的web api初始化对象
func GetApiInitializers() []Initializer {
	return apiInitializerRegister.Initializers
}

type WebApiSterter struct {
	BaseStarter
}

func (w *WebApiSterter) Setup(ctx StarterContext) {
	for _, v := range GetApiInitializers() {
		v.Init()
	}
}
