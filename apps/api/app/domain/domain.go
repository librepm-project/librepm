package domain

import (
	"libs/mysql_utils"
)

type Domain struct {
	ProjectService   ProjectServiceInterface
	IssueService     IssueServiceInterface
	FilterService    FilterServiceInterface
	BoardService     BoardServiceInterface
	DashboardService DashboardServiceInterface
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
		FilterService: FilterService{
			FilterRepository: FilterRepository{
				DB: DB,
			},
		},
		BoardService: BoardService{
			BoardRepository: BoardRepository{
				DB: DB,
			},
		},
		DashboardService: DashboardService{
			DashboardRepository: DashboardRepository{
				DB: DB,
			},
		},
	}
}
