export const actions = {
    login: async({request}) => {
        const body = Object.fromEntries(await request.formData())
        const response = await fetch("http://127.0.0.1:8000/login", {
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
                throw redirect(303, "/game")
            }
    }
}