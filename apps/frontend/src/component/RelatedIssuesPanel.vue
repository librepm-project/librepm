<template>
  <div class="mb-6">

    <!-- Title + Add button -->
    <div class="d-flex align-center mb-3">
      <p class="text-subtitle2 font-weight-medium mb-0 flex-grow-1">
        <v-icon x-small class="mr-1">mdi-link-variant</v-icon>
        {{ t('issue.related_issues') }}
      </p>
      <v-btn icon size="x-small" variant="text" @click="openAddForm">
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </div>

    <!-- Add form -->
    <div v-if="showAddForm" class="mb-4">
      <v-autocomplete
        v-model="newTargetIssueId"
        :items="availableIssues"
        item-title="label"
        item-value="id"
        :label="t('issue.summary')"
        density="compact"
        variant="underlined"
        bg-color="transparent"
        hide-details
        class="mb-2"
      />
      <v-select
        v-model="newRelationType"
        :items="relationTypes"
        item-title="label"
        item-value="value"
        :label="t('issue.relation_type')"
        density="compact"
        variant="underlined"
        bg-color="transparent"
        hide-details
        class="mb-3"
      />
      <div class="d-flex gap-2">
        <v-btn
          size="small"
          color="primary"
          :disabled="!newTargetIssueId || !newRelationType"
          @click="addRelation"
        >
          {{ t('issue.add_relation') }}
        </v-btn>
        <v-btn size="small" variant="text" @click="showAddForm = false">
          {{ t('global.cancel') }}
        </v-btn>
      </div>
    </div>

    <!-- Relations list -->
    <div v-for="group in groupedByDirection" :key="group.direction" class="mb-4">
      <p class="text-caption text-disabled font-weight-medium mb-2">
        {{ getDirectionLabel(group.direction) }}
      </p>
      <div class="d-flex flex-column gap-2">
        <div
          v-for="item in group.items"
          :key="item.issue.id"
          class="d-flex align-center gap-2 pa-2 rounded hover-bg"
        >
          <v-chip
            size="small"
            variant="outlined"
            class="font-weight-bold cursor-pointer"
            @click="navigateToIssue(item.issue.key)"
          >
            {{ item.issue.key }}
          </v-chip>
          <span
            class="text-body2 flex-grow-1 cursor-pointer"
            @click="navigateToIssue(item.issue.key)"
          >
            {{ item.issue.summary }}
          </span>
          <v-btn
            icon
            size="x-small"
            variant="text"
            color="error"
            @click="deleteRelation(item.relationId)"
          >
            <v-icon size="small">mdi-close</v-icon>
          </v-btn>
        </div>
      </div>
    </div>

    <p v-if="relations.length === 0 && !showAddForm" class="text-caption text-disabled mb-0">
      {{ t('global.no_data') }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { RelatedIssue } from '@/lib/interfaces/related-issue.interface';
import { Issue } from '@/lib/interfaces/issue.interface';
import { useRelatedIssueStore } from '@/store/related-issue.store';
import { useIssueStore } from '@/store/issue.store';

interface Props {
  relations: RelatedIssue[];
  currentIssueId: string;
}

interface GroupedItem {
  issue: Issue;
  relationId: string;
}

const props = defineProps<Props>();
const router = useRouter();
const { t } = useI18n();
const relatedIssueStore = useRelatedIssueStore();
const issueStore = useIssueStore();

// Add form state
const showAddForm = ref(false);
const newTargetIssueId = ref<string | null>(null);
const newRelationType = ref('');

const relationTypes = computed(() => [
  { value: 'related',    label: t('issue.direction.related') },
  { value: 'blocks',     label: t('issue.direction.blocks') },
  { value: 'implements', label: t('issue.direction.implements') },
  { value: 'duplicates', label: t('issue.direction.duplicates') },
]);

const availableIssues = computed(() => {
  const relatedIds = new Set(
    props.relations.flatMap(r => [r.issueA?.id, r.issueB?.id]).filter(Boolean)
  );
  return issueStore.index
    .filter(issue => issue.id !== props.currentIssueId && !relatedIds.has(issue.id))
    .map(issue => ({ id: issue.id!, label: `${issue.key} – ${issue.summary}` }));
});

const openAddForm = async () => {
  if (issueStore.index.length === 0) {
    await issueStore.getIssues();
  }
  showAddForm.value = true;
};

const addRelation = async () => {
  if (!newTargetIssueId.value || !newRelationType.value) return;
  await relatedIssueStore.create(props.currentIssueId, newTargetIssueId.value, newRelationType.value);
  showAddForm.value = false;
  newTargetIssueId.value = null;
  newRelationType.value = '';
};

const deleteRelation = async (relationId: string) => {
  await relatedIssueStore.delete(relationId);
};

// Grouped display
const groupedByDirection = computed(() => {
  const grouped: Record<string, GroupedItem[]> = {};

  props.relations.forEach(relation => {
    const direction = relation.directionFromCurrent || relation.type;
    if (!grouped[direction]) grouped[direction] = [];
    const otherIssue = relation.issueA?.id !== props.currentIssueId ? relation.issueA : relation.issueB;
    if (otherIssue && relation.id) {
      grouped[direction].push({ issue: otherIssue, relationId: relation.id });
    }
  });

  return Object.entries(grouped)
    .sort((a, b) => getSortOrder(a[0]) - getSortOrder(b[0]))
    .map(([direction, items]) => ({ direction, items }));
});

const getDirectionLabel = (direction: string): string => {
  const key = `issue.direction.${direction}`;
  const translated = t(key);
  return translated !== key ? translated : direction;
};

const getSortOrder = (direction: string): number => {
  const order: Record<string, number> = {
    'blocks': 1,
    'blocked_by': 2,
    'implements': 3,
    'implemented_by': 4,
    'duplicates': 5,
    'is_duplicated_by': 6,
    'related': 7,
  };
  return order[direction] || 99;
};

const navigateToIssue = (issueKey: string | undefined) => {
  if (issueKey) router.push(`/issue/key/${issueKey}`);
};
</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

.cursor-pointer {
  cursor: pointer;
}

.hover-bg:hover {
  background-color: rgba(63, 81, 181, 0.05);
}
</style>
