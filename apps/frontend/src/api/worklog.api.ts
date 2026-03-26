import { Worklog } from '@/lib/interfaces/worklog.interface';
import api from '@/api/api';

const index = async (issueId: string): Promise<Worklog[]> => {
  const response = await api.apiCall().get(`/issue/${issueId}/worklog`);
  return response.data as Worklog[];
};

const create = async (issueId: string, payload: Partial<Worklog>): Promise<Worklog> => {
  const response = await api.apiCall().post(`/issue/${issueId}/worklog`, payload);
  return response.data as Worklog;
};

const update = async (issueId: string, worklogId: string, payload: Partial<Worklog>): Promise<Worklog> => {
  const response = await api.apiCall().put(`/issue/${issueId}/worklog/${worklogId}`, payload);
  return response.data as Worklog;
};

const destroy = async (issueId: string, worklogId: string): Promise<void> => {
  await api.apiCall().delete(`/issue/${issueId}/worklog/${worklogId}`);
};

export default { index, create, update, destroy };
