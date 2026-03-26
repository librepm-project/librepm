import { IssueAuditLog } from '@/lib/interfaces/issue-audit-log.interface';
import api from '@/api/api';

const index = async (issueId: string): Promise<IssueAuditLog[]> => {
  const response = await api.apiCall().get(`/issue/${issueId}/audit-log`);
  return response.data as IssueAuditLog[];
};

export default { index };
