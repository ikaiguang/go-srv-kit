package setuppkg

// options 配置可选项
//type options struct {
//	configPath string
//}
//
//// Option is config option.
//type Option func(*options)
//
//// WithConfigPath 配置路径
//func WithConfigPath(configPath string) Option {
//	return func(o *options) {
//		o.configPath = configPath
//	}
//}
//
//// New 启动与配置
//func New(setupOpts ...Option) (Engine, error) {
//	// parses the command-line flags
//	if !flag.Parsed() {
//		flag.Parse()
//	}
//
//	// 初始化配置手柄
//	configHandler, err := newConfigWithFiles(setupOpts...)
//	if err != nil {
//		return nil, err
//	}
//
//	// 开始配置
//	stdlog.Println("|==================== 配置程序 开始 ====================|")
//	defer stdlog.Println("|==================== 配置程序 结束 ====================|")
//
//	return newEngine(configHandler)
//}
//
//// newEngine 启动与配置
//func newEngine(configHandler Config) (Engine, error) {
//	// 初始化手柄
//	var (
//		err          error
//		setupHandler = initEngine(configHandler)
//	)
//
//	// 设置调试工具
//	if err = setupHandler.loadingDebugUtil(); err != nil {
//		return nil, err
//	}
//
//	// 设置日志工具
//	if _, err = setupHandler.loadingLogHelper(); err != nil {
//		return nil, err
//	}
//
//	// 覆盖 重写 json 响应
//	apputil.RewriteJSONEncoding()
//
//	// 注册和发现
//	if err = setupHandler.InitRegistry(); err != nil {
//		return nil, err
//	}
//
//	// redis 客户端
//	if cfg := setupHandler.Config.RedisConfig(); cfg != nil && cfg.Enable {
//		_, err := setupHandler.GetRedisClient()
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	// mongo 客户端
//	if cfg := setupHandler.Config.MongoConfig(); cfg != nil && cfg.Enable {
//		_, err := setupHandler.GetMongoDB()
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return setupHandler, nil
//}
//
//// newConfigWithFiles 初始化配置手柄
//func newConfigWithFiles(setupOpts ...Option) (Config, error) {
//	// parses the command-line flags
//	if !flag.Parsed() {
//		flag.Parse()
//	}
//
//	// 启动选项
//	setupOpt := &options{}
//	for i := range setupOpts {
//		setupOpts[i](setupOpt)
//	}
//
//	stdlog.Println("|==================== 加载配置文件 开始 ====================|")
//	defer stdlog.Println()
//	defer stdlog.Println("|==================== 加载配置文件 结束 ====================|")
//
//	// 配置路径
//	confPath := _configFilepath
//	if setupOpt.configPath != "" {
//		confPath = setupOpt.configPath
//	} else if confPath == "" {
//		confPath = _defaultConfigFilepath
//	}
//
//	var opts []config.Option
//	stdlog.Println("|*** 加载：配置文件路径: ", confPath)
//	opts = append(opts, config.WithSource(file.NewSource(confPath)))
//
//	return NewConfiguration(opts...)
//}
//
//// NewConfiguration 配置处理手柄
//func NewConfiguration(opts ...config.Option) (Config, error) {
//	handler := &configuration{}
//	if err := handler.init(opts...); err != nil {
//		return nil, err
//	}
//	return handler, nil
//}
