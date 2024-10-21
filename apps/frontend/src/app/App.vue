<template>
  <v-app id="librepm">
    <v-app-bar flat>
      <v-container class="mx-auto d-flex align-center justify-center">
        <v-avatar
          class="me-4 "
          color="grey-darken-1"
          size="32"
        ></v-avatar>

        <v-btn
          v-for="link in links"
          :key="link.key"
          :text="link.text"
          :to="link.to"
          variant="text"
        ></v-btn>

        <v-spacer></v-spacer>

        <v-responsive max-width="160">
          <v-text-field
            density="compact"
            label="Search"
            rounded="lg"
            variant="solo-filled"
            flat
            hide-details
            single-line
          ></v-text-field>
        </v-responsive>
      </v-container>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="2">
            <v-sheet rounded="lg">
              <v-list rounded="lg">
                <v-list-item
                  v-for="n in 5"
                  :key="n"
                  :title="`List Item ${n}`"
                  link="/profile"
                ></v-list-item>
                <v-divider class="my-2"></v-divider>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet
              min-height="70vh"
              rounded="lg"
            >
              <v-main>
                <RouterView />
              </v-main>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { RouterLink, RouterView } from 'vue-router';

  const links = [
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
      text: 'Profile',
      key: 'profile',
      to: '/profile'
    },
    {
      text: 'Logout',
      key: 'logout',
      to: '/logout'
    },    
  ]
</script>
