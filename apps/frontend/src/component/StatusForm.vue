<template>
    <v-form @submit.prevent="submit" ref="form" class="form-container">
        <v-card elevation="0" class="rounded-xl pa-6 mb-6">
            <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
                {{ props.submitButtonText === 'global.create' ? 'New Status' : 'Edit Status' }}
            </v-card-title>

            <v-row>
                <v-col cols="12" md="8">
                    <v-text-field
                        v-model="name"
                        :rules="[requiredRule]"
                        :label="$t('global.name')"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-row>
                <v-col cols="12">
                    <label class="text-subtitle-2 font-weight-bold mb-3">
                        {{ $t('status.color') || 'Color' }}
                    </label>
                    <v-color-picker
                        v-model="color"
                        mode="hex"
                        dot-size="30"
                        width="300"
                        hide-canvas
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

const router = useRouter();

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
    status: Object as any,
})

const name = ref("");
const color = ref("#3F51B5");
const form = ref(null);

onMounted(() => {
    if (props.status) {
        name.value = props.status.name || "";
        color.value = props.status.color || "#3F51B5";
    }
})

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        const statusData = {
            name: name.value,
            color: color.value,
        };
        props.onSubmit(statusData);
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
