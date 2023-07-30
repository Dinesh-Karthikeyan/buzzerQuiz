export const actions = {
    buzzer: async ({ request }) => {
        const body = Object.fromEntries(await request.formData())
        const response = await fetch("http://127.0.0.1:8000/game", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({})
        })
        const responseJsonBody = await response.json()
        return {
            status: 200,
            body: responseJsonBody,
            headers: {
                'Content-Type': 'application/json',
            },
        };
    },
    evaluate: async({request}) => {
        const body = Object.fromEntries(await request.formData())
        const response = await fetch("http://127.0.0.1:8000/game", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id: body.id,
                option: body.option
            })
        })
    }
}

export async function load({ page }) {
    const { responseJsonBody } = page
    let id = responseJsonBody.id
    let question = responseJsonBody.question
    let a = responseJsonBody.A
    let b = responseJsonBody.B
    let c = responseJsonBody.C
    let d = responseJsonBody.D
    let reward = responseJsonBody.reward
    return {
        id,
        question,
        a,
        b,
        c,
        d,
        reward,
        display: true
    }
}