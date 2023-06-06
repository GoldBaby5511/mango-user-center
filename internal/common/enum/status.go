package enum

// 用户账号状态
const (
	UserStateOk      = iota + 1 // 正常
	UserStateFrozen             // 冻结
	UserStateUncheck            // 账号未经验证
)

// 第三方钱包类型
const (
	WalletCategoryMetaMask = iota + 1 // metamask
)

// 验证码类型
const (
	VerifyCodeTypeSignUp         = iota + 1 // 注册
	VerifyCodeTypeChangePassword            // 修改密码
)
