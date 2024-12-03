<template>
  <v-app id="librepm">
    <Header />

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="2">
            <v-sheet rounded="lg">
              <v-list rounded="lg">
                <v-list-item v-for="n in 5" :key="n" :title="`List Item ${n}`"></v-list-item>
                <v-divider class="my-2"></v-divider>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet min-height="70vh" rounded="lg">
              <v-main>
                <h2 class="text-h2">{{$route.meta.title}}</h2>
                <RouterView />
              </v-main>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router';
import Header from '@/layout/Header.vue';
</script>
