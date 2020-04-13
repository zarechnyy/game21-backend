package model

type GameModelResponse struct {
	Type int `json:"type"` //socket command type
	Model interface{} `json:"model"`
}

type ScoreModel struct {
	Score int `json:"score"`
}