package service

import (
	"context"
	basepayload "stepp-backend/src/model/base"
	response2 "stepp-backend/src/model/response"
)

func (s *service) GetPost(ctx context.Context, postId int) (*basepayload.BasePayload, error) {

	response := &response2.PostResponse{}
	postDb, err := s.repository.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}
	{
		response.PostId = postDb.PostId
		response.Caption = postDb.Caption
		response.CoverUrl = postDb.CoverUrl
		response.CreatedAt = postDb.CreatedAt
		response.LikeCount = postDb.LikeCount
	}

	postOwnerDb, err := s.repository.GetUser(ctx, int(postDb.OwnerId))
	if err != nil {
		return nil, err
	}
	{
		response.Owner.UserId = postOwnerDb.UserId
		response.Owner.Name = postOwnerDb.Name
		response.Owner.ProfileImgUrl = postOwnerDb.ProfileImgUrl
	}

	originalDb, err := s.repository.GetOriginal(ctx, int(postDb.OriginalId))
	if err != nil {
		return nil, err
	}
	{
		response.Original.OriginalId = originalDb.OriginalId
		response.Original.Name = originalDb.Name
		response.Original.SteppCount = originalDb.SteppCount
		response.Original.TotalStepp = originalDb.TotalStepp
		response.Original.OriginalPostId = originalDb.OriginalPostId
	}

	originalOwnerDb, err := s.repository.GetUser(ctx, int(originalDb.OwnerId))
	if err != nil {
		return nil, err
	}
	{
		response.Original.Owner.UserId = originalOwnerDb.UserId
		response.Original.Owner.Name = originalOwnerDb.Name
		response.Original.Owner.ProfileImgUrl = originalOwnerDb.ProfileImgUrl
	}

	originalLocationDb, err := s.repository.GetLocation(ctx, int(originalDb.LocationId))
	if err != nil {
		return nil, err
	}
	{
		response.Original.Location.LocationId = originalLocationDb.LocationId
		response.Original.Location.Name = originalLocationDb.Name
		response.Original.Location.GeoPoint = originalLocationDb.GeoPoint
		response.Original.Location.SteppCount = originalLocationDb.SteppCount
	}

	return basepayload.GetSuccessResponse(response), nil
}
