interface GoResponse {
    message?: string
    error?: string
}

interface GoTokenResponse extends GoResponse {
    token?: string
}