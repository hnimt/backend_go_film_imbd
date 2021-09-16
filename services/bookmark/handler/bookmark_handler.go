package handler

import (
	"context"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/repo"
	"micro_backend_film/services/bookmark/pb"

	"github.com/google/uuid"
)

type BookmarkHandler struct {
	BMRepo *repo.BookmarkRepo
}

func (bh *BookmarkHandler) AddBookmark(ctx context.Context, req *pb.ReqAddBookmark) (*pb.ResAddBookmark, error) {
	bookmark := entity.Bookmark{
		BID:    uuid.NewString(),
		UserID: req.UserID,
		FilmID: req.FilmID,
	}

	res, err := bh.BMRepo.AddBM(bookmark)
	if err != nil {
		return nil, err
	}

	result := &pb.ResAddBookmark{
		Film: &pb.Film{
			Name:   res.Film.Name,
			Year:   res.Film.Year,
			Rating: res.Film.Rating,
		},
	}

	return result, nil
}

func (bh *BookmarkHandler) DelBookmark(ctx context.Context, req *pb.ReqDelBookmark) (*pb.ResDelBookmark, error) {
	err := bh.BMRepo.RemoveByUserAndFilm(req.UserID, req.FilmID)
	if err != nil {
		return nil, err
	}

	return &pb.ResDelBookmark{}, nil
}

func (bh *BookmarkHandler) FindByUser(ctx context.Context, req *pb.ReqFindByUser) (*pb.ResFindByUser, error) {
	bookmarks, err := bh.BMRepo.FindBMByUser(req.UserID)
	if err != nil {
		return nil, err
	}

	res := &pb.ResFindByUser{}

	for _, v := range bookmarks {
		res.Films = append(res.Films, &pb.Film{
			Name:   v.Film.Name,
			Year:   v.Film.Year,
			Rating: v.Film.Rating,
		})
	}

	return res, nil
}
