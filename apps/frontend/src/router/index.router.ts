import { createRouter, createWebHistory } from 'vue-router';
import { issueRouter } from './issue.router';
import { projectRouter } from './project.router';
import { adminRouter } from './admin.router';
import { boardRouter } from './board.router';
import { dashboardRouter } from './dashboard.router';
import { filterRouter } from './filter.router';
import { sessionRouter } from './session.router';

export default createRouter({
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
