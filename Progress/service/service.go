package service

import (
	"context"
	"log"

	pb "progress/genproto/progress"
	stg "progress/storage"
)

type ProgressService struct {
    stg stg.InitRoot
    pb.UnimplementedProgressServiceServer // Embed UnimplementedProgressServiceServer
}

func NewProgressService(stg stg.InitRoot) *ProgressService {
    return &ProgressService{stg: stg}
}

// Implementing ProgressService interface

func (s *ProgressService) GetProgressByLanguage(ctx context.Context, req *pb.ProgressByLanguageReq) (*pb.ProgressByLanguageResponse, error) {
    resp, err := s.stg.Progress().GetProgressByLanguage(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetLeaderboardByLanguage(ctx context.Context, req *pb.ProgressByLanguageRequest) (*pb.LeaderboardResponse, error) {
    resp, err := s.stg.Progress().GetLeaderboardByLanguage(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetSkillsByLanguage(ctx context.Context, req *pb.ProgressByLanguageReq) (*pb.SkillsResponse, error) {
    resp, err := s.stg.Progress().GetSkillsByLanguage(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetDailyProgress(ctx context.Context, req *pb.UserId) (*pb.DailyProgressResponse, error) {
    resp, err := s.stg.Progress().GetDailyProgress(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetWeeklyProgress(ctx context.Context, req *pb.UserId) (*pb.WeeklyProgressResponse, error) {
    resp, err := s.stg.Progress().GetWeeklyProgress(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetMonthlyProgress(ctx context.Context, req *pb.UserId) (*pb.MonthlyProgressResponse, error) {
    resp, err := s.stg.Progress().GetMonthlyProgress(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}

func (s *ProgressService) GetAchievements(ctx context.Context, req *pb.UserId) (*pb.AchievementsResponse, error) {
    resp, err := s.stg.Progress().GetAchievements(req)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return resp, nil
}
