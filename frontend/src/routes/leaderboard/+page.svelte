<script>
	import { goto } from '$app/navigation';
	
	/*Ez csak egy példa array egyelőre, hogy lássam nincs kötött sorszám,
	majd mikor megbeszéljük, pontosan hogyan kerülnek az adatok
	beltöltésre a táblázatba, megírom ide rendesen a kódot*/
	let numberOfRows = 100;
	let tableContent = Array.from({ length: numberOfRows }, (_, i) => i + 1);
</script>

<body class="bg-gradient-to-b from-fuchsia-400 to-rose-600 bg-fixed">
	<div class="relative p-4 flex items-center">
		<div class="absolute left-1/2 transform -translate-x-1/2 text-center text-white text-3xl font-bold">Leaderboard</div>
		<div class="ml-auto">
			<button class="bg-teal-400 text-white rounded-lg px-4 py-2" on:click={() => goto('/')}>Play</button>
		</div>
	</div>	
	<div class="flex justify-center items-center">
		<div class="grid grid-cols-2 w-70 mx-auto">
			<div class="border-0 p-2 text-white text-lg font-bold">Name:</div>
			<div class="border-0 p-2 text-white text-lg font-bold text-right">Points:</div>
			{#each tableContent as row}
				<div class="border-0 p-2 text-white">Column 1, Row {row}</div>
				<div class="border-0 p-2 text-white text-right">Column 2, Row {row}</div>
			{/each}
		</div>
	</div>
</body>
