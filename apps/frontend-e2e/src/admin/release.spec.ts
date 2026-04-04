import { test, expect, showAllRows } from '../support/fixtures';

test.describe('Admin > Releases', () => {
  test('shows releases index page', async ({ adminPage: page }) => {
    await page.goto('/admin/release');
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL('/admin/release');
    await expect(page.locator('table, [class*="empty"], [class*="no-data"]').first()).toBeVisible();
  });

  test('navigates to release create page', async ({ adminPage: page }) => {
    await page.goto('/admin/release');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/release/create');
  });

  test('creates a new planned release', async ({ adminPage: page }) => {
    const releaseName = `E2E Release ${Date.now()}`;
    await page.goto('/admin/release/create');

    await page.getByLabel('Release Name').fill(releaseName);

    await page.getByLabel('Status').click({ force: true });
    await page.getByRole('option', { name: 'Planned' }).click();

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/release');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: releaseName })).toBeVisible();
  });

  test('creates a released release', async ({ adminPage: page }) => {
    const releaseName = `E2E Released ${Date.now()}`;
    await page.goto('/admin/release/create');

    await page.getByLabel('Release Name').fill(releaseName);

    await page.getByLabel('Status').click({ force: true });
    await page.getByRole('option', { name: 'Released' }).click();

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/release');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: releaseName })).toBeVisible();
  });

  test('opens release detail', async ({ adminPage: page }) => {
    const releaseName = `E2E Detail ${Date.now()}`;
    await page.goto('/admin/release/create');
    await page.getByLabel('Release Name').fill(releaseName);
    await page.getByLabel('Status').click({ force: true });
    await page.getByRole('option', { name: 'Planned' }).click();
    await page.getByRole('button', { name: 'Create' }).click();
    await expect(page).toHaveURL('/admin/release');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);

    await page.getByRole('cell', { name: releaseName }).click();
    await expect(page).toHaveURL(/\/admin\/release\/.+/);
    await expect(page.getByLabel('Release Name')).toHaveValue(releaseName);
  });
});
