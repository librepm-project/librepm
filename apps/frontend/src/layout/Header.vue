<template>
  <v-app-bar elevation="4" color="primary" class="px-3" density="compact">
    <v-app-bar-nav-icon 
      class="d-md-none text-white mr-1" 
      @click.stop="layoutStore.toggleDrawer()"
    />
    <v-container class="d-flex align-center justify-space-between pa-0" fluid>
      <div class="d-flex align-center gap-2 cursor-pointer" @click="router.push('/')">
        <v-img :src="logo" alt="LibrePM" height="32" width="32" contain />
        <span class="text-h6 font-weight-bold text-white d-none d-sm-flex">LibrePM</span>
      </div>

      <v-spacer />

      <nav v-if="userCurrentStore.current" class="d-none d-md-flex align-center gap-2">
        <template v-for="link in navigationLinks" :key="link.key">
          <v-btn
            v-if="!link.sublinks"
            :prepend-icon="link.icon"
            :text="link.text"
            :to="link.to"
            variant="text"
            class="text-white"
          />
        </template>
      </nav>

      <v-spacer />

      <div class="d-flex align-center gap-3">
        <template v-if="userCurrentStore.current">
          <v-btn
            to="/issue/create"
            prepend-icon="mdi-plus"
            variant="outlined"
            class="text-white"
            size="small"
          >
            {{ t('title.issueCreate') }}
          </v-btn>
          <notification-bell />
          <v-menu v-for="link in navigationLinks.filter(l => l.sublinks)" :key="link.key" open-on-hover location="bottom">
            <template v-slot:activator="{ props }">
              <v-btn icon variant="text" class="text-white" v-bind="props">
                <v-icon>{{ link.icon }}</v-icon>
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
                <v-list-item-title class="font-weight-medium">{{ sublink.text }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
          <v-btn
            variant="text"
            class="text-white text-body2 text-no-wrap text-capitalize"
            to="/profile"
            :ripple="false"
          >
            {{ userCurrentStore.current.firstName }} {{ userCurrentStore.current.lastName }}
          </v-btn>
          <v-btn
            icon
            variant="text"
            class="text-white"
            :title="t('login.logout')"
            @click="logout"
          >
            <v-icon>mdi-logout</v-icon>
          </v-btn>
        </template>
        <template v-else>
          <v-btn
            to="/login"
            prepend-icon="mdi-login"
            variant="text"
            class="text-white"
          >
            {{ t('login.submit') }}
          </v-btn>
        </template>
      </div>
    </v-container>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import logo from '@/assets/logo.png';
import NotificationBell from '@/component/NotificationBell.vue';
import { useUserCurrentStore } from '@/store/user-current.store';
import { useUserSessionStore } from '@/store/user-session.store';
import { useLayoutStore } from '@/store/layout.store';
import { navigationLinks } from '@/lib/navigation';

const { t } = useI18n();
const router = useRouter();
const userCurrentStore = useUserCurrentStore();
const userSessionStore = useUserSessionStore();
const layoutStore = useLayoutStore();

const logout = () => {
  userSessionStore.logout();
  router.push('/login');
};
</script>

<style scoped>
.gap-4 {
  gap: 1rem;
}

.gap-2 {
  gap: 0.5rem;
}

.gap-3 {
  gap: 0.75rem;
}
</style>
