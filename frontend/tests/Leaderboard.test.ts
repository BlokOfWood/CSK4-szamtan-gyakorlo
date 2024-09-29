import { render, screen, waitFor } from '@testing-library/svelte';
import '@testing-library/jest-dom';
import '@testing-library/jest-dom'; 
import Leaderboard from '../src/routes/leaderboard/+page.svelte';

global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve({
        results: [
          { id: 1, name: 'Alice', score: 100 },
          { id: 2, name: 'Bob', score: 90 },
        ],
      }),
    })
  ) as jest.Mock;
  
  test('fetches and displays results', async () => {
    render(Leaderboard);
  
    await waitFor(() => {
      expect(screen.getByText('Alice')).toBeInTheDocument();
      expect(screen.getByText('100')).toBeInTheDocument();
      expect(screen.getByText('Bob')).toBeInTheDocument();
      expect(screen.getByText('90')).toBeInTheDocument();
    });
  });