package snowflake

import (
	"sync"
	"time"
)

var (
	Epoch         int64 = 1597075200000 //2020年8月11号0:00 时刻的毫秒级时间戳
	machineID     int64                 //机器id
	sn            int64                 //序列号
	lastTimeStamp int64                 //记录上次的时间戳(毫秒级)
	mutex         sync.Mutex
)

func init() {
	lastTimeStamp = time.Now().UnixNano()/1e6 - Epoch
}
func SetMachineId(mid int64) {
	machineID = mid << 12
}

func GetId() int64 {
	mutex.Lock()
	curTimeStamp := time.Now().UnixNano()/1e6 - Epoch
	if curTimeStamp == lastTimeStamp {
		sn++
		if sn > 4095 {
			// 序列号超出，则重置序列号。这也意味着每毫秒最多能生成4096个id值
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano()/1e6 - Epoch
			lastTimeStamp = curTimeStamp
			sn = 0
		} //与运算 对应位全为1时，则为1.否则为0
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		mutex.Unlock()
		return rightBinValue | machineID | sn
	} else if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		mutex.Unlock()
		return rightBinValue | machineID | sn
	}
	mutex.Unlock()
	return 0
}
