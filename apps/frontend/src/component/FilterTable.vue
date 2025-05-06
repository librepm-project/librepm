<template>
  <v-data-table density="compact" :items="items" :headers="headers">
  </v-data-table>
</template>

<script setup lang="ts">
import { Filter } from '@/lib/interfaces/filter.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Filter[]
}>()

const headers = [
  { key: 'name', title: t('filter.name') },
  { key: 'description', title: t('filter.description') },
]

</script>