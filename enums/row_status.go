package enums

type RowStatus int

const (
	// NORMAL 正常
	NORMAL RowStatus = iota
	// ARCHIVED 归档
	ARCHIVED
)

// String 转换为字符串
func (status RowStatus) String() string {
	switch status {
	case NORMAL:
		return "NORMAL"
	case ARCHIVED:
		return "ARCHIVED"
	}
	return "NORMAL"

}
