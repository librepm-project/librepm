import { test, expect } from './support/fixtures';

test.describe('Issue – detail page', () => {
  test('shows related issues panel', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-1');
    await page.waitForLoadState('networkidle');
    // WED-1 (Book Venue) blocks WED-2 (Choose Menu)
    await expect(page.getByText('Related Issues')).toBeVisible();
    await expect(page.getByText('Choose Menu')).toBeVisible();
  });

  test('shows comments tab', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await page.getByRole('tab', { name: 'Comments' }).click();
    // Check for the comment input textarea (always visible in the panel)
    await expect(page.getByPlaceholder('Add Comment')).toBeVisible();
  });

  test('adds a comment to an issue', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await page.getByRole('tab', { name: 'Comments' }).click();

    await page.getByPlaceholder('Add Comment').fill('E2E test comment');
    await page.getByRole('button', { name: 'Add Comment' }).click();

    await expect(page.getByText('E2E test comment').first()).toBeVisible();
  });

  test('shows audit log tab', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await page.getByRole('tab', { name: 'History' }).click();
    // Check the panel heading (always visible; empty-state text is unreliable due to test data pollution)
    await expect(page.locator('.v-window-item--active').getByText('History')).toBeVisible();
  });

  test('shows worklog tab', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await page.getByRole('tab', { name: 'Work Log' }).click();
    // Check the panel heading (always visible; empty-state text is unreliable due to test data pollution)
    await expect(page.locator('.v-window-item--active').getByText('Work Log')).toBeVisible();
  });

  test('adds a worklog entry', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await page.getByRole('tab', { name: 'Work Log' }).click();

    // Click the + icon button in the worklog panel header to open the add form
    // The mdi-plus inside the worklog window item (not the attachments panel above)
    const worklogPanel = page.locator('.v-window-item--active');
    await worklogPanel.locator('.mdi-plus').click();

    await page.getByLabel('Hours').fill('2');
    await page.getByRole('button', { name: 'Add Work Log' }).click();

    await expect(page.getByText('2h', { exact: true }).first()).toBeVisible();
  });

  test('shows attachments panel', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    // Attachments panel is above the tabs, check it's visible
    await expect(page.locator('.v-card-text').getByText('Attachments').first()).toBeVisible();
  });
});
