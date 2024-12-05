<template>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('issue.key') }}
          </th>
          <th class="text-left">
            {{ t('issue.summary') }}
          </th>
          <th class="text-left">
            {{ t('issue.status') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
        >
          <td>{{ item.key }}</td>
          <td><router-link :to="`/issue/${item.id}`">{{ item.summary }}</router-link></td>
          <td>{{ item.status.name }}</td>
        </tr>
      </tbody>
    </v-table>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: Issue[]
}>()

</script>