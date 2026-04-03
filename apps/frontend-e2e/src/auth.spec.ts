import { test, expect } from '@playwright/test';

test.describe('Authentication', () => {
  test('shows login page', async ({ page }) => {
    await page.goto('/login');
    await expect(page.getByRole('button', { name: 'Login' })).toBeVisible();
  });

  test('logs in successfully with valid credentials', async ({ page }) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('testuser@example.com');
    await page.getByLabel('Password').fill('password123');
    await page.getByRole('button', { name: 'Login' }).click();
    await expect(page).toHaveURL('/');
  });

  test('shows error with invalid credentials', async ({ page }) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('wrong@example.com');
    await page.getByLabel('Password').fill('wrongpassword');
    await page.getByRole('button', { name: 'Login' }).click();
    await expect(page).toHaveURL('/login');
  });

  test('logs out successfully', async ({ page }) => {
    await page.goto('/login');
    await page.getByLabel('E-mail').fill('testuser@example.com');
    await page.getByLabel('Password').fill('password123');
    await page.getByRole('button', { name: 'Login' }).click();
    await page.waitForURL('/');
    await page.getByRole('button', { name: /logout/i }).click();
    await expect(page).toHaveURL('/login');
  });
});
