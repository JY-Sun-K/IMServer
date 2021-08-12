package services

import (

	"imdemo/imtwo/letterscheduler/pb"
	"imdemo/imtwo/letterscheduler/timeline"
	"io"
	"log"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

func (s *StreamService)WriteStream(stream pb.StreamService_WriteStreamServer) error {
	for  {

		r,err:=stream.Recv()
		log.Println("WriteStream Receive:",r)
		l:=&timeline.Letter{
			IPAddress: r.MP.AddressIp,
			From:      r.MP.From,
			To:        r.MP.To,
			Message:   r.MP.Msg,
			SendTime:  r.MP.SendTime,
		}
		WS.WriteChan<- l
		if err == io.EOF {
			return nil
		}
		if err != nil {
			_=stream.Send(&pb.WriteStreamResponse{
				Err:  err.Error(),
				Code: 500,
			})

			return err
		}
		err=stream.Send(&pb.WriteStreamResponse{
			Err:  "",
			Code: 0,
		})
		if err != nil {
			return err
		}

	}

}

//func (s *StreamService) ReadStream(stream pb.StreamService_ReadStreamServer) error{
//	for  {
//
//		err:=stream.Send(&pb.ReadStreamResponse{MP: })
//
//		_,err=stream.Recv()
//		if err == io.EOF {
//			return nil
//		}
//		if err != nil {
//
//			return err
//		}
//	}
//}
