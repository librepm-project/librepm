<template>
  <div>
    <div class="text-right">
      <v-btn :text="t('global.create')" to="/user/create" color="primary" prepend-icon="mdi-plus"></v-btn>
    </div>
    <user-table :items="userStore.index" />
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/user.store';
import UserTable from '@/component/UserTable.vue';
import { onMounted } from 'vue';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const userStore = useUserStore();

onMounted(async () => {
  await userStore.getUsers();
});
</script>
