<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.fullName="{ item }">
      {{ item.firstName }} {{ item.lastName }}
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import { User } from '@/lib/interfaces/user.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: User[]
}>()

const headers = [
  { key: 'email', title: t('user.email') },
  { key: 'fullName', title: t('user.fullName') },
]

</script>