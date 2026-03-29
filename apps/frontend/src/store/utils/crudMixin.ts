interface SimpleCrudApi<T> {
  create: (item: unknown) => Promise<T>;
  update: (id: string, item: unknown) => Promise<T>;
  destroy: (id: string) => Promise<void>;
  show: (id: string) => Promise<T>;
  index: () => Promise<T[]>;
}

export function createCrudActions<T extends { id?: string }>(api: SimpleCrudApi<T>) {
  return {
    async post(item: Omit<T, 'id'>) {
      const newItem = await api.create(item);
      this.index.push(newItem);
      return newItem;
    },

    async put(id: string, item: Omit<T, 'id'>) {
      const updated = await api.update(id, item);
      const idx = this.index.findIndex((e: T) => e.id === id);
      if (idx !== -1) {
        this.index[idx] = updated;
      }
      return updated;
    },

    async delete(id: string) {
      await api.destroy(id);
      this.index = this.index.filter((e: T) => e.id !== id);
    },

    async get(id: string) {
      this.current = await api.show(id);
    },

    async getAll() {
      this.index = await api.index();
    },
  };
}
