import { test, expect } from '../support/fixtures';

test.describe('Admin > Projects', () => {
  test('shows project list', async ({ authenticatedPage: page }) => {
    await page.goto('/admin/project');
    await expect(page).toHaveURL('/admin/project');
  });

  test('navigates to project create page', async ({ authenticatedPage: page }) => {
    await page.goto('/admin/project');
    await page.getByRole('button', { name: /create/i }).click();
    await expect(page).toHaveURL('/admin/project/create');
  });
});
