package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func OddDtoToOddModel(oddDto dto.OddsDto) models.Odds {
	odd := models.Odds{}
	for _, ob := range oddDto.Response[0].Bookmakers[0].Bets {
		switch ob.Name {
		case "Both Teams Score":
			for _, ov := range ob.Values {
				switch ov.Value {
				case "Yes":
					oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
					odd.All_Teams_Scores_odd = float32(oddIn64)
				case "No":
					oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
					odd.Not_All_Teams_Scores_odd = float32(oddIn64)
				}
			}
		case "Match Winner":
			for _, ov := range ob.Values {
				switch ov.Value {
				case "Away":
					oddIn64, _ := strconv.ParseFloat(ov.Odd, 32)
					odd.Team_Away_Win_odd = float32(oddIn64)

				case "Home":
					oddIn64, _ := strconv.ParseFloat(ov.Odd, 32)
					odd.Team_Home_Win_odd = float32(oddIn64)
				case "Draw":
					oddIn64, _ := strconv.ParseFloat(ov.Odd, 32)
					odd.Draw_odd = float32(oddIn64)
				}

			}
		}
	}

	return odd

}