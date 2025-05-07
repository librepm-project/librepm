<template>
  <board-table :items="boardStore.index" />
</template>

<script lang="ts" setup>
import { useBoardStore } from '@/store/board.store';
import BoardTable from '@/component/BoardTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
const boardStore = useBoardStore();

onMounted(async () => {
  await boardStore.getBoards();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/board/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
