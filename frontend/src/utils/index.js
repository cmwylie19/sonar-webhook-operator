import axios from "axios";

const idTokenHelper = () => {
  const { token } = JSON.parse(localStorage['token'] || '{"token":"test"}')
  return token
}

export default class API {

  constructor(){
     this.instance = axios.create({
      baseURL: process.env.REACT_APP_BACKEND_URL,
      timeout: 5000,
      responseType: 'json'
      // headers: {'Authorization': `Bearer ${idTokenHelper}`}
    });
    
  }

  fetchResults(){
    return this.instance.get("/webhook/x/results")
  }

  checkToken(){
    alert("Check Token")
  }

   async pingHealthCheck() {
     let response = await this.instance.get("/webhook/healthz")
     return response
   }

   async createUser(email, password) {
     let response = await this.instance.post("/webhook/create",{email, password})
     return response
   }

  async authenticateUser(email, password){
    this.instance.post("/webhook/authenticate",{
      email,
      password
    })
  }
}
