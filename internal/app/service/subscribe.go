package service

import (
	"errors"
	"github.com/bianjieai/icndev-server/internal/app/model/dto"
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/model/vo"
	"github.com/bianjieai/icndev-server/internal/app/repository/cache"
	"github.com/bianjieai/icndev-server/utils"
	"sort"
	"time"
)

type SubscribeService struct {
	subscribeRepo     *cache.SubscribeEmailCacheRepo
	challengeRepo     *cache.ChallengeScoreCacheRepo
	specialAwardsRepo *cache.SpecialAwardsCacheRepo
}

func NewSubscribeService(subscribeRepo *cache.SubscribeEmailCacheRepo, challengeRepo *cache.ChallengeScoreCacheRepo, specialAwardsRepo *cache.SpecialAwardsCacheRepo) *SubscribeService {
	return &SubscribeService{subscribeRepo: subscribeRepo, challengeRepo: challengeRepo, specialAwardsRepo: specialAwardsRepo}
}

func (svc *SubscribeService) SubscribeEmail(req *vo.CreateSubscribeEmailReq) error {
	var se entity.SubscribeEmail
	se.Email = req.Email
	se.CreateAt = time.Now().Unix()
	err := svc.subscribeRepo.Create(&se)
	return err
}

type ScoreNum struct {
	Score int
	Num   int
}

type ScoreNums []ScoreNum

func (p ScoreNums) Len() int {
	return len(p)
}

func (p ScoreNums) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ScoreNums) Less(i, j int) bool {
	return p[i].Score > p[j].Score
}

func (svc *SubscribeService) ChallengeScore(req vo.ChallengeScoreReq) (*dto.ChallengeScoreDTO, int64, error) {
	var (
		offset, limit int
	)
	if req.TeamName == "" {
		offset, limit, _ = utils.ParsePage(req.Page.Page, req.Page.Size, req.Page.Total)
	}
	res, total, err := svc.challengeRepo.FindAll()
	if err != nil {
		return nil, 0, err
	}
	var challengeScore dto.ChallengeScoreDTO
	if len(res) > 0 {
		var scoreNums ScoreNums
		sMap := make(map[int]int)
		for _, v := range res {
			sMap[v.FinalScore] += 1
		}
		for k, v := range sMap {
			var scoreNum ScoreNum
			scoreNum.Score = k
			scoreNum.Num = v
			scoreNums = append(scoreNums, scoreNum)
		}
		sort.Sort(scoreNums)

		challengeScore.UpdateTime = res[0].UpdateTime
		var scores dto.Scores
		for _, v := range res {
			var scoreRank dto.ScoreRank
			rank := 1
			for _, sn := range scoreNums {
				if v.FinalScore == sn.Score {
					scoreRank.Rank = rank
					break
				}
				rank += sn.Num
			}
			//scoreRank.Rank = v.Rank
			scoreRank.TeamName = v.TeamName
			scoreRank.TaskCompleted = v.TaskCompleted
			scoreRank.FinalScore = v.FinalScore
			scores = append(scores, scoreRank)
		}
		sort.Sort(&scores)

		if req.TeamName != "" {
			index := -1
			for i, v := range scores {
				if v.TeamName == req.TeamName {
					index = i
					break
				}
			}
			if index == -1 {
				return nil, 0, nil
			}

			page := index / req.Page.Size
			req.Page.Page = page + 1
			offset, limit, _ = utils.ParsePage(page+1, req.Page.Size, true)
			if offset+limit >= len(res) {
				challengeScore.ScoreRank = scores[offset:len(res)]
			} else {
				challengeScore.ScoreRank = scores[offset : offset+limit]
			}
			return &challengeScore, total, nil
		} else {
			if offset+limit >= len(res) {
				challengeScore.ScoreRank = scores[offset:len(res)]
			} else {
				challengeScore.ScoreRank = scores[offset : offset+limit]
			}
			return &challengeScore, total, nil
		}
	} else {
		return nil, 0, errors.New("no data in db")
	}
}

func (svc *SubscribeService) SpecialAwards() (*dto.ChallengeScoreDTO, error) {
	res, _, err := svc.specialAwardsRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var challengeScore dto.ChallengeScoreDTO
	if len(res) > 0 {
		var scoreNums ScoreNums
		sMap := make(map[int]int)
		for _, v := range res {
			sMap[v.FinalScore] += 1
		}
		for k, v := range sMap {
			var scoreNum ScoreNum
			scoreNum.Score = k
			scoreNum.Num = v
			scoreNums = append(scoreNums, scoreNum)
		}
		sort.Sort(scoreNums)

		challengeScore.UpdateTime = res[0].UpdateTime
		var scores dto.Scores
		for _, v := range res {
			var scoreRank dto.ScoreRank
			rank := 1
			for _, sn := range scoreNums {
				if v.FinalScore == sn.Score {
					scoreRank.Rank = rank
					break
				}
				rank += sn.Num
			}
			//scoreRank.Rank = v.Rank
			scoreRank.TeamName = v.TeamName
			scoreRank.TaskCompleted = v.TaskCompleted
			scoreRank.FinalScore = v.FinalScore
			scores = append(scores, scoreRank)
		}
		sort.Sort(&scores)
		challengeScore.ScoreRank = scores
		return &challengeScore, nil
	} else {
		return nil, errors.New("no data in db")
	}
}
