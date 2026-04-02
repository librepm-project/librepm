import api from '@/api/api';

const index = async (userId: string): Promise<string[]> => {
  const response = await api.apiCall().get(`/user/${userId}/role`);
  return (response.data as { roles: string[] }).roles;
};

const create = async (userId: string, role: string): Promise<void> => {
  await api.apiCall().post(`/user/${userId}/role`, { role });
};

const destroy = async (userId: string, role: string): Promise<void> => {
  await api.apiCall().delete(`/user/${userId}/role/${role}`);
};

export default { index, create, destroy };
