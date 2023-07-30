<script>
	let teamname = ""
	let password = ""
	let confPass = ""
	
	async function hi() {
		if (password!=confPass) alert("Both password and confirm password field must match")
		try {
			const response = await fetch("http://127.0.0.1:8000/register",  {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						team_name: teamname,
						password: password
					})
				})
		} catch (err) {
			
		}
	}
</script>

<div class="authContainer">
	<main>
		<h1>Register</h1>
		<form method="post">
			<label for="name">
				<input type="text" name="team_name" id="name" placeholder="Enter your team name"  bind:value={teamname}/>
			</label>
			<label for="password">
				<input type="password" name="password" id="password" placeholder="Enter your password" bind:value={password}/>
			</label>
			<label for="passwordConfirm">
				<input
					type="password"
					name="passwordConfirm"
					id="passwordConfirm"
					placeholder="Confirm your password"
					bind:value={confPass}
				/>
			</label>
			<button type="submit">Submit</button>
		</form>
	</main>
</div>

<style>
	.authContainer {
		margin-top: 50px;
		display: flex;
		flex-direction: column;
		flex: 1;
		align-items: center;
		height: 100vh;
	}
	main {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		align-items: center;
		/* border: 1px solid black; */
		border: none;
		padding: 10px;
		border-radius: 1em;
		/* width: 70vw; */
		padding-inline: 7rem;
		padding-block: 2rem;
		box-shadow: 0 10px 20px rgba(0, 0, 0, 0.3);
	}
	main h1 {
		font-size: 2.5rem;
	}
	form {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
		gap: 1rem;
		margin-bottom: 30px;
	}
	form button {
		width: 100%;
		font-size: 1.2rem;
		padding: 1rem;
		border-radius: 5px;
		border: none;
		box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);
		transition-duration: 250ms;
		color: white;
		background-color: black;
	}
	form button:hover {
		cursor: pointer;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
		color: black;
		background-color: white;
	}
	form label {
		padding: 5px;
		border: 1px solid black;
		border-radius: 5px;
	}
	form label input {
		width: 100%;
		border: none;
		font-size: 1.3rem;
	}

	form label input:focus {
		outline: none;
	}

</style>