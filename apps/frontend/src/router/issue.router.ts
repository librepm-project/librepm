import { RouteRecordRaw } from 'vue-router';
import IssueIndexPage from '@/page/issue/IssueIndexPage.vue';
import IssueShowPage from '@/page/issue/IssueShowPage.vue';
import IssueCreatePage from '@/page/issue/IssueCreatePage.vue';
import IssueUpdatePage from '@/page/issue/IssueUpdatePage.vue';

export const issueRouter: RouteRecordRaw[] = [
  {
    path: '/issue',
    name: 'issueIndex',
    component: IssueIndexPage,
  },
  {
    path: '/issue/create',
    name: 'issueCreate',
    component: IssueCreatePage,
  },
  {
    path: '/issue/:issueId',
    name: 'issueShow',
    component: IssueShowPage,
  },
  {
    path: '/issue/:issueId/edit',
    name: 'issueUpdate',
    component: IssueUpdatePage,
  },
];
