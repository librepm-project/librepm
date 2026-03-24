<template>
  <div v-if="boardStore.current">
    <board-form :onSubmit="update" :initialData="boardStore.current" submitButtonText="global.update" />
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
    await boardStore.getBoard(route.params.boardId.toString());
});

const update = async (board: Partial<Board>) => {
  await boardStore.update(route.params.boardId.toString(), board);
  router.push(`/board/${route.params.boardId}`);
}

</script>
