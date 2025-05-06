<template>
    <v-data-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('user.email') }}
          </th>
          <th class="text-left">
            {{ t('user.fullName') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item.email"
        >
          <td>{{ item.email }}</td>
          <td>{{ item.firstName }} {{ item.lastName }}</td>
        </tr>
      </tbody>
    </v-data-table>
</template>

<script setup lang="ts">
import { User } from '@/lib/interfaces/user.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
    items: User[]
}>()

</script>
