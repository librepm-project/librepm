package seed

import (
	"apps/api/app/domain"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-playground/validator/v10"
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
	RelatedIssueService                domain.RelatedIssueServiceInterface
	RelatedIssueRepository             domain.RelatedIssueRepositoryInterface
	PurgeRepository                    PurgeRepositoryInterface
}

func (s SeedService) Seed(filePath string) []error {
	var errors []error

	seedData, err := s.getSeedData(filePath)

	if err != nil {
		panic(err)
	}
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

	err = s.createIssue(seedData.Issues)
	errors = append(errors, err)

	err = s.createRelatedIssues(seedData.RelatedIssues)
	errors = append(errors, err)

	slog.Info("seed completed successfully")
	return errors
}

func (s SeedService) Purge() {
	s.PurgeRepository.Purge()
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

	validate := validator.New()
	if err := validate.Struct(seedData); err != nil {
		return seedData, fmt.Errorf("validation failed: %w", err)
	}

	return seedData, nil
}
