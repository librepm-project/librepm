<template>
  <div>
    <h1>Boards</h1>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Description
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in boards"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </v-table>
  </div>
</template>

<script lang="ts" setup>
  import Board from '../../lib/interfaces/board.interface.ts';
  import { useBoardStore } from '../../store/board.store.ts';

  useBoardStore().getBoards();

  const boards: Board[] = useBoardStore().index;

</script>