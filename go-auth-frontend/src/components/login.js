import React, {useCallback} from 'react'
import doc from "../oauth-config.json"
import axios from 'axios'

const Login = () => {
    const click = useCallback(async (d) => {
        try {
           const data =  await axios.get(`http://localhost:4000/${d.backend_url}`)

           console.log(data.data)
           window.location.assign(data.data)
        } catch (error) {
            
        }
    }, []) 
    const getButton = (d) => {
        return (
          <span>
            <img src={d.icon} onClick={() => click(d)} height="70" width="70" style={{ padding: 20 }} />
            <h3>{d.provider}</h3>
          </span>
        )
      }

    const elts = doc.filter(d=> d.status==="active").map(d => getButton(d))
    console.log(doc)
    return (
        <div>
            <div style={{ display: "flex", justifyContent: "center", marginTop: 100 }}>
            <p>Login Options</p>
                {elts}
        </div>
        </div>
    )
}

export default Login
