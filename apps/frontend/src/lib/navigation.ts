export interface Sublink {
  text: string;
  key: string;
  to: string;
  icon: string;
}

export interface Link {
  text: string;
  key: string;
  to?: string;
  icon: string;
  sublinks?: Sublink[];
}

export const navigationLinks: Link[] = [
  {
    text: 'Dashboard',
    key: 'dashboard',
    to: '/',
    icon: 'mdi-view-dashboard-outline',
  },
  {
    text: 'Issues',
    key: 'issues',
    to: '/issue',
    icon: 'mdi-bug-outline',
  },
  {
    text: 'Boards',
    key: 'boards',
    to: '/board',
    icon: 'mdi-table-large',
  },
  {
    text: 'Filters',
    key: 'filters',
    to: '/filter',
    icon: 'mdi-filter-outline',
  },
  {
    text: 'Administration',
    key: 'admin',
    icon: 'mdi-cog-outline',
    sublinks: [
      {
        text: 'Projects',
        key: 'projects',
        to: '/admin/project',
        icon: 'mdi-folder-outline',
      },
      {
        text: 'Statuses',
        key: 'statuses',
        to: '/admin/status',
        icon: 'mdi-layers-outline',
      },
      {
        text: 'Trackers',
        key: 'trackers',
        to: '/admin/tracker',
        icon: 'mdi-compass-outline',
      },
      {
        text: 'Users',
        key: 'users',
        to: '/admin/user',
        icon: 'mdi-account-multiple',
      },
      {
        text: 'Boards',
        key: 'boards',
        to: '/admin/board',
        icon: 'mdi-table-large',
      },
    ],
  },
];
