<template>
    <div>
        <div v-for="(condition, index) in conditions" :key="index" class="d-flex align-center gap-2 mb-3">
            <v-select
                v-model="condition.field"
                :items="fieldOptions"
                item-title="label"
                item-value="key"
                :label="$t('filter.field')"
                density="compact"
                hide-details
                style="min-width: 140px; max-width: 160px"
                @update:modelValue="onFieldChange(condition)"
            />
            <v-select
                v-model="condition.op"
                :items="getOperatorsForField(condition.field)"
                item-title="label"
                item-value="key"
                :label="$t('filter.operator')"
                density="compact"
                hide-details
                style="min-width: 120px; max-width: 140px"
            />
            <v-select
                v-if="getFieldDef(condition.field)?.valueType === 'select'"
                v-model="condition.value"
                :items="getValueOptions(condition.field)"
                item-title="name"
                item-value="id"
                :label="$t('filter.value')"
                density="compact"
                hide-details
                class="flex-grow-1"
            />
            <v-text-field
                v-else
                v-model="condition.value"
                :label="$t('filter.value')"
                density="compact"
                hide-details
                class="flex-grow-1"
            />
            <v-btn
                icon="mdi-minus"
                variant="text"
                color="error"
                density="compact"
                @click="removeCondition(index)"
            />
        </div>

        <v-btn
            variant="tonal"
            color="primary"
            size="small"
            prepend-icon="mdi-plus"
            @click="addCondition"
        >
            {{ $t('filter.add_condition') }}
        </v-btn>
    </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue';
import filterApi from '@/api/filter.api';
import { useProjectStore } from '@/store/project.store';
import { useTrackerStore } from '@/store/tracker.store';
import { useStatusStore } from '@/store/status.store';
import { useUserStore } from '@/store/user.store';
import { FilterCondition, FilterConditionField, FilterConditionOptions } from '@/lib/interfaces/filter.interface';

const props = defineProps<{
    modelValue: FilterCondition[];
}>();

const emit = defineEmits<{
    (e: 'update:modelValue', value: FilterCondition[]): void;
}>();

const conditions = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val),
});

const conditionOptions = ref<FilterConditionOptions>({ fields: [] });
const projectStore = useProjectStore();
const trackerStore = useTrackerStore();
const statusStore = useStatusStore();
const userStore = useUserStore();

onMounted(async () => {
    const [opts] = await Promise.all([
        filterApi.conditionOptions(),
        projectStore.getAllItems(),
        trackerStore.getAllItems(),
        statusStore.getAllItems(),
        userStore.getAllItems(),
    ]);
    conditionOptions.value = opts;
});

const fieldOptions = computed(() => conditionOptions.value.fields);

function getFieldDef(fieldKey: string): FilterConditionField | undefined {
    return conditionOptions.value.fields.find(f => f.key === fieldKey);
}

function getOperatorsForField(fieldKey: string) {
    return getFieldDef(fieldKey)?.operators ?? [];
}

function getValueOptions(fieldKey: string): any[] {
    const def = getFieldDef(fieldKey);
    if (!def?.valueEndpoint) return [];
    if (def.valueEndpoint === '/project') return projectStore.index;
    if (def.valueEndpoint === '/tracker') return trackerStore.index;
    if (def.valueEndpoint === '/status') return statusStore.index;
    if (def.valueEndpoint === '/user') return userStore.index.map(u => ({ id: u.id, name: `${u.firstName} ${u.lastName}` }));
    return [];
}

function onFieldChange(condition: FilterCondition) {
    const def = getFieldDef(condition.field);
    condition.op = def?.operators[0]?.key ?? '';
    condition.value = '';
}

function addCondition() {
    const firstField = conditionOptions.value.fields[0];
    const newCondition: FilterCondition = {
        field: firstField?.key ?? '',
        op: firstField?.operators[0]?.key ?? '',
        value: '',
    };
    emit('update:modelValue', [...props.modelValue, newCondition]);
}

function removeCondition(index: number) {
    const updated = props.modelValue.filter((_, i) => i !== index);
    emit('update:modelValue', updated);
}
</script>
