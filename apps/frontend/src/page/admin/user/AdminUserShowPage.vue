<template>
  <div>
    <user-form
      v-if="userStore.current"
      :user="userStore.current"
      :on-submit="handleSubmit"
      :on-delete="handleDelete"
      submit-button-text="global.update"
    />
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/user.store';
import { useRoute, useRouter } from 'vue-router';
import UserForm from '@/component/UserForm.vue';
import { User } from '@/lib/interfaces/user.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const handleSubmit = async (user: User) => {
  try {
    const userId = route.params.userId as string;
    await userStore.put(userId, user);
    router.push('/admin/user');
  } catch (error) {
    console.error('Error updating user:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this user?')) {
    try {
      const userId = route.params.userId as string;
      await userStore.delete(userId);
      router.push('/admin/user');
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  }
};

onMounted(async () => {
  const userId = route.params.userId as string;
  await userStore.get(userId);
});
</script>
  