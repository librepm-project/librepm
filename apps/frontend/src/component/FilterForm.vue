<template>
    <v-form @submit.prevent="submit" ref="form" class="form-container">
        <v-card elevation="0" class="rounded-xl pa-6 mb-6">
            <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
                {{ props.submitButtonText === 'global.create' ? 'New Filter' : 'Edit Filter' }}
            </v-card-title>

            <v-row>
                <v-col cols="12" md="8">
                    <v-text-field
                        v-model="name"
                        :rules="[requiredRule]"
                        :label="$t('filter.name')"
                        hint="e.g., Open Issues, High Priority"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-row>
                <v-col cols="12">
                    <v-textarea
                        v-model="description"
                        :label="$t('global.description')"
                        hint="Describe what this filter is used for"
                        outlined
                        dense
                        rows="3"
                    />
                </v-col>
            </v-row>

            <v-divider class="my-6" />

            <v-card-subtitle class="px-0 mb-4 text-subtitle-2 font-weight-bold">
                {{ $t('filter.conditions') }}
            </v-card-subtitle>

            <filter-condition-builder v-model="conditions" />

            <v-divider class="my-6" />

            <v-row class="mt-4">
                <v-col cols="12" class="d-flex gap-3">
                    <v-btn
                        type="submit"
                        color="primary"
                        size="large"
                        prepend-icon="mdi-check"
                        rounded="lg"
                        class="font-weight-bold"
                    >
                        {{ $t(props.submitButtonText) }}
                    </v-btn>
                    <v-btn
                        type="button"
                        variant="outlined"
                        color="default"
                        size="large"
                        rounded="lg"
                        @click="$router.back()"
                    >
                        Cancel
                    </v-btn>
                </v-col>
            </v-row>
        </v-card>
    </v-form>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { Filter, FilterCondition } from '@/lib/interfaces/filter.interface';
import FilterConditionBuilder from '@/component/FilterConditionBuilder.vue';

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
    filter: Object as any,
});

const name = ref('');
const description = ref('');
const conditions = ref<FilterCondition[]>([]);
const form = ref(null);

watch(
    () => props.filter,
    (filter) => {
        if (filter) {
            name.value = filter.name || '';
            description.value = filter.description || '';
            conditions.value = filter.conditions ? [...filter.conditions] : [];
        }
    },
    { immediate: true },
);

const submit = async () => {
    const { valid } = await (form.value as any).validate();
    if (valid) {
        const filterData: Omit<Filter, 'id'> = {
            name: name.value,
            description: description.value,
            conditions: conditions.value,
        };
        props.onSubmit!(filterData);
    }
};
</script>

<style scoped>
.form-container {
    max-width: 100%;
}

.gap-3 {
    gap: 1rem;
}
</style>
