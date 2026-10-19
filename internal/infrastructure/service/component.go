package service

import (
	"context"
	"fmt"
	"heating-prediction/internal/currentuser"
	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/domain/ports/repo"
	"heating-prediction/internal/domain/rules"
	"heating-prediction/internal/infrastructure/dto"
	"io"
	"strings"

	"github.com/google/uuid"
)

type ComponentService struct {
	componentRepo repo.ComponentRepo
	likeRepo      repo.LikeRepo
	fileRepo      repo.S3Repo
}

func NewComponentService(
	componentRepo repo.ComponentRepo,
	likeRepo repo.LikeRepo,
	fileRepo repo.S3Repo,
) *ComponentService {
	return &ComponentService{
		componentRepo: componentRepo,
		likeRepo:      likeRepo,
		fileRepo:      fileRepo,
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

func fileName(original string) string {
	ext := ""
	if i := strings.LastIndex(original, "."); i >= 0 {
		ext = original[i:]
	}
	return uuid.NewString() + ext
}

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
func (s *ComponentService) uploadFile(
	ctx context.Context,
	fileName string,
	file io.Reader,
	contentType models.ContentType,
	size int64,
) error {
	validatedFile, err := rules.ValidateContent(contentType, file, size)
	if err != nil {
		return fmt.Errorf("validate file: %w", err)
	}

	metadata := &models.ContentMetadata{
		URL:      fileName,
		File:     validatedFile.Reader,
		Type:     contentType,
		MimeType: validatedFile.FileType,
		Size:     validatedFile.Size,
	}

	if err := s.fileRepo.Upload(ctx, metadata); err != nil {
		return fmt.Errorf("upload file: %w", err)
	}

	return nil
}

func (s *ComponentService) CreateComponent(ctx context.Context, req *dto.ComponentRequest, video io.Reader, image io.Reader) error {
	videoName := fileName(req.VideoFile.Filename)
	if err := s.uploadFile(ctx, videoName, video, models.Video, req.VideoFile.Size); err != nil {
		return fmt.Errorf("upload video: %w", err)
	}

	imageName := fileName(req.ImageFile.Filename)
	if err := s.uploadFile(ctx, imageName, image, models.Image, req.ImageFile.Size); err != nil {
		return fmt.Errorf("upload image: %w", err)
	}
	component := &models.Component{
		ComponentName:        req.ComponentName,
		ComponentDescription: req.ComponentDescription,
		ComponentStatus:      models.StatusPublished,
		ImageURL:             imageName,
		VideoURL:             videoName,
		ComponentType:        req.ComponentType,
		CreatorID:            1, //TODO сделать идентификацию
		PowerDissipation:     req.PowerDissipation,
		ThermalResistance:    req.ThermalResistance,
	}

	if err := s.componentRepo.CreateComponent(ctx, component); err != nil {
		return fmt.Errorf("create component: %w", err)
	}
	return nil
}
