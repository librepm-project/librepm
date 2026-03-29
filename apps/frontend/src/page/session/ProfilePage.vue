<template>
  <div>
    <user-form
      v-if="userCurrentStore.current"
      :user="userCurrentStore.current"
      :on-submit="handleSubmit"
      submit-button-text="global.save"
    />
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { useUserCurrentStore } from '@/store/user-current.store';
import UserForm from '@/component/UserForm.vue';
import { User } from '@/lib/interfaces/user.interface';

const userCurrentStore = useUserCurrentStore();

const handleSubmit = async (user: User) => {
  await userCurrentStore.update(user);
};

onMounted(async () => {
  if (!userCurrentStore.current) {
    await userCurrentStore.get();
  }
});
</script>
