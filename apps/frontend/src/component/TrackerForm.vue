<template>
    <v-form @submit.prevent="submit" ref="form">
        <v-text-field v-model="name" :rules="[requiredRule]" :label="$t('global.name')"></v-text-field>
        <v-btn class="mt-2" type="submit" block>{{ $t(props.submitButtonText) }}</v-btn>
    </v-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
})

const name = ref("");
const description = ref("");
const form = ref(null);

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        props.onSubmit({ name, codeName });
    }
};
</script>