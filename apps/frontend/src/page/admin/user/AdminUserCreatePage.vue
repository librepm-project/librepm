<template>
  <div>
    <user-form
      :on-submit="handleSubmit"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/user.store';
import { useRouter } from 'vue-router';
import UserForm from '@/component/UserForm.vue';
import { User } from '@/lib/interfaces/user.interface';

const router = useRouter();
const userStore = useUserStore();

const handleSubmit = async (user: User) => {
  try {
    await userStore.post(user);
    router.push('/admin/user');
  } catch (error) {
    console.error('Error creating user:', error);
  }
};
</script>
