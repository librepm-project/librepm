import { RouteRecordRaw } from 'vue-router';
import IssueIndexView from '@/pages/issue/IssueIndexPage.vue';
import IssueShowView from '@/pages/issue/IssueShowPage.vue';

export const issueRouter: RouteRecordRaw[] = [
  {
    path: '/issue',
    name: 'issueIndex',
    component: IssueIndexView,
  },
  {
    path: '/issue/{issueid}',
    name: 'issueShow',
    component: IssueShowView,
  },
];
