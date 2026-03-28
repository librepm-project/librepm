import { defineStore } from 'pinia';
import { IssueAuditLog } from '@/lib/interfaces/issue-audit-log.interface';
import issueAuditLogApi from '@/api/issue-audit-log.api';

interface IssueAuditLogStore {
  auditLogs: IssueAuditLog[];
}

export const useIssueAuditLogStore = defineStore('issueAuditLog', {
  state: (): IssueAuditLogStore => ({
    auditLogs: [],
  }),
  actions: {
    async get(issueId: string) {
      this.auditLogs = await issueAuditLogApi.index(issueId);
    },
  },
});
