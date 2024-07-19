package service

import (
	"context"
	"log"

	pb "users/genproto/users"
	"users/kafka"
	stg "users/storage"
)

type UserService struct {
	stg           stg.InitRoot
	kafkaProducer *kafka.KafkaProducer
	pb.UnimplementedUsersServiceServer
}

func NewUserService(stg stg.InitRoot, kafkaProducer *kafka.KafkaProducer) *UserService {
	return &UserService{
		stg:           stg,
		kafkaProducer: kafkaProducer,
	}
}

// sendKafkaMessage sends a message to a specified Kafka topic
func (s *UserService) sendKafkaMessage(topic, message string) {
	err := s.kafkaProducer.SendMessage(topic, message)
	if err != nil {
		log.Printf("Error sending Kafka message to topic %s: %v", topic, err)
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	resp, err := s.stg.Users().RegisterUser(req)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return nil, err
	}

	// Send Kafka message for user registration
	message := "User registered: " + req.Email
	s.sendKafkaMessage("user_events_topic", message)

	return resp, nil
}

func (s *UserService) LogIn(ctx context.Context, req *pb.LogInReq) (*pb.LogInResp, error) {
	resp, err := s.stg.Users().LogIn(req)
	if err != nil {
		log.Printf("Error logging in user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) GetProfileInfo(ctx context.Context, req *pb.GetProfileInfoReq) (*pb.GetProfileInfoResp, error) {
	resp, err := s.stg.Users().GetProfileInfo(req)
	if err != nil {
		log.Printf("Error getting profile info: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileResp, error) {
	resp, err := s.stg.Users().UpdateProfile(req)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Empty, error) {
	resp, err := s.stg.Users().ChangePassword(req)
	if err != nil {
		log.Printf("Error changing password: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.Empty, error) {
	resp, err := s.stg.Users().ResetPassword(req)
	if err != nil {
		log.Printf("Error resetting password: %v", err)
		return nil, err
	}

	// Send Kafka message for password reset
	message := "Password reset requested for: " + req.Email
	s.sendKafkaMessage("user_events_topic", message)

	return resp, nil
}

func (s *UserService) PasswordRecovery(ctx context.Context, req *pb.RecoveryPasswordReq) (*pb.Empty, error) {
	resp, err := s.stg.Users().PasswordRecovery(req)
	if err != nil {
		log.Printf("Error recovering password: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) GetUserSettings(ctx context.Context, req *pb.GetUsersSettingsReq) (*pb.GetUserSettingsResp, error) {
	resp, err := s.stg.Users().GetUserSettings(req)
	if err != nil {
		log.Printf("Error getting user settings: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateSettings(ctx context.Context, req *pb.UpdateSettingsReq) (*pb.Empty, error) {
	resp, err := s.stg.Users().UpdateSettings(req)
	if err != nil {
		log.Printf("Error updating settings: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) RefreshToken(ctx context.Context, t *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	res, err := s.stg.Users().RefreshToken(t)
	if err != nil {
		log.Printf("Error refreshing token: %v", err)
		return nil, err
	}
	return res, nil
}
