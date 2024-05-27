package models

type DashBoardModel struct {
	Name string     `json:"name"`
	Data StageData `json:"data"`
}

type StageData struct {
	MachineDetails    []MachineDetail `json:"machineDetails"`
	PerformanceDetails Performance     `json:"performanceDetails"`
}

type MachineDetail struct {
	MachineName string `json:"machineName"`
	Model       string `json:"model"`
	Status      string `json:"status"`
}

type Performance struct {
	Efficiency     float64 `json:"efficiency"`
	ProductionRate int     `json:"productionRate"`
	Downtime       int     `json:"downtime"`
}
