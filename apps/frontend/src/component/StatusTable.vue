<template>
    <v-data-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('status.name') }}
          </th>
          <th class="text-left">
            {{ t('status.color') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
        >
          <td>{{ item.name }}</td>
          <td><v-chip :color="item.color">{{ item.color }}</v-chip></td>
        </tr>
      </tbody>
    </v-data-table>
</template>

<script setup lang="ts">
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: Status[]
}>()

</script>