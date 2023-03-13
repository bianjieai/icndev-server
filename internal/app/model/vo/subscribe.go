package vo

type CreateSubscribeEmailReq struct {
	Email string `form:"email" json:"email"`
}

type ChallengeScoreReq struct {
	TeamName string `form:"team_name" json:"team_name"`
	Page     *PageVO
}
