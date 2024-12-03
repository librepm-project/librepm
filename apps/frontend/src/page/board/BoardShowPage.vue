<template>
  <div>
    <v-select
      label="Select"
      :items="boardStore.index"
      item-title="name"
      item-value="id"
    ></v-select>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useBoardStore } from '@/store/board.store';

  const boardStore = useBoardStore();

  onMounted(async () => {
    await boardStore.getBoards();
  })

  let selectedBoardId = "0";
  let selectedBoard = boardStore.index[Number(selectedBoardId)];
 
</script>