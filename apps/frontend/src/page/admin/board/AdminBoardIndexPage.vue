<template>
  <div>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Description
          </th>
          <th>
            Operations
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
    </v-table>
  </div>
</template>

<script lang="ts" setup>
import { Board } from '@/lib/interfaces/board.interface';
import { useBoardStore } from '@/store/board.store';
import { RouterLink } from 'vue-router';


  useBoardStore().getBoards();

  const boards: Board[] = useBoardStore().index;

</script>