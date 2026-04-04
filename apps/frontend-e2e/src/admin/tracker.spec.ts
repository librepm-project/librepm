import { test, expect } from '../support/fixtures';

test.describe('Admin > Trackers', () => {
  test('shows seeded trackers in list', async ({ adminPage: page }) => {
    await page.goto('/admin/tracker');
    await page.waitForLoadState('networkidle');
    // Use items confirmed to be on page 1 (Bug, Epic are seeded items on first page)
    await expect(page.getByRole('cell', { name: 'Bug', exact: true })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Epic', exact: true })).toBeVisible();
  });

  test('navigates to tracker create page', async ({ adminPage: page }) => {
    await page.goto('/admin/tracker');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/tracker/create');
  });

  test('creates a new tracker', async ({ adminPage: page }) => {
    const trackerName = `E2E Tracker ${Date.now()}`;
    await page.goto('/admin/tracker/create');

    await page.getByLabel('Name').fill(trackerName);

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/tracker');
  });

  test('opens existing tracker detail', async ({ adminPage: page }) => {
    await page.goto('/admin/tracker');
    await page.waitForLoadState('networkidle');
    // Bug is confirmed to be on page 1
    await page.getByRole('cell', { name: 'Bug', exact: true }).click();
    await expect(page).toHaveURL(/\/admin\/tracker\/.+/);
    await expect(page.getByLabel('Name')).toHaveValue('Bug');
  });
});
