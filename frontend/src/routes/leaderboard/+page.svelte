<script lang="ts">
	import { onMount } from 'svelte';

	interface Result {
		id: number;
		name: string;
		score: number;
	}

	let results: Result[] = [];

	onMount(async () => {
		try {
			const response = await fetch('http://localhost:7995/v1/results',
			{method: 'GET'});
			const data = await response.json();
			results = data.results;
		} catch (error) {
			console.error('Error fetching results:', error);
		}
	});
</script>

<div
	class="absolute inset-0 bg-gradient-to-b from-backgroundFrom to-backgroundTo bg-fixed overflow-auto"
>
	<div class="p-4 text-center text-white text-3xl font-bold">Leaderboard</div>
	<div class="absolute right-1 top-1 p-4 flex items-center justify-end z-10">
		<a href="/" class="bg-accent text-white rounded-md px-4 py-2">Play</a>
	</div>
	<div class="flex justify-center items-center">
		<div class="grid grid-cols-2 w-70">
			<div class="border-0 p-2 text-white text-lg font-bold">Name:</div>
			<div class="border-0 p-2 text-white text-lg font-bold text-right">Points:</div>
			{#each results as result}
				<div class="border-0 p-2 text-white">{result.name}</div>
				<div class="border-0 p-2 text-white text-right">{result.score}</div>
			{/each}
		</div>
	</div>
</div>
