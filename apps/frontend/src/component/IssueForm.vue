<template>
    <v-container>
        <v-form @submit.prevent="createIssue">
            <v-row>
                <v-col cols="12">
                    <v-text-field v-model="newIssue.summary" :label="t('issue.summary')" required></v-text-field>
                </v-col>

                <v-col cols="12">
                    <v-textarea v-model="newIssue.description" :label="t('issue.description')"></v-textarea>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="newIssue.project" :label="t('issue.project')" :items="projects" item-title="name"
                        item-value="id" required></v-select>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="newIssue.status" :label="t('issue.status')" :items="statuses" item-title="name"
                        item-value="id" required></v-select>
                </v-col>

                <v-col cols="12" md="4">
                    <v-select v-model="newIssue.tracker" :label="t('issue.tracker')" :items="trackers" item-title="name"
                        item-value="id" required></v-select>
                </v-col>
            </v-row>

            <v-btn type="submit" color="primary">{{ t(props.submitButtonText) }}</v-btn>
        </v-form>
    </v-container>
</template>

<script lang="ts" setup>
import { Issue } from '@/lib/interfaces/issue.interface';
import { Project } from '@/lib/interfaces/project.interface';
import { Status } from '@/lib/interfaces/status.interface';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const props = defineProps<{
    submitButtonText: string;
    onSubmit: (item: Issue) => void
}>();

const projects = ref<Project[]>([
    { id: '1', name: 'Projekt A', codeName: 'PROJ_A' },
    { id: '2', name: 'Projekt B', codeName: 'PROJ_B' },
]);

const statuses = ref<Status[]>([
    { id: '10', name: 'Új', color: 'grey' },
    { id: '20', name: 'Folyamatban', color: 'blue' },
    { id: '30', name: 'Kész', color: 'green' },
]);

const trackers = ref<Tracker[]>([
    { id: '100', name: 'Feladat', color: 'light-blue' },
    { id: '200', name: 'Hiba', color: 'red' },
    { id: '300', name: 'Javaslat', color: 'amber' },
]);

const newIssue = ref<Omit<Issue, 'id' | 'key'>>({
    tracker: {} as Tracker,
    summary: '',
    description: '',
    status: {} as Status,
    project: {} as Project,
});

const createIssue = async () => {
    props.onSubmit(newIssue.value as Issue);

    newIssue.value = {
        tracker: {} as Tracker,
        summary: '',
        description: '',
        status: {} as Status,
        project: {} as Project,
    };
};
</script>