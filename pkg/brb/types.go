package brb

type deliveries struct {
	Deliveries []struct {
		Title  string `json:"title"`
		Params struct {
			Matrix []struct {
				ID                string `json:"id"`
				IsExpressDelivery bool   `json:"isExpressDelivery"`
				Day               string `json:"day"`
				DayShort          string `json:"dayShort"`
				Hours             []struct {
					ID                              string      `json:"id"`
					DeliveryTime                    string      `json:"deliveryTime"`
					Hour                            string      `json:"hour"`
					Price                           float64     `json:"price"`
					Available                       bool        `json:"available"`
					IsUnavailableAlcSellingTime     bool        `json:"isUnavailableAlcSellingTime"`
					IsUnavailableAlcOrEnergySelling bool        `json:"isUnavailableAlcOrEnergySelling"`
					IsLockerOrPup                   bool        `json:"isLockerOrPup"`
					SalesCoefficient                float64     `json:"salesCoefficient"`
					DeliveryWave                    string      `json:"deliveryWave"`
					PickingHour                     int         `json:"pickingHour"`
					ChangeTimeslotShop              interface{} `json:"changeTimeslotShop"`
				} `json:"hours"`
			} `json:"matrix"`
		} `json:"params"`
	} `json:"deliveries"`
	ReservationValidForSeconds int         `json:"reservationValidForSeconds"`
	Messages                   interface{} `json:"messages"`
}
