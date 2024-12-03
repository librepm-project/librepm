import { defineStore } from 'pinia';

interface SidebarStore {
  items: {
    key: string;
    title: string;
    link: string;
  }[];
}

export const useSidebarStore = defineStore('sidebar', {
  state: (): SidebarStore => {
    return {
      items: [],
    };
  }
});
