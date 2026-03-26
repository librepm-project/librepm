<template>
  <v-app id="librepm">
    <Header />
    
    <v-navigation-drawer
      v-model="layoutStore.isDrawerOpen"
      temporary
      class="d-md-none"
    >
      <v-list>
        <div class="pa-4 d-flex align-center gap-2 mb-4">
          <v-icon color="primary" size="24">mdi-rhombus-split</v-icon>
          <span class="text-h6 font-weight-black text-primary">LibrePM</span>
        </div>
        
        <template v-for="link in navigationLinks" :key="link.key">
          <v-list-item
            v-if="!link.sublinks"
            :to="link.to"
            :prepend-icon="link.icon"
            :title="link.text"
            rounded="lg"
            class="ma-2"
          />
          
          <v-list-group v-else :value="link.key">
            <template v-slot:activator="{ props }">
              <v-list-item
                v-bind="props"
                :prepend-icon="link.icon"
                :title="link.text"
                rounded="lg"
                class="ma-2"
              />
            </template>

            <v-list-item
              v-for="sublink in link.sublinks"
              :key="sublink.key"
              :to="sublink.to"
              :prepend-icon="sublink.icon"
              :title="sublink.text"
              rounded="lg"
              class="ma-2 ml-8"
            />
          </v-list-group>
        </template>
      </v-list>
      
      <v-divider class="ma-2" />
      
      <v-list v-if="userCurrentStore.current">
        <v-list-item
          prepend-icon="mdi-logout"
          title="Logout"
          rounded="lg"
          class="ma-2 text-error"
          @click="logout"
        />
      </v-list>
    </v-navigation-drawer>

    <v-main class="bg-surface">
      <v-container fluid class="pa-2 pa-sm-6">
        <v-row>
          <Sidebar v-if="layoutStore.hasSidebar" class="d-none d-md-flex" />
          <Main :cols="12" :md="layoutStore.hasSidebar ? 10 : 12" />
        </v-row>
      </v-container>
      <Footer />
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import Header from '@/layout/Header.vue';
import Sidebar from '@/layout/Sidebar.vue';
import Main from '@/layout/Main.vue';
import Footer from '@/layout/Footer.vue';
import { useLayoutStore } from '@/store/layout.store';
import { useUserCurrentStore } from '@/store/userCurrent.store';
import { useUserSessionStore } from '@/store/userSession.store';
import { getToken } from '@/lib/cookie';
import { navigationLinks } from '@/lib/navigation';
import { useRouter } from 'vue-router';

const layoutStore = useLayoutStore();
const userCurrentStore = useUserCurrentStore();
const userSessionStore = useUserSessionStore();
const router = useRouter();

const logout = () => {
  userSessionStore.logout();
  router.push('/login');
  layoutStore.toggleDrawer(false);
};

onMounted(async () => {
  if (getToken()) {
    await userCurrentStore.getUser();
  }
});
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