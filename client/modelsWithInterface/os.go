package modelsWithInterface

import(
	"github.com/shirou/gopsutil/host"
    "fmt"
)

type OSMetrics struct {
	Uptime          uint64 `json:"uptime"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

func getOSMetrics() (OSMetrics, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return OSMetrics{}, err
	}

	return OSMetrics{
		Uptime:          hostInfo.Uptime,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
	}, nil
}


func (d1 *OSMetrics) CollectMetrics() error{
   temp, err := getOSMetrics()
   if err != nil{
	fmt.Println(err)
	return err
   }
   //d1 = &temp not working

   d1.Platform = temp.Platform
   d1.PlatformVersion = temp.PlatformVersion
   d1.Uptime = temp.Uptime
   return nil
}