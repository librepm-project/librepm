import { RouteRecordRaw } from 'vue-router';
import IssueIndexPage from '@/page/issue/IssueIndexPage.vue';
import IssueShowPage from '@/page/issue/IssueShowPage.vue';
import IssueCreatePage from '@/page/issue/IssueCreatePage.vue';

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
    name: 'issueShowById',
    component: IssueShowPage,
  },
  {
    path: '/issue/key/:key',
    name: 'issueShow',
    component: IssueShowPage,
  },
];
