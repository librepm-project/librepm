import { defineStore } from 'pinia';
import { Issue } from '../lib/interfaces/issue.interface';

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
      this.current = {
        key: 'SW-1',
        summary: 'First Issue',
        status: 'In progress',
        description: 'Lorem ipsum dolor amet',
      };
    },

    getIssues() {
      this.index = [
        {
          key: 'SW-1',
          summary: 'First Issue',
          status: 'In progress',
          description: 'Lorem ipsum dolor amet',
        },
        {
          key: 'SW-2',
          summary: 'Second Issue',
          status: 'To Do',
          description: 'Lorem ipsum dolor amet',
        },
      ];
    },
  },
});
