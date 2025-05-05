package seed

import (
	"apps/api/app/domain"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SeedServiceInterface interface {
	Seed() []error
	Purge()
}

type SeedService struct {
	BoardRepository                    domain.BoardRepositoryInterface
	BoardColumnRepository              domain.BoardColumnRepositoryInterface
	BoardColumnStatusRepository        domain.BoardColumnStatusRepositoryInterface
	DashboardRepository                domain.DashboardRepositoryInterface
	FilterRepository                   domain.FilterRepositoryInterface
	FilterConditionRepository          domain.FilterConditionRepositoryInterface
	IssueRepository                    domain.IssueRepositoryInterface
	ProjectRepository                  domain.ProjectRepositoryInterface
	ProjectTrackerRepository           domain.ProjectTrackerRepositoryInterface
	ProjectTrackerStatusRepository     domain.ProjectTrackerStatusRepositoryInterface
	ProjectTrackerTransitionRepository domain.ProjectTrackerTransitionRepositoryInterface
	ProjectUserRepository              domain.ProjectUserRepositoryInterface
	StatusRepository                   domain.StatusRepositoryInterface
	TrackerRepository                  domain.TrackerRepositoryInterface
	TransitionRepository               domain.TransitionRepositoryInterface
	UserRepository                     domain.UserRepositoryInterface
	PurgeRepository                    PurgeRepositoryInterface
}

func (s SeedService) Seed(filePath string) []error {
	var errors []error

	seedData, err := s.getSeedData(filePath)
	errors = append(errors, err)

	err = s.createTracker(seedData.Trackers)
	errors = append(errors, err)

	err = s.createStatus(seedData.Statuses)
	errors = append(errors, err)

	err = s.createUser(seedData.Users)
	errors = append(errors, err)

	err = s.createProject(seedData.Projects)
	errors = append(errors, err)

	err = s.createBoard(seedData.Boards)
	errors = append(errors, err)

	err = s.createDashboard(seedData.Dashboards)
	errors = append(errors, err)

	err = s.createFilter(seedData.Filters)
	errors = append(errors, err)

	fmt.Println("Seed completed successfully")
	return errors
}

func (s SeedService) Purge() {
	s.PurgeRepository.Purge()
}

type SeedData struct {
	Trackers   []TrackerData   `yaml:"trackers"`
	Statuses   []StatusData    `yaml:"statuses"`
	Boards     []BoardData     `yaml:"boards"`
	Users      []UserData      `yaml:"users"`
	Dashboards []DashboardData `yaml:"dashboards"`
	Filters    []FilterData    `yaml:"filters"`
	Projects   []ProjectData   `yaml:"projects"`
}

func (s SeedService) getSeedData(filePath string) (SeedData, error) {
	var seedData SeedData

	file, err := os.ReadFile(filePath)

	if err != nil {
		return seedData, fmt.Errorf("failed to read seed data YAML: %w", err)
	}

	if err = yaml.Unmarshal(file, &seedData); err != nil {
		return seedData, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	return seedData, nil
}
