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
    await use(page);
  },

  adminPage: async ({ page }, use) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('testuser@example.com');
    await page.getByLabel('Password').fill('password123');
    await page.getByRole('button', { name: 'Login' }).click();
    await expect(page).toHaveURL(/\/dashboard/, { timeout: 30000 });
    await use(page);
  },
});

export { expect } from '@playwright/test';
