import { defineStore } from 'pinia';

interface SidebarItem {
  key: string;
    title: string;
    link: string;
}

interface SidebarStore {
  items: SidebarItem[];
}

export const useSidebarStore = defineStore('sidebar', {
  state: (): SidebarStore => {
    return {
      items: [],
    };
  },
  actions: {
    set(items: SidebarItem[]) {
      this.items = items;
    },
    reset() {
      this.items = [];
    },
  }
});
