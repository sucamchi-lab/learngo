package main

type Engine struct {
	Fuel  int
	Boost int
}

type Pilot struct {
	CallSign string
	Ship     string
	Engine   Engine
}

type Mission struct {
	Zone   string
	Threat int
}

type FlightReport struct {
	PilotName    string
	ShipName     string
	Zone         string
	StartingFuel int
	EndingFuel   int
	DangerScore  int
}

func buildFlightReport(pilot Pilot, mission Mission) FlightReport {
	endingFuel := pilot.Engine.Fuel - mission.Threat
	if endingFuel < 0 {
		endingFuel = 0
	}

	return FlightReport{
		PilotName:    pilot.CallSign,
		ShipName:     pilot.Ship,
		Zone:         mission.Zone,
		StartingFuel: pilot.Engine.Fuel,
		EndingFuel:   endingFuel,
		DangerScore:  mission.Threat + pilot.Engine.Boost,
	}
}
