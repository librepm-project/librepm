<template>
  <v-data-table density="compact" :items="items" :headers="headers">
  </v-data-table>
</template>

<script setup lang="ts">
import { Project } from '@/lib/interfaces/project.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Project[]
}>()

const headers = [
  { key: 'codeName', title: t('project.codeName') },
  { key: 'name', title: t('global.name') },
]

</script>