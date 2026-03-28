<template>
  <div v-if="boardStore.current">
    <board-form :onSubmit="update" :onDelete="handleDelete" :initialData="boardStore.current" submitButtonText="global.update" />
  </div>
</template>

<script lang="ts" setup>
import { useBoardStore } from '@/store/board.store';
import BoardForm from '@/component/BoardForm.vue';
import { Board } from '@/lib/interfaces/board.interface';
import { useRouter, useRoute } from 'vue-router';
import { onMounted } from 'vue';

const router = useRouter();
const route = useRoute();
const boardStore = useBoardStore();

onMounted(async () => {
    await boardStore.get(route.params.boardId.toString());
});

const update = async (board: Partial<Board>) => {
  await boardStore.update(route.params.boardId.toString(), board);
  router.push(`/board/${route.params.boardId}`);
}

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this board?')) {
    try {
      const boardId = route.params.boardId.toString();
      await boardStore.destroy(boardId);
      router.push('/admin/board');
    } catch (error) {
      console.error('Error deleting board:', error);
    }
  }
};

</script>
