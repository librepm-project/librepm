<template>
  <div>
    <board-form :onSubmit="boardStore.postBoard" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useBoardStore } from '@/store/board.store';

import BoardForm from '@/component/BoardForm.vue';

const boardStore = useBoardStore();

</script>