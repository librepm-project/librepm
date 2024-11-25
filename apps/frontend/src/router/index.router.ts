import { createRouter, createWebHistory } from 'vue-router';
import { issueRouter } from './issue.router';
import { projectRouter } from './project.router';
import { adminRouter } from './admin.router';
import { boardRouter } from './board.router';
import { dashboardRouter } from './dashboard.router';
import { filterRouter } from './filter.router';
import { sessionRouter } from './session.router';

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
  ],
});

router.beforeEach((to, _from, next) => {
  const accessToken = localStorage.getItem("accessToken");
  const requiresAuth = to.matched.some((record) => record.meta.authRequired);
  if (requiresAuth && !accessToken) {
    next("login");
  } else {
    next();
  }
});

export default  router;