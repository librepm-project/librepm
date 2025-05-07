import { defineStore } from 'pinia';

interface SidebarItem {
  key: string;
    title: string;
    link: string;
}


interface ActionButton {
  text: string;
  to: string;
  color: string;
  icon: string;  
}

interface LayoutStore {
  sidebarItems: SidebarItem[];
  actionButtons: ActionButton[];
}

export const useLayoutStore = defineStore('sidebar', {
  state: (): LayoutStore => {
    return {
      sidebarItems: [],
      actionButtons: [],
    };
  },
  actions: {
    setSidebar(items: SidebarItem[]) {
      this.sidebarItems = items;
    },
    resetSidebar() {
      this.sidebarItems = [];
    },
    setActions(items: ActionButton[]) {
      this.actionButtons = items;
    },
    resetActions() {
      this.actionButtons = [];
    },
  }
});
