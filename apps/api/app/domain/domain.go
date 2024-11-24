package domain

import (
	"libs/mysql_utils"
)

type Domain struct {
	ProjectService ProjectServiceInterface
	IssueService   IssueServiceInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateProductDatabase(DB)
	return Domain{
		ProjectService: ProjectService{
			ProjectRepository: ProjectRepository{
				DB: DB,
			},
		},
		IssueService: IssueService{
			IssueRepository: IssueRepository{
				DB: DB,
			},
		},
	}
}
