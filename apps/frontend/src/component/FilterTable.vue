<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.actions="{ item }">
      <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { Filter } from '@/lib/interfaces/filter.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

defineProps<{
  items: Filter[],
  onEdit: (item: Filter) => void,
  onDelete: (item: Filter) => void,
}>();

const headers = [
  { key: 'name', title: t('filter.name') },
  { key: 'description', title: t('filter.description') },
  { key: 'actions', title: '', sortable: false },
];
</script>