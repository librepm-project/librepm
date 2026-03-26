<template>
  <v-container fluid v-if="boardStore.current">
    <div class="d-flex align-center mb-5">
      <v-menu transition="slide-y-transition">
        <template #activator="{ props, isActive }">
          <v-btn
            v-bind="props"
            variant="tonal"
            color="primary"
            rounded="xl"
            size="large"
            prepend-icon="mdi-view-dashboard"
            :append-icon="isActive ? 'mdi-chevron-up' : 'mdi-chevron-down'"
            class="board-switcher-btn"
          >
            {{ boardStore.current.name }}
          </v-btn>
        </template>
        <v-card rounded="xl" elevation="4" min-width="240" class="mt-1">
          <v-list>
            <v-list-subheader class="text-caption font-weight-bold text-uppercase">
              {{ t('board.boards') }}
            </v-list-subheader>
            <v-list-item
              v-for="board in boardStore.index"
              :key="board.id"
              :to="`/board/${board.id}`"
              rounded="lg"
              :active="board.id === boardStore.current?.id"
              active-color="primary"
              class="mx-2 mb-1"
            >
              <template #prepend>
                <v-icon :color="board.id === boardStore.current?.id ? 'primary' : undefined">
                  mdi-view-dashboard-outline
                </v-icon>
              </template>
              <v-list-item-title class="font-weight-medium">{{ board.name }}</v-list-item-title>
              <template v-if="board.id === boardStore.current?.id" #append>
                <v-icon size="small" color="primary">mdi-check</v-icon>
              </template>
            </v-list-item>
          </v-list>
        </v-card>
      </v-menu>
    </div>

    <v-row class="board-row flex-nowrap overflow-x-auto">
      <v-col v-if="sortedColumns.length > 0"
        v-for="column in sortedColumns" :key="column.id" class="board-column" cols="12" sm="6" md="4" lg="3">
        <v-card
          variant="flat"
          color="grey-lighten-4"
          class="h-100 d-flex flex-column column-drop-zone"
          :class="getColumnClass(column)"
          @dragover="onColumnDragOver($event, column)"
          @dragleave="onColumnDragLeave($event, column)"
          @drop.prevent="onColumnDrop(column)"
        >
          <v-card-title class="d-flex align-center py-2 bg-grey-lighten-3">
            <span class="text-subtitle-1 font-weight-bold">{{ column.label }}</span>
            <v-chip size="x-small" class="ms-2" variant="flat">{{ getIssuesForColumn(column).length }}</v-chip>
          </v-card-title>

          <v-card-text class="flex-grow-1 pa-2 overflow-y-auto">
            <div v-for="issue in getIssuesForColumn(column)" :key="issue.id" class="mb-2">
              <v-card
                variant="flat"
                class="pa-3 issue-card border"
                :class="{ 'is-dragging': draggingIssue?.id === issue.id }"
                draggable="true"
                @dragstart="onIssueDragStart($event, issue)"
                @dragend="onIssueDragEnd"
                @click="router.push(`/issue/key/${issue.key}`)"
              >
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
import { computed, onMounted, onUnmounted, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useBoardStore } from '@/store/board.store';
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import TrackerChip from '@/component/TrackerChip.vue';
import { BoardColumn } from '@/lib/interfaces/board.interface';
import { Issue } from '@/lib/interfaces/issue.interface';
import issueApi from '@/api/issue.api';

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
  layoutStore.resetActions();
});

const sortedColumns = computed(() =>
  [...(boardStore.current?.boardColumns ?? [])].sort((a, b) => (a.weight ?? 0) - (b.weight ?? 0))
);

const getIssuesForColumn = (column: BoardColumn) => {
  if (!column.statuses || !issueStore.index) return [];
  const statusIds = column.statuses.map(s => s.id);
  return issueStore.index.filter(issue => issue.status && statusIds.includes(issue.status.id));
};

// --- Drag & Drop issue status change ---
const draggingIssue = ref<Issue | null>(null);
const dragOverColumnId = ref<string | null>(null);
const dragOverValid = ref(false);

const onIssueDragStart = (event: DragEvent, issue: Issue) => {
  draggingIssue.value = issue;
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move';
    event.dataTransfer.setData('text/plain', issue.id || '');
  }
};

const onIssueDragEnd = () => {
  draggingIssue.value = null;
  dragOverColumnId.value = null;
  dragOverValid.value = false;
};

const isValidDrop = (column: BoardColumn): boolean => {
  if (!draggingIssue.value) return false;
  return column.statuses != null && column.statuses.length > 0;
};

const onColumnDragOver = (event: DragEvent, column: BoardColumn) => {
  event.preventDefault();
  dragOverColumnId.value = column.id;
  dragOverValid.value = isValidDrop(column);
  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = dragOverValid.value ? 'move' : 'none';
  }
};

const onColumnDragLeave = (event: DragEvent, column: BoardColumn) => {
  const relatedTarget = event.relatedTarget as Node | null;
  if (relatedTarget && (event.currentTarget as Element).contains(relatedTarget)) return;
  if (dragOverColumnId.value === column.id) {
    dragOverColumnId.value = null;
    dragOverValid.value = false;
  }
};

const getColumnClass = (column: BoardColumn) => {
  if (dragOverColumnId.value !== column.id) return '';
  return dragOverValid.value ? 'drop-valid' : 'drop-invalid';
};

const onColumnDrop = async (column: BoardColumn) => {
  dragOverColumnId.value = null;
  dragOverValid.value = false;

  const issue = draggingIssue.value;
  draggingIssue.value = null;

  if (!issue || !column.statuses || column.statuses.length === 0) return;

  const newStatus = column.statuses[0];
  if (issue.status?.id === newStatus.id) return;

  // Optimistic update
  const issueInIndex = issueStore.index.find(i => i.id === issue.id);
  const prevStatus = issue.status;
  const prevStatusId = issue.statusId;

  if (issueInIndex) {
    issueInIndex.status = newStatus;
    issueInIndex.statusId = newStatus.id;
  }

  try {
    await issueApi.update(issue.id!, { statusId: newStatus.id });
  } catch {
    if (issueInIndex) {
      issueInIndex.status = prevStatus;
      issueInIndex.statusId = prevStatusId;
    }
  }
};
</script>

<style scoped>
.board-switcher-btn {
  font-size: 1.1rem;
  font-weight: 700;
  letter-spacing: 0;
  text-transform: none;
  padding-inline: 20px;
}

.board-row {
  height: calc(100vh - 180px);
  min-height: 500px;
}

.board-column {
  min-width: 300px;
  max-width: 400px;
}

.issue-card {
  cursor: grab;
  transition: transform 0.1s, box-shadow 0.1s, opacity 0.15s;
  user-select: none;
  border: 1px solid rgba(0, 0, 0, 0.12) !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05) !important;
  background-color: white !important;
}

.issue-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1) !important;
}

.issue-card:active {
  cursor: grabbing;
}

.issue-card.is-dragging {
  opacity: 0.4;
}

.column-drop-zone {
  transition: outline 0.1s, background-color 0.1s;
}

.drop-valid {
  outline: 2px solid rgb(var(--v-theme-primary));
  background-color: rgba(var(--v-theme-primary), 0.04) !important;
}

.drop-invalid {
  outline: 2px dashed rgb(var(--v-theme-error));
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
