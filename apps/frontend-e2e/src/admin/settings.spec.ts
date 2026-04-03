import { test, expect } from '../support/fixtures';

test.describe('Admin > Settings', () => {
  test('shows settings page via gear icon', async ({ adminPage: page }) => {
    await page.getByRole('link', { name: 'Settings' }).click();
    await expect(page).toHaveURL('/admin/settings');
  });

  test('shows settings in admin sidebar', async ({ adminPage: page }) => {
    await page.goto('/admin/settings');
    await expect(page.getByRole('list').getByRole('link', { name: 'Settings' })).toBeVisible();
  });
});
