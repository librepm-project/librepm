package domain

import (
	"github.com/google/uuid"
)

type TrackerServiceInterface interface {
	All() *[]TrackerModel
	Show(tracker_id uuid.UUID) *TrackerModel
	Create(tracker *TrackerModel)
	Update(tracker_id uuid.UUID, tracker *TrackerModel)
	Destroy(tracker_id uuid.UUID)
}

type TrackerService struct {
	TrackerRepository TrackerRepositoryInterface
}

func (s TrackerService) All() *[]TrackerModel {
	trackers := s.TrackerRepository.All()
	return trackers
}

func (s TrackerService) Show(tracker_id uuid.UUID) *TrackerModel {
	return s.TrackerRepository.FindByID(tracker_id)
}

func (s TrackerService) Create(tracker *TrackerModel) {
	s.TrackerRepository.Create(tracker)
}

func (s TrackerService) Update(tracker_id uuid.UUID, tracker *TrackerModel) {
	s.TrackerRepository.Update(tracker_id, tracker)
}

func (s TrackerService) Destroy(tracker_id uuid.UUID) {
	s.TrackerRepository.Destroy(tracker_id)
}
