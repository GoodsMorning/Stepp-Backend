package service

import (
	"context"
	basepayload "stepp-backend/src/model/base"
)

func (s *service) GetPost(ctx context.Context, postId int) (*basepayload.BasePayload, error) {

	response, err := s.repository.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	return basepayload.GetSuccessResponse(response), nil
}
