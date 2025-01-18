package domain

import (
	"github.com/google/uuid"
)

type TrackerServiceInterface interface {
	All() (*[]TrackerModel, error)
	Show(tracker_id uuid.UUID) (*TrackerModel, error)
	Create(tracker *TrackerModel) error
	Update(tracker_id uuid.UUID, tracker *TrackerModel) error
	Destroy(tracker_id uuid.UUID) error
}

type TrackerService struct {
	TrackerRepository TrackerRepositoryInterface
}

func (s TrackerService) All() (*[]TrackerModel, error) {
	trackers, err := s.TrackerRepository.All()
	return trackers, err
}

func (s TrackerService) Show(tracker_id uuid.UUID) (*TrackerModel, error) {
	return s.TrackerRepository.FindByID(tracker_id)
}

func (s TrackerService) Create(tracker *TrackerModel) error {
	return s.TrackerRepository.Create(tracker)
}

func (s TrackerService) Update(tracker_id uuid.UUID, tracker *TrackerModel) error {
	return s.TrackerRepository.Update(tracker_id, tracker)
}

func (s TrackerService) Destroy(tracker_id uuid.UUID) error {
	return s.TrackerRepository.Destroy(tracker_id)
}
