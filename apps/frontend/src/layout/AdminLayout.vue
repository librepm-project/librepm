<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import AdminSidebarNav from '@/layout/AdminSidebarNav.vue';

const layoutStore = useLayoutStore();

onMounted(() => {
  layoutStore.setSidebarComponent(AdminSidebarNav);
});

onUnmounted(() => {
  layoutStore.resetSidebarComponent();
});
</script>
