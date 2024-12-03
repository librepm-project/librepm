<template>
    <v-app-bar flat>
        <v-container class="mx-auto d-flex align-center justify-center">
        <v-avatar class="me-4 " color="grey-darken-1" size="32"></v-avatar>

        <v-btn v-for="link in links" :key="link.key" :text="link.text" :to="link.to" variant="text"></v-btn>

        <v-spacer></v-spacer>

        <v-responsive max-width="160">
            <v-text-field density="compact" label="Search" rounded="lg" variant="solo-filled" flat hide-details
            single-line></v-text-field>
        </v-responsive>
        </v-container>
    </v-app-bar>
</template>

<script lang="ts" setup>

interface Link {
  text: string;
  key: string;
  to: string;
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
    text: 'Statuses',
    key: 'statuses',
    to: '/status'
  },
  {
    text: 'Trackers',
    key: 'trackers',
    to: '/tracker'
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