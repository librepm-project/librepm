import { RouteRecordRaw } from 'vue-router';
import IssueIndexPage from '@/page/issue/IssueIndexPage.vue';
import IssueShowPage from '@/page/issue/IssueShowPage.vue';
import IssueCreatePage from '@/page/issue/IssueCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const issueRouter: RouteRecordRaw[] = [
  {
    path: '/issue',
    name: 'issueIndex',
    component: IssueIndexPage,
    meta: { permission: Permissions.IssueRead },
  },
  {
    path: '/issue/create',
    name: 'issueCreate',
    component: IssueCreatePage,
    meta: { permission: Permissions.IssueCreate },
  },
  {
    path: '/issue/:issueId',
    name: 'issueShowById',
    component: IssueShowPage,
    meta: { permission: Permissions.IssueRead },
  },
  {
    path: '/issue/key/:key',
    name: 'issueShow',
    component: IssueShowPage,
    meta: { permission: Permissions.IssueRead },
  },
];
