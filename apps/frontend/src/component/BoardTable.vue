<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.actions="{ item }">
      <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { Board } from '@/lib/interfaces/board.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Board[];
  onEdit: (item: Board) => void;
  onDelete: (item: Board) => void;
}>()

const headers = [
  { key: 'name', title: t('global.name') },
  { key: 'description', title: t('global.description') },
  { key: 'actions', title: '' },
]

</script>