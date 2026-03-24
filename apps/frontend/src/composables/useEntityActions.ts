import { Router } from 'vue-router';

export function useEntityActions(
  store: any,
  router: Router,
  basePath: string,
  deleteConfirmMessage: string = 'Are you sure?'
) {
  const handleEdit = (item: any) => {
    if (!item.id) {
      console.warn('Item has no ID');
      return;
    }
    router.push(`${basePath}/${item.id}`);
  };

  const handleDelete = async (item: any) => {
    if (!item.id) {
      console.warn('Item has no ID');
      return;
    }

    if (!confirm(deleteConfirmMessage)) {
      return;
    }

    try {
      await store.deleteItem(item.id);
    } catch (error) {
      console.error('Error deleting item:', error);
      alert('Failed to delete item');
    }
  };

  const handleViewDetails = (item: any) => {
    if (!item.id) {
      console.warn('Item has no ID');
      return;
    }
    router.push(`${basePath}/${item.id}`);
  };

  return {
    handleEdit,
    handleDelete,
    handleViewDetails,
  };
}
