package service

import (
	"context"
	"log"

	pb "learning/genproto/learning"
	stg "learning/storage"
)

type LearningService struct {
	stg stg.InitRoot
	pb.UnimplementedLearningServiceServer
}

func NewLearningService(stg stg.InitRoot) *LearningService {
	return &LearningService{stg: stg}
}

func (s *LearningService) GetAllLanguages(ctx context.Context, req *pb.EmptyReq) (*pb.GetAllLanguagesResp, error) {
	resp, err := s.stg.Learning().GetAllLanguages(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) StartLearningLanguage(ctx context.Context, req *pb.StartLearningReq) (*pb.StartLearningResp, error) {
	resp, err := s.stg.Learning().StartLearningLanguage(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetLessonInfo(ctx context.Context, req *pb.ParticipateLessonReq) (*pb.ParticipateLessonResp, error) {
	resp, err := s.stg.Learning().GetLessonInfo(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) CompleteLesson(ctx context.Context, req *pb.CompleteLessonReq) (*pb.CompleteLessonResp, error) {
	resp, err := s.stg.Learning().CompleteLesson(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) StartExercise(ctx context.Context, req *pb.StartExerciceReq) (*pb.StartExerciceResp, error) {
	resp, err := s.stg.Learning().StartExercise(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) AnswerExercise(ctx context.Context, req *pb.AnswerExerciseReq) (*pb.AnswerExerciseResp, error) {
	resp, err := s.stg.Learning().AnswerExercise(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetAllVocabulary(ctx context.Context, req *pb.GetAllVocabularyReq) (*pb.GetAllVocabularyResp, error) {
	resp, err := s.stg.Learning().GetAllVocabulary(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) AddNewWord(ctx context.Context, req *pb.AddNewWordReq) (*pb.AddNewWordResp, error) {
	resp, err := s.stg.Learning().AddNewWord(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
