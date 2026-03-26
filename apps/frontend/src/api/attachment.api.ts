import { Attachment } from '@/lib/interfaces/attachment.interface';
import api from '@/api/api';

const indexByIssue = async (issueId: string): Promise<Attachment[]> => {
  const response = await api.apiCall().get(`/issue/${issueId}/attachment`);
  return response.data as Attachment[];
};

const createForIssue = async (issueId: string, file: File): Promise<Attachment> => {
  const formData = new FormData();
  formData.append('file', file);
  const response = await api.apiCall().post(`/issue/${issueId}/attachment`, formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  });
  return response.data as Attachment;
};

const destroy = async (attachmentId: string): Promise<void> => {
  await api.apiCall().delete(`/attachment/${attachmentId}`);
};

const download = async (attachmentId: string, filename: string): Promise<void> => {
  const response = await api.apiCall().get(`/attachment/${attachmentId}/download`, {
    responseType: 'blob',
  });
  const url = URL.createObjectURL(response.data);
  const link = document.createElement('a');
  link.href = url;
  link.download = filename;
  link.click();
  URL.revokeObjectURL(url);
};

export default { indexByIssue, createForIssue, destroy, download };
