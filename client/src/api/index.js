import Axios from "axios";

const axios = Axios.create({
    baseURL: "http://127.0.0.1:5000/api"
})

export default class API {
    static async sendMessage(text) {
        await axios.post("/messages", JSON.stringify({text}))
    }

    static async getMessage(id) {
        return (await axios.get(`/messages/${id}`)).data
    }
}
