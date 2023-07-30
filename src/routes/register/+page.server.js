import { redirect } from '@sveltejs/kit'

export const load = async({cookies}) => {
    if (cookies.get("jwt") != "") {
        throw new redirect(300, "/game")
    }
}