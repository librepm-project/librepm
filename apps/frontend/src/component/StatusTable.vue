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
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Status[];
  onEdit: (item: Status) => void;
  onDelete: (item: Status) => void;
}>()

const headers = [
  { key: 'name', title: t('status.name') },
  { key: 'color', title: t('status.color') },
  { key: 'actions', title: '' },
]

</script>