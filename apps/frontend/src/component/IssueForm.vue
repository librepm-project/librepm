<template>
    <v-container>
        <v-form @submit.prevent="createIssue">
            <v-row>
                <v-col cols="12">
                    <v-text-field v-model="form.summary" :label="t('issue.summary')" required></v-text-field>
                </v-col>

                <v-col cols="12">
                    <v-textarea v-model="form.description" :label="t('issue.description')"></v-textarea>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="form.projectId" :label="t('issue.project')" :items="projectStore.index"
                        item-title="name" item-value="id" required></v-select>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="form.trackerId" :label="t('issue.tracker')" :items="trackers" item-title="name"
                        item-value="id" required></v-select>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="form.statusId" :label="t('issue.status')" :items="statuses" item-title="name"
                        item-value="id" required></v-select>
                </v-col>
            </v-row>

            <v-btn type="submit" color="primary">{{ t(props.submitButtonText) }}</v-btn>
        </v-form>
    </v-container>
</template>

<script lang="ts" setup>
import { Issue } from '@/lib/interfaces/issue.interface';
import { useProjectStore } from '@/store/project.store';
import { ref, onMounted, watch } from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const props = defineProps<{
    submitButtonText: string;
    onSubmit: (item: Issue) => void
}>();

const projectStore = useProjectStore();

onMounted(async () => {
    await projectStore.getProjects();
});

const form = ref<Issue>({
    summary: '',
    description: '',
    trackerId: '',
    statusId: '',
    projectId: '',
});

const trackers = ref<any[]>([]);
const statuses = ref<any[]>([]);

watch(() => form.value.projectId, async (newProjectId, oldProjectId) => {
    if (newProjectId === oldProjectId || typeof newProjectId !== 'string') {
        return
    }
    await projectStore.getIssueProperty(newProjectId);
    trackers.value = projectStore.currentIssueProperty?.trackers || [];
});

watch(() => form.value.trackerId, async (newTrackerId, oldTrackerId) => {
    if (newTrackerId === oldTrackerId || typeof newTrackerId !== 'string' || !projectStore.currentIssueProperty) {
        return;
    }
    const tracker = projectStore.currentIssueProperty.trackers.find(t => t.id === newTrackerId);
    statuses.value = tracker?.statuses || [];
});

const createIssue = async () => {
    props.onSubmit(form.value as Issue);

    form.value = {
        trackerId: '',
        summary: '',
        description: '',
        statusId: '',
        projectId: '',
    };
};

</script>