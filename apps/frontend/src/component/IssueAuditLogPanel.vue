<template>
  <div class="mb-6">
    <!-- Fejléc -->
    <div class="d-flex align-center mb-3">
      <p class="text-subtitle2 font-weight-medium mb-0 flex-grow-1">
        <v-icon
          x-small
          class="mr-1"
        >
          mdi-history
        </v-icon>
        {{ t('audit_log.history') }}
      </p>
    </div>

    <!-- Lista -->
    <div
      v-for="log in auditLogs"
      :key="log.id"
      class="d-flex align-start gap-2 pa-2 rounded hover-bg mb-1"
    >
      <v-icon
        size="small"
        class="text-disabled mt-1"
        style="flex-shrink: 0"
      >
        {{ eventIcon(log) }}
      </v-icon>
      <div class="flex-grow-1">
        <span class="text-body2">
          <span class="font-weight-medium">{{ userName(log) }}</span>
          {{ ' ' }}{{ eventDescription(log) }}
        </span>
      </div>
      <span
        class="text-caption text-disabled"
        style="white-space: nowrap"
      >{{ formatDate(log.createdAt) }}</span>
    </div>

    <p
      v-if="auditLogs.length === 0"
      class="text-caption text-disabled mb-0"
    >
      {{ t('audit_log.no_history') }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { IssueAuditLog } from '@/lib/interfaces/issue-audit-log.interface';

interface Props {
  issueId: string;
  auditLogs: IssueAuditLog[];
}

defineProps<Props>();
const { t } = useI18n();

function userName(log: IssueAuditLog): string {
  if (!log.user) return 'Unknown';
  return `${log.user.firstName} ${log.user.lastName}`.trim() || log.user.email;
}

function parseMeta(log: IssueAuditLog): Record<string, string> {
  if (!log.meta) return {};
  try {
    return JSON.parse(log.meta);
  } catch {
    return {};
  }
}

function formatSeconds(s: string): string {
  const n = parseInt(s, 10);
  if (isNaN(n)) return s;
  const h = Math.floor(n / 3600);
  const m = Math.floor((n % 3600) / 60);
  if (h > 0 && m > 0) return `${h}h ${m}m`;
  if (h > 0) return `${h}h`;
  if (m > 0) return `${m}m`;
  return `${n}s`;
}

function formatDate(iso: string): string {
  return new Date(iso).toLocaleString();
}

function eventIcon(log: IssueAuditLog): string {
  switch (log.eventType) {
    case 'field_changed': return 'mdi-pencil-outline';
    case 'attachment_added':
    case 'attachment_removed': return 'mdi-paperclip';
    case 'worklog_added':
    case 'worklog_updated':
    case 'worklog_removed': return 'mdi-clock-outline';
    case 'related_issue_added':
    case 'related_issue_removed': return 'mdi-link-variant';
    default: return 'mdi-information-outline';
  }
}

function eventDescription(log: IssueAuditLog): string {
  const meta = parseMeta(log);
  switch (log.eventType) {
    case 'field_changed':
      return `changed ${log.field}: "${log.oldValue}" → "${log.newValue}"`;
    case 'attachment_added':
      return `added attachment: ${meta.filename ?? log.newValue}`;
    case 'attachment_removed':
      return `removed attachment: ${meta.filename ?? log.oldValue}`;
    case 'worklog_added':
      return `logged ${formatSeconds(log.newValue)}${meta.comment ? ` – ${meta.comment}` : ''}`;
    case 'worklog_updated':
      return `updated worklog: ${formatSeconds(log.oldValue)} → ${formatSeconds(log.newValue)}`;
    case 'worklog_removed':
      return `removed worklog (${formatSeconds(log.oldValue)})`;
    case 'related_issue_added':
      return `added ${meta.relation_type ?? ''} relation: ${meta.issue_key ?? ''}`;
    case 'related_issue_removed':
      return `removed ${meta.relation_type ?? ''} relation: ${meta.issue_key ?? ''}`;
    default:
      return log.eventType;
  }
}
</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

.hover-bg:hover {
  background-color: rgba(63, 81, 181, 0.05);
}
</style>
