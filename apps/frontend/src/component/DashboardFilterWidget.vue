<template>
    <v-card variant="outlined" rounded="xl" class="h-100">
        <v-card-title class="d-flex align-center justify-space-between py-3 px-4">
            <span class="text-subtitle-1 font-weight-bold">{{ widget.title }}</span>
            <v-btn
                icon="mdi-close"
                variant="text"
                density="compact"
                size="small"
                @click="emit('remove', widget.id)"
            />
        </v-card-title>

        <v-divider />

        <v-card-text class="pa-0">
            <div v-if="loading" class="d-flex justify-center py-6">
                <v-progress-circular indeterminate color="primary" size="24" />
            </div>

            <v-list v-else-if="issues.length > 0" density="compact" lines="one">
                <v-list-item
                    v-for="issue in issues"
                    :key="issue.id"
                    :to="`/issue/${issue.id}`"
                    rounded="lg"
                    class="mx-2 mb-1"
                >
                    <template #prepend>
                        <span class="text-caption text-grey mr-2">{{ issue.key }}</span>
                    </template>
                    <v-list-item-title class="text-body-2">{{ issue.summary }}</v-list-item-title>
                    <template #append>
                        <status-chip :status="issue.status" size="x-small" />
                    </template>
                </v-list-item>
            </v-list>

            <div v-else class="text-center py-6 text-grey">
                <v-icon size="large">mdi-tray-blank</v-icon>
                <p class="text-caption mt-1">No issues match this filter</p>
            </div>
        </v-card-text>
    </v-card>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { DashboardFilterWidget } from '@/lib/interfaces/dashboard.interface';
import { Issue } from '@/lib/interfaces/issue.interface';
import issueApi from '@/api/issue.api';
import StatusChip from '@/component/StatusChip.vue';

const props = defineProps<{
    widget: DashboardFilterWidget;
}>();

const emit = defineEmits<{
    (e: 'remove', widgetId: string): void;
}>();

const issues = ref<Issue[]>([]);
const loading = ref(true);

onMounted(async () => {
    try {
        const response = await issueApi.index({ filterId: props.widget.filterId });
        issues.value = response;
    } finally {
        loading.value = false;
    }
});
</script>
