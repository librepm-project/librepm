<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('status.name') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
        >
          <td>{{ item.name }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: Status[]
}>()

</script>