<template>
  <user-table :items="userStore.index" />
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/user.store';
import UserTable from '@/component/UserTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
const userStore = useUserStore();

onMounted(async () => {
  await userStore.getUsers();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/user/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
