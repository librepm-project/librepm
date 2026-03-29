<template>
  <div>
    <board-form
      :on-submit="create"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useBoardStore } from '@/store/board.store';
import BoardForm from '@/component/BoardForm.vue';
import { Board } from '@/lib/interfaces/board.interface';
import { useRouter } from 'vue-router';

const router = useRouter();
const boardStore = useBoardStore();

const create = async (board: Partial<Board>) => {
  await boardStore.create(board);
  router.push('/admin/board');
}

</script>
