package song

import (
	"GetSongsDataService/internal/song/dtos"
	"context"
)

type MusixmatchUseCase interface {
	GetSongData(ctx context.Context, groupName, song string) (*dtos.GetSongDataResponse, error)
	GetLyrics(ctx context.Context, ip string) (string, error)
}
