import { createRouter, createWebHistory } from 'vue-router';
import { issueRouter } from '@/router/issue.router';
import { adminProjectRouter } from '@/router/admin.project.router';
import { adminBoardRouter } from '@/router/admin.board.router';
import { adminUserRouter } from '@/router/admin.user.router';
import { boardRouter } from '@/router/board.router';
import { dashboardRouter } from '@/router/dashboard.router';
import { filterRouter } from '@/router/filter.router';
import { sessionRouter } from '@/router/session.router';
import { adminStatusRouter } from '@/router/admin.status.router';
import { adminTrackerRouter } from '@/router/admin.tracker.router';
import { adminPriorityRouter } from '@/router/admin.priority.router';
import { adminReleaseRouter } from '@/router/admin.release.router';
import { authorizationMiddleware } from '@/router/middleware/authorization.middleware';
import { pageTitleMiddleware } from '@/router/middleware/page-title.middleware';
import { layoutMiddleware } from '@/router/middleware/layout.middleware';
import { permissionMiddleware } from '@/router/middleware/permission.middleware';
import AdminSettingPage from '@/page/admin/setting/AdminSettingPage.vue'; // Import the new page

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    ...boardRouter,
    ...dashboardRouter,
    ...filterRouter,
    ...issueRouter,
    ...sessionRouter,
    ...adminProjectRouter,
    ...adminStatusRouter,
    ...adminTrackerRouter,
    ...adminPriorityRouter,
    ...adminReleaseRouter,
    ...adminUserRouter,
    ...adminBoardRouter,
    {
      path: '/admin/settings',
      name: 'admin-settings',
      component: AdminSettingPage,
      meta: { title: 'admin.settings', permission: 'setting:read' as const },
    },
  ],
});

router.beforeEach(layoutMiddleware);
router.beforeEach(pageTitleMiddleware);
router.beforeEach(authorizationMiddleware);
router.beforeEach(permissionMiddleware);

export default router;