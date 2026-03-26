import { Comment } from '@/lib/interfaces/comment.interface';
import api from '@/api/api';

const indexByIssue = async (issueId: string): Promise<Comment[]> => {
  const response = await api.apiCall().get(`/issue/${issueId}/comment`);
  return response.data as Comment[];
};

const create = async (issueId: string, content: string): Promise<Comment> => {
  const response = await api.apiCall().post(`/issue/${issueId}/comment`, { content });
  return response.data as Comment;
};

const update = async (commentId: string, content: string): Promise<Comment> => {
  const response = await api.apiCall().put(`/comment/${commentId}`, { content });
  return response.data as Comment;
};

const destroy = async (commentId: string): Promise<void> => {
  await api.apiCall().delete(`/comment/${commentId}`);
};

export default { indexByIssue, create, update, destroy };
