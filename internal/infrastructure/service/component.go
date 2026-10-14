package service

import (
	"context"
	"fmt"
	"heating-prediction/internal/currentuser"
	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/domain/ports/repo"
)

type ComponentService struct {
	componentRepo repo.ComponentRepo
	likeRepo      repo.LikeRepo
}

func NewComponentService(
	componentRepo repo.ComponentRepo,
	likeRepo repo.LikeRepo,
) *ComponentService {
	return &ComponentService{
		componentRepo: componentRepo,
		likeRepo:      likeRepo,
	}
}

func (s *ComponentService) PublishDraft(
	ctx context.Context,
	componentID uint16,
	status models.ServiceStatusEnum,
) error {
	component, err := s.componentRepo.ReadComponentByID(ctx, componentID)
	if err != nil {
		return fmt.Errorf("can't access component by id %v: %w", componentID, err)
	}

	if status == models.StatusDraft {
		return fmt.Errorf("can't change status to черновик")
	}

	if component.ComponentStatus != models.StatusDraft {
		return fmt.Errorf("only черновик can be published")
	}

	component.ComponentStatus = status

	if err := s.componentRepo.UpdateComponent(ctx, component); err != nil {
		return fmt.Errorf("can't update component: %w", err)
	}

	return nil
}

// func fileName(original string) string {
// 	ext := ""
// 	if i := strings.LastIndex(original, "."); i >= 0 {
// 		ext = original[i:]
// 	}s
// 	return uuid.NewString() + ext
// }

func (s *ComponentService) GetFeedById(ctx context.Context, componentID uint16, next bool) (*models.Component, error) {
	if next {
		result, err := s.componentRepo.ReadNextPublishedByID(ctx, componentID)
		if err != nil {
			return nil, fmt.Errorf("can't read next published after %d: %w", componentID, err)
		}
		return &result, nil
	}

	result, err := s.componentRepo.ReadPublishedByID(ctx, componentID)
	if err != nil {
		return nil, fmt.Errorf("can't read published component %d: %w", componentID, err)
	}
	return &result, nil
}

func (s *ComponentService) GetFiltredComponentsList(ctx context.Context, minPower float64) ([]models.Component, error) {
	componentsList, err := s.componentRepo.ReadComponentsList(ctx, minPower)
	if err != nil {
		return []models.Component{}, err
	}
	return componentsList, nil
}

func (s *ComponentService) GetDraft(ctx context.Context) (models.Component, error) {
	currentUserID := currentuser.GetCurrentUserID()
	draft, err := s.componentRepo.ReadDraftByUser(ctx, currentUserID)
	if err != nil {
		return models.Component{}, err
	}
	return draft, nil
}

func (s *ComponentService) DeleteComponent(ctx context.Context, id uint16) error {
	component, err := s.componentRepo.ReadComponentByID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't read component %d: %w", id, err)
	}
	if component.ComponentStatus == models.StatusDeleted {
		return fmt.Errorf("component %d already deleted", id)
	}
	return s.componentRepo.DeleteComponentSQL(ctx, int(id))
}

func (s *ComponentService) LikeComponent(ctx context.Context, componentID uint16, value int) error {
	if value != 0 && value != 1 {
		return fmt.Errorf("value must be 0 or 1")
	}
	uid := int(currentuser.GetCurrentUserID())
	cid := int(componentID)

	if value == 1 {
		return s.likeRepo.AddLike(ctx, uid, cid)
	}
	return s.likeRepo.RemoveLike(ctx, uid, cid)
}
