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
  title: string;
  onTitleClick: (() => void) | null;
  isEditingTitle: boolean;
  titleDraft: string;
  onTitleSave: ((newTitle: string) => void) | null;
  onTitleCancel: (() => void) | null;
  isDrawerOpen: boolean;
  sidebarItems: SidebarItem[];
  actionButtons: ActionButton[];
  sidebarComponent: Component | null;
  sidebarProps: Record<string, any>;
  hideLayout: boolean;
}

export const useLayoutStore = defineStore('sidebar', {
  state: (): LayoutStore => {
    return {
      title: '',
      onTitleClick: null,
      isEditingTitle: false,
      titleDraft: '',
      onTitleSave: null,
      onTitleCancel: null,
      isDrawerOpen: false,
      sidebarItems: [],
      actionButtons: [],
      sidebarComponent: null,
      sidebarProps: {},
      hideLayout: false,
    };
  },
  getters: {
    hasSidebar: (state) => state.sidebarItems.length > 0 || state.sidebarComponent !== null,
  },
  actions: {
    toggleDrawer(value: boolean | null = null) {
      this.isDrawerOpen = value !== null ? value : !this.isDrawerOpen;
    },
    setTitle(title: string, onTitleClick: (() => void) | null = null) {
      this.title = title;
      this.onTitleClick = onTitleClick;
    },
    setTitleEditing(isEditing: boolean, draft: string = '', onSave: ((newTitle: string) => void) | null = null, onCancel: (() => void) | null = null) {
      this.isEditingTitle = isEditing;
      this.titleDraft = draft;
      this.onTitleSave = onSave;
      this.onTitleCancel = onCancel;
    },
    resetTitle() {
      this.title = '';
      this.onTitleClick = null;
      this.isEditingTitle = false;
      this.titleDraft = '';
      this.onTitleSave = null;
      this.onTitleCancel = null;
    },
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
    setHideLayout(value: boolean) {
      this.hideLayout = value;
    },
  }
});
