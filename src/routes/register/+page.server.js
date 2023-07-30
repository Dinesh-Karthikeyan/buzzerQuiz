import { error, redirect } from '@sveltejs/kit'

export const actions = {
    register: async({request, cookies}) => {
        const body = Object.fromEntries(await request.formData())
        try {
            const response = await fetch("http://127.0.0.1:8000/register",  {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    team_name: body.team_name,
                    password: body.password
                })
            })

            if (cookies.get("jwt") != "") {
                throw new redirect(300, "/game")
            }
        } catch (err) {
            throw new error(500, err)
        }
    }
}