import { test, expect } from './support/fixtures';

test.describe('Issues – CRUD', () => {
  test('shows seeded issues in list', async ({ authenticatedPage: page }) => {
    await page.goto('/issue');
    await expect(page.getByRole('table')).toBeVisible();
    await expect(page.getByRole('cell', { name: 'WED-0' })).toBeVisible();
  });

  test('opens issue by key from URL', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await expect(page.getByText('Wedding Planning')).toBeVisible();
  });

  test('opens issue by clicking row in list', async ({ authenticatedPage: page }) => {
    await page.goto('/issue');
    await page.getByRole('cell', { name: 'WED-1' }).first().click();
    await expect(page).toHaveURL(/\/issue\/key\/WED-1/);
    await expect(page.getByText('Book Venue')).toBeVisible();
  });

  test('creates a new issue', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/create');
    await page.waitForLoadState('networkidle');

    await page.getByLabel('Project').click({ force: true });
    await page.getByRole('option', { name: 'Wedding Project' }).click();
    await page.waitForLoadState('networkidle'); // wait for tracker options to load async

    await page.getByLabel('Tracker').click({ force: true });
    await page.getByRole('option', { name: 'Planning' }).click();

    await page.getByLabel('Status').click({ force: true });
    await page.getByRole('option', { name: 'Planned' }).click();

    await page.getByLabel('Summary').fill('E2E Created Issue');

    await page.getByRole('button', { name: 'Create' }).click();

    await expect(page).toHaveURL(/\/issue\/key\//);
    await page.waitForLoadState('networkidle');
    await expect(page.getByText('E2E Created Issue')).toBeVisible();
  });

  test('issue show page displays project, status and tracker', async ({ authenticatedPage: page }) => {
    await page.goto('/issue/key/WED-0');
    await page.waitForLoadState('networkidle');
    await expect(page.getByText('Wedding Planning')).toBeVisible();
    await expect(page.getByText('Wedding Project')).toBeVisible();
    await expect(page.getByText('Confirmed')).toBeVisible();
    await expect(page.getByText('Planning', { exact: true })).toBeVisible();
  });
});
