<div class="absolute right-1 top-1 p-4 flex items-center justify-end z-10">
	<a href="/leaderboard" class="bg-accent text-white rounded-md px-4 py-2">Leaderboard</a>
</div>
<div
	class="bg-gradient-to-b from-fuchsia-400 to-rose-600 absolute inset-0 flex flex-col justify-center items-center"
>
	<div class="text-white text-6xl pb-2 z-0">10 * x = 40</div>
	<div class="flex items-center gap-2">
		<progress class="prog bg-background rounded-full h-4 w-[291px]" max="100" value="70"
			>70</progress
		>
		<div class="text-2xl text-background">07/10</div>
	</div>

	<div class="pt-6 flex gap-x-2">
		<input class="text-xs w-72 px-4 py-2 rounded-md flex-grow" placeholder="Answer" type="number" />
		<button class="text-xs text-white bg-accent rounded-md px-4 py-2">Next</button>
	</div>
</div>

<style>
	progress::-webkit-progress-bar {
		@apply bg-background;
	}

	progress::-webkit-progress-value {
		@apply bg-accent;
	}

	progress::-moz-progress-bar {
		@apply bg-accent;
	}

	/* Chrome, Safari, Edge, Opera */
	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	/* Firefox */
	input[type='number'] {
		-moz-appearance: textfield;
	}
</style>
