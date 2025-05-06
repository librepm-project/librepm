<template>
    <v-data-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('project.codeName') }}
          </th>
          <th class="text-left">
            {{ t('global.name') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.codeName"
        >
          <td>{{ item.codeName }}</td>
          <td>{{ item.name }}</td>
        </tr>
      </tbody>
    </v-data-table>
</template>

<script setup lang="ts">
import { Project } from '@/lib/interfaces/project.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: Project[]
}>()

</script>