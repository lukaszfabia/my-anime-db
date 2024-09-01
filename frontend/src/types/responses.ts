// add generic data to GoRespose
export interface GoResponse {
    code: number
    msg: string
    data?: any
}