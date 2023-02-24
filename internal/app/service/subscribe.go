package service

import (
	"errors"
	"github.com/bianjieai/icndev-server/internal/app/model/dto"
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/model/vo"
	"github.com/bianjieai/icndev-server/internal/app/repository/cache"
	"github.com/bianjieai/icndev-server/utils"
	"time"
)

type SubscribeService struct {
	subscribeRepo *cache.SubscribeEmailCacheRepo
	challengeRepo *cache.ChallengeScoreCacheRepo
}

func NewSubscribeService(subscribeRepo *cache.SubscribeEmailCacheRepo, challengeRepo *cache.ChallengeScoreCacheRepo) *SubscribeService {
	return &SubscribeService{subscribeRepo: subscribeRepo, challengeRepo: challengeRepo}
}

func (svc *SubscribeService) SubscribeEmail(req *vo.CreateSubscribeEmailReq) error {
	var se entity.SubscribeEmail
	se.Email = req.Email
	se.CreateAt = time.Now().Unix()
	err := svc.subscribeRepo.Create(&se)
	return err
}

func (svc *SubscribeService) ChallengeScore(req vo.PageVO) (*dto.ChallengeScoreDTO, int64, error) {
	offset, limit, _ := utils.ParsePage(req.Page, req.Size, req.Total)
	res, total, err := svc.challengeRepo.FindByLimit(offset, limit)
	if err != nil {
		return nil, 0, err
	}
	var challengeScore dto.ChallengeScoreDTO
	if len(res) > 0 {
		challengeScore.UpdateTime = res[0].UpdateTime
		for _, v := range res {
			var scoreRank dto.ScoreRank
			scoreRank.Rank = v.Rank
			scoreRank.TeamName = v.TeamName
			scoreRank.TaskCompleted = v.TaskCompleted
			scoreRank.FinalScore = v.FinalScore
			challengeScore.ScoreRank = append(challengeScore.ScoreRank, scoreRank)
		}
		return &challengeScore, total, nil
	} else {
		return nil, 0, errors.New("no data in db")
	}
}
