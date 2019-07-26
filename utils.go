package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type nodeStats struct {
	Mem struct {
		Current int     `json:"current"`
		Total   int     `json:"total"`
		Load    float64 `json:"load"`
	} `json:"mem"`
	Disk struct {
		Current int `json:"current"`
		Total   int `json:"total"`
		Used    int `json:"used"`
	} `json:"disk"`
	CPU struct {
		Load float64 `json:"load"`
	} `json:"cpu"`
}

func getStatsFromNode() (float64, float64, int){
	response, err := exec.Command("/bin/sh", "stats.sh").Output()
	if err != nil {
		fmt.Println("unable to get stats")
	}
	output := nodeStats{}
	_ = json.Unmarshal(response, &output)
	cpuLoad := output.CPU.Load
	memLoad := output.Mem.Load
	diskCapacity := output.Disk.Used

	return cpuLoad,memLoad,diskCapacity

}