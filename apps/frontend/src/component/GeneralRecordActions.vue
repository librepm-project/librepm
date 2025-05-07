<template>
  <v-menu location="bottom end">
    <template #activator="{ props }">
      <v-btn v-bind="props" icon variant="text" size="small">
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </template>

    <v-list>
      <v-list-item @click="onEdit(item)">
        <v-list-item-title>
          <v-icon start class="me-2">mdi-pencil</v-icon>
          {{ t('global.edit') }}
        </v-list-item-title>
      </v-list-item>

      <v-list-item @click="onDelete(item)">
        <v-list-item-title>
          <v-icon start class="me-2" color="error">mdi-delete</v-icon>
          {{ t('global.delete') }}
        </v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

defineProps<{
  item: any,
  onEdit: (item: any) => void,
  onDelete: (item: any) => void,
}>();
</script>