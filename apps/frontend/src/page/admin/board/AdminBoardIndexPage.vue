<template>
  <board-table :items="boardStore.index" :onEdit="onEdit" />
</template>

<script lang="ts" setup>
import { useBoardStore } from '@/store/board.store';
import BoardTable from '@/component/BoardTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Board } from '@/lib/interfaces/board.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const router = useRouter();
const layoutStore = useLayoutStore();
const boardStore = useBoardStore();

const onEdit = (item: Board) => {
  router.push(`/admin/board/${item.id}/edit`);
}

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
