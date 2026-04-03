import { test, expect } from './support/fixtures';

test.describe('Dashboard', () => {
  test('shows dashboard on root route', async ({ authenticatedPage: page }) => {
    await page.goto('/');
    await expect(page).toHaveURL('/');
  });
});
