import React from 'react'
import doc from "./oauth-config.json"
import "axios"
import axios from 'axios'
const BACKEND_URI = "http://localhost:4000/"


const App = () => {
  const [uri, setUri] = React.useState(null)
  const click = async (d) => {
    console.log(d)
    try {

      const resp = await axios.get(BACKEND_URI + d.backend_url)

      console.log(resp.data)
      setUri(resp.data)


    } catch (error) {
      console.log(error)
    }
  }

  const getButton = (d) => {
    return (
      <span>
        <img src={d.icon} onClick={() => click(d)} height="70" width="70" style={{ padding: 20 }} />
        <h3>{d.provider}</h3>
      </span>
    )
  }
  const elts = doc.map(d => getButton(d))
  return (
    <div >
      {uri ?
        <>
          <h2>Click me</h2>
          <br />
          <a href={uri}>{uri}</a>
        </>
        :

        <div style={{ display: "flex", justifyContent: "center", marginTop: 100 }}>
          <p>Login Options</p>
          {elts}
        </div>
      }
    </div>
  )
}

export default App
