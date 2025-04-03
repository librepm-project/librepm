import { i18n } from '@/i18n';

export const minimumLengthRule = (len: number) => (value: any) => {
  return (
    (value && value.length >= len) ||
    i18n.global.t('rule.minimumLength', { len: len })
  );
};

export const requiredRule = (value: any) => {
  return !!value || i18n.global.t('rule.required');
};

export const emailRule = (value: any) => {
  return /.+@.+\..+/.test(value) || i18n.global.t('rule.email');
};
