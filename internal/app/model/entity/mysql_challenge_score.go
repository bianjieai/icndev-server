package entity

const TableNameChallengeScore = "challenge_score"

type ChallengeScore struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Rank          int    `gorm:"column:rank" json:"rank"`
	TeamName      string `gorm:"column:team_name" json:"team_name"`
	TaskCompleted string `gorm:"column:task_completed" json:"task_completed"`
	FinalScore    int    `gorm:"column:final_score" json:"final_score"`
	UpdateTime    int64  `gorm:"column:update_time" json:"update_time"`
}

func (*ChallengeScore) TableName() string {
	return TableNameChallengeScore
}
