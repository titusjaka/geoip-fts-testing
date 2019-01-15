package geoip_fts_testing

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
)

func IpToInt(ipStr string) *big.Int {
	ip := net.ParseIP(ipStr)
	newInt := big.NewInt(0)
	ipBytes := []byte(ip)
	if ip.To4() != nil {
		if len(ip) == 16 {
			ipBytes = ip[12:16]
		}
	}
	return newInt.SetBytes(ipBytes)
}

func IntToIp(ipInt *big.Int) string {
	return net.IP(ipInt.Bytes()).String()
}

func GetIdFromIpRange(startIP, endIP string) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%s", startIP, endIP)))
	return hex.EncodeToString(hasher.Sum(nil))
}
