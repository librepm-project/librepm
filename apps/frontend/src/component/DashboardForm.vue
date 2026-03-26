<template>
    <v-form @submit.prevent="submit" ref="form">
        <v-text-field v-model="name" :rules="[requiredRule]" :label="$t('global.name')"></v-text-field>
        <v-text-field v-model="description" :rules="[]" :label="$t('global.description')"></v-text-field>
        <v-checkbox v-model="isPublic" :label="$t('global.public')"></v-checkbox>
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
import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps({
    onSubmit: Function,
    onDelete: Function as any,
    submitButtonText: String,
})

const name = ref("");
const description = ref("");
const isPublic = ref(false);
const form = ref(null);

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        props.onSubmit({ name: name.value, description: description.value, public: isPublic.value });
    }
};
</script>