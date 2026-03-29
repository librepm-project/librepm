import api from '@/api/api';

/**
 * Factory function to create a generic entity API client
 * Reduces boilerplate for standard CRUD operations
 */
export function createEntityAPI<T extends { id: string }>(basePath: string) {
  return {
    async index(entityId: string): Promise<T[]> {
      const response = await api.apiCall().get(`${basePath}/${entityId}`);
      return response.data as T[];
    },

    async indexByEntity(entityId: string): Promise<T[]> {
      const response = await api.apiCall().get(`${basePath}/${entityId}`);
      return response.data as T[];
    },

    async show(itemId: string): Promise<T> {
      const response = await api.apiCall().get(`${basePath}/${itemId}`);
      return response.data as T;
    },

    async create(entityId: string, payload: unknown): Promise<T> {
      const response = await api.apiCall().post(`${basePath}/${entityId}`, payload);
      return response.data as T;
    },

    async update(entityId: string, itemId: string, payload: unknown): Promise<T> {
      const response = await api.apiCall().put(`${basePath}/${entityId}/${itemId}`, payload);
      return response.data as T;
    },

    async destroy(itemId: string): Promise<void> {
      await api.apiCall().delete(`${basePath}/${itemId}`);
    },
  };
}
