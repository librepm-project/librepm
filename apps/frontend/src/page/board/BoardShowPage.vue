<template>
  <v-container fluid v-if="boardStore.current">
    <div class="d-flex align-center mb-4">
      <h2 class="text-h4">{{ boardStore.current.name }}</h2>
      <v-spacer></v-spacer>
    </div>

    <v-row class="board-row flex-nowrap overflow-x-auto">
      <v-col v-if="boardStore.current.boardColumns && boardStore.current.boardColumns.length > 0"
        v-for="column in boardStore.current.boardColumns" :key="column.id" class="board-column" cols="12" sm="6" md="4" lg="3">
        <v-card variant="flat" color="grey-lighten-4" class="h-100 d-flex flex-column">
          <v-card-title class="d-flex align-center py-2 bg-grey-lighten-3">
            <span class="text-subtitle-1 font-weight-bold">{{ column.label }}</span>
            <v-chip size="x-small" class="ms-2" variant="flat">{{ getIssuesForColumn(column).length }}</v-chip>
          </v-card-title>
          
          <v-card-text class="flex-grow-1 pa-2 overflow-y-auto">
            <div v-for="issue in getIssuesForColumn(column)" :key="issue.id" class="mb-2">
              <v-card variant="flat" class="pa-3 issue-card border" @click="router.push(`/issue/${issue.id}`)">
                <div class="d-flex justify-space-between align-start mb-1">
                  <span class="text-caption text-grey">{{ issue.key }}</span>
                  <tracker-chip :tracker="issue.tracker" size="x-small" />
                </div>
                <div class="text-body-2 font-weight-medium line-clamp-2">{{ issue.summary }}</div>
              </v-card>
            </div>
            <div v-if="getIssuesForColumn(column).length === 0" class="text-center py-4 text-grey-lighten-1">
              <v-icon size="large">mdi-tray-blank</v-icon>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col v-else cols="12" class="text-center py-8">
        <v-icon icon="mdi-alert-circle-outline" size="x-large" color="grey-lighten-1"></v-icon>
        <p class="text-h6 text-grey-lighten-1 mt-2">{{ t('board.no_columns_defined') }}</p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useBoardStore } from '@/store/board.store';
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import TrackerChip from '@/component/TrackerChip.vue';
import { BoardColumn } from '@/lib/interfaces/board.interface';

const boardStore = useBoardStore();
const issueStore = useIssueStore();
const layoutStore = useLayoutStore();
const route = useRoute();
const router = useRouter();
const { t } = useI18n();

const loadData = async () => {
  const boardId = route.params.boardId?.toString();
  if (boardId) {
    await Promise.all([
      boardStore.getBoard(boardId),
      issueStore.getIssues(),
      boardStore.getBoards()
    ]);

    console.log('Board Data:', JSON.stringify(boardStore.current)); // Log board data
    console.log('Issue Data:', JSON.stringify(issueStore.index));   // Log issue data

    layoutStore.setSidebar(boardStore.index.map(board => ({
      key: board.id,
      title: board.name,
      link: `/board/${board.id}`,
    })));

    layoutStore.setActions([
      {
        text: 'global.edit',
        to: `/admin/board/${boardId}/edit`,
        color: 'primary',
        icon: 'mdi-pencil'
      }
    ]);
  }
};

onMounted(loadData);

watch(() => route.params.boardId, loadData);

onUnmounted(() => {
  layoutStore.resetSidebar();
  layoutStore.resetActions();
});

const getIssuesForColumn = (column: BoardColumn) => {
  console.log('Column data for filtering:', JSON.stringify(column)); // Log column data
  if (!column.statuses || !issueStore.index) return [];
  const statusIds = column.statuses.map(s => s.id);
  return issueStore.index.filter(issue => issue.status && statusIds.includes(issue.status.id));
};
</script>

<style scoped>
.board-row {
  height: calc(100vh - 180px);
  min-height: 500px;
}

.board-column {
  min-width: 300px;
  max-width: 400px;
}

.issue-card {
  cursor: pointer;
  transition: transform 0.1s, box-shadow 0.1s;
}

.issue-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.1) !important;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
