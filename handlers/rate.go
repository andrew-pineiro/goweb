package handlers

import (
	"time"
)

type Rate struct {
	IP              string
	Count           int
	LastBanDate     time.Time
	LastConnectDate time.Time
}

const (
	MAX_RATEMAP_SIZE  = 1024
	RATE_LIMIT        = 1000
	BAN_LIMIT_SECONDS = 10
)

var GlobalRateMap []Rate
var GlobalBanList []Rate

func banIP(index int) {
	GlobalRateMap[index].LastBanDate = time.Now()
	GlobalBanList = append(GlobalBanList, GlobalRateMap[index])
}
func checkBanList(index int) bool {
	rate := GlobalRateMap[index]
	for i, bannedMF := range GlobalBanList {
		if rate.IP == bannedMF.IP {
			if time.Now().After(rate.LastBanDate.Add(BAN_LIMIT_SECONDS * time.Second)) {
				GlobalBanList = append(GlobalBanList[:i], GlobalBanList[i+1:]...) //strip element from slice
				GlobalRateMap[i].Count = 0
				GlobalRateMap[i].LastBanDate = time.Time{}
				return false
			}
			return true
		}
	}
	return false
}
func getRateIndex(ip string) int {
	for i, rate := range GlobalRateMap {
		if rate.IP == ip {
			GlobalRateMap[i].Count++
			return i
		}
	}
	// if no existing rate exists
	rate := Rate{
		IP:              ip,
		Count:           1,
		LastConnectDate: time.Now(),
	}
	i := addRateToMap(rate)
	return i
}

func addRateToMap(rate Rate) int {
	if len(GlobalRateMap)+1 >= MAX_RATEMAP_SIZE {
		GlobalRateMap = append(GlobalRateMap[:0], GlobalRateMap[0+1:]...)
	}
	GlobalRateMap = append(GlobalRateMap, rate)
	return getRateIndex(rate.IP)
}

func CheckRateCount(ip string) bool {
	currRateIndex := getRateIndex(ip)
	if currRateIndex < 0 {
		return true
	}

	isBanned := checkBanList(currRateIndex)
	if isBanned {
		return false
	}
	if GlobalRateMap[currRateIndex].Count >= RATE_LIMIT {
		banIP(currRateIndex)
		return false
	}
	return true
}
