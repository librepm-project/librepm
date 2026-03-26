package seed

import (
	"apps/api/app/domain"
	"fmt"
	"log/slog"
)

func (s SeedService) createRelatedIssues(items []RelatedIssueData) error {
	issueMap := make(map[string]*domain.IssueModel)

	// First, get all issues into a map by summary
	allIssues, err := s.IssueRepository.All()
	if err != nil {
		return fmt.Errorf("failed to fetch issues: %w", err)
	}

	for i := range *allIssues {
		issueMap[(*allIssues)[i].Summary] = &(*allIssues)[i]
	}

	for _, item := range items {
		issueA, okA := issueMap[item.IssueSummaryA]
		issueB, okB := issueMap[item.IssueSummaryB]

		if !okA {
			slog.Warn("issue not found for related issue", "summary", item.IssueSummaryA)
			continue
		}
		if !okB {
			slog.Warn("issue not found for related issue", "summary", item.IssueSummaryB)
			continue
		}

		_, err := s.RelatedIssueService.Create(issueA.ID, issueB.ID, item.Type)
		if err != nil {
			slog.Warn("failed to create related issue relationship", "error", err)
		}
	}

	return nil
}
