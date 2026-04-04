import { test, expect, showAllRows } from '../support/fixtures';

test.describe('Admin > Projects – CRUD', () => {
  test('shows seeded projects in list', async ({ adminPage: page }) => {
    await page.goto('/admin/project');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: 'Wedding Project' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'WED', exact: true })).toBeVisible();
  });

  test('navigates to project create page via Create button', async ({ adminPage: page }) => {
    await page.goto('/admin/project');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/project/create');
  });

  test('creates a new project', async ({ adminPage: page }) => {
    const ts = Date.now();
    await page.goto('/admin/project/create');

    await page.getByLabel('Name', { exact: true }).fill(`E2E Project ${ts}`);
    await page.getByLabel('Code name').fill(`E${ts.toString().slice(-4)}`);

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/project');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: `E2E Project ${ts}` })).toBeVisible();
  });

  test('opens seeded project detail', async ({ adminPage: page }) => {
    await page.goto('/admin/project');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await page.getByRole('cell', { name: 'Wedding Project' }).click();
    await expect(page).toHaveURL(/\/admin\/project\/.+/);
    await expect(page.getByLabel('Name', { exact: true })).toHaveValue('Wedding Project');
    await expect(page.getByLabel('Code name')).toHaveValue('WED');
  });

  test('updates a project name', async ({ adminPage: page }) => {
    const ts = Date.now();
    await page.goto('/admin/project/create');
    await page.getByLabel('Name', { exact: true }).fill(`Update Me ${ts}`);
    await page.getByLabel('Code name').fill(`U${ts.toString().slice(-4)}`);
    await page.getByRole('button', { name: 'Create' }).click();
    await expect(page).toHaveURL('/admin/project');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);

    await page.getByRole('cell', { name: `Update Me ${ts}` }).click();
    await expect(page).toHaveURL(/\/admin\/project\/.+/);

    await page.getByLabel('Name', { exact: true }).clear();
    await page.getByLabel('Name', { exact: true }).fill(`Update Me ${ts} (Updated)`);
    await page.getByRole('button', { name: 'Update' }).click();

    await expect(page).toHaveURL('/admin/project');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: `Update Me ${ts} (Updated)` })).toBeVisible();
  });
});
