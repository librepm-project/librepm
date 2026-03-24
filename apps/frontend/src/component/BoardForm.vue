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

        <v-card v-for="(column, index) in form.boardColumns" :key="index" variant="outlined" class="mb-4 pa-4">
            <v-row align="center">
                <v-col cols="12" md="4">
                    <v-text-field v-model="column.label" :label="t('board.column_label')" :rules="[requiredRule]" hide-details></v-text-field>
                </v-col>
                <v-col cols="12" md="7">
                    <v-select v-model="column.statusIds" :items="statusStore.index" item-title="name" item-value="id"
                        :label="t('board.column_statuses')" multiple chips closable-chips hide-details></v-select>
                </v-col>
                <v-col cols="12" md="1" class="text-right">
                    <v-btn icon color="error" variant="text" @click="removeColumn(index)">
                        <v-icon>mdi-delete</v-icon>
                    </v-btn>
                </v-col>
            </v-row>
        </v-card>

        <v-btn class="mt-4" type="submit" color="primary" prepend-icon="mdi-content-save" block>
            {{ t(props.submitButtonText) }}
        </v-btn>
    </v-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { useI18n } from 'vue-i18n';
import { useStatusStore } from '@/store/status.store';
import { Board, BoardColumn } from '@/lib/interfaces/board.interface';

const { t } = useI18n();
const statusStore = useStatusStore();

const props = defineProps<{
    onSubmit: (item: Partial<Board>) => void;
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

// If initialData has columns, map them to have statusIds for the select
if (props.initialData?.boardColumns) {
    form.value.boardColumns = props.initialData.boardColumns.map(col => ({
        ...col,
        statusIds: col.statuses?.map(s => s.id) || col.statusIds || []
    }));
}

onMounted(async () => {
    await statusStore.getStatuses();
});

const addColumn = () => {
    if (!form.value.boardColumns) form.value.boardColumns = [];
    form.value.boardColumns.push({
        id: '',
        label: '',
        statuses: [],
        statusIds: []
    });
};

const removeColumn = (index: number) => {
    form.value.boardColumns?.splice(index, 1);
};

const submit = async () => {
    const { valid } = await formRef.value.validate();
    if (valid) {
        // Clean up the object before sending
        const payload = {
            ...form.value,
            boardColumns: form.value.boardColumns?.map(col => ({
                id: col.id,
                label: col.label,
                statusIds: col.statusIds
            }))
        };
        props.onSubmit(payload);
    }
};
</script>
