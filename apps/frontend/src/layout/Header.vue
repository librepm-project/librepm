<template>
  <v-app-bar flat>
    <v-container class="mx-auto d-flex align-center justify-center">
      <v-img :src="logo" alt="LibrePM" height="70" />

      <template v-for="link in links" :key="link.key">
        <v-btn :prepend-icon="link.icon" :text="link.text" :to="link.to" variant="text" v-if="!link.sublinks"></v-btn>
        <template v-else>
          <v-menu open-on-hover>
            <template v-slot:activator="{ props }">
              <v-btn :prepend-icon="link.icon" v-bind="props" variant="text">{{ link.text }}</v-btn>
            </template>

            <v-list>
              <v-list-item v-for="(sublink, index) in link.sublinks" :key="index" :value="index" :to="sublink.to">

                <v-list-item-title><v-icon start size="small" class="me-2">{{ sublink.icon }}</v-icon>{{ sublink.text
                  }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </template>
      </template>
      <v-spacer></v-spacer>

      <v-responsive max-width="160">
        <v-text-field :label="t('global.search')" rounded="lg" variant="solo-filled" flat hide-details
          single-line></v-text-field>
      </v-responsive>
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
        key: 'userss',
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