<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.actions="{ item }">
      <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { Project } from '@/lib/interfaces/project.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Project[]
  onEdit: (item: Project) => void,
  onDelete: (item: Project) => void,
}>()

const headers = [
  { key: 'codeName', title: t('project.codeName') },
  { key: 'name', title: t('global.name') },
  { key: 'actions', title: '' },
]

</script>