package enums

// AccessLevel 访问级别
type AccessLevel int

const (
	// PRIVATE 仅自己可见
	PRIVATE AccessLevel = iota
	// PUBLIC 公开
	PUBLIC
	// PROTECTED 仅登录用户可见
	PROTECTED
)

// String 转换为字符串
func (a AccessLevel) String() string {
	switch a {
	case PRIVATE:
		return "PRIVATE"
	case PUBLIC:
		return "PUBLIC"
	case PROTECTED:
		return "PROTECTED"
	}
	return "PRIVATE"
}
