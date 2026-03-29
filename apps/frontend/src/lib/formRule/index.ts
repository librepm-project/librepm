import { i18n } from '@/i18n';

export const minimumLengthRule = (len: number) => (value: unknown) => {
  const str = value != null ? String(value) : '';
  return (
    str.length >= len ||
    i18n.global.t('rule.minimumLength', { len: len })
  );
};

export const requiredRule = (value: unknown) => {
  return !!value || i18n.global.t('rule.required');
};

export const emailRule = (value: unknown) => {
  return /.+@.+\..+/.test(value != null ? String(value) : '') || i18n.global.t('rule.email');
};
