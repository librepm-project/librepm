import { test, expect } from '../support/fixtures';

test.describe('Admin > Settings', () => {
  test('shows settings page via gear icon', async ({ authenticatedPage: page }) => {
    await page.getByRole('link', { name: 'Settings' }).click();
    await expect(page).toHaveURL('/admin/settings');
  });

  test('shows settings in admin sidebar', async ({ authenticatedPage: page }) => {
    await page.goto('/admin/settings');
    await expect(page.getByRole('list').getByRole('link', { name: 'Settings' })).toBeVisible();
  });
});
