import { test as base, expect, Page } from '@playwright/test';

type Fixtures = {
  authenticatedPage: Page;
  adminPage: Page;
};

export const test = base.extend<Fixtures>({
  authenticatedPage: async ({ page }, use) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('testuser@example.com');
    await page.getByLabel('Password').fill('password123');
    await page.getByRole('button', { name: 'Login' }).click();
    await expect(page).toHaveURL(/\/dashboard/, { timeout: 30000 });
    await page.waitForLoadState('networkidle');
    await use(page);
  },

  adminPage: async ({ page }, use) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('testuser@example.com');
    await page.getByLabel('Password').fill('password123');
    await page.getByRole('button', { name: 'Login' }).click();
    await expect(page).toHaveURL(/\/dashboard/, { timeout: 30000 });
    await page.waitForLoadState('networkidle');
    await use(page);
  },
});

/**
 * Switches the v-data-table footer "Items per page" selector to "All"
 * so that all rows are visible regardless of accumulated test data.
 */
export async function showAllRows(page: Page): Promise<void> {
  const footer = page.locator('.v-data-table-footer');
  if (await footer.count() === 0) return;
  // Click the dropdown arrow — the inner input is covered by v-select__selection-text
  await footer.locator('.v-field__append-inner').click();
  await page.getByRole('option', { name: 'All' }).click();
}

export { expect } from '@playwright/test';
