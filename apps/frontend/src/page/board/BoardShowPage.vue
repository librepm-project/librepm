<template>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useBoardStore } from '@/store/board.store';
import { useLayoutStore } from '@/store/layout.store';

const boardStore = useBoardStore();

onMounted(async () => {
  await boardStore.getBoards();
})

const layoutStore = useLayoutStore();

onMounted(async () => {
  await boardStore.getBoards();
  layoutStore.setSidebar(boardStore.index.map(board => ({
    key: board.id,
    title: board.name,
    link: `/board/${board.id}`,
  })));
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/board/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(() => {
  layoutStore.resetSidebar();
  layoutStore.resetActions();
});
</script>