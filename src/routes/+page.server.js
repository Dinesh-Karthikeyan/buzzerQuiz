import { redirect } from '@sveltejs/kit'

export function load({cookies}) {
    const jwt = cookies.get("jwt")
    if (jwt != "") new redirect(300, "/game")
}