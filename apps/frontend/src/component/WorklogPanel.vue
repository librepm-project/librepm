<template>
  <div class="mb-6">
    <!-- Fejléc -->
    <div class="d-flex align-center mb-3">
      <p class="text-subtitle2 font-weight-medium mb-0 flex-grow-1">
        <v-icon
          x-small
          class="mr-1"
        >
          mdi-clock-outline
        </v-icon>
        {{ t('worklog.work_log') }}
        <span
          v-if="totalSeconds > 0"
          class="text-caption text-disabled ml-2"
        >
          {{ t('worklog.total') }}: {{ formatSeconds(totalSeconds) }}
        </span>
      </p>
      <v-btn
        icon
        size="x-small"
        variant="text"
        @click="openAddForm"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </div>

    <!-- Hozzáadás form -->
    <div
      v-if="showAddForm"
      class="mb-4"
    >
      <div class="d-flex gap-2 mb-2">
        <v-text-field
          v-model.number="addForm.hours"
          :label="t('worklog.hours')"
          type="number"
          min="0"
          density="compact"
          variant="underlined"
          hide-details
          style="max-width: 80px"
        />
        <v-text-field
          v-model.number="addForm.minutes"
          :label="t('worklog.minutes')"
          type="number"
          min="0"
          max="59"
          density="compact"
          variant="underlined"
          hide-details
          style="max-width: 80px"
        />
        <v-text-field
          v-model="addForm.loggedAt"
          :label="t('worklog.logged_at')"
          type="date"
          density="compact"
          variant="underlined"
          hide-details
          class="flex-grow-1"
        />
      </div>
      <v-text-field
        v-model="addForm.comment"
        :label="t('worklog.comment')"
        density="compact"
        variant="underlined"
        hide-details
        class="mb-3"
      />
      <div class="d-flex gap-2">
        <v-btn
          size="small"
          color="primary"
          :disabled="addFormSeconds <= 0"
          @click="submitAdd"
        >
          {{ t('worklog.add_worklog') }}
        </v-btn>
        <v-btn
          size="small"
          variant="text"
          @click="showAddForm = false"
        >
          {{ t('global.cancel') }}
        </v-btn>
      </div>
    </div>

    <!-- Lista -->
    <div
      v-for="worklog in worklogs"
      :key="worklog.id"
      class="mb-3"
    >
      <!-- Szerkesztés form -->
      <div
        v-if="editingId === worklog.id"
        class="pa-2 rounded edit-bg"
      >
        <div class="d-flex gap-2 mb-2">
          <v-text-field
            v-model.number="editForm.hours"
            :label="t('worklog.hours')"
            type="number"
            min="0"
            density="compact"
            variant="underlined"
            hide-details
            style="max-width: 80px"
          />
          <v-text-field
            v-model.number="editForm.minutes"
            :label="t('worklog.minutes')"
            type="number"
            min="0"
            max="59"
            density="compact"
            variant="underlined"
            hide-details
            style="max-width: 80px"
          />
          <v-text-field
            v-model="editForm.loggedAt"
            :label="t('worklog.logged_at')"
            type="date"
            density="compact"
            variant="underlined"
            hide-details
            class="flex-grow-1"
          />
        </div>
        <v-text-field
          v-model="editForm.comment"
          :label="t('worklog.comment')"
          density="compact"
          variant="underlined"
          hide-details
          class="mb-2"
        />
        <div class="d-flex gap-2">
          <v-btn
            size="small"
            color="primary"
            :disabled="editFormSeconds <= 0"
            @click="submitEdit(worklog.id!)"
          >
            {{ t('global.update') }}
          </v-btn>
          <v-btn
            size="small"
            variant="text"
            @click="editingId = null"
          >
            {{ t('global.cancel') }}
          </v-btn>
        </div>
      </div>

      <!-- Sor nézet -->
      <div
        v-else
        class="d-flex align-center gap-2 pa-2 rounded hover-bg"
      >
        <div class="flex-grow-1">
          <div class="d-flex align-center gap-2">
            <span class="text-body2 font-weight-medium">{{ formatSeconds(worklog.seconds) }}</span>
            <span class="text-caption text-disabled">{{ formatDate(worklog.loggedAt) }}</span>
            <span
              v-if="worklog.user"
              class="text-caption text-disabled"
            >
              {{ worklog.user.firstName }} {{ worklog.user.lastName }}
            </span>
          </div>
          <div
            v-if="worklog.comment"
            class="text-caption text-medium-emphasis mt-1"
          >
            {{ worklog.comment }}
          </div>
        </div>
        <v-btn
          icon
          size="x-small"
          variant="text"
          @click="startEdit(worklog)"
        >
          <v-icon size="small">
            mdi-pencil
          </v-icon>
        </v-btn>
        <v-btn
          icon
          size="x-small"
          variant="text"
          color="error"
          @click="remove(worklog.id!)"
        >
          <v-icon size="small">
            mdi-close
          </v-icon>
        </v-btn>
      </div>
    </div>

    <p
      v-if="worklogs.length === 0 && !showAddForm"
      class="text-caption text-disabled mb-0"
    >
      {{ t('worklog.no_worklogs') }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { Worklog } from '@/lib/interfaces/worklog.interface';
import { useWorklogStore } from '@/store/worklog.store';

interface Props {
  issueId: string;
  worklogs: Worklog[];
}

const props = defineProps<Props>();
const { t } = useI18n();
const worklogStore = useWorklogStore();

// ── Segédfüggvények ──────────────────────────────────────────────────────────

function formatSeconds(s: number): string {
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  if (h > 0 && m > 0) return `${h}h ${m}m`;
  if (h > 0) return `${h}h`;
  if (m > 0) return `${m}m`;
  return `${s}s`;
}

function formatDate(iso: string): string {
  return new Date(iso).toLocaleDateString();
}

function todayIso(): string {
  return new Date().toISOString().split('T')[0];
}

// ── Összesítő ────────────────────────────────────────────────────────────────

const totalSeconds = computed(() => props.worklogs.reduce((sum, w) => sum + w.seconds, 0));

// ── Hozzáadás ────────────────────────────────────────────────────────────────

const showAddForm = ref(false);
const addForm = ref({ hours: 0, minutes: 0, loggedAt: todayIso(), comment: '' });
const addFormSeconds = computed(() => addForm.value.hours * 3600 + addForm.value.minutes * 60);

const openAddForm = () => {
  addForm.value = { hours: 0, minutes: 0, loggedAt: todayIso(), comment: '' };
  showAddForm.value = true;
};

const submitAdd = async () => {
  await worklogStore.create(props.issueId, {
    seconds: addFormSeconds.value,
    comment: addForm.value.comment,
    loggedAt: new Date(addForm.value.loggedAt).toISOString(),
  });
  showAddForm.value = false;
};

// ── Szerkesztés ──────────────────────────────────────────────────────────────

const editingId = ref<string | null>(null);
const editForm = ref({ hours: 0, minutes: 0, loggedAt: todayIso(), comment: '' });
const editFormSeconds = computed(() => editForm.value.hours * 3600 + editForm.value.minutes * 60);

const startEdit = (worklog: Worklog) => {
  editingId.value = worklog.id!;
  editForm.value = {
    hours: Math.floor(worklog.seconds / 3600),
    minutes: Math.floor((worklog.seconds % 3600) / 60),
    loggedAt: worklog.loggedAt.split('T')[0],
    comment: worklog.comment ?? '',
  };
};

const submitEdit = async (worklogId: string) => {
  await worklogStore.update(props.issueId, worklogId, {
    seconds: editFormSeconds.value,
    comment: editForm.value.comment,
    loggedAt: new Date(editForm.value.loggedAt).toISOString(),
  });
  editingId.value = null;
};

// ── Törlés ───────────────────────────────────────────────────────────────────

const remove = async (worklogId: string) => {
  await worklogStore.destroy(props.issueId, worklogId);
};
</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

.hover-bg:hover {
  background-color: rgba(63, 81, 181, 0.05);
}

.edit-bg {
  background-color: rgba(0, 0, 0, 0.03);
}
</style>
