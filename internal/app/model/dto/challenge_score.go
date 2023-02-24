package dto

type ChallengeScoreDTO struct {
	ScoreRank  []ScoreRank `json:"score_rank"`
	UpdateTime int64       `json:"update_time"`
}

type ScoreRank struct {
	Rank          int    `json:"rank"`
	TeamName      string `json:"team_name"`
	TaskCompleted string `json:"task_completed"`
	FinalScore    int    `json:"final_score"`
}
