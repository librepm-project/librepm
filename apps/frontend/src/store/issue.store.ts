import { defineStore } from 'pinia';
import { Issue } from '@/lib/interfaces/issue.interface';
import issueApi from '@/api/issue.api';


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
    async getIssue(issueId: string) {
      this.current = await issueApi.show(issueId);
    },

    async getIssues() {
      this.index = await issueApi.index()
    },

    async create(issue: Partial<Issue>) {
      return issueApi.create(issue);
    }
  },
});
