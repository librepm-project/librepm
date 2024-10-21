import { RouteRecordRaw } from 'vue-router';
import IssueIndexView from '../views/issue/IssueIndexView.vue';
import IssueShowView from '../views/issue/IssueShowView.vue';

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
