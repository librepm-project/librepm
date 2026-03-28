import { defineStore } from 'pinia';
import { Worklog } from '@/lib/interfaces/worklog.interface';
import worklogApi from '@/api/worklog.api';

interface WorklogStore {
  worklogs: Worklog[];
}

export const useWorklogStore = defineStore('worklog', {
  state: (): WorklogStore => ({
    worklogs: [],
  }),
  actions: {
    async get(issueId: string) {
      this.worklogs = await worklogApi.index(issueId);
    },
    async create(issueId: string, payload: Partial<Worklog>) {
      const created = await worklogApi.create(issueId, payload);
      this.worklogs.unshift(created);
      return created;
    },
    async update(issueId: string, worklogId: string, payload: Partial<Worklog>) {
      const updated = await worklogApi.update(issueId, worklogId, payload);
      const idx = this.worklogs.findIndex(w => w.id === worklogId);
      if (idx !== -1) this.worklogs[idx] = updated;
      return updated;
    },
    async destroy(issueId: string, worklogId: string) {
      await worklogApi.destroy(issueId, worklogId);
      this.worklogs = this.worklogs.filter(w => w.id !== worklogId);
    },
  },
});
