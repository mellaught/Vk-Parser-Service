package models

// Config from config.json file.
type Config struct {
	Token       []string
	Version     string
	URL         string
	ServiceHost string
	ServicePort string
}

type IntersecReq struct {
	Id      int64 `json:"id"`               // User id
	N       int   `json:"intersect_number"` // current minimum number(N) of occurrences.
	Sex     int   `json:"sex"`              // sex: 1 - woman, 2 - man
	Message bool  `json:"message"`          // can write private message
}

type Answer struct {
	Text     string      `json:"text"`
	Responce interface{} `json:"responce"`
}

// Post object describes a wall post and contains the following fields:
type Post struct {
	Id       int64  `json:"id"`       // 	Post ID on the wall.
	Owner_id int64  `json:"owner_id"` // Wall owner ID.
	Date     int64  `json:"date"`     // Date (in Unix time) the post was added.
	Text     string `json:"text"`     // 	Text of the post.
	Comments struct {
		Count int64 `json:"count"` // Number of comments.
	} `json:"comments"` //Information about comments to the post; an object containing:
	Like struct {
		Count int64 `json:"count"` // Number of users who liked the post.
	} `json:"likes"` // Information about likes to the post; an object containing
	Views struct {
		Count int64 `json:"count"` // Number of users who viewed the post.
	}
	PostType string `json:"post_type"` // Type of the post, can be: post, copy, reply, postpone, suggest.

}

type Members struct {
	Data struct {
		Count int    `json:"count"`
		Users []User `json:"items"`
	} `json:"response"`
}

type User struct {
	Id              int    `json:"id"`
	Sex             int    `json:"sex"`
	CanWrite        int    `json:"can_write_private_message"`
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
