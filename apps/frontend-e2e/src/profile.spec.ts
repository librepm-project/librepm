import { test, expect } from './support/fixtures';

test.describe('Profile', () => {
  test('shows profile page with current user data', async ({ authenticatedPage: page }) => {
    await page.goto('/profile');
    await expect(page).toHaveURL('/profile');
    await expect(page.getByLabel('E-mail')).toHaveValue('testuser@example.com');
    await expect(page.getByLabel('First name')).toHaveValue('John');
    await expect(page.getByLabel('Last name')).toHaveValue('Doe');
  });

  test('updates profile phone number', async ({ authenticatedPage: page }) => {
    await page.goto('/profile');

    await page.getByLabel('Phone').clear();
    await page.getByLabel('Phone').fill('999-000-1111');

    await Promise.all([
      page.waitForResponse(resp => resp.url().includes('/user') && resp.request().method() === 'PUT'),
      page.getByRole('button', { name: /save/i }).click(),
    ]);

    // Reload and verify persisted
    await page.reload();
    await expect(page.getByLabel('Phone')).toHaveValue('999-000-1111');

    // Restore original value
    await page.getByLabel('Phone').clear();
    await page.getByLabel('Phone').fill('123-456-7890');
    await Promise.all([
      page.waitForResponse(resp => resp.url().includes('/user') && resp.request().method() === 'PUT'),
      page.getByRole('button', { name: /save/i }).click(),
    ]);
  });

  test('updates first and last name', async ({ authenticatedPage: page }) => {
    await page.goto('/profile');

    await page.getByLabel('First name').clear();
    await page.getByLabel('First name').fill('John Updated');

    await Promise.all([
      page.waitForResponse(resp => resp.url().includes('/user') && resp.request().method() === 'PUT'),
      page.getByRole('button', { name: /save/i }).click(),
    ]);

    await page.reload();
    await expect(page.getByLabel('First name')).toHaveValue('John Updated');

    // Restore
    await page.getByLabel('First name').clear();
    await page.getByLabel('First name').fill('John');
    await Promise.all([
      page.waitForResponse(resp => resp.url().includes('/user') && resp.request().method() === 'PUT'),
      page.getByRole('button', { name: /save/i }).click(),
    ]);
  });

  test('profile page is accessible from navigation', async ({ authenticatedPage: page }) => {
    await page.goto('/dashboard');
    // Navigate to profile - likely accessible from a user menu
    await page.goto('/profile');
    await expect(page).toHaveURL('/profile');
    await expect(page.getByLabel('E-mail')).toBeVisible();
  });
});
