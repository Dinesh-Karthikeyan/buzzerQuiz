<div class="authContainer">
	<main>
		<h1>Login</h1>
		<form action="?/login" method="post">
			<label for="team_name">
				<input type="text" name="team_name" id="team_name" placeholder="Enter your team name" />
			</label>
			<label for="password">
				<input type="password" name="password" id="password" placeholder="Enter your password" />
			</label>
			<button type="submit">Submit</button>
		</form>
	</main>
</div>

<style>
    .authContainer {
        display: flex;
        flex-direction: column;
        margin-top: 50px;
        flex: 1;
        align-items: center;
        height: 100vh;
    }
    main {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        border:none;
        padding: 10px;
        border-radius: 1em;
        padding-inline: 7rem;
        padding-block: 2rem;
        box-shadow: 0 10px 20px rgba(0, 0, 0, 0.3);
    }
    main form {
        display:flex;
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