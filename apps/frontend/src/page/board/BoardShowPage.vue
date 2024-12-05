<template>
  <div>
    <v-select
      :label="t('global.select')"
      :items="boardStore.index"
      item-title="name"
      item-value="id"
    ></v-select>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useBoardStore } from '@/store/board.store';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();
const boardStore = useBoardStore();

onMounted(async () => {
  await boardStore.getBoards();
})

let selectedBoardId = "0";
let selectedBoard = boardStore.index[Number(selectedBoardId)];
 
</script>