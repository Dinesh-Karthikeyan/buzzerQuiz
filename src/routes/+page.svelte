<main>
    <a href="/login">
        <button>Log IN</button>
    </a>
    <a href="/register">
        <button>Sign UP</button>
    </a>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 4rem;
        height: 100vh;
        width: 100vw;
    }
    button {
        padding: 1rem;
        font-size: 1.5rem;
        border: 2px solid black;
        background-color: white;
        border-radius: 1rem;
        transition-duration: 250ms;
    }
    button:hover {
        background-color: black;
        color: white;
    }
</style>