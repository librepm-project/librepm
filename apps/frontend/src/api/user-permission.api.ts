import api from '@/api/api';

interface UserPermissionResponse {
  permissions: string[];
}

const index = async (): Promise<string[]> => {
  const response = await api.apiCall().get('/user/permission');
  return (response.data as UserPermissionResponse).permissions;
};

export default { index };
