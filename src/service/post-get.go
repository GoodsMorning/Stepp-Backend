package service

import (
	"context"
	basepayload "stepp-backend/src/model/base"
)

func (s *service) GetPost(ctx context.Context, postId int) (*basepayload.BasePayload, error) {

	//response := &response2.PostResponse{}
	//postDb, err := s.repository.GetPost(ctx, postId)

	response, err := s.repository.GetTestPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	return basepayload.GetSuccessResponse(response), nil
}
