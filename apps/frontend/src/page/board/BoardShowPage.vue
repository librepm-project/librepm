<template>
  <div class="text-right">
    <v-btn :text="t('global.create')" to="/board/create" color="primary" prepend-icon="mdi-plus"></v-btn>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useBoardStore } from '@/store/board.store';
import { useSidebarStore } from '@/store/sidebar.store';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();
const boardStore = useBoardStore();

onMounted(async () => {
  await boardStore.getBoards();
})

const sidebarStore = useSidebarStore();

onMounted(async () => {
  await boardStore.getBoards();
  sidebarStore.set(boardStore.index.map(board => ({
    key: board.id,
    title: board.name,
    link: `/board/${board.id}`,
  })))
});

onUnmounted(() => {
  sidebarStore.reset()
});
</script>