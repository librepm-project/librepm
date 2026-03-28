<template>
  <user-table 
    :items="userStore.index" 
    :onEdit="handleEdit" 
  />
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/user.store';
import UserTable from '@/component/UserTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { User } from '@/lib/interfaces/user.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const userStore = useUserStore();

const handleEdit = (user: User) => {
  router.push(`/admin/user/${user.id}`);
};

onMounted(async () => {
  await userStore.getAll();
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

