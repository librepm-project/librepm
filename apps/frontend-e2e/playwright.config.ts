import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
  testDir: './src',
  // Serial execution locally — tests share a single DB (WED-0, WED-1, etc.)
  // and parallel writes/reads cause non-deterministic failures.
  fullyParallel: false,
  workers: process.env.CI ? 2 : 1,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 1,
  reporter: process.env.CI ? 'html' : 'list',
  timeout: 60000,
  use: {
    baseURL: process.env.BASE_URL ?? 'http://localhost:4200',
    actionTimeout: 15000,
    navigationTimeout: 20000,
    trace: 'on-first-retry',
    screenshot: 'only-on-failure',
  },
  projects: [
    { name: 'chromium', use: { ...devices['Desktop Chrome'] } },
  ],
  outputDir: 'playwright-report',
});
