<template>
  <v-app id="librepm">
    <Header />

    <v-main>
      <v-container>
        <v-row>
          <Sidebar v-if="layoutStore.sidebarItems.length > 0"/>
          <Main />
        </v-row>
        <v-row>
          <Footer />
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import Header from '@/layout/Header.vue';
import Sidebar from '@/layout/Sidebar.vue';
import Main from '@/layout/Main.vue';
import Footer from '@/layout/Footer.vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();

</script>

<style>
.v-data-table thead tr {
  background-color: rgba(var(--v-theme-secondary));
}
</style>