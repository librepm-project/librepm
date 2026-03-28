<template>
    <v-form @submit.prevent="submit" ref="formRef">
        <v-text-field v-model="form.name" :rules="[requiredRule]" :label="t('global.name')"></v-text-field>
        <v-text-field v-model="form.description" :label="t('global.description')"></v-text-field>

        <v-divider class="my-4"></v-divider>
        <div class="d-flex justify-space-between align-center mb-4">
            <h3>{{ t('board.columns') }}</h3>
            <v-btn color="secondary" size="small" @click="addColumn" prepend-icon="mdi-plus">
                {{ t('board.add_column') }}
            </v-btn>
        </div>

        <v-card
            v-for="(column, index) in form.boardColumns"
            :key="column.id || index"
            variant="outlined"
            class="mb-4 column-card"
            :class="{ 'drag-over': dragOverIndex === index, 'is-dragging': dragSourceIndex === index }"
            draggable="true"
            @dragstart="onColumnDragStart($event, index)"
            @dragover="onColumnDragOver($event, index)"
            @dragleave="onColumnDragLeave($event)"
            @drop.prevent="onColumnDrop(index)"
            @dragend="onColumnDragEnd"
        >
            <v-row align="center" class="pa-4" no-gutters>
                <v-col cols="auto" class="d-flex align-center pe-3">
                    <v-icon class="drag-handle" color="grey-lighten-1">mdi-drag-vertical</v-icon>
                </v-col>
                <v-col cols="12" md="4" class="pe-2">
                    <v-text-field v-model="column.label" :label="t('board.column_label')" :rules="[requiredRule]" hide-details></v-text-field>
                </v-col>
                <v-col cols="12" md="6" class="pe-2">
                    <v-select v-model="column.statusIds" :items="statusStore.index" item-title="name" item-value="id"
                        :label="t('board.column_statuses')" multiple chips closable-chips hide-details></v-select>
                </v-col>
                <v-col cols="auto" class="d-flex align-center">
                    <v-btn icon color="error" variant="text" @click="removeColumn(index)">
                        <v-icon>mdi-delete</v-icon>
                    </v-btn>
                </v-col>
            </v-row>
        </v-card>

        <v-divider class="my-6" />

        <v-row class="mt-4">
            <v-col cols="12" class="d-flex gap-3 align-center">
                <v-btn
                    type="submit"
                    color="primary"
                    size="large"
                    prepend-icon="mdi-check"
                    rounded="lg"
                    class="font-weight-bold"
                >
                    {{ t(props.submitButtonText) }}
                </v-btn>
                <v-btn
                    v-if="onDelete"
                    type="button"
                    color="error"
                    variant="text"
                    size="large"
                    prepend-icon="mdi-delete"
                    rounded="lg"
                    class="font-weight-bold"
                    @click="onDelete"
                >
                    {{ t('global.delete') }}
                </v-btn>
                <v-spacer />
                <v-btn
                    type="button"
                    variant="outlined"
                    color="default"
                    size="large"
                    rounded="lg"
                    @click="router.back()"
                >
                    Cancel
                </v-btn>
            </v-col>
        </v-row>
    </v-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { useI18n } from 'vue-i18n';
import { useStatusStore } from '@/store/status.store';
import { Board } from '@/lib/interfaces/board.interface';
import { useRouter } from 'vue-router';

const { t } = useI18n();
const statusStore = useStatusStore();
const router = useRouter();

const props = defineProps<{
    onSubmit: (item: Partial<Board>) => void;
    onDelete?: () => void;
    submitButtonText: string;
    initialData?: Partial<Board>;
}>();

const formRef = ref<any>(null);

const form = ref<Partial<Board>>({
    name: '',
    description: '',
    boardColumns: [],
    ...props.initialData
});

if (props.initialData?.boardColumns) {
    form.value.boardColumns = props.initialData.boardColumns.map(col => ({
        ...col,
        statusIds: col.statuses?.map(s => s.id) || col.statusIds || []
    }));
}

onMounted(async () => {
    await statusStore.getAll();
});

const addColumn = () => {
    if (!form.value.boardColumns) form.value.boardColumns = [];
    form.value.boardColumns.push({ id: '', label: '', statuses: [], statusIds: [] });
};

const removeColumn = (index: number) => {
    form.value.boardColumns?.splice(index, 1);
};

// --- Drag & Drop column reordering ---
const dragSourceIndex = ref<number | null>(null);
const dragOverIndex = ref<number | null>(null);

const onColumnDragStart = (event: DragEvent, index: number) => {
    dragSourceIndex.value = index;
    if (event.dataTransfer) event.dataTransfer.effectAllowed = 'move';
};

const onColumnDragOver = (event: DragEvent, index: number) => {
    if (dragSourceIndex.value === null || dragSourceIndex.value === index) return;
    event.preventDefault();
    dragOverIndex.value = index;
    if (event.dataTransfer) event.dataTransfer.dropEffect = 'move';
};

const onColumnDragLeave = (event: DragEvent) => {
    const relatedTarget = event.relatedTarget as Node | null;
    if (relatedTarget && (event.currentTarget as Element).contains(relatedTarget)) return;
    dragOverIndex.value = null;
};

const onColumnDrop = (targetIndex: number) => {
    if (dragSourceIndex.value === null || dragSourceIndex.value === targetIndex) return;
    const columns = form.value.boardColumns!;
    const [moved] = columns.splice(dragSourceIndex.value, 1);
    columns.splice(targetIndex, 0, moved);
    dragSourceIndex.value = null;
    dragOverIndex.value = null;
};

const onColumnDragEnd = () => {
    dragSourceIndex.value = null;
    dragOverIndex.value = null;
};

const submit = async () => {
    const { valid } = await formRef.value.validate();
    if (valid) {
        const payload = {
            ...form.value,
            boardColumns: form.value.boardColumns?.map((col, index) => ({
                id: col.id,
                label: col.label,
                statusIds: col.statusIds,
                weight: index
            }))
        };
        props.onSubmit(payload);
    }
};
</script>

<style scoped>
.column-card {
    transition: opacity 0.15s, border-color 0.15s;
}

.column-card.is-dragging {
    opacity: 0.35;
}

.column-card.drag-over {
    border-color: rgb(var(--v-theme-primary)) !important;
    border-width: 2px !important;
}

.drag-handle {
    cursor: grab;
}

.drag-handle:active {
    cursor: grabbing;
}
</style>
