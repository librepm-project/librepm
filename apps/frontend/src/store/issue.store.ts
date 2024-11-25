import { defineStore } from 'pinia';
import { Issue } from '../lib/interfaces/issue.interface';
import { issueMock } from '../lib/mock-data';
import issueApi from '../api/issue.api';

interface IssueStore {
  current: Issue | null;
  index: Issue[];
}

export const useIssueStore = defineStore('issue', {
  state: (): IssueStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    getIssue(issueId: number) {
      this.current = issueMock[0];
    },

    async getIssues() {
      this.index = await issueApi.index()
    },
  },
});
