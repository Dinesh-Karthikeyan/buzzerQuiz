<script>
	export let data;
	let { id, question, a, b, c, d, reward, display } = data;
	let show = display ? display : false;
	function hi() {
		show = false;
	}
</script>

<main>
	{#if show}
		<form action="?/evaluate" method="post">
			<div>
				<div>REWARD: {reward}</div>
				<p>{question}</p>
				<input type="text" value={id} name="id" hidden />
				<label for="option">
					<input type="radio" name="option" id="b" value="b" />{b}
					<input type="radio" name="option" id="a" value="a" />{a}
					<input type="radio" name="option" id="c" value="c" />{c}
					<input type="radio" name="option" id="d" value="d" />{d}
				</label>
				<button type="submit" on:click={hi}>Submit</button>
			</div>
		</form>
	{:else}
		<button>Buzzer</button>
	{/if}
</main>
