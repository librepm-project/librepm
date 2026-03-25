<template>
  <div v-if="relations.length > 0" class="mb-6">
    <p class="text-subtitle2 font-weight-medium mb-3">
      <v-icon x-small class="mr-1">mdi-link-variant</v-icon>
      Related Issues
    </p>

    <div v-for="group in groupedByDirection" :key="group.direction" class="mb-4">
      <p class="text-caption text-disabled font-weight-medium mb-2">
        {{ getDirectionLabel(group.direction) }}
      </p>
      <div class="d-flex flex-column gap-2">
        <div
          v-for="issue in group.issues"
          :key="issue.id"
          class="d-flex align-center gap-2 pa-2 rounded cursor-pointer hover-bg"
          @click="navigateToIssue(issue.id)"
        >
          <v-chip
            size="small"
            variant="outlined"
            class="font-weight-bold"
          >
            {{ issue.key }}
          </v-chip>
          <span class="text-body2">{{ issue.summary }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { RelatedIssue } from '@/lib/interfaces/related-issue.interface';

interface Props {
  relations: RelatedIssue[];
  currentIssueId: string;
}

const props = defineProps<Props>();
const router = useRouter();

const groupedByDirection = computed(() => {
  const grouped: Record<string, any[]> = {};

  props.relations.forEach(relation => {
    const direction = relation.directionFromCurrent || relation.type;
    if (!grouped[direction]) grouped[direction] = [];

    const otherIssue = relation.issueA?.id !== props.currentIssueId ? relation.issueA : relation.issueB;
    if (otherIssue) {
      grouped[direction].push(otherIssue);
    }
  });

  return Object.entries(grouped)
    .sort((a, b) => getSortOrder(a[0]) - getSortOrder(b[0]))
    .map(([direction, issues]) => ({ direction, issues }));
});

const getDirectionLabel = (direction: string): string => {
  const labels: Record<string, string> = {
    'blocks': 'Blocks',
    'blocked_by': 'Blocked By',
    'implements': 'Implements',
    'implemented_by': 'Implemented By',
    'duplicates': 'Duplicates',
    'is_duplicated_by': 'Duplicated By',
    'related': 'Related',
  };
  return labels[direction] || direction;
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

const navigateToIssue = (issueId: string) => {
  router.push(`/issue/${issueId}`);
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

