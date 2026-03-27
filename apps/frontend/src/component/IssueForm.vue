<template>
    <v-container>
        <v-form @submit.prevent="submit">
            <v-row>
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
                <v-col cols="12" md="4">
                    <v-select v-model="form.priorityId" :label="t('issue.priority')" :items="priorityStore.index"
                        item-title="name" item-value="id" clearable></v-select>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="form.assignedUserId" :label="t('issue.assignee')" :items="userStore.index"
                        :item-title="(u: any) => `${u.firstName} ${u.lastName}`" item-value="id" clearable></v-select>
                </v-col>

                <v-col cols="12">
                    <v-text-field v-model="form.summary" :label="t('issue.summary')" required></v-text-field>
                </v-col>

                <v-col cols="12">
                    <RichTextField v-model="form.description" :label="t('issue.description')" />
                </v-col>
            </v-row>
            <v-btn type="submit" color="primary">{{ t(props.submitButtonText) }}</v-btn>
        </v-form>
    </v-container>
</template>

<script lang="ts" setup>
import { Issue } from '@/lib/interfaces/issue.interface';
import { useProjectStore } from '@/store/project.store';
import { useUserStore } from '@/store/user.store';
import { usePriorityStore } from '@/store/priority.store';
import RichTextField from '@/component/RichTextField.vue';
import { ref, onMounted, watch } from 'vue';
import { useI18n } from 'vue-i18n';


const { t } = useI18n();

const props = defineProps<{
    submitButtonText: string;
    onSubmit: (item: Partial<Issue>) => void;
    initialData?: Partial<Issue>;
}>();

const projectStore = useProjectStore();
const userStore = useUserStore();
const priorityStore = usePriorityStore();

const initForm = () => ({
    description: '',
    ...props.initialData
})

onMounted(async () => {
    await Promise.all([projectStore.getProjects(), userStore.getAllItems(), priorityStore.getPriorities()]);
    if (props.initialData?.projectId) {
        await projectStore.getIssueProperty(props.initialData.projectId);
        trackers.value = projectStore.currentIssueProperty?.trackers || [];
        
        if (props.initialData?.trackerId) {
            const tracker = trackers.value.find(t => t.id === props.initialData!.trackerId);
            statuses.value = tracker?.statuses || [];
        }
    }
});

const form = ref<Partial<Issue>>(initForm());

const trackers = ref<any[]>([]);
const statuses = ref<any[]>([]);

watch(() => form.value.projectId, async (newProjectId, oldProjectId) => {
    if (newProjectId === oldProjectId || typeof newProjectId !== 'string') {
        return
    }
    await projectStore.getIssueProperty(newProjectId);
    trackers.value = projectStore.currentIssueProperty?.trackers || [];
    
    // Only reset if it's a real change by the user, not initial load
    if (oldProjectId !== undefined) {
        delete form.value.trackerId;
        delete form.value.statusId;
        statuses.value = [];
    }
});

watch(() => form.value.trackerId, async (newTrackerId, oldTrackerId) => {
    if (newTrackerId === oldTrackerId || typeof newTrackerId !== 'string' || !projectStore.currentIssueProperty) {
        return;
    }
    const tracker = projectStore.currentIssueProperty.trackers.find(t => t.id === newTrackerId);
    statuses.value = tracker?.statuses || [];
    
    // Only reset if it's a real change by the user
    if (oldTrackerId !== undefined) {
        delete form.value.statusId;
    }
});

const submit = async () => {
    props.onSubmit(form.value);
    if (!props.initialData) {
        form.value = initForm();
    }
};

</script>
