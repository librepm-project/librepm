<template>
    <v-card variant="outlined" rounded="xl" class="h-100">
        <v-card-title class="d-flex align-center py-3 px-4 gap-2">
            <v-icon class="drag-handle text-grey-lighten-1 cursor-grab" size="small">mdi-drag-vertical</v-icon>
            <span class="text-subtitle-1 font-weight-bold flex-grow-1">{{ widget.title }}</span>
            <v-menu>
                <template #activator="{ props }">
                    <v-btn
                        icon="mdi-dots-vertical"
                        variant="text"
                        density="compact"
                        size="small"
                        v-bind="props"
                    />
                </template>
                <v-list density="compact" min-width="180">
                    <v-list-subheader class="text-caption font-weight-bold text-uppercase">
                        {{ $t('dashboard.width_full').replace('Full width', 'Width') }}
                    </v-list-subheader>
                    <v-list-item
                        v-for="opt in widthOptions"
                        :key="opt.value"
                        :active="widget.width === opt.value"
                        active-color="primary"
                        rounded="lg"
                        class="mx-1"
                        @click="emit('resize', opt.value)"
                    >
                        <template #prepend>
                            <v-icon size="small">{{ opt.icon }}</v-icon>
                        </template>
                        <v-list-item-title>{{ $t(opt.label) }}</v-list-item-title>
                        <template v-if="widget.width === opt.value" #append>
                            <v-icon size="x-small" color="primary">mdi-check</v-icon>
                        </template>
                    </v-list-item>
                    <v-divider class="my-1" />
                    <v-list-item
                        rounded="lg"
                        class="mx-1"
                        base-color="error"
                        @click="emit('remove', widget.id!)"
                    >
                        <template #prepend>
                            <v-icon size="small">mdi-delete-outline</v-icon>
                        </template>
                        <v-list-item-title>{{ $t('global.delete') }}</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
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
    (e: 'resize', width: string): void;
}>();

const widthOptions = [
    { value: '1/3', label: 'dashboard.width_third', icon: 'mdi-view-column-outline' },
    { value: '1/2', label: 'dashboard.width_half', icon: 'mdi-view-column' },
    { value: '2/3', label: 'dashboard.width_two_thirds', icon: 'mdi-view-sequential-outline' },
    { value: 'full', label: 'dashboard.width_full', icon: 'mdi-view-stream-outline' },
];

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

<style scoped>
.drag-handle {
    cursor: grab;
}
.drag-handle:active {
    cursor: grabbing;
}
</style>
