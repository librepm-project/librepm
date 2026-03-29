import { defineStore, StoreDefinition } from 'pinia';

/**
 * Factory function to create a generic entity store
 * Generates boilerplate Pinia store for CRUD operations
 */
export function createEntityStore<T extends { id: string }>(
  storeName: string,
  api: {
    indexByEntity?: (entityId: string) => Promise<T[]>;
    index?: (entityId: string) => Promise<T[]>;
    create?: (entityId: string, payload: unknown) => Promise<T>;
    update?: (entityId: string, itemId: string, payload: unknown) => Promise<T>;
    destroy?: (itemId: string) => Promise<void>;
  }
): StoreDefinition {
  return defineStore(storeName, {
    state: () => ({
      items: [] as T[],
    }),

    actions: {
      async getItems(entityId: string) {
        const indexFn = api.indexByEntity || api.index;
        if (indexFn) {
          this.items = await indexFn(entityId);
        }
      },

      async createItem(entityId: string, payload: unknown) {
        if (!api.create) throw new Error('Create not supported');
        const created = await api.create(entityId, payload);
        this.items.unshift(created);
        return created;
      },

      async updateItem(entityId: string, itemId: string, payload: unknown) {
        if (!api.update) throw new Error('Update not supported');
        const updated = await api.update(entityId, itemId, payload);
        const idx = this.items.findIndex((item) => item.id === itemId);
        if (idx !== -1) this.items[idx] = updated;
        return updated;
      },

      async destroyItem(itemId: string) {
        if (!api.destroy) throw new Error('Destroy not supported');
        await api.destroy(itemId);
        this.items = this.items.filter((item) => item.id !== itemId);
      },
    },
  });
}
