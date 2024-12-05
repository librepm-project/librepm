<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('filter.name') }}
          </th>
          <th class="text-left">
            {{ t('filter.description') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Filter } from '@/lib/interfaces/filter.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: Filter[]
}>()

</script>