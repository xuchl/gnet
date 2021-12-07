package main
import (
	"fmt"
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"time"
)

func main() {
	tapconfig := water.Config{
		DeviceType: water.TAP,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: "tun0",

		},
	}

	ifce, err := water.New(tapconfig)
	if err != nil {
		fmt.Printf("创建失败：%v \n", err)
		return
	}

	var buf ethernet.Frame

	for {
		_, err := ifce.Read(buf)
		if err != nil {
			fmt.Printf("读取数据：%v \n", err)
			return
		}

		time.Sleep(time.Second)
	}
}