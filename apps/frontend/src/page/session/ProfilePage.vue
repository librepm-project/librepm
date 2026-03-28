<template>
  <div>
    <user-form
      v-if="userCurrentStore.current"
      :user="userCurrentStore.current"
      :onSubmit="handleSubmit"
      submitButtonText="global.save"
    />
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { useUserCurrentStore } from '@/store/userCurrent.store';
import UserForm from '@/component/UserForm.vue';
import { User } from '@/lib/interfaces/user.interface';

const userCurrentStore = useUserCurrentStore();

const handleSubmit = async (user: User) => {
  await userCurrentStore.updateCurrent(user);
};

onMounted(async () => {
  if (!userCurrentStore.current) {
    await userCurrentStore.get();
  }
});
</script>
