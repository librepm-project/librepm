import { RouteRecordRaw } from 'vue-router';
import IssueIndexPage from '@/page/issue/IssueIndexPage.vue';
import IssueShowPage from '@/page/issue/IssueShowPage.vue';
import { reactive } from 'vue';

export const issueRouter: RouteRecordRaw[] = [
  {
    path: '/issue',
    name: 'issueIndex',
    component: IssueIndexPage,
    meta: reactive({
      title: 'Issues',
    }),
  },
  {
    path: '/issue/{issueid}',
    name: 'issueShow',
    component: IssueShowPage,
    meta: reactive({
      title: '',
    }),
  },
];
