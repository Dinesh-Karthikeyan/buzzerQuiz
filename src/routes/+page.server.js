import { redirect } from '@sveltejs/kit'

export async function load({ cookies }) {
    const jwt = cookies.get('jwt')
    if (jwt != undefined && jwt != "") {
        
        new redirect(303, '/game')
    }
}