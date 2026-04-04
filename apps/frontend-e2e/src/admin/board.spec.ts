import { test, expect, showAllRows } from '../support/fixtures';

test.describe('Admin > Boards', () => {
  test('shows seeded boards in list', async ({ adminPage: page }) => {
    await page.goto('/admin/board');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: 'Wedding Project Board' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Startup Launch Board' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'London Trip Board' })).toBeVisible();
  });

  test('navigates to board create page', async ({ adminPage: page }) => {
    await page.goto('/admin/board');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/board/create');
  });

  test('creates a new board', async ({ adminPage: page }) => {
    await page.goto('/admin/board/create');

    await page.getByLabel('Name').fill('E2E Test Board');
    await page.getByLabel('Description').fill('Created by e2e test');

    await page.getByRole('button', { name: 'Create' }).click();

    // After board create, redirects to board show or admin/board
    await expect(page).toHaveURL(/\/(admin\/board|board\/.+)/);
  });

  test('opens board edit page from admin list', async ({ adminPage: page }) => {
    await page.goto('/admin/board');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await page.getByRole('cell', { name: 'Wedding Project Board' }).click();
    await expect(page).toHaveURL(/\/admin\/board\/.+\/edit/);
  });

  test('board edit page shows existing columns', async ({ adminPage: page }) => {
    await page.goto('/admin/board');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await page.getByRole('cell', { name: 'Wedding Project Board' }).click();
    await expect(page).toHaveURL(/\/admin\/board\/.+\/edit/);

    // Wedding Project Board has Planned, Contacted, Confirmed, Delivered columns
    await expect(page.getByText('Planned')).toBeVisible();
    await expect(page.getByText('Contacted')).toBeVisible();
  });
});
