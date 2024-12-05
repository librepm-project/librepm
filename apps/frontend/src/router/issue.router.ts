import { RouteRecordRaw } from 'vue-router';
import IssueIndexPage from '@/page/issue/IssueIndexPage.vue';
import IssueShowPage from '@/page/issue/IssueShowPage.vue';

export const issueRouter: RouteRecordRaw[] = [
  {
    path: '/issue',
    name: 'issueIndex',
    component: IssueIndexPage,
  },
  {
    path: '/issue/:issueId',
    name: 'issueShow',
    component: IssueShowPage,
  },
];
