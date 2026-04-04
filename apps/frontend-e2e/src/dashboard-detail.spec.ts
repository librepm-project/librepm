import { test, expect } from './support/fixtures';

const navigateToWeddingDashboard = async (page: any) => {
  await page.goto('/dashboard');
  await page.waitForURL(/\/dashboard\/.*/);
  await page.waitForLoadState('networkidle');
  await page.locator('.dashboard-switcher-btn').click();
  await page.getByRole('link', { name: 'Wedding Dashboard' }).click();
  await page.waitForURL(/\/dashboard\/.+/);
  await page.waitForLoadState('networkidle');
};

test.describe('Dashboard – detail', () => {
  test('shows seeded dashboard with widgets', async ({ authenticatedPage: page }) => {
    await navigateToWeddingDashboard(page);
    await expect(page.getByText('Wedding Tasks')).toBeVisible();
    await expect(page.getByText('All Open Issues')).toBeVisible();
  });

  test('shows dashboard switcher with all seeded dashboards', async ({ authenticatedPage: page }) => {
    await page.goto('/dashboard');
    await page.waitForURL(/\/dashboard\/.*/);
    await page.waitForLoadState('networkidle');

    await page.locator('.dashboard-switcher-btn').click();

    await expect(page.getByRole('link', { name: 'Wedding Dashboard' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Startup Dashboard' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'London Trip Dashboard' })).toBeVisible();
  });

  test('switches between dashboards', async ({ authenticatedPage: page }) => {
    await page.goto('/dashboard');
    await page.waitForURL(/\/dashboard\/.*/);
    await page.waitForLoadState('networkidle');

    await page.locator('.dashboard-switcher-btn').click();
    await page.getByRole('link', { name: 'Startup Dashboard' }).click();
    await page.waitForURL(/\/dashboard\/.+/);

    await expect(page.locator('.dashboard-switcher-btn')).toContainText('Startup Dashboard');
  });

  test('adds a widget to a dashboard', async ({ authenticatedPage: page }) => {
    await page.goto('/dashboard');
    await page.waitForURL(/\/dashboard\/.*/);
    await page.waitForLoadState('networkidle');

    await page.locator('.dashboard-switcher-btn').click();
    await page.getByRole('link', { name: 'Retro Gaming Party Dashboard' }).click();
    await page.waitForURL(/\/dashboard\/.+/);
    await page.waitForLoadState('networkidle');

    await page.getByRole('button', { name: 'Add Widget' }).click();

    const dialog = page.getByRole('dialog');
    await expect(dialog).toBeVisible();

    await dialog.getByLabel('Title').fill('E2E Widget');

    // Widget Type is pre-selected as 'filter', click the second v-select (Select Filter)
    // Using the v-select container directly avoids the data-no-activator overlay interception
    await dialog.locator('.v-select').nth(1).click();
    await page.getByRole('option', { name: 'Open Issues' }).click();

    await dialog.getByRole('button', { name: 'Create' }).click();

    await expect(page.getByText('E2E Widget')).toBeVisible();
  });

  test('creates a new dashboard', async ({ authenticatedPage: page }) => {
    await page.goto('/dashboard/create');

    await page.getByLabel('Name').fill('E2E Dashboard');
    await page.getByLabel('Description').fill('Created by e2e test');

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL(/\/dashboard\/.+/);
  });
});
