import { Cookies } from "@sveltejs/kit"
if (Cookies.get("jwt") != "") {
    throw new redirect(300, "/game")
}