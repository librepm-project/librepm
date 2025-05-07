<template>
  <v-data-table :items="items" :headers="headers">
  </v-data-table>
</template>

<script setup lang="ts">
import { Board } from '@/lib/interfaces/board.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Board[]
}>()

const headers = [
  { key: 'name', title: t('global.name') },
  { key: 'description', title: t('global.description') },
]

</script>