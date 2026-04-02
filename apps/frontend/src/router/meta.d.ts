import type { Permission } from '@/lib/permissions';

declare module 'vue-router' {
  interface RouteMeta {
    permission?: Permission;
    title?: string;
    hideLayout?: boolean;
  }
}
