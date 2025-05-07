<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.fullName="{ item }">
      {{ item.firstName }} {{ item.lastName }}
    </template>
    <template #item.actions="{ item }">
      <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { User } from '@/lib/interfaces/user.interface';
import { useI18n } from 'vue-i18n'
GeneralRecordActions

const { t } = useI18n();

defineProps<{
  items: User[];
  onEdit: (item: User) => void;
  onDelete: (item: User) => void;
}>()

const headers = [
  { key: 'email', title: t('user.email') },
  { key: 'fullName', title: t('user.fullName') },
  { key: 'actions', title: '' },
]

</script>