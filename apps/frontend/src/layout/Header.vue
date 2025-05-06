<template>
  <v-app-bar flat>
    <v-container class="mx-auto d-flex align-center justify-center">
      <v-img :src="logo" alt="LibrePM" height="70" />

      <template v-for="link in links" :key="link.key">
        <v-btn :text="link.text" :to="link.to" variant="text" v-if="!link.sublinks"></v-btn>
        <template v-else>
          <v-menu open-on-hover>
            <template v-slot:activator="{ props }">
              <v-btn v-bind="props" variant="text">{{ link.text }}</v-btn>
            </template>

            <v-list>
              <v-list-item v-for="(sublink, index) in link.sublinks" :key="index" :value="index" :to="sublink.to">
                <v-list-item-title>{{ sublink.text }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </template>
      </template>
      <v-spacer></v-spacer>

      <v-responsive max-width="160">
        <v-text-field density="compact" :label="t('global.search')" rounded="lg" variant="solo-filled" flat hide-details
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
};


interface Link {
  text: string;
  key?: string;
  to?: string;
  sublinks?: Sublink[];
};


const links: Link[] = [
  {
    text: 'Projects',
    key: 'projects',
    to: '/project'
  },
  {
    text: 'Dashboard',
    key: 'dashboard',
    to: '/'
  },
  {
    text: 'Issues',
    key: 'issues',
    to: '/issue'
  },
  {
    text: 'Boards',
    key: 'boards',
    to: '/board'
  },
  {
    text: 'Filters',
    key: 'filters',
    to: '/filter'
  },
  {
    text: 'Administration',
    sublinks: [
      {
        text: 'Statuses',
        key: 'statuses',
        to: '/status'
      },
      {
        text: 'Trackers',
        key: 'trackers',
        to: '/tracker'
      },
    ],
  },

  /*  {
      text: 'Profile',
      key: 'profile',
      to: '/profile'
    },
    {
      text: 'Logout',
      key: 'logout',
      to: '/logout'
    }, */
]
</script>