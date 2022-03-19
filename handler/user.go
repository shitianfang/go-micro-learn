package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"user/domain/model"
	"user/domain/service"
	user "user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, req *user.Request, rsp *user.Response) error {
	user := &model.User{
		UserName: req.Name,
	}
	_, err := u.UserDataService.AddUser(user)
	if err != nil {
		return err
	}
	rsp.Msg = "add user success" // 不用返回resp，只需要把proto上的resp对应信息或整体替换掉就自动返回
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
	log.Infof("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&user.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
