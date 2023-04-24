export const TOKEN_STORAGE_KEY = "worldchatGoJWT";

export async function isAuthorized() : Promise<boolean> {
    const token = localStorage.getItem(TOKEN_STORAGE_KEY)
    if (token) {
        const response = await fetch("http://localhost:8080/auth", {
            method: "post",
            headers: {
                "Accept": "application/json"
            },
            body: JSON.stringify({
                token,
            })
        })
        if (response.ok) {
            return true
        }
    }
    return false
}