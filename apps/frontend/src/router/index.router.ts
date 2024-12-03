import { createRouter, createWebHistory } from 'vue-router';
import { issueRouter } from '@/router/issue.router';
import { projectRouter } from '@/router/project.router';
import { adminRouter } from '@/router/admin.router';
import { boardRouter } from '@/router/board.router';
import { dashboardRouter } from '@/router/dashboard.router';
import { filterRouter } from '@/router/filter.router';
import { sessionRouter } from '@/router/session.router';
import { statusRouter } from '@/router/status.router';
import { trackerRouter } from '@/router/tracker.router';
import { authorizationMiddleware } from '@/router/middleware/authorization.middleware';
import { pageTitleMiddleware } from '@/router/middleware/page-title.middleware';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    ...adminRouter,
    ...adminRouter,
    ...boardRouter,
    ...dashboardRouter,
    ...filterRouter,
    ...issueRouter,
    ...sessionRouter,
    ...projectRouter,
    ...statusRouter,
    ...trackerRouter,
  ],
});

router.beforeEach(pageTitleMiddleware)
router.beforeEach(authorizationMiddleware);

export default router;