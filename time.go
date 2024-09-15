package dash

import "time"

func Now() int64 {
	return time.Now().Unix()
}

func NowMilli() int64 {
	return time.Now().UnixMilli()
}

func NowMicro() int64 {
	return time.Now().UnixMicro()
}
