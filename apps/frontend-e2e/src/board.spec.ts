import { test, expect } from './support/fixtures';

test.describe('Board', () => {
  test('redirects /board to a board page', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await expect(page).toHaveURL(/\/board\/.+/);
  });

  test('shows board columns', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    // At least one column header must be visible
    await expect(page.locator('.v-card-title').first()).toBeVisible();
  });

  test('shows Wedding Project Board columns via switcher', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await page.locator('.board-switcher-btn').click();
    await page.getByRole('link', { name: 'Wedding Project Board' }).click();
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await expect(page.getByText('Planned')).toBeVisible();
    await expect(page.getByText('Contacted')).toBeVisible();
    await expect(page.getByText('Confirmed')).toBeVisible();
    await expect(page.getByText('Delivered')).toBeVisible();
  });

  test('board switcher lists all seeded boards', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await page.locator('.board-switcher-btn').click();

    await expect(page.getByRole('link', { name: 'Wedding Project Board' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Startup Launch Board' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'London Trip Board' })).toBeVisible();
  });

  test('switches to Startup Launch Board', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await page.locator('.board-switcher-btn').click();
    await page.getByRole('link', { name: 'Startup Launch Board' }).click();
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await expect(page.getByText('Not Started')).toBeVisible();
    await expect(page.getByText('Researching')).toBeVisible();
  });

  test('shows issue cards and navigates to issue from card', async ({ authenticatedPage: page }) => {
    await page.goto('/board');
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    await page.locator('.board-switcher-btn').click();
    await page.getByRole('link', { name: 'Wedding Project Board' }).click();
    await page.waitForURL(/\/board\/.+/);
    await page.waitForLoadState('networkidle');

    // WED-0 (Wedding Planning) is in Confirmed column
    await expect(page.locator('.issue-card').filter({ hasText: 'Wedding Planning' })).toBeVisible();

    await page.locator('.issue-card').filter({ hasText: 'Wedding Planning' }).click();
    await expect(page).toHaveURL(/\/issue\/key\/WED-0/);
  });
});
