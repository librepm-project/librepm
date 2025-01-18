package domain

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SeedServiceInterface interface {
	Seed() []error
}

type SeedService struct {
	BoardRepository                    BoardRepositoryInterface
	BoardColumnRepository              BoardColumnRepositoryInterface
	BoardColumnStatusRepository        BoardColumnStatusRepositoryInterface
	DashboardRepository                DashboardRepositoryInterface
	FilterRepository                   FilterRepositoryInterface
	FilterConditionRepository          FilterConditionRepositoryInterface
	IssueRepository                    IssueRepositoryInterface
	ProjectRepository                  ProjectRepositoryInterface
	ProjectTrackerRepository           ProjectTrackerRepositoryInterface
	ProjectTrackerStatusRepository     ProjectTrackerStatusRepositoryInterface
	ProjectTrackerTransitionRepository ProjectTrackerTransitionRepositoryInterface
	ProjectUserRepository              ProjectUserRepositoryInterface
	StatusRepository                   StatusRepositoryInterface
	TrackerRepository                  TrackerRepositoryInterface
	TransitionRepository               TransitionRepositoryInterface
	UserRepository                     UserRepositoryInterface
}

func (s SeedService) Seed() []error {
	var errors []error

	seedData, err := s.getSeedData()
	errors = append(errors, err)

	err = s.createStatus(seedData.Statuses)
	errors = append(errors, err)

	err = s.createBoard(seedData.Boards)
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

	fmt.Println("Seed completed successfully")
	return errors
}

type SeedData struct {
	Statuses   []StatusData    `yaml:"statuses"`
	Boards     []BoardData     `yaml:"boards"`
	Users      []UserData      `yaml:"users"`
	Dashboards []DashboardData `yaml:"dashboards"`
	Filters    []FilterData    `yaml:"filters"`
	Projects   []ProjectData   `yaml:"projects"`
}

func (s SeedService) getSeedData() (SeedData, error) {
	var seedData SeedData

	file, err := os.ReadFile("/app/apps/api/seed.yaml")

	if err != nil {
		return seedData, fmt.Errorf("failed to read seed data YAML: %w", err)
	}

	if err = yaml.Unmarshal(file, &seedData); err != nil {
		return seedData, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	return seedData, nil
}
