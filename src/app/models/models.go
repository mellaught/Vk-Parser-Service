package models

// Config from config.json file.
type Config struct {
	Token       string
	Version     string
	URL         string
	ServiceHost string
	ServicePort string
	Timeout     int
}

type IntersecReq struct {
	Id      int64 `json:"id"`               // User id
	N       int   `json:"intersect_number"` // current minimum number(N) of occurrences.
	Sex     int   `json:"sex"`              // sex: 1 - woman, 2 - man
	Message bool  `json:"message"`
}

type Answer struct {
	Text     string      `json:"text"`
	Responce interface{} `json:"responce"`
}

type Members struct {
	Data struct {
		Count int     `json:"count"`
		Users []int64 `json:"items"`
	} `json:"response"`
}

type User struct {
	Id              int    `json:"id"`
	FirstName       string `json:"first_name"`
	SecondName      string `json:"last_name"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
}

type Groups struct {
	Data struct {
		G struct {
			Count int     `json:"count"`
			Items []int64 `json:"items"`
		} `json:"groups"`
	} `json:"response"`
}
