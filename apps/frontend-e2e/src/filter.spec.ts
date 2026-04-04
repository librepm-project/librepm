import { test, expect, showAllRows } from './support/fixtures';

test.describe('Filters', () => {
  test('shows seeded filters in list', async ({ authenticatedPage: page }) => {
    await page.goto('/filter');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: 'Open Issues' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Wedding Tasks' })).toBeVisible();
  });

  test('navigates to filter create page via Create link', async ({ authenticatedPage: page }) => {
    await page.goto('/filter');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/filter/create');
  });

  test('creates a new filter', async ({ authenticatedPage: page }) => {
    const filterName = `E2E Filter ${Date.now()}`;
    await page.goto('/filter/create');

    await page.getByLabel('Name').fill(filterName);
    await page.getByLabel('Description').fill('Created by e2e test');

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/filter');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: filterName })).toBeVisible();
  });

  test('opens filter view page by clicking row', async ({ authenticatedPage: page }) => {
    await page.goto('/filter');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await page.getByRole('cell', { name: 'Open Issues' }).click();
    await expect(page).toHaveURL(/\/filter\/[^/]+$/);
    await expect(page.getByRole('table')).toBeVisible();
  });

  test('opens filter edit page via pencil icon', async ({ authenticatedPage: page }) => {
    await page.goto('/filter');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    const row = page.getByRole('row').filter({ hasText: 'Open Issues' });
    await row.locator('.mdi-pencil').click();
    await expect(page).toHaveURL(/\/filter\/.+\/edit/);
    await expect(page.getByLabel('Name')).toHaveValue('Open Issues');
  });

  test('creates and deletes a filter', async ({ authenticatedPage: page }) => {
    const filterName = `E2E Delete Filter ${Date.now()}`;
    await page.goto('/filter/create');
    await page.getByLabel('Name').fill(filterName);
    await page.getByRole('button', { name: 'Create' }).click();
    await expect(page).toHaveURL('/filter');
    await page.waitForLoadState('networkidle');
    await showAllRows(page);
    await expect(page.getByRole('cell', { name: filterName })).toBeVisible();

    // Open edit page via pencil icon
    const row = page.getByRole('row').filter({ hasText: filterName });
    await row.locator('.mdi-pencil').click();
    await expect(page).toHaveURL(/\/filter\/.+\/edit/);

    // handleDelete uses native confirm() — accept it before clicking
    page.once('dialog', d => d.accept());
    await page.getByRole('button', { name: 'Delete' }).click();

    await expect(page).toHaveURL('/filter');
    await expect(page.getByRole('cell', { name: filterName })).not.toBeVisible();
  });
});
