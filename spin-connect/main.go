package main

import (
	"encoding/json"
	"log"
	"net/http"
	"spin-connect/models"
)

var DashboardModels []models.DashBoardModel

func main() {
	DashboardModels = []models.DashBoardModel{
		{
			Name: "dashboard",
			Data: models.StageData{
				MachineDetails: []models.MachineDetail{
					{
						MachineName: "Blowroom Machine 1",
						Model:       "Model X",
						Status:      "Running",
					},
					{
						MachineName: "Blowroom Machine 2",
						Model:       "Model Y",
						Status:      "Stopped",
					},
				},
				PerformanceDetails: models.Performance{
					Efficiency:     85.6,
					ProductionRate: 1200,
					Downtime:       120,
				},
			},
		},
	}

	handleRequests()
}

func handleRequests() {
	
	http.HandleFunc("/getAllDashboardData", returnAllDashboardData)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func returnAllDashboardData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(DashboardModels)
}
