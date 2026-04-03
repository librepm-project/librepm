import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router';
import { usePermissionStore } from '@/store/permission.store';
import type { Permission } from '@/lib/permissions';

export const permissionMiddleware = async (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext
) => {
  const requiredPermission = to.meta?.permission as Permission | undefined;
  if (!requiredPermission) {
    return next();
  }
  const permissionStore = usePermissionStore();
  if (!permissionStore.loaded) {
    try {
      await permissionStore.fetch();
    } catch {
      return next({ name: 'forbidden' });
    }
  }
  if (!permissionStore.can(requiredPermission)) {
    return next({ name: 'forbidden' });
  }
  next();
};
