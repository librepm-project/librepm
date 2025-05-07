<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.color="{ item }">
      <v-chip :color="item.color">{{ item.color }}</v-chip>
    </template>
    <template #item.actions="{ item }">
      <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Tracker[];
  onEdit: (item: Tracker) => void;
  onDelete: (item: Tracker) => void;
}>()

const headers = [
  { key: 'name', title: t('tracker.name') },
  { key: 'color', title: t('tracker.color') },
  { key: 'actions', title: '' },
]

</script>