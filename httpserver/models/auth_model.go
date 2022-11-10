package models

type JwtDecodedModel struct {
	Data Data  `json:"data"`
	Exp  int64 `json:"exp"`
}

type Data struct {
	CreatedAt    string      `json:"CreatedAt"`
	DeletedAt    interface{} `json:"DeletedAt"`
	ID           int64       `json:"ID"`
	UpdatedAt    string      `json:"UpdatedAt"`
	Age          int64       `json:"age"`
	Comments     interface{} `json:"comments"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	Photos       interface{} `json:"photos"`
	SocialMedias interface{} `json:"social_medias"`
	Username     string      `json:"username"`
}
