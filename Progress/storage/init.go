package progress

import (
	pb "progress/genproto/progress"
)

type InitRoot interface {
	Progress() ProgressService
}

type ProgressService interface {
    GetProgressByLanguage(req *pb.ProgressByLanguageReq) (*pb.ProgressByLanguageResponse, error)
    GetLeaderboardByLanguage(req *pb.ProgressByLanguageRequest) (*pb.LeaderboardResponse, error)
    GetSkillsByLanguage(req *pb.ProgressByLanguageReq) (*pb.SkillsResponse, error)
    GetDailyProgress(req *pb.UserId) (*pb.DailyProgressResponse, error)
    GetWeeklyProgress(req *pb.UserId) (*pb.WeeklyProgressResponse, error)
    GetMonthlyProgress(req *pb.UserId) (*pb.MonthlyProgressResponse, error)
    GetAchievements(req *pb.UserId) (*pb.AchievementsResponse, error)
}

