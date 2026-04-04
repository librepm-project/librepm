import { test, expect, showAllRows } from '../support/fixtures';

test.describe('Admin > Statuses', () => {
  test('shows seeded statuses in list', async ({ adminPage: page }) => {
    await page.goto('/admin/status');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: 'Done' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Open' })).toBeVisible();
  });

  test('navigates to status create page', async ({ adminPage: page }) => {
    await page.goto('/admin/status');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/status/create');
  });

  test('creates a new status', async ({ adminPage: page }) => {
    const statusName = `E2E Status ${Date.now()}`;
    await page.goto('/admin/status/create');

    await page.getByLabel('Name').fill(statusName);

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/status');
    await page.waitForLoadState('networkidle');
  });

  test('opens existing status detail', async ({ adminPage: page }) => {
    await page.goto('/admin/status');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await page.getByRole('cell', { name: 'Done' }).click();
    await expect(page).toHaveURL(/\/admin\/status\/.+/);
    await expect(page.getByLabel('Name')).toHaveValue('Done');
  });
});
