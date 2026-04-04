import { test, expect } from '../support/fixtures';

test.describe('Admin > Users', () => {
  test('shows seeded users in list', async ({ adminPage: page }) => {
    await page.goto('/admin/user');
    await page.waitForLoadState('networkidle');
    await expect(page.getByRole('cell', { name: 'testuser@example.com' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'testuser2@example.com' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'John' })).toBeVisible();
    await expect(page.getByRole('cell', { name: 'Jane' })).toBeVisible();
  });

  test('navigates to user create page', async ({ adminPage: page }) => {
    await page.goto('/admin/user');
    await page.waitForLoadState('networkidle');
    await page.getByRole('link', { name: 'Create', exact: true }).click();
    await expect(page).toHaveURL('/admin/user/create');
  });

  test('creates a new user', async ({ adminPage: page }) => {
    const ts = Date.now();
    await page.goto('/admin/user/create');

    await page.getByLabel('E-mail').fill(`e2e-${ts}@example.com`);
    await page.getByLabel('First name').fill('E2E');
    await page.getByLabel('Last name').fill('User');

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL('/admin/user');
    await page.waitForLoadState('networkidle');
    await expect(page.getByRole('cell', { name: `e2e-${ts}@example.com` })).toBeVisible();
  });

  test('opens existing user detail', async ({ adminPage: page }) => {
    await page.goto('/admin/user');
    await page.waitForLoadState('networkidle');
    await page.getByRole('cell', { name: 'testuser@example.com' }).click();
    await expect(page).toHaveURL(/\/admin\/user\/.+/);
    await expect(page.getByLabel('E-mail')).toHaveValue('testuser@example.com');
    await expect(page.getByLabel('First name')).toHaveValue('John');
    await expect(page.getByLabel('Last name')).toHaveValue('Doe');
  });
});
