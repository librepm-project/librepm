import { ref, computed, Ref } from 'vue';
import { useRouter } from 'vue-router';

export function useFormHandler<T extends { id?: string }>(
  modelValue: Ref<T | null>,
  onSubmitFn: (data: Omit<T, 'id'>) => Promise<void>,
  submitButtonText: string
) {
  const form = ref(null);
  const router = useRouter();
  const loading = ref(false);
  const error = ref<string | null>(null);

  const isEdit = computed(() => !!modelValue.value?.id);
  const formTitle = computed(() => {
    if (submitButtonText === 'global.create') {
      return 'New Item';
    }
    return 'Edit Item';
  });

  const submit = async (formData: Omit<T, 'id'>) => {
    try {
      const { valid } = await form.value?.validate();
      if (!valid) return;

      loading.value = true;
      error.value = null;
      
      await onSubmitFn(formData);
      
      router.back();
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred';
      console.error('Form submission error:', err);
    } finally {
      loading.value = false;
    }
  };

  const handleCancel = () => {
    router.back();
  };

  return {
    form,
    isEdit,
    formTitle,
    loading,
    error,
    submit,
    handleCancel,
  };
}
