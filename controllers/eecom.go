package controllers

import (
	"github.com/willitlaunch/willitlaunch-go/games"
)

type EECOMController struct {
	FlightControllerBase
}

func (c *EECOMController) Init() {
	c.FlightControllerBase.Init()
	c.Name = "EECOM"

	CLGame1 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 0}}
	CLGame1.Init()
	CLGame1.CLDial.Label = "Cyro Level Left"
	c.Games = append(c.Games, &CLGame1)

	CPGame1 := games.CabinPressureGame{GameBase: games.GameBase{Gid: 1}}
	CPGame1.Init()
	c.Games = append(c.Games, &CPGame1)

	CLGame2 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 2}}
	CLGame2.Init()
	CLGame1.CLDial.Label = "Cyro Level Right"
	c.Games = append(c.Games, &CLGame2)

	CPGame2 := games.CabinPressureGame{GameBase: games.GameBase{Gid: 3}}
	CPGame2.Init()
	CPGame2.CLDial.Label = "External Pressure"
	c.Games = append(c.Games, &CPGame2)

	// BPGame := games.BatteryPowerGame{Gid: 2}
	// BPGame.Init()
	// c.Games = append(c.Games, &BPGame)
}
