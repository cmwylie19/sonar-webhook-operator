import axios from 'axios'

const jwtConfig = (token) => {
    headers: {
       Authorization: "Bearer " + token
    }
 }
const fetchSonarResults = () => {
    return axios.get(URL, config)
}

const fetchToken = () => {
    let body = {
        
    }
    return axios.post(URL, config, body)
}