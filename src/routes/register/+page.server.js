import { redirect } from '@sveltejs/kit'

export const actions = {
    register: async ({ request, cookies, locals}) => {
        const body = Object.fromEntries(await request.formData())
            const response = await fetch("http://127.0.0.1:8000/register", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    team_name: body.team_name,
                    password: body.password
                })
            })
            // console.log(response)
            locals.team_name = body.team_name

            if (cookies.get("jwt") != "") {
                console.log("heheheh")
                throw redirect(303, "/game")
            }
    }
}