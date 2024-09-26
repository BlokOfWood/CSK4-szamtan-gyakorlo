<script>
	/*Ez csak egy példa array egyelőre, hogy lássam nincs kötött sorszám,
	majd mikor megbeszéljük, pontosan hogyan kerülnek az adatok
	beltöltésre a táblázatba, megírom ide rendesen a kódot*/
	let numberOfRows = 100;
	let tableContent = Array.from({ length: numberOfRows }, (_, i) => i + 1);
</script>

<div class="absolute inset-0 bg-gradient-to-b from-backgroundFrom to-backgroundTo bg-fixed overflow-auto">
	<div class="p-4 text-center text-white text-3xl font-bold">Leaderboard</div>
		<div class="absolute right-1 top-1 p-4 flex items-center justify-end z-10">
			<a href="/" class="bg-accent text-white rounded-md px-4 py-2">Play</a>
		</div>
	<div class="flex justify-center items-center">
		<div class="grid grid-cols-2 w-70">
			<div class="border-0 p-2 text-white text-lg font-bold">Name:</div>
			<div class="border-0 p-2 text-white text-lg font-bold text-right">Points:</div>
			{#each tableContent as row}
				<div class="border-0 p-2 text-white">Column 1, Row {row}</div>
				<div class="border-0 p-2 text-white text-right">Column 2, Row {row}</div>
			{/each}
		</div>
	</div>
</div>