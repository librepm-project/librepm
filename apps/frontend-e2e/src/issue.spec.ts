import { test, expect } from './support/fixtures';

test.describe('Issues', () => {
  test('shows issue list', async ({ authenticatedPage: page }) => {
    await page.goto('/issue');
    await expect(page.getByRole('heading', { name: /issues/i })).toBeVisible();
  });

  test('navigates to issue create page', async ({ authenticatedPage: page }) => {
    await page.goto('/issue');
    await page.getByRole('button', { name: /create issue/i }).click();
    await expect(page).toHaveURL('/issue/create');
  });
});
