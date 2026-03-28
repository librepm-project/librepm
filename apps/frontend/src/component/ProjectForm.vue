<template>
    <v-form @submit.prevent="submit" ref="form" class="form-container">
        <v-card elevation="0" class="rounded-xl pa-6 mb-6">
            <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
                {{ props.submitButtonText === 'global.create' ? 'New Project' : 'Edit Project' }}
            </v-card-title>

            <v-row>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="name"
                        :rules="[requiredRule]"
                        :label="$t('global.name')"
                        outlined
                        dense
                    />
                </v-col>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="codeName"
                        :rules="[requiredRule]"
                        :label="$t('project.codeName')"
                        hint="e.g., PROJ, DEV, DEMO"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-divider class="my-6" />

            <v-row>
                <v-col cols="12" md="6">
                    <v-select
                        v-model="defaultStatusId"
                        :items="statuses"
                        item-title="name"
                        item-value="id"
                        :label="$t('project.default_status')"
                        variant="outlined"
                        clearable
                        dense
                    />
                </v-col>
                <v-col cols="12" md="6">
                    <v-select
                        v-model="defaultTrackerId"
                        :items="trackers"
                        item-title="name"
                        item-value="id"
                        :label="$t('project.default_tracker')"
                        variant="outlined"
                        clearable
                        dense
                    />
                </v-col>
            </v-row>

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
                        {{ $t(props.submitButtonText) }}
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
                        {{ $t('global.delete') }}
                    </v-btn>
                    <v-spacer />
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
import { ref, watch, onMounted } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { useRouter } from 'vue-router';
import statusApi from '@/api/status.api';
import trackerApi from '@/api/tracker.api';
import { Status } from '@/lib/interfaces/status.interface';
import { Tracker } from '@/lib/interfaces/tracker.interface';

const router = useRouter();

const props = defineProps({
    onSubmit: { type: Function, required: true },
    onDelete: Function as any,
    submitButtonText: String,
    project: Object as any,
})

const name = ref("");
const codeName = ref("");
const defaultStatusId = ref<string | null>(null);
const defaultTrackerId = ref<string | null>(null);
const form = ref(null);

const statuses = ref<Status[]>([]);
const trackers = ref<Tracker[]>([]);

onMounted(async () => {
    const [s, t] = await Promise.all([
        statusApi.index(),
        trackerApi.index()
    ]);
    statuses.value = s;
    trackers.value = t;
});

watch(() => props.project, (project) => {
    if (project) {
        name.value = project.name || "";
        codeName.value = project.codeName || "";
        defaultStatusId.value = project.defaultStatusId || null;
        defaultTrackerId.value = project.defaultTrackerId || null;
    }
}, { immediate: true })

const submit = async () => {
    const { valid } = await (form.value as any).validate();
    if (valid) {
        props.onSubmit({ 
            name: name.value, 
            codeName: codeName.value,
            defaultStatusId: defaultStatusId.value,
            defaultTrackerId: defaultTrackerId.value
        });
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
</style>