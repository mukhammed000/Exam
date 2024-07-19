package storage

import (
	pb "learning/genproto/learning"
)

type InitRoot interface {
	Learning() LearningService
}

type LearningService interface {
	GetAllLanguages(req *pb.EmptyReq) (*pb.GetAllLanguagesResp, error)
	StartLearningLanguage(req *pb.StartLearningReq) (*pb.StartLearningResp, error)
	GetLessonInfo(req *pb.ParticipateLessonReq) (*pb.ParticipateLessonResp, error)
	CompleteLesson(req *pb.CompleteLessonReq) (*pb.CompleteLessonResp, error)
	StartExercise(req *pb.StartExerciceReq) (*pb.StartExerciceResp, error)
	AnswerExercise(req *pb.AnswerExerciseReq) (*pb.AnswerExerciseResp, error)
	GetAllVocabulary(req *pb.GetAllVocabularyReq) (*pb.GetAllVocabularyResp, error)
	AddNewWord(req *pb.AddNewWordReq) (*pb.AddNewWordResp, error)
}
