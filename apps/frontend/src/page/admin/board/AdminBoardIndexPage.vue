<template>
  <div>
    <v-data-table>
      <thead>
        <tr>
          <th class="text-left">
            {{ t('dashboard.name') }}
          </th>
          <th class="text-left">
            {{ t('dashboard.description') }}
          </th>
          <th>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in boards"
          :key="item.id"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
          <td><router-link to="/admin/board/{item.id}/edit">Configure</router-link></td>
        </tr>
      </tbody>
    </v-data-table>
  </div>
</template>

<script lang="ts" setup>
import { Board } from '@/lib/interfaces/board.interface';
import { useBoardStore } from '@/store/board.store';
import { RouterLink } from 'vue-router';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();
useBoardStore().getBoards();

const boards: Board[] = useBoardStore().index;

</script>