import { test, expect } from '../support/fixtures';

test.describe('Admin > Projects', () => {
  test('shows project list', async ({ adminPage: page }) => {
    await page.goto('/admin/project');
    await expect(page).toHaveURL('/admin/project');
  });

  test('navigates to project create page', async ({ adminPage: page }) => {
    await page.goto('/admin/project');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/project/create');
  });
});
