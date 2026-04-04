import { test, expect } from '../support/fixtures';

test.describe('Admin > Priorities', () => {
  test('shows seeded priorities in list', async ({ adminPage: page }) => {
    await page.goto('/admin/priority');
    await page.waitForLoadState('networkidle');
    await expect(page.getByRole('cell', { name: 'High' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Normal' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Low' })).toBeVisible();
  });

  test('navigates to priority create page', async ({ adminPage: page }) => {
    await page.goto('/admin/priority');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/priority/create');
  });

  test('creates a new priority', async ({ adminPage: page }) => {
    const priorityName = `E2E Priority ${Date.now()}`;
    await page.goto('/admin/priority/create');

    await page.getByLabel('Name').fill(priorityName);

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/priority');
    await page.waitForLoadState('networkidle');
    await expect(page.getByRole('cell', { name: priorityName })).toBeVisible();
  });

  test('opens existing priority detail', async ({ adminPage: page }) => {
    await page.goto('/admin/priority');
    await page.waitForLoadState('networkidle');
    await page.getByRole('cell', { name: 'High' }).click();
    await expect(page).toHaveURL(/\/admin\/priority\/.+/);
    await expect(page.getByLabel('Name')).toHaveValue('High');
  });
});
