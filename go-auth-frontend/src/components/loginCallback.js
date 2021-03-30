import axios from 'axios'
import React, {useEffect} from 'react'

const LoginCallback = ({location, uri}) => {
    const [render, setRender] = React.useState([])

    useEffect(()=>{
        axios.get(`http://localhost:4000/${uri}${location.search}`).then(res=>{

        console.log(res.data)
            setRender([
            
            <div style={{color:"red"}}>Response received from backend</div>,
             <p>{JSON.stringify(res.data).split(',').map(d=><h3 style={{color:"blue" }}>{d}</h3>)}</p>
            
            ])
        }
        )

        console.log(render)
    },[])
    
    return (
        <>
        <p>
            <span style={{color:"red"}}>Authorization Code received</span> <br/>
            {location.search}
        </p>
        {render}
        </>
    )
}

export default LoginCallback
