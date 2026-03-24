<template>
  <user-table 
    :items="userStore.index" 
    :onEdit="handleEdit" 
    :onDelete="handleDelete"
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

const handleDelete = async (user: User) => {
  if (confirm('Are you sure?')) {
    try {
      await userStore.deleteUser(user.id);
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  }
};

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

