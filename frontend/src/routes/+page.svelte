<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	type Operator = '+' | '-' | '*' | '/';
	const operators: Operator[] = ['+', '-', '*', '/'];

	let currentPoints = 0;
	let currentTask: { problem: string; solution: number } = { problem: '', solution: 0 };
	let remainingTime: number = 10;

	let givenAnswer: string = '';
	let answerField: HTMLInputElement;

	let givenUsername = '';

	let gameOver = false;

	onMount(() => {
		currentTask = generateProblem(3);
		requestAnimationFrame(startCountdown);
	});

	let lastFrame: number;
	function startCountdown(now: DOMHighResTimeStamp) {
		const elapsed = now - (lastFrame ?? now);
		lastFrame = now;

		remainingTime = Math.max(remainingTime - elapsed * 0.001, 0);

		if (remainingTime !== 0) {
			requestAnimationFrame(startCountdown);
		} else {
			gameOver = true;
		}
	}

	function generateProblem(problemDigits: number): { problem: string; solution: number } {
		// Generate how many of the number of total digits will belong to the first constant in the problem.
		let firstNumDigitCount = genInt(1, problemDigits);

		let firstConstant = generateNumberWithDigits(firstNumDigitCount);
		let secondConstant = generateNumberWithDigits(problemDigits - firstNumDigitCount);

		/**
		 * The problems are equations with two operands and an answer.
		 * We randomize which slot in the equation is unknown.
		 * 0 and 1 would represent the first and second operands, 2 represents the result.
		 */
		let indexOfUnknown = genInt(0, 3);

		let equationOperator: Operator = operators[genInt(0, 4)];
		let equationMembers: number[] = [];

		switch (equationOperator) {
			case '+':
				equationMembers.push(firstConstant);
				equationMembers.push(secondConstant);
				equationMembers.push(firstConstant + secondConstant);
			case '-':
				equationMembers.push(firstConstant);
				equationMembers.push(secondConstant);
				equationMembers.push(firstConstant - secondConstant);
			case '*':
				equationMembers.push(firstConstant);
				equationMembers.push(secondConstant);
				equationMembers.push(firstConstant * secondConstant);
			case '/':
				equationMembers.push(firstConstant * secondConstant);
				equationMembers.push(firstConstant);
				equationMembers.push(secondConstant);
		}

		let finalEquationMembers = equationMembers.map((member, index) => {
			if (index == indexOfUnknown) return 'x';
			return member;
		});
		let problem = `${finalEquationMembers[0]} ${equationOperator} ${finalEquationMembers[1]} = ${finalEquationMembers[2]}`;
		let solution = equationMembers[indexOfUnknown];

		return {
			problem,
			solution
		};
	}

	/**
	 * Generates an integer with a number of digits equal to the number provided.
	 * @param digitCount The number of digits the number should have
	 * @return Generated integer number with given number of digits
	 */
	function generateNumberWithDigits(digitCount: number): number {
		// There is no such thing as half a digit or 0.471 digits.
		if (!Number.isInteger(digitCount)) throw new Error('Digit count is not an integer.');
		// There is no number with 0 digits or negative digits.
		if (digitCount <= 0) throw new Error('Digit count cannot be 0 or lower.');

		return genInt(1, digitCount * 10);
	}

	/**
	 * Generate an integer in the range given by the parameters.
	 * @param min The minimum possible number (inclusive)
	 * @param max The maximum possible number (exclusive)
	 * @returns Generated random integer number
	 */
	function genInt(min: number, max: number): number {
		if (min >= max) throw new Error('Minimum value is greater than or equal to the maximum value.');

		let result = Math.floor(Math.random() * (max - min)) + min;

		// If Math.random() produces a value equal to 1, it is possible to reach the max given, since the number won't be rounded down.
		if (result == max) return max - 1;

		return result;
	}

	function submitAnswer() {
		let givenAnswerInt = 0;

		try {
			givenAnswerInt = parseInt(givenAnswer);
		} catch {
			throw new Error('Incorrect answer given');
		}

		if (givenAnswerInt === currentTask.solution) {
			currentPoints++;
			remainingTime = 10;
		}

		givenAnswer = '';
		currentTask = generateProblem(3 + Math.floor(currentPoints / 10));
		answerField.focus();
	}

	async function submitResults() {
		await fetch('http://localhost:7995/v1/results', {
			method: 'POST',
			body: JSON.stringify({
				name: givenUsername,
				score: currentPoints
			}),
		});

		await goto('leaderboard');
	}
</script>

<div class="absolute right-1 top-1 p-4 flex items-center justify-end z-10">
	<a href="/leaderboard" class="bg-accent text-white rounded-md px-4 py-2">Leaderboard</a>
</div>

<div
	class="bg-gradient-to-b from-backgroundFrom to-backgroundTo absolute inset-0 flex flex-col gap-y-2 justify-center items-center"
>
	{#if !gameOver}
		<div class=" text-center pb-6 text-6xl text-white">
			{currentPoints}
		</div>
		<div class="text-white text-6xl pb-2 z-0">{currentTask.problem}</div>
		<div class="flex items-center gap-2">
			<progress class="prog bg-secondary rounded-full h-4 w-[291px]" max="10" value={remainingTime}
			></progress>
			<div class="text-2xl text-white">{remainingTime.toFixed(0)}/10</div>
		</div>

		<div class="pt-8 flex gap-x-2">
			<input
				bind:value={givenAnswer}
				bind:this={answerField}
				on:keydown={(e) => {
					if (e.key === 'Enter') {
						submitAnswer();
					}
				}}
				class="text-xs w-72 px-4 py-2 rounded-md flex-grow"
				placeholder="Answer"
				type="number"
			/>
			<button class="text-xs text-white bg-accent rounded-md px-4 py-2" on:click={submitAnswer}
				>Next</button
			>
		</div>
	{:else}
		<div class="flex-col flex gap-8 items-center justify-center">
			<div class="font-bold text-white text-4xl">Post your score on the leaderboard</div>
			<input
				bind:value={givenUsername}
				type="text"
				placeholder="Username"
				class="text-xs w-72 px-4 py-2 rounded-md"
			/>
			<div class="flex gap-4">
				<button class="text-xs bg-white rounded-md px-4 py-2">Cancel</button>
				<button
					class="text-xs text-white bg-accent rounded-md px-4 py-2"
					on:mousedown={submitResults}>Post</button
				>
			</div>
		</div>
	{/if}
</div>

<style>
	progress::-webkit-progress-bar {
		@apply bg-secondary;
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
