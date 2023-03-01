package dto

import (
	"strings"
)

type ChallengeScoreDTO struct {
	ScoreRank  Scores `json:"score_rank"`
	UpdateTime int64  `json:"update_time"`
}

func (p Scores) Len() int { return len(p) }
func (p Scores) Less(i, j int) bool {
	if p[i].Rank > p[j].Rank {
		return false
	} else if p[i].Rank < p[j].Rank {
		return true
	}

	s3 := strings.ToUpper(p[i].TeamName)
	s4 := strings.ToUpper(p[j].TeamName)

	temp1 := []byte(p[i].TeamName)
	temp2 := []byte(p[j].TeamName)

	lcc1 := []byte(s3)
	lcc2 := []byte(s4)
	for i, _ := range lcc1 {
		if lcc1[i] >= 48 && lcc1[i] <= 57 {
			lcc1[i] += 100
		}
	}

	for i, _ := range lcc2 {
		if lcc2[i] >= 48 && lcc2[i] <= 57 {
			lcc2[i] += 100
		}
	}

	less := false
	for i, _ := range lcc1 {
		if i >= len(lcc2) {
			less = false
			break
		} else if lcc1[i] > lcc2[i] {
			less = false
			break
		} else if lcc1[i] < lcc2[i] {
			less = true
			break
		} else {
			if temp1[i] > temp2[i] {
				less = false
				break
			} else if temp1[i] < temp2[i] {
				less = true
				break
			}
			continue
		}
	}

	return less
}
func (p Scores) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Scores []ScoreRank

type ScoreRank struct {
	Rank          int    `json:"rank"`
	TeamName      string `json:"team_name"`
	TaskCompleted string `json:"task_completed"`
	FinalScore    int    `json:"final_score"`
}
