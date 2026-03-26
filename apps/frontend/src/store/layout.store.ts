import { defineStore } from 'pinia';
import { Component } from 'vue';

interface SidebarItem {
  key: string;
    title: string;
    link: string;
}


interface ActionButton {
  text: string;
  to?: string;
  onClick?: () => void;
  color: string;
  icon: string;  
}

interface LayoutStore {
  sidebarItems: SidebarItem[];
  actionButtons: ActionButton[];
  sidebarComponent: Component | null;
  sidebarProps: Record<string, any>;
}

export const useLayoutStore = defineStore('sidebar', {
  state: (): LayoutStore => {
    return {
      sidebarItems: [],
      actionButtons: [],
      sidebarComponent: null,
      sidebarProps: {},
    };
  },
  getters: {
    hasSidebar: (state) => state.sidebarItems.length > 0 || state.sidebarComponent !== null,
  },
  actions: {
    setSidebar(items: SidebarItem[]) {
      this.sidebarItems = items;
    },
    resetSidebar() {
      this.sidebarItems = [];
    },
    setSidebarComponent(component: Component | null, props: Record<string, any> = {}) {
      this.sidebarComponent = component;
      this.sidebarProps = props;
    },
    resetSidebarComponent() {
      this.sidebarComponent = null;
      this.sidebarProps = {};
    },
    setActions(items: ActionButton[]) {
      this.actionButtons = items;
    },
    resetActions() {
      this.actionButtons = [];
    },
  }
});
