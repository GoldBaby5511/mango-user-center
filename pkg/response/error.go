package response

// 鸭子类型，自定义error
var (
	_ error = Msg("HelloWorld")
)

type Msg string

func (f Msg) Error() string {
	return string(f)
}
