package songGRPC

import (
	"GetSongsDataService/internal/song"
	logrusCustom "GetSongsDataService/pkg/logger"
	"context"
	"fmt"
	songProtoBuf "github.com/DanKo-code/Protobuf-For-Songs-Library-Upgraded/protos/gen/go/song"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type serverGRPC struct {
	songProtoBuf.UnimplementedSongDataServer
	usecase song.MusixmatchUseCase
}

func Register(gRPC *grpc.Server, usecase song.MusixmatchUseCase) {
	songProtoBuf.RegisterSongDataServer(gRPC, &serverGRPC{
		usecase: usecase,
	})
}

func (s serverGRPC) GetSongData(ctx context.Context, request *songProtoBuf.GetSongDataRequest) (*songProtoBuf.GetSongDataResponse, error) {

	logrusCustom.LogWithLocation(logrus.InfoLevel, fmt.Sprintf("Entered GetSongData gRPC handler with parameters: groupName:%s, song:%s", request.GetGroup(), request.GetSongName()))

	reqContext, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	getSongDataResponse, err := s.usecase.GetSongData(reqContext, request.GetGroup(), request.GetSongName())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	getSongDataResponseProto := songProtoBuf.GetSongDataResponse{
		Ip:          getSongDataResponse.Ip,
		Link:        getSongDataResponse.Link,
		ReleaseDate: getSongDataResponse.ReleaseDate,
		TrackName:   getSongDataResponse.TrackName,
		ArtistName:  getSongDataResponse.ArtistName,
	}

	return &getSongDataResponseProto, nil
}
