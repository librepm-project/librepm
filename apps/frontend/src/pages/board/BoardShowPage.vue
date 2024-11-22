<template>
  <div>
    <h1>Board</h1>
    <v-select
  label="Select"
  :items="boards"
  item-title="name"
  item-value="id"
  :name="selectedBoardId"
></v-select>
  </div>
</template>

<script setup lang="ts">
import { Board } from '../../lib/interfaces/board.interface';
import { useBoardStore } from '../../store/board.store';

  useBoardStore().getBoards();

  const boards: Board[] = useBoardStore().index;
  let selectedBoardId = "0";
  let selectedBoard = boards[Number(selectedBoardId)];
 
</script>