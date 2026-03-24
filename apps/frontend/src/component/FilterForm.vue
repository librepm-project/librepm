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
                Filter Conditions (JSON)
            </v-card-subtitle>

            <v-row>
                <v-col cols="12">
                    <v-textarea
                        v-model="conditionsJson"
                        label="Conditions (JSON format)"
                        hint='e.g., [{"field":"status","op":"eq","value":"open"}]'
                        outlined
                        dense
                        rows="6"
                        monospace
                    />
                </v-col>
            </v-row>

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
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { useRouter } from 'vue-router';
import { Filter } from '@/lib/interfaces/filter.interface';

const router = useRouter();

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
    filter: Object as any,
})

const name = ref("");
const description = ref("");
const conditionsJson = ref("[]");
const form = ref(null);

onMounted(() => {
    if (props.filter) {
        name.value = props.filter.name || "";
        description.value = props.filter.description || "";
        conditionsJson.value = JSON.stringify(props.filter.conditions || [], null, 2);
    }
})

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        try {
            const conditions = JSON.parse(conditionsJson.value);
            const filterData: Omit<Filter, 'id'> = {
                name: name.value,
                description: description.value,
                conditions: conditions,
            };
            props.onSubmit(filterData);
        } catch (error) {
            alert("Invalid JSON in conditions: " + error);
        }
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

:deep(.v-field) {
    background-color: rgba(255, 255, 255, 0.8) !important;
}

:deep(.v-field__input) {
    font-family: 'Courier New', monospace;
}
</style>