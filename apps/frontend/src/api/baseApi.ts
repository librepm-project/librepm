import api from '@/api/api';

export interface CrudApi<T extends { id?: string }> {
  index(): Promise<T[]>;
  show(id: string): Promise<T>;
  create(item: Omit<T, 'id'>): Promise<T>;
  update(id: string, item: Omit<T, 'id'>): Promise<T>;
  destroy(id: string): Promise<void>;
}

export function createCrudApi<T extends { id?: string }>(
  endpoint: string
): CrudApi<T> {
  return {
    async index(): Promise<T[]> {
      const response = await api.apiCall().get(`/${endpoint}`);
      return response.data as T[];
    },

    async show(id: string): Promise<T> {
      const response = await api.apiCall().get(`/${endpoint}/${id}`);
      return response.data as T;
    },

    async create(item: Omit<T, 'id'>): Promise<T> {
      const response = await api.apiCall().post(`/${endpoint}`, item);
      return response.data as T;
    },

    async update(id: string, item: Omit<T, 'id'>): Promise<T> {
      await api.apiCall().put(`/${endpoint}/${id}`, item);
      return this.show(id);
    },

    async destroy(id: string): Promise<void> {
      await api.apiCall().delete(`/${endpoint}/${id}`);
    },
  };
}
