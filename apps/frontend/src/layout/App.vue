<template>
  <v-app id="librepm">
    <Header />
    <v-main class="bg-surface">
      <v-container fluid class="pa-6">
        <v-row>
          <Sidebar v-if="layoutStore.sidebarItems.length > 0 || layoutStore.sidebarComponent" class="d-none d-md-flex" />
          <Main />
        </v-row>
      </v-container>
      <Footer />
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

<style scoped>
#librepm {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  min-height: 100vh;
}

.bg-surface {
  background-color: rgba(245, 247, 250, 0.5);
}
</style>