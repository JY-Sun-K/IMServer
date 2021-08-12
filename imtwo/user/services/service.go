package services

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"imdemo/imtwo/user/dao"
	"imdemo/imtwo/user/pb"
	"imdemo/imtwo/user/protocol"
	"imdemo/imtwo/user/timeline"
	"imdemo/imtwo/user/utils"
	"log"
)

type UserService struct {
	UserDAO dao.UserDAO
	pb.UnimplementedUserServiceServer
}



func(s *UserService) Login(ctx context.Context,r *pb.LoginRequest) (*pb.LoginResponse, error){
	if r.Email==""||r.Password=="" {
		return &pb.LoginResponse{
			UserId:   0,
			UserName: "",
			Email:    "",
			Err:      "参数错误",
			Code:     400,
			Token:    "",
		},nil
	}
	userInfo,err:=s.UserDAO.IsExistUser(r.Email)
	if userInfo==nil || err!=nil {
		return &pb.LoginResponse{
			UserId:   0,
			UserName: "",
			Email:    "",
			Err:      "参数错误",
			Code:     404,
			Token:    "",
		},nil
	}
	//比较密码是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(r.Password)); err!=nil {
		return nil,err
	}

	user:=timeline.MakeTimeLine(userInfo.UserId)
	dao.RD.AddUser(userInfo.UserId,user)
	token:= utils.CreateToken(userInfo.UserId)
	return &pb.LoginResponse{
		UserId:   userInfo.UserId,
		UserName: userInfo.UserName,
		Email:    userInfo.Email,
		Err:      "",
		Code:     200,
		Token:    token,
	},nil





}


func (s *UserService)Register(ctx context.Context,r *pb.RegisterRequest) (*pb.RegisterResponse, error){
	if r.Email==""||r.Password==""||r.UserName=="" {
		return &pb.RegisterResponse{
			UserId:   0,
			UserName: "",
			Email:    "",
			Err:      "参数错误",
			Code:     400,
			Token:    "",
		},nil
	}
	userInfo,err:=s.UserDAO.IsExistUser(r.Email)
	//log.Println(userInfo,err)
	if !(userInfo==nil && err==nil) {
		return &pb.RegisterResponse{
			UserId:   0,
			UserName: "",
			Email:    "",
			Err:      "参数错误",
			Code:     404,
			Token:    "",
		},nil
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}
	userid, err := s.UserDAO.CreateUser(r.UserName, string(hasedPassword), r.Email)
	if err != nil {
		return nil,err
	}
	token:= utils.CreateToken(userid)
	return &pb.RegisterResponse{
		UserId:   userid,
		UserName: r.UserName,
		Email:    r.Email,
		Err:      "",
		Code:     200,
		Token:    token,
	},nil
}

func (s *UserService) GetUserInfo(ctx context.Context,r *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	_= utils.CheckAuth(ctx)
	userInfo,err:=s.UserDAO.IsExistUser(r.Email)
	if userInfo==nil || err!=nil {
		return &pb.GetUserInfoResponse{
			UserId:   0,
			UserName: "",
			Email:    "",
			Friends:  nil,
			Err:      "认证错误",
			Code:     401,
		},nil
	}

	UserInfo, err := s.UserDAO.GetUserInfo(userInfo.UserId)
	if err != nil {
		return nil,err
	}
	return &pb.GetUserInfoResponse{
		UserId:   UserInfo.UserId,
		UserName: userInfo.UserName,
		Email:    userInfo.Email,
		Friends:  UserInfo.Friends,
		Err:      "",
		Code:     200,
	},nil

}

func (s *UserService) SearchUser(ctx context.Context,r *pb.SearchUserRequest) (*pb.SearchUserResponse, error) {
	_= utils.CheckAuth(ctx)
	users,err:=s.UserDAO.SearchUser(r.UserName)
	if err != nil {
		return nil, err
	}
	n:=len(users)
	if n<=0{
		return &pb.SearchUserResponse{SearchUsers: nil},nil
	}
	userinfo:=make([]*pb.SearchUserInfo,0,n)
	for i := 0; i < n; i++ {
		u:=pb.SearchUserInfo{
			UserId:   users[i].UserId,
			UserName: users[i].UserName,
		}
		userinfo=append(userinfo,&u)
	}

	return &pb.SearchUserResponse{SearchUsers: userinfo}, nil
}

func (s *UserService) AddUser(ctx context.Context,r *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	userId:= utils.CheckAuth(ctx)
	err:=s.UserDAO.AddUser(userId,r.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.AddUserResponse{
		Err:  "",
		Code: 200,
	},nil

}

func (s *UserService) SendMsg(ctx context.Context,r *pb.SendMsgRequest) (*pb.SendMsgResponse, error) {


	l:=&protocol.Letter{
		IPAddress: r.MP.AddressIp,
		From:      r.MP.From,
		To:        r.MP.To,
		Message:   r.MP.Msg,
		SendTime:  r.MP.SendTime,
	}
	WriteChan<-l
	return &pb.SendMsgResponse{
		Err:  "",
		Code: 200,
	},nil

}

func (s *UserService) ReceiveMsg(ctx context.Context,r *pb.ReceiveMsgRequest) (*pb.ReceiveResponse, error) {
	l:=&timeline.Letter{
		IPAddress: r.MP.AddressIp,
		From:      r.MP.From,
		To:        r.MP.To,
		Message:   r.MP.Msg,
		SendTime:  r.MP.SendTime,
	}
	WS.WriteChan<-l
	log.Println("receive msg:",l)
	return &pb.ReceiveResponse{
		Err:  "",
		Code: 200,
	},nil
}