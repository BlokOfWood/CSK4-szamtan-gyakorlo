import { test, expect } from '@playwright/test';

test('Leaderboard page rendered', async ({ page }) => {
  // Mock the API response
  await page.route('http://localhost:7995/v1/results', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        results: [
          { id: 1, name: 'Alice', score: 100 },
          { id: 2, name: 'Bob', score: 90 },
        ],
      }),
    });
  });

  // Go to the page containing the Svelte component
  await page.goto('http://localhost:5174/leaderboard'); // Adjust the URL to your app's URL

  // Check if the leaderboard is populated correctly
  await expect(page.locator('text=Leaderboard')).toBeVisible();
  await expect(page.locator('text=Alice')).toBeVisible();
  await expect(page.locator('text=100')).toBeVisible();
  await expect(page.locator('text=Bob')).toBeVisible();
  await expect(page.locator('text=90')).toBeVisible();
});

test('Leaderboard displays fetched results', async ({ page }) => {
  // Mock the API response
  await page.route('http://localhost:7995/v1/results', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        results: [
          { id: 1, name: 'Alice', score: 100 },
          { id: 2, name: 'Bob', score: 90 },
        ],
      }),
    });
  });

  // Go to the page containing the Svelte component
  await page.goto('http://localhost:5174/leaderboard'); // Adjust the URL to your app's URL

  // Check if the leaderboard is populated correctly
  await expect(page.locator('text=Leaderboard')).toBeVisible();
  await expect(page.locator('text=Alice')).toBeVisible();
  await expect(page.locator('text=100')).toBeVisible();
  await expect(page.locator('text=Bob')).toBeVisible();
  await expect(page.locator('text=90')).toBeVisible();
});

test('Clicking the "Play" button navigates to the homepage', async ({ page }) => {
  // Navigate to the leaderboard page
  await page.goto('http://localhost:5174/leaderboard');

  // Wait for the "Play" button to be visible
  await page.waitForSelector('a.bg-accent:text("Play")');

  // Click the "Play" button
  await page.click('a.bg-accent:text("Play")');

  // Wait for the homepage to load
  await page.waitForURL('http://localhost:5174/');

  // Assert that the URL has changed
  expect(page.url()).toBe('http://localhost:5174/');
});
