<template>
  <v-app-bar elevation="4" color="primary" class="px-6">
    <v-container class="d-flex align-center justify-space-between" fluid>
      <div class="d-flex align-center gap-4">
        <v-img :src="logo" alt="LibrePM" height="50" width="50" contain />
        <span class="text-h6 font-weight-bold text-white">LibrePM</span>
      </div>

      <v-spacer />

      <nav class="d-none d-md-flex align-center gap-2">
        <template v-for="link in links" :key="link.key">
          <v-btn
            v-if="!link.sublinks"
            :prepend-icon="link.icon"
            :text="link.text"
            :to="link.to"
            variant="text"
            class="text-white"
          />
          <template v-else>
            <v-menu open-on-hover location="bottom">
              <template v-slot:activator="{ props }">
                <v-btn
                  :prepend-icon="link.icon"
                  v-bind="props"
                  variant="text"
                  class="text-white"
                >
                  {{ link.text }}
                </v-btn>
              </template>

              <v-list>
                <v-list-item
                  v-for="(sublink, index) in link.sublinks"
                  :key="index"
                  :to="sublink.to"
                  class="rounded-lg ma-1"
                >
                  <template v-slot:prepend>
                    <v-icon size="small">{{ sublink.icon }}</v-icon>
                  </template>
                  <v-list-item-title class="font-weight-medium">
                    {{ sublink.text }}
                  </v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </template>
        </template>
      </nav>

      <v-spacer />

      <v-text-field
        :label="t('global.search')"
        density="compact"
        variant="solo-filled"
        single-line
        hide-details
        prepend-inner-icon="mdi-magnify"
        class="search-field"
      />
    </v-container>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
import logo from '@/assets/logo.png'

const { t } = useI18n();

interface Sublink {
  text: string;
  key: string;
  to: string;
  icon: string;
}

interface Link {
  text: string;
  key: string;
  to?: string;
  icon: string;
  sublinks?: Sublink[];
}

const links: Link[] = [
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
</script>

<style scoped>
.search-field {
  max-width: 300px;
  background-color: rgba(255, 255, 255, 0.95) !important;
}

.gap-4 {
  gap: 1rem;
}

.gap-2 {
  gap: 0.5rem;
}
</style>