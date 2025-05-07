<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.color="{ item }">
      <v-chip :color="item.color">{{ item.color }}</v-chip>
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Status[]
}>()

const headers = [
  { key: 'name', title: t('status.name') },
  { key: 'color', title: t('status.color') },
]

</script>